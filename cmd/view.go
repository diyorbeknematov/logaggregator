/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"logaggregator/models"

	"github.com/spf13/cobra"
)

var flags models.Filter

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View logs with filtering options",
	Long: `View logs with filtering options based on various criteria such as timestamp,
log level, user ID, and service name. For example:

logaggregator view -f "2024-09-01" -t "2024-09-13" -l "ERROR" -u "123" -s "auth-service"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("view called", flags.TimestampFrom, flags.TimestampTo, flags.Level, flags.Service, flags.UserID)

		logs, err := SQLite.LogRepository().GetLogs(flags)
		if err != nil {
			log.Println("Loglarni olishda xatolik kelib chiqdi: ", err)
			return
		}

		for _, log := range logs {
			fmt.Printf("Time: %s\nService Name: %s\nUser ID: %s\nLevel: %s\nMessage: %s\nError: %s\n\n", log.Timestamp, log.Service, log.UserID, log.Level, log.Message, log.Error)
		}
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)

	// Here you will define your flags and configuration settings.
	
	viewCmd.Flags().StringVarP(&flags.TimestampFrom, "from", "f", "", "Start date for filtering logs (e.g., 2024-09-01)")
	viewCmd.Flags().StringVarP(&flags.TimestampTo, "to", "t", "", "End date for filtering logs (e.g., 2024-09-13)")
	viewCmd.Flags().StringVarP(&flags.Level, "level", "l", "", "Filter logs by log level (e.g., INFO, ERROR)")
	viewCmd.Flags().StringVarP(&flags.UserID, "user", "u", "", "Filter logs by user ID")
	viewCmd.Flags().StringVarP(&flags.Service, "service", "s", "", "Filter logs by service name")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
