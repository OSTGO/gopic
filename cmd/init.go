/*
Copyright © 2022 https://longtao.fun

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopic/cmdIementaion"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init env",
	Long:  `init env,please use root privilege`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmdIementaion.CmdInit())
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
