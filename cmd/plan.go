/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// planCmd represents the plan command
var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Manages the plans you have in setup.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("plan called")
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a plan",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("plan create")
		fmt.Println("test")
	},
}

func init() {
	rootCmd.AddCommand(planCmd)

	planCmd.AddCommand(createCmd)
}
