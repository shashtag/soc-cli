/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package createemail

import (
	"fmt"

	"github.com/spf13/cobra"
)

var parentFolder string

// wsgmCmd represents the wsgm command
var wsgmCmd = &cobra.Command{
	Use:   "wsgm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wsgm called")
	},
}

func addCmd() {
	CreateEmailCmd.AddCommand(wsgmCmd)
}

func init() {
	addCmd()
	wsgmCmd.Flags().StringVarP(&parentFolder, "parentFolder", "p", "", "Full path to the folder where the excel files are located")

	if err := wsgmCmd.MarkFlagRequired("parentFolder"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wsgmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wsgmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}