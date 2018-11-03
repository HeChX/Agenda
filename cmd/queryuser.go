package cmd

import (
	"fmt"

	"github.com/HeChX/agenda/service"
	"github.com/spf13/cobra"
)

var queryuserCmd = &cobra.Command{
	Use:   "queryuser -n [name]",
	Short: "query a user through username",
	Long:  `Query a user through username`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if service.Logined {
			service.QueryUser(name)
		} else {
			fmt.Println("Please log in first!")
		}
	},
}

func init() {
	rootCmd.AddCommand(queryuserCmd)
	queryuserCmd.Flags().StringP("name", "n", "", "user name")
}
