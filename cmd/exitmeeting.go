package cmd

import (
	"fmt"

	"agenda/service"

	"github.com/spf13/cobra"
)

var exitmeetingCmd = &cobra.Command{
	Use:   "exitmeeting -t [title]",
	Short: "exit a meeting",
	Long:  `Provite the title of a meeting to exit it`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		if service.Logined {
			service.ExitMeeting(title)
		} else {
			fmt.Println("Please log in first!")
		}

	},
}

func init() {
	rootCmd.AddCommand(exitmeetingCmd)
	exitmeetingCmd.Flags().StringP("title", "t", "", "meeting title")

}
