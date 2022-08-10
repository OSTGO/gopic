/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopic/cmdIementaion"
)

var recurse bool // 递归查找
var covertPath, outDir string

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert pics",
	Long:  `convert pics`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmdIementaion.CmdConvert(covertPath, outDir, outFormat, allStorage, nameReserve, recurse, storageList)
		if err != nil {
			panic(err)
		}
		fmt.Println("convert success!")
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringVarP(&covertPath, "covertPath", "c", "./", "")
	convertCmd.Flags().BoolVarP(&allStorage, "all", "a", false, "")
	convertCmd.Flags().BoolVarP(&recurse, "recurse", "r", false, "")
	convertCmd.Flags().StringSliceVarP(&storageList, "storage", "s", nil, "")
	convertCmd.Flags().StringVarP(&outFormat, "format", "f", "", "")
	convertCmd.Flags().StringVarP(&outDir, "dir", "d", "", "")
	convertCmd.Flags().BoolVarP(&nameReserve, "name", "n", false, "")
	convertCmd.MarkFlagRequired("dir")
}
