package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
	"time"

	"github.com/kr/fs"
)

func main() {
	h := os.Getenv("HOME")
	w := fs.Walk(h)
	now := time.Now()
	var files []string

	log := os.Getenv("HOME") + "/.pomelo.log"
	msg := fmt.Sprintf("\nNew Entry Date %s\n", now)
	err := ioutil.WriteFile(log, []byte(msg), 0644)
	if err != nil {
		fmt.Println("Pomelo could not write to the log file.")
		os.Exit(1)
	}

	for w.Step() {
		if err := w.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		f := w.Path()
		info, err := os.Stat(f)
		if err != nil {
			fmt.Printf("Could not retrieve file info for %s\n", f)
			continue
		} else {
			fmt.Printf("Retrieved %s\n", f)
		}

		stat := info.Sys().(*syscall.Stat_t)
		atime := time.Unix(int64(stat.Atimespec.Sec), int64(stat.Atimespec.Nsec))
		diff := now.Sub(atime)

		if diff.Hours() > 720 {
			files = append(files, f)
		}
	}

	numFiles := fmt.Sprintf("Number of files 30 days and older: %d\n", len(files))
	fmt.Println(numFiles)

	f, err := os.OpenFile(log, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Pomelo could not write to the log file.")
		os.Exit(1)
	}
	defer f.Close()

	for k, v := range files {
		msg := fmt.Sprintf("%d. %s\n", k+1, v)
		if _, err := f.WriteString(msg); err != nil {
			fmt.Printf("Pomelo could not write %s to the log file.", v)
			continue
		}
	}

	fmt.Println("Pomelo complete")
}
