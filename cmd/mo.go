/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var namespaces = []string{
	"kube-system",
	"suraj",
}

// moCmd represents the mo command
var moCmd = &cobra.Command{
	Use:   "mo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("mo called")

		fmt.Println("args", args)
		nn, err := cmd.Flags().GetStringSlice("nn")
		if err != nil {
			panic(err)
		}
		// fmt.Println("nn", nn)
		// if err := cmd.Flags().Set("nn", ""); err != nil {
		// 	panic(err)
		// }
		// cmd.Flags().Visit(func(f *pflag.Flag) {
		// 	if f.Name == "nn" {
		// 		f.Value.Set("")
		// 	}
		// })

		for _, n := range nn {
			var stdout bytes.Buffer
			var stderr bytes.Buffer
			a := []string{}
			a = append(a, "-n", n)
			a = append(a, os.Args[2:]...)
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

	},
}

func init() {
	moCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.SetHelpCommand(&cobra.Command{
		Hidden: true,
	})

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// moCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// moCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	moCmd.Flags().StringSliceP("nn", "", []string{}, "Specify comma separated list of namespaces")
}
