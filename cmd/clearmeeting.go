package cmd

import (
	"fmt"

	"agenda/service"

	"github.com/spf13/cobra"
)

var clearmeetingCmd = &cobra.Command{
	Use:   "clearmeeting",
	Short: "clear all meeting you call",
	Long:  `clear all meeting you call`,
	Run: func(cmd *cobra.Command, args []string) {
		if service.Logined {
			service.ClearMeeting()
		} else {
			fmt.Println("You don't log in!")
		}

	},
}

func init() {
	rootCmd.AddCommand(clearmeetingCmd)
}
