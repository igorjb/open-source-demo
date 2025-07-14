/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getLicenseCmd represents the getLicense command
var getLicenseCmd = &cobra.Command{
	Use:   "getLicense",
	Short: "Retrieve the open source license for the project",
	Long: `This is a command that can help you retrieve the open source license for the project.
It provides details about the license type and any relevant information.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getLicense called")
	},
}

func init() {
	rootCmd.AddCommand(getLicenseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getLicenseCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getLicenseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
