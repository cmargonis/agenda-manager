
package configuration
import (
	"strings"
	"fmt"
	"github.com/spf13/viper"
	"flag"
	"github.com/spf13/pflag"
)

type FileConfiguration struct {
	token       string
	channelName string
}

func (fc *FileConfiguration) GetToken() string {
	if fc.token != "" {
		return fc.token
	}

	// using standard library "flag" package
	c, err := getConfiguration()
	if err != nil {
		panic(fmt.Errorf("Invalid or no configuration provided. Use -c flag to pass filename. %s \n", err))
	}
	fc.token = c.token
	return fc.token
}

func NewFileConfiguration() *FileConfiguration {
	return &FileConfiguration{}
}

func getConfiguration() (*FileConfiguration, error) {
	initializeViper()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &FileConfiguration{token: viper.GetString("token")}, nil
}

func initializeViper() {
	flag.String("c", "settings.yaml", "Supply file containing configuration parameters.")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	fileName := viper.GetString("c")
	viper.SetConfigName(strings.TrimSuffix(fileName, ".yaml"))
	viper.AddConfigPath(".")
}
