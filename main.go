package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var version = "1.0.0"
var configFile = "config.yaml"

func config(name string) string {
	vi := viper.New()
	vi.SetConfigFile(configFile)
	vi.ReadInConfig()
	return vi.GetString(name)
}

func wingsBase(use string, short string, long string, run []string) *cobra.Command {
	wingsBase := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(cmd *cobra.Command, args []string) {
			command := exec.Command(run[0], run[1], run[2])
			out, err := command.CombinedOutput()
			fmt.Printf("%s\n", out)
			if err != nil {
				fmt.Println(err)
			}

		},
	}
	return wingsBase
}

var (
	// Persistent Flags
	verboseFlag bool

	// Local Flags

	// General root command, returns information.
	rootCmd = &cobra.Command{
		Use:   "ptero",
		Short: "Pterodactyl CLI",
		Long:  "A CLI for the Pterodactyl panel & daemon that allows you to easily control and manage it.\nCreated by: INfoUpgraders#0001 - https://github.com/infoupgraders/pterodactyl-cli/",
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Script version",
		Long:  "Displays the scripts current version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Current version:", version)
		},
	}

	wingsCmd = &cobra.Command{
		Use:   "wings",
		Short: "All wings commands",
	}

	wingsStatusCmd = wingsBase("status", "Status of wings", "Displays the status of wings", []string{"systemctl", "status", config("daemon_name")})

	wingsRestartCmd = wingsBase("restart", "Restart wings", "Restarts the wings", []string{"systemctl", "restart", config("daemon_name")})
)

func init() {
	rootCmd.AddCommand(wingsCmd)
	rootCmd.AddCommand(versionCmd)
	wingsCmd.AddCommand(wingsRestartCmd)
	wingsCmd.AddCommand(wingsStatusCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
