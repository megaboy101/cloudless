package main

import (
	"fmt"

	"github.com/megaboy101/cloudless-cli/pkg/auth"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := auth.Signup(args[0], args[1]); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Registration successful!")
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
