/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package info

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"net/http"
	"os/exec"
)

var (
	Input string
)

// AddrCmd represents the addr command
var AddrCmd = &cobra.Command{
	Use:   "addr",
	Short: "Your address and related information",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		AddrInfo()
	},
}

func AddrInfo() {
	byte, err := exec.Command("curl", "ifconfig.me").Output()
	if err != nil {
		color.Red("Couldn't get address. Please install 'curl'")
		return
	}

	ipAddress := string(byte)
	resp, err := http.Get("http://ip-api.com/json/" + ipAddress)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var location LocationInfo
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		color.Red("Error decoding JSON:", err)
		return
	}
	fmt.Printf("Country: %s\n", location.Country)
	color.Green("Your IP Address: %s\n", location.Query)
	fmt.Printf("City: %s\n", location.City)
	fmt.Printf("Region: %s\n", location.Region)
	fmt.Printf("Latitude: %.4f\n", location.Lat)
	fmt.Printf("Longitude: %.4f\n", location.Lon)
	fmt.Printf("ISP: %s\n", location.ISP)
	fmt.Printf("Timezone: %s\n", location.Timezone)
	fmt.Printf("Zip Code: %s\n", location.ZipCode)

}
func init() {
	AddrCmd.Flags().StringVarP(&Input, "address", "a", "", "getting address information")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addrCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addrCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
