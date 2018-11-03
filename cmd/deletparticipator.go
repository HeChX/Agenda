package cmd

import (
	"fmt"

	"agenda/service"

	"github.com/spf13/cobra"
)

var deleteparticipatorCmd = &cobra.Command{
	Use:   "deleteparticipator -n [name] -t [meeting title]",
	Short: "remove a participator from a meeting",
	Long:  `Provide the name of the user to romove it from the meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		title, _ := cmd.Flags().GetString("title")
		if service.Logined {
			service.DeleteParticipator(title, name)
		} else {
			fmt.Println("Please log in first!")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteparticipatorCmd)
	deleteparticipatorCmd.Flags().StringP("name", "n", "", "user name")
	deleteparticipatorCmd.Flags().StringP("title", "t", "", "meeting title")

}
