package cmd

import (
	"fmt"

	"github.com/HeChX/agenda/service"
	"github.com/spf13/cobra"
)

var querymeetingCmd = &cobra.Command{
	Use:   "querymeeting -s [start time] -e [end time]",
	Short: "query meeting",
	Long:  `View meetings within a specified time period`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		starttime, _ := cmd.Flags().GetString("start")
		endtime, _ := cmd.Flags().GetString("end")
		if service.Logined {
			service.QueryMeeting(starttime, endtime)
		} else {
			fmt.Println("Please log in first!")
		}

	},
}

func init() {
	rootCmd.AddCommand(querymeetingCmd)
	querymeetingCmd.Flags().StringP("start", "s", "", "start time")
	querymeetingCmd.Flags().StringP("end", "e", "", "end time")

}
