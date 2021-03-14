package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/ghodss/yaml"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ncsnw/serpentor"
)

const (
	ConfigType     = "yaml"
	ConfigFileName = "serpentor-example"
)

func main() {
	serpentor.NewCmd("example").
		Short("An example program using serpentor with cobra and viper").
		Long("This is an example program that demonstrates how to use serpentor to construct a cobra program that usese viper for configuration.").
		AddCommand(
			serpentor.NewCmd("config").
				Short("Commands for interacting with configuration").
				AddCommand(
					serpentor.NewCmd("show").
						Short("Print the current config to stdout").
						Run(showConfig),
					serpentor.NewCmd("get key").
						Short("Get the key's value from the config").
						Run(getConfig),
					serpentor.NewCmd("set key value").
						Short("Set the config key to the specified value").
						Run(setConfig),
				),
		).
		Execute()
}

func showConfig(_ *cobra.Command, _ []string) {
	file, err := ioutil.ReadFile(GetConfigFileFullPath())
	cobra.CheckErr(err)
	fmt.Println(string(file))
}

func getConfig(_ *cobra.Command, args []string) {
	cobra.ExactArgs(1)
	bitties, err := yaml.Marshal(viper.Get(args[0]))
	cobra.CheckErr(err)
	fmt.Println(string(bitties))
}

func setConfig(_ *cobra.Command, args []string) {
	cobra.ExactArgs(2)
	fmt.Printf("Setting %s=%s\n", args[0], args[1])
	viper.Set(args[0], args[1])
	cobra.CheckErr(viper.WriteConfig())
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func GetHome() string {
	home, err := homedir.Dir()
	cobra.CheckErr(err)
	return home
}

func GetConfigFileFullPath() string {
	return path.Join(GetHome(), fmt.Sprintf("%s.%s", ConfigFileName, ConfigType))
}

func init() {
	if !Exists(GetConfigFileFullPath()) {
		_, err := os.Create(GetConfigFileFullPath())
		cobra.CheckErr(err)
	}
	viper.AddConfigPath(GetHome())
	viper.SetConfigType(ConfigType)
	viper.SetConfigName(ConfigFileName)
	cobra.CheckErr(viper.ReadInConfig())
}
