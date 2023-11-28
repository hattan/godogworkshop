package cmd

import (
	client "github.com/hattan/godogcli/pkg/httpclient"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the doggos",
	Long:  `This command calls the godog api and shows a listing of all dogs`,
	Run: func(cmd *cobra.Command, args []string) {
		dogs := client.GetDogs()
		for _, dog := range dogs {
			dog.Display()
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
