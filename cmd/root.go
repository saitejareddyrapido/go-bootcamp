package main

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "rate",
	Short: "A simple CLI to manage product ratings",
}

func init() {
	rootCmd.AddCommand(addRatingCmd)
	rootCmd.AddCommand(getRatingsCmd)
}
