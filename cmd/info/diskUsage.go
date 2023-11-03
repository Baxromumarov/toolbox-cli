// Package info /*
package info

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/ricochet2200/go-disk-usage/du"
	"github.com/spf13/cobra"
)

// DiskUsageCmd represents the diskUsage command
var DiskUsageCmd = &cobra.Command{
	Use:   "diskUsage",
	Short: "Information about current directory usage and total",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		usage := du.NewDiskUsage(".")
		color.Yellow("Disk usage of current directory:")
		fmt.Printf("Currency directory usage: %f %%\nAvailable: %.f GB\nTotal Size: %.f\nTotal Usage: %.f GB\n",
			usage.Usage(),
			float64(usage.Available())/float64(1024*1024*1024),
			float64(usage.Size())/float64(1024*1024*1024),
			float64(usage.Used())/float64(1024*1024*1024),
		)
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diskUsageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diskUsageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
