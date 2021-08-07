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
	abs, _ := filepath.Abs(".")
	viper.SetDefault("path.download", abs)
	viper.SetDefault("log", "info")
	return viper.WriteConfigAs(filepath.Join(configPath, "config.yaml"))
}

func LoadConfig() error {
	// 查看全部环境变量
	// environ := os.Environ()
	// for i := range environ {
	// 	fmt.Println(environ[i])
	// }
	configPath := os.Getenv("CRYPT_SYSTEM_CONFIG_PATH")
	if configPath == "" {
		configPath = ".crypto-system/"
	}
	viper.SetConfigType("yaml")
	viper.SetConfigFile(filepath.Join(configPath, "config.yaml"))

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&App.Config); err != nil {
		return err
	}

	return nil
}
