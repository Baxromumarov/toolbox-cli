/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package info

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

var Source string

// WeatherCmd represents the weather command
var WeatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "All countries weather information",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		GetWeather(os.Args[len(os.Args)-1])
	},
}

func init() {
	WeatherCmd.Flags().StringVarP(&Source, "country", "c", "", "getting weather information by country name")
	WeatherCmd.Flags().StringVarP(&Source, "Location", "l", "", "getting weather information by city name")

	//if err := WeatherCmd.MarkFlagRequired("weather"); err != nil {
	//	fmt.Println(err)
	//}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// weatherCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags, which will only run when this command
	// is called directly, e.g.:
	// weatherCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func GetWeather(q string) {
	if q == "" || strings.ToLower(q) == "weather" {
		q = "Tashkent"
	}
	URL := "http://api.weatherapi.com/v1/forecast.json?key=82f5a4e121014bb09cc82302230211&q=" + q + "&days=1&aqi=yes&alerts=yes"
	res, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("Weather API not available")
		return
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var weather Weather
	if err := json.Unmarshal(body, &weather); err != nil {
		panic(err)
	}
	location, current, hours := weather.Location, weather.Current, weather.Forecast.ForecastDat[0].Hour
	color.Green(fmt.Sprintf(
		"%s, %s: %.0fC, %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
	))

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)
		if date.Before(time.Now()) {
			continue
		}
		message := fmt.Sprintf(
			"%s - %.0fC, %.0f%%, %s",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)
		if hour.ChanceOfRain < 40 {
			fmt.Println(message)
		} else {
			color.Red(message)
		}
	}
}
