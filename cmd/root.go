package cmd

import (
	"fmt"
	"os"

	"github.com/kczapla/learning-cobra/cmd/calc"
	"github.com/kczapla/learning-cobra/cmd/greet"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "learning-cobra",
		Short: "learning cobra because i want to dig into k8s code",
		Long:  "I am just going to add bunch of commands to get a feeling of the framework",
		Run: func(cmd *cobra.Command, args []string) {
			//random stuff
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	fmt.Println("root cmd init cnfg")
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name of copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")

	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE")
	viper.SetDefault("license", "apache")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(greet.GreetCmd)
	rootCmd.AddCommand(calc.CalcCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")

		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}
}
