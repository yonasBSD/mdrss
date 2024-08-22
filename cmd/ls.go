/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
  mdrss "github.com/TimoKats/mdrss/lib"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
    config_file, _ := cmd.Flags().GetString("config")
    config, _ := mdrss.ReadConfig(config_file)
    files, fileErr := mdrss.GetArticles(config)
    if fileErr == nil {
      for _, file := range files {
        mdrss.Info.Println(file)
      }
    }
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
