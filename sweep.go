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

var multiflag strslice

func main() {
	flag.Var(&multiflag, "v", "Specify the branch you want to exclude")
	flag.Parse()

	if len(multiflag) == 0 {
		multiflag = append(multiflag, "master")
	}

	var b, e string

	for _, s := range multiflag {
		b = s
		e = e + " | grep -v " + s
	}
	cmdstr := "git checkout " + b + " && git branch" + e + " | xargs git branch -D"

	err := exec.Command("sh", "-c", cmdstr).Run()
	if err != nil {
		fmt.Printf("Error! Failed to sweep branches.\n")
	}
}
