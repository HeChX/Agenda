package cmd

import (
	"fmt"

	"agenda/service"
	"github.com/spf13/cobra"
)

var logoffCmd = &cobra.Command{
	Use:   "logoff",
	Short: "log off your account",
	Long:  `Log off your account`,
	Run: func(cmd *cobra.Command, args []string) {
		if service.Logined {
			service.DeleteUser()
		} else {
			fmt.Println("Please log in first!")
		}
	},
}

func init() {
	rootCmd.AddCommand(logoffCmd)
}
