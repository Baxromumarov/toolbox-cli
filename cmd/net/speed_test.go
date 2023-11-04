/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// netCmd represents the net command
var SpeedCmd = &cobra.Command{
	Use:   "speed",
	Short: "Net is a palette that contains network based commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// Create a new timer
	timer := time.NewTimer(time.Second * 10)

	// Create a new HTTP client
	client := &http.Client{}

	// Make a request to the speedtest.net server
	req, err := http.NewRequest("GET", "https://speedtest.net/random1000MB.zip", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Start the timer
	// timer.Start()

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Close the response body
	defer resp.Body.Close()
	numDots := 1
	numIterations := 10 // Total number of iterations
	for i := 0; i < numIterations; i++ {
		if numDots == 3 {
			fmt.Print("\r   \r") // Erase 3 dots with spaces using carriage return (\r)
			fmt.Print(".")
			numDots = 1
		} else {
			fmt.Print(".")
			numDots++
		}

		time.Sleep(1 * time.Second)
	}
	fmt.Print("\r   \r")

	// Wait for the timer to expire
	<-timer.C

	// Reset the timer
	timer.Reset(time.Second * 10)

	// Get the current time
	current := time.Now()

	// Calculate the download speed
	size := resp.ContentLength
	elapsed := time.Since(current)
	speed := float64(size) / elapsed.Seconds() / 1024 / 32

	// Print the download speed
	fmt.Printf("Your internet speed: %.2f Mbps\n", speed*(-1))
}
