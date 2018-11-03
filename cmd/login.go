package cmd

import (
	"agenda/service"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login -n [name] -p [password]",
	Short: "log in",
	Long:  `Provide registered, correct user name and password for login`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		password, _ := cmd.Flags().GetString("password")
		service.Login(name, password)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("name", "n", "", "name")
	loginCmd.Flags().StringP("password", "p", "", "password")
}
