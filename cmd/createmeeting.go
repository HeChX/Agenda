package cmd

import (
	"fmt"

	"agenda/service"

	"github.com/spf13/cobra"
)

var createmeetingCmd = &cobra.Command{
	Use:   "createmeeting -t [title] -s [starttime] -e [endtime] -p [participator]",
	Short: "create a meeting",
	Long:  `Provide the information of the meeting to create it`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		starttime, _ := cmd.Flags().GetString("start")
		endtime, _ := cmd.Flags().GetString("end")
		participator, _ := cmd.Flags().GetStringArray("participator")
		if service.Logined {
			service.CreateMeeting(title, starttime, endtime, participator)
		} else {
			fmt.Println("Please log in first!")
		}
	},
}

func init() {
	rootCmd.AddCommand(createmeetingCmd)
	createmeetingCmd.Flags().StringP("title", "t", "", "meeting title")
	createmeetingCmd.Flags().StringP("start", "s", "", "meeting start time")
	createmeetingCmd.Flags().StringP("end", "e", "", "meeting end time")
	createmeetingCmd.Flags().StringArrayP("participator", "p", []string{}, "meeting participator")
}
