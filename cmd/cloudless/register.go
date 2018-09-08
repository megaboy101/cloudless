package main

import (
	"log"

	"github.com/spf13/viper"

	"github.com/megaboy101/cloudless/pkg/user"
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
		// Initialize user service
		us, err := user.New()

		if err != nil {
			log.Fatalf("Error initializing user service: %v", err)
		}

		// Register a new user externally
		user, err := us.Register(args[0], args[1], args[2])

		if err != nil {
			log.Fatalf("Error creating user: %v", err)
		}

		// Save user credentials locally to a config file
		viper.Set("User", user)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		SaveConfig()

		log.Println("New user created and saved locally!")
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
