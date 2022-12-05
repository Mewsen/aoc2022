package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	mkDailyDir()
	goModInit(getCurrentDayDirName() + "/1")
	goModInit(getCurrentDayDirName() + "/2")
	copyFile(getCurrentDayDirName() + "/1")
	copyFile(getCurrentDayDirName() + "/2")
}

func goModInit(path string) error {
	cmd := exec.Command("go", "mod", "init")
	cmd.Dir = path
	_, err := cmd.CombinedOutput()
	return err
}

func mkDailyDir() (err error) {
	err = os.MkdirAll(getCurrentDayDirName()+"/1", 0777)
	err = os.Mkdir(getCurrentDayDirName()+"/2", 0777)
	return
}

func getCurrentDayDirName() string {
	t := time.Now()
	return fmt.Sprint("day", t.Day())
}

func copyFile(out string) (int64, error) {
	i, err := os.Open("template/main.go")
	if err != nil {
		return 0, err
	}
	defer i.Close()
	o, err := os.Create(fmt.Sprint(out, "/main.go"))
	if err != nil {
		return 0, err
	}
	defer o.Close()
	return o.ReadFrom(i)
}
