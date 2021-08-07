package context

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Server *Server `yaml:"server"`
	Path   *Path   `yaml:"path"`
	Log    string  `yaml:"log"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Path struct {
	Config   string `yaml:"config"`
	Download string `yaml:"download"`
}

func GenerateConfig() error {
	configPath := os.Getenv("CRYPT_SYSTEM_CONFIG_PATH")
	viper.AddConfigPath(configPath)
	viper.SetDefault("server.host", "192.168.50.219")
	viper.SetDefault("server.port", 9090)
	viper.SetDefault("path.config", configPath)
	viper.SetDefault("path.download", "./")
	viper.SetDefault("log", "info")
	return viper.WriteConfigAs("config.yaml")
}

func LoadConfig() error {
	configPath := os.Getenv("CRYPT_SYSTEM_CONFIG_PATH")
	viper.SetConfigType("yaml")
	viper.SetConfigFile(filepath.Join(configPath, "config.yml"))

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return viper.Unmarshal(App.Config)

}
