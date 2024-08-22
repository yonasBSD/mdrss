/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
  "reflect"

  mdrss "github.com/TimoKats/mdrss/lib"
	"github.com/spf13/cobra"
)

// confCmd represents the conf command
var confCmd = &cobra.Command{
	Use:   "conf",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
    config, _ := cmd.Flags().GetString("config")
    configValues := reflect.ValueOf(config)
    typeOfS := configValues.Type()
    for i := 0; i < configValues.NumField(); i++ {
      if len(typeOfS.Field(i).Name) < 8 {
        mdrss.Info.Printf("%s\t\t%v\n", typeOfS.Field(i).Name, configValues.Field(i).Interface())
      } else {
        mdrss.Info.Printf("%s\t%v\n", typeOfS.Field(i).Name, configValues.Field(i).Interface())
      }
    }
	},
}

func init() {
	rootCmd.AddCommand(confCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// confCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// confCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
