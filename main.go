package main

import (
	"github.com/spf13/cobra"
	"github.com/zhangga/ximalaya/cmd"
	"log"
)

var rootCmd = &cobra.Command{
	Use:     "ximalaya",
	Short:   "A simple CLI for ximalaya",
	Long:    "A simple CLI for ximalaya",
	Version: "0.0.1",
	Run:     rootRun,
}

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	rootCmd.AddCommand(cmd.RenameCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func rootRun(cmd *cobra.Command, args []string) {

}
