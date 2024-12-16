package greet

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	GreetCmd = &cobra.Command{
		Use:   "greet",
		Short: "echo greet message",
		Long:  "just print boring greeting message",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			greeted := strings.Join(args, " ")
			if viper.GetBool("greet.uppercase") {
				greeted = strings.ToUpper(greeted)
			} else if viper.GetBool("lowercase") {
				greeted = strings.ToLower(greeted)
			}
			fmt.Println("hello", greeted)
		},
	}

	Uppercase bool
	Lowercase bool
)

func init() {
	fmt.Println("greet init")
	GreetCmd.Flags().BoolVar(&Uppercase, "uppercase", false, "uppercase greeting")
	GreetCmd.Flags().BoolVar(&Lowercase, "lowercase", false, "lowercase greeting")
	GreetCmd.MarkFlagsMutuallyExclusive("uppercase", "lowercase")

	viper.BindPFlag("greet.uppercase", GreetCmd.Flags().Lookup("uppercase"))
	viper.BindPFlag("greet.lowercase", GreetCmd.PersistentFlags().Lookup("lowercase"))
}
