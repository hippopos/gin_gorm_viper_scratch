package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "bear [subcommand]",
	Short:   "bear",
	Long:    `A Restful API server serve for data, and provide an easy way to view logs`,
	Version: "0.1 ",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	v, _ := cmd.PersistentFlags().GetBool("version")
	// 	fmt.Println(v, args)
	// 	// Do Stuff Here
	// },
}

// Execute
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().BoolVarP(&version, "version", "v", false, "Print the version number of scheduler")
	rootCmd.AddCommand(newServeCmd())
}
