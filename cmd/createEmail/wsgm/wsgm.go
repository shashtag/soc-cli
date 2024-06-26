/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package wsgm

import (
	"fmt"
	"log"

	createemail "github.com/shashtag/email-cli/internal/createEmail"
	"github.com/spf13/cobra"
)

// wsgmCmd represents the wsgm command
var WsgmCmd = &cobra.Command{
	Use:   "wsgm",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Combining all the CSV files")
		createemail.CombineCSV()
		fmt.Println("Creating the email")
		table, _, err := createemail.CreateEmail()
		if err != nil {
			log.Fatal("could not create email : ", err)
		}
		fmt.Println("Sending Email")
		createemail.SendEmail(table)
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wsgmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wsgmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
