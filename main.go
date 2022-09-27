/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var namespaces = []string{
	// "kube-system",
	// "suraj",
}

// var nn

// func init() {
// 	flag.String("n", "", "specify list of namespaces")
// }

func main() {
	// command := cmd.NewDefaultKubectlCommand()
	// if err := cli.RunNoErrOutput(command); err != nil {
	// 	// Pretty-print the error and exit with an error.
	// 	util.CheckErr(err)
	// }

	// cmd.Execute()
	args := []string{}
	for _, arg := range os.Args[2:] {
		if strings.Contains(arg, "--namespaces") {
			a := strings.Replace(arg, "--namespaces=", "", 1)
			for _, n := range strings.Split(a, ",") {
				namespaces = append(namespaces, strings.TrimSpace(n))
			}
			args = append(args, "")
			continue
		}

		args = append(args, arg)
	}

	for _, n := range namespaces {
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		a := []string{}
		a = append(a, "-n", n)
		a = append(a, args...)
		c := exec.Command("kubectl", a...)
		fmt.Println("executed command", c.String())
		c.Stderr = &stderr
		c.Stdout = &stdout

		if err := c.Run(); err != nil {
			fmt.Println(stdout.String())

			fmt.Println(stderr.String())
			continue
		}
		fmt.Println(stdout.String())

		fmt.Println("namespace:", n)
	}
}
