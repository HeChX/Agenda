package cmd

import (
	"fmt"

	"agenda/service"

	"github.com/spf13/cobra"
)

var cancelmeetingCmd = &cobra.Command{
	Use:   "cancelmeeting -t [title]",
	Short: "cancel a meeting",
	Long:  `Provide the title of the meeting to cancel it`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		if service.Logined {
			service.CancelMeeting(title)
		} else {
			fmt.Println("Please log in first!")
		}
	},
}

func init() {
	rootCmd.AddCommand(cancelmeetingCmd)
	cancelmeetingCmd.Flags().StringP("title", "t", "", "meeting title")
}
