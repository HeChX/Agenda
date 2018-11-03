package cmd

import (
	"fmt"

	"agenda/service"

	"github.com/spf13/cobra"
)

var addparticipatorCmd = &cobra.Command{
	Use:   "addpaticipator -n [name] -t [title]",
	Short: "add a new participator to a meeting",
	Long:  `Please provide the title of the meeting and the name of the participator`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		title, _ := cmd.Flags().GetString("title")
		if service.Logined {
			service.AddParticipator(title, name)
		} else {
			fmt.Println("Please log in first!")
		}
	},
}

func init() {
	rootCmd.AddCommand(addparticipatorCmd)
	addparticipatorCmd.Flags().StringP("name", "n", "", "user name")
	addparticipatorCmd.Flags().StringP("title", "t", "", "meeting title")
}
