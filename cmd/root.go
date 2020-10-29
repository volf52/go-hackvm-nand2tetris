package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gohvm",
	Short: "HackVM Compiler - Nand2Tetris",
	Long:  "Go version of the HackVM Intermediate Compiler",
	//Run: func(cmd *cobra.Command, args []string){
	//
	//},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
