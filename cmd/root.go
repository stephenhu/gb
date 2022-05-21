package cmd

import (

	"github.com/spf13/cobra"
)


var (

	fLocation			string

	rootCmd = &cobra.Command{
		Use: "gb",
		Short: "gb command line tool",
		Long: "gb is a command line tool for crawling a site and downloading media files",
		Version: "0.1",
	}

)


func init() {

	cobra.OnInitialize()

	rootCmd.AddCommand(crawlCmd)

} // init


func Execute() error {
	return rootCmd.Execute()
} // Execute
