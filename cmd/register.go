// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"agenda/service"

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register -n [username] -p [password] -e [email] -t [phone]",
	Short: "Register a new user",
	Long:  `You can register a new user with a unique user name`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("name")
		userpassword, _ := cmd.Flags().GetString("password")
		useremail, _ := cmd.Flags().GetString("email")
		userphone, _ := cmd.Flags().GetString("phone")
		if service.Logined {
			fmt.Println("Logged in users cannot create new users, please log out first!")
		} else {
			service.RegisterUser(username, userpassword, useremail, userphone)
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	registerCmd.Flags().StringP("name", "n", "Anonymous", "name")
	registerCmd.Flags().StringP("password", "p", "Anonymous", "password")
	registerCmd.Flags().StringP("email", "e", "Anonymous", "email")
	registerCmd.Flags().StringP("phone", "t", "Anonymous", "phone")

}
