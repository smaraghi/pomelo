package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
	"time"

	"github.com/kr/fs"
)

func main() {
	now := time.Now()
	var files []string
	var dir string
	var term int

	flag.StringVar(&dir, "dir", os.Getenv("HOME"), "Specify a directory to recursively search.")
	flag.IntVar(&term, "term", 30, "Specify a term length for checking your files' access times. Default is 30.")
	flag.Parse()

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Println("User specified path does not exist.")
		os.Exit(1)
	}

	log := os.Getenv("HOME") + "/.pomelo.log"
	msg := fmt.Sprintf("\nNew Entry Date %s\n", now)
	err := ioutil.WriteFile(log, []byte(msg), 0600)
	if err != nil {
		fmt.Println("Pomelo could not write to the log file.")
		os.Exit(1)
	}

	w := fs.Walk(dir)
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

		if diff.Hours() > 24*float64(term) {
			files = append(files, f)
		}
	}

	numFiles := fmt.Sprintf("Number of files %d days and older: %d\n", term, len(files))
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

	if len(files) == 0 {
		fmt.Println("No files to delete.\nPomelo complete.")
		os.Exit(0)
	}

	var res string
	fmt.Println("Would you like to delete ALL of these files? [y|N]")
	fmt.Scanln(&res)
	if res == "y" {
		fmt.Println("Are you sure? [y|N]")
		fmt.Scanln(&res)
		if res == "y" {
			fmt.Println("Deleting all files.")
			for _, v := range files {
				os.Remove(v)
			}
		} else {
			fmt.Println("Skipping.")
		}

	} else if res == "n" {
		fmt.Println("Files will not be deleted.")
	} else {
		fmt.Println("Unknown answer, exiting now.")
	}

	fmt.Println("Pomelo complete")
}
