package main

import (
	"flag"
	"fmt"
	"github.com/holygeek/timetext"
	"os"
	"os/exec"
	"time"
)

func main() {
	var file bool
	var noPath bool
	var since int64
	var delta int64
	var compact bool
	var terse bool

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, os.Args[0]+" [-h] [-f|-s|-d] [-n] <arg>")
		flag.PrintDefaults()
	}
	flag.BoolVar(&file, "f", false, "Show age of given file")
	flag.BoolVar(&noPath, "n", false, "Suppress file fullpath output")
	flag.Int64Var(&delta, "d", -1, "Show age for the given seconds")
	flag.Int64Var(&since, "s", -1, "Show age since the given seconds from epoch")
	flag.BoolVar(&compact, "c", false, "Show age in compact format (3m 1s)")
	flag.BoolVar(&terse, "t", false, "Terse output - 3m 0s is shown as only 3m")
	flag.Parse()

	getDuration := timetext.LongDuration
	if terse {
		getDuration = timetext.TerseLongDuration
	}
	if compact {
		getDuration = timetext.Duration
		if terse {
			getDuration = timetext.TerseDuration
		}
	}

	var age int64
	if since >= 0 {
		age = time.Now().Unix() - since
	} else if delta >= 0 {
		age = delta
	} else {
		if flag.NArg() == 0 {
			flag.Usage()
			return
		}

		for _, arg := range flag.Args() {
			var filename string
			if file {
				filename = arg
			} else {
				exe := arg
				cmd := exec.Command("which", exe)
				if output, err := cmd.Output(); err != nil {
					fmt.Fprintln(os.Stderr, err)
					return
				} else {
					filename = string(output[0 : len(output)-1])
				}
			}

			if fi, err := os.Stat(filename); err == nil {
				age = time.Now().Unix() - fi.ModTime().Unix()
				if !noPath {
					fmt.Print(filename + ": ")
				}
				fmt.Println(getDuration(age))
			} else {
				fmt.Fprintln(os.Stderr, err)
				return
			}
		}
		return
	}
	fmt.Println(getDuration(age))
}
