/*
Copyright © 2022 https://longtao.fun

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopic/cmdIementaion"
)

// confCmd represents the conf command
var confCmd = &cobra.Command{
	Use:   "conf",
	Short: "config env",
	Long:  `config env`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmdIementaion.CmdConf())
	},
}

func init() {
	rootCmd.AddCommand(confCmd)
}
