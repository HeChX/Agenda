package cmd

import (
	"fmt"

	"github.com/HeChX/agenda/service"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "log out",
	Long:  `Log out your account`,
	Run: func(cmd *cobra.Command, args []string) {
		if service.Logined {
			service.Logout()
		} else {
			fmt.Println("Please log in first!")
		}

	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
