package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {
	var (
		b = flag.String("exclude", "master", "string flag")
	)
	flag.Parse()
	cmdstr := "git branch | grep -v " + *b + " | xargs git branch -D"
	err := exec.Command("sh", "-c", cmdstr).Run()
	if err != nil {
		fmt.Printf("Error! Failed to sweep branches.\n")
	}
}
