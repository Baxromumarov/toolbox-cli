// Package info /*
package info

import (
	"fmt"
	"github.com/showwin/speedtest-go/speedtest"
	"github.com/spf13/cobra"
	"net/http"
)

// InternetSpeedCmd internetSpeedCmd represents the internetSpeed command
var InternetSpeedCmd = &cobra.Command{
	Use:   "internetSpeed",
	Short: "This command check your internetSpeed",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if IsDeviceConnectedToInternet() {
			CalculateInternetSpeed()

		}
	},
}

func CalculateInternetSpeed() {
	var speedTestClient = speedtest.New()
	serverList, _ := speedTestClient.FetchServers()
	targets, _ := serverList.FindServer([]int{})

	for _, s := range targets {
		// Please make sure your host can access this test server,
		// otherwise you will get an error.
		// It is recommended to replace a server at this time
		s.PingTest(nil)
		s.DownloadTest()
		s.UploadTest()
		fmt.Printf("Latency: %s, Download: %f Mbps, Upload: %f Mbps\n", s.Latency, s.DLSpeed, s.ULSpeed)
		s.Context.Reset() // reset counter
	}

}
func IsDeviceConnectedToInternet() bool {
	if _, err := http.Get("http://www.google.com"); err != nil {
		return false
	} else {
		return true
	}
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
