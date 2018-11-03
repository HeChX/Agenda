package cmd

import (
	"fmt"

	"github.com/HeChX/agenda/service"
	"github.com/spf13/cobra"
)

var queryalluserCmd = &cobra.Command{
	Use:   "queryuser -a",
	Short: "query all uses",
	Long:  `Query the information of all users`,
	Run: func(cmd *cobra.Command, args []string) {
		if service.Logined {
			service.QueryUserAllUser()
		} else {
			fmt.Println("Please log in first!")
		}
	},
}

func init() {
	rootCmd.AddCommand(queryalluserCmd)
}
