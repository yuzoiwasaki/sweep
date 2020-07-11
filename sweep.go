package main

import (
	"flag"
	"fmt"
	"os/exec"
)

type strslice []string

func (s *strslice) String() string {
	return fmt.Sprintf("%v", multiflag)
}

func (s *strslice) Set(v string) error {
	*s = append(*s, v)
	return nil
}

var multiflag strslice = strslice{"master"}

func main() {
	flag.Var(&multiflag, "exclude", "string flag")
	flag.Parse()

	var excludeCmd string

	for _, s := range multiflag {
		excludeCmd = excludeCmd + " | grep -v " + s
	}
	cmdstr := "git branch" + excludeCmd + " | xargs git branch -D"

	err := exec.Command("sh", "-c", cmdstr).Run()
	if err != nil {
		fmt.Printf("Error! Failed to sweep branches.\n")
	}
}
