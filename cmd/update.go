/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
  mdrss "github.com/TimoKats/mdrss/lib"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
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
      config.Articles = mdrss.ReadMarkdown(config, files)
      rssXml := mdrss.CreateRSS(config)
      rssErr := mdrss.WriteRSS(rssXml, config)
      if rssErr != nil {
        return
      }
      mdrss.Info.Printf("Content written to %s", config.OutputFile)
    }
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
