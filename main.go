/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

var namespaces = []string{
	"kube-system",
	"suraj",
}

func main() {
	// command := cmd.NewDefaultKubectlCommand()
	// if err := cli.RunNoErrOutput(command); err != nil {
	// 	// Pretty-print the error and exit with an error.
	// 	util.CheckErr(err)
	// }

	// cmd.Execute()

	for _, n := range namespaces {
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		a := []string{}
		a = append(a, "-n", n)
		a = append(a, os.Args[2:]...)
		fmt.Println("hello")
		c := exec.Command("kubectl", a...)
		fmt.Println("executed command", c.String())
		c.Stderr = &stderr
		c.Stdout = &stdout
		if err := c.Run(); err != nil {
			fmt.Println(stderr.String())
			os.Exit(1)
		}
		fmt.Println("namespace:", n)
		fmt.Println(stdout.String())
	}
}
