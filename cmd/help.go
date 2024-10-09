/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Shows help information for commands",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("help called")

		if len(args) == 0 {
			// Show general help if no specific command is provided
			showGeneralHelp()
		} else {
			// Show specific help based on the provided command
			switch args[0] {
			case "view":
				showViewHelp()
			case "delete":
				showDeleteHelp()
			case "tail":
				showTailHelp()
			default:
				fmt.Printf("Unknown command: %s\n", args[0])
				showGeneralHelp()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func showGeneralHelp() {
	fmt.Println("Usage: logaggregator <command> [flags]")
	fmt.Println("Available commands:")
	fmt.Println("  view    - View logs with filtering options")
	fmt.Println("  delete  - Delete logs based on filters")
	fmt.Println("  help    - Show this help message or help for a specific command")
	fmt.Println("  tail    - Show live log updates")
	fmt.Println("  monitor - Start ")
}

func showViewHelp() {
	fmt.Println("Usage: logaggregator view [flags]")
	fmt.Println("View logs with filtering options.")
	fmt.Println("Flags:")
	fmt.Println("  --from-date <date>    Start date for filtering")
	fmt.Println("  --to-date <date>      End date for filtering")
	fmt.Println("  --service <service>   Filter logs by service name")
	fmt.Println("  --level <level>       Filter logs by log level")
}

func showDeleteHelp() {
	fmt.Println("Usage: logaggregator delete [flags]")
	fmt.Println("Delete logs based on filters.")
	fmt.Println("Flags:")
	fmt.Println("  --service <service>   Delete logs for the given service")
	fmt.Println("  --level <level>       Delete logs with the specified log level")
}

func showTailHelp() {
	fmt.Println("Usage: logaggregator tail [flags]")
	fmt.Println("Show live log updates.")
	fmt.Println("Flags:")
	fmt.Println("  --follow              Follow new log entries")
	fmt.Println("  --lines <number>      Number of lines to show")
}
