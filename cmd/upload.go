/*
Copyright © 2022 https://longtao.fun

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopic/cmdIementaion"
)

var path string
var outFormat string
var allStorage bool
var storageList []string
var nameReserve bool

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload pic list",
	Long:  `upload pic list`,
	Run: func(cmd *cobra.Command, args []string) {
		outURL := cmdIementaion.CmdUpload(storageList, args, allStorage, nameReserve, path, outFormat)
		fmt.Print(outURL)
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	uploadCmd.Flags().StringVarP(&path, "path", "p", "", "")
	uploadCmd.Flags().BoolVarP(&allStorage, "all", "a", false, "")
	uploadCmd.Flags().StringSliceVarP(&storageList, "storage", "s", nil, "")
	uploadCmd.Flags().StringVarP(&outFormat, "format", "f", "", "")
	uploadCmd.Flags().BoolVarP(&nameReserve, "name", "n", false, "")
	uploadCmd.MarkFlagRequired("pathList")
}
