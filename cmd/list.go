package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		listTemplates()
	},
}

func listTemplates() {
	fmt.Printf("Here is the list of templates\n")
	GetContents("https://api.github.com/repos/toptal/gitignore/contents/templates?ref=master")

}

func addInt(args []string) {
	var sum int

	for _, value := range args {
		itemp, err := strconv.Atoi(value)

		if err != nil {
			fmt.Println(err)
		}
		sum += itemp
	}
	fmt.Printf("addition of %s is %d \n", args, sum)
}

func init() {
	rootCmd.AddCommand(listCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
