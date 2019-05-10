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

	for w.Step() {
		if err := w.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		f := w.Path()
		info, err := os.Stat(f)
		if err != nil {
			fmt.Printf("Could not retrieve file info for %s", f)
			continue
		}

		stat := info.Sys().(*syscall.Stat_t)
		atime := time.Unix(int64(stat.Atimespec.Sec), int64(stat.Atimespec.Nsec))
		// printing to avoid lint error
		fmt.Println(atime)
	}
}
