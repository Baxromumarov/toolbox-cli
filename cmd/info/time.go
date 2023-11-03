/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package info

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

// TimeCmd represents the time command
var TimeCmd = &cobra.Command{
	Use:   "time",
	Short: "Time command for current time information",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		GetTimeInfo()
	},
}

func GetTimeInfo() {

	// Time information
	currentTime := time.Now()
	fmt.Printf("Year: %d, Month: %s, Day: %d\n", currentTime.Year(), currentTime.Month(), currentTime.Month())
	fmt.Printf("Hour: %d, minute: %d, second: %d\n", currentTime.Hour(), currentTime.Minute(), currentTime.Second())
	timezoneName, offset := currentTime.Zone()
	fmt.Printf("Time Zone: %s, Offset: %d seconds\n", timezoneName, offset)

}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags, which will work for this command,
	// and all subcommands, e.g.:
	// timeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags, which will only run when this command
	// is called directly, e.g.:
	// timeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
