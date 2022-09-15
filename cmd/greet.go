package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/eriicafes/goctl/utils"
)

// flags
var (
	uppercase bool
)

// args
type Args struct {
	Message string
}

var args = utils.ParseArgs(func(cmd *cobra.Command, rawArgs []string) (Args, error) {
	err := cobra.ExactArgs(1)(cmd, rawArgs)

	if err != nil {
		return Args{}, err
	}

	return Args{Message: rawArgs[0]}, nil
})

// command
var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "A short greeting",
	Args:  args.Validate,
	Run: func(cmd *cobra.Command, _ []string) {
		message := args.Fields.Message

		// capitalise greeting if uppercase flag was set to true
		if uppercase {
			message = strings.ToUpper(message)
		}

		fmt.Println("Hello", message)
	},
}

func init() {
	// declare flags
	greetCmd.Flags().BoolVarP(&uppercase, "upper", "u", false, "convert message to uppercase")

	// bind command
	rootCmd.AddCommand(greetCmd)
}
