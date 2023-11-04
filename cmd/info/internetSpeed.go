// Package info /*
package info

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/showwin/speedtest-go/speedtest"
	"github.com/spf13/cobra"
)

// InternetSpeedCmd internetSpeedCmd represents the internetSpeed command
var InternetSpeedCmd = &cobra.Command{
	Use:   "internetSpeed",
	Short: "This command check your internetSpeed",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if IsDeviceConnectedToInternet() {
			fmt.Println(CalculateInternetSpeed())

		}
	},
}

func CalculateInternetSpeed() InternetSpeedResult {
	go Animation()
	var speedTestClient = speedtest.New()
	serverList, _ := speedTestClient.FetchServers()
	targets, _ := serverList.FindServer([]int{})
	var result = InternetSpeedResult{}
	//fmt.Println(targets)
	for _, s := range targets {
		// Please make sure your host can access this test server,
		// otherwise you will get an error.
		// It is recommended to replace a server at this time
		s.PingTest(nil)
		s.DownloadTest()
		s.UploadTest()

		//color.Blue(fmt.Sprintf("Download: %.2f Mbps", s.DLSpeed))
		//color.Green(fmt.Sprintf("Upload: %.2f Mbps", s.ULSpeed))
		//color.Red(fmt.Sprintf("Latency: %v ns", s.Latency))
		result.DownloadSpeed = s.DLSpeed
		result.UploadSpeed = s.ULSpeed
		result.Latency = s.Latency
		result.Country = s.Country
		result.Name = s.Name
		result.Sponsor = s.Sponsor
		s.Context.Reset() // reset counter
	}

	return result
}
func IsDeviceConnectedToInternet() bool {
	if _, err := http.Get("http://www.google.com"); err != nil {
		return false
	} else {
		return true
	}
}

func Animation() {
	loadingCharacters := []string{"|", "/", "-", "\\"}
	connectingMessages := []string{

		"Configuring DNS...",
		"Obtaining IP Address...",
		"Establishing Connection...",
	}

	for _, message := range connectingMessages {
		color.Yellow(fmt.Sprintf("%s\n", message))
		time.Sleep(2 * time.Second)
	}
	// fmt.Printf("\r                                 \r")

	for i := 0; i < 15000; i++ {
		character := loadingCharacters[i%len(loadingCharacters)]
		fmt.Printf("\rCalculating %s", character)
		time.Sleep(100 * time.Millisecond)

	}

	fmt.Printf("\r                                 \r")

}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags, which will work for this command,
	// and all subcommands, e.g.:
	// internetSpeedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// internetSpeedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
