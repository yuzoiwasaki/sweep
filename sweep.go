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
	flag.Var(&multiflag, "exclude", "Set exclude branch")
	flag.Parse()

	var excludecmd string

	for _, s := range multiflag {
		excludecmd = excludecmd + " | grep -v " + s
	}
	cmdstr := "git branch" + excludecmd + " | xargs git branch -D"

	err := exec.Command("sh", "-c", cmdstr).Run()
	if err != nil {
		fmt.Printf("Error! Failed to sweep branches.\n")
	}
}
