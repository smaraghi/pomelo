package main

import (
	"fmt"
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

	fmt.Printf("Number of files 30 days and older: %d\n", len(files))
}
