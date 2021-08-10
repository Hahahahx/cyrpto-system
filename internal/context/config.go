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
	Ipfs   *Ipfs   `yaml:"ipfs"`
}

type Ipfs struct {
	Host string `yaml:"host"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Path struct {
	Config string `yaml:"config"`
}

func (p *Path) Download() string {
	return filepath.Join(p.Config, "files")
}

func (p *Path) Log() string {
	return filepath.Join(p.Config, "logs")
}
func (p *Path) Cache() string {
	return filepath.Join(p.Config, "caches")
}

func (s *Server) URL(route string) string {
	return s.Host + ":" + s.Port + "/" + route
}

func (i *Ipfs) Api() string {
	return i.Host + ":5001"
}

func (i *Ipfs) Gateway() string {
	return i.Host + ":8080"
}

func (i *Ipfs) GetFileURL(cid string) string {
	return i.Host + ":8080/ipfs/" + cid
}

func GenerateConfig() error {
	configPath := os.Getenv("CRYPT_SYSTEM_CONFIG_PATH")
	viper.AddConfigPath(configPath)
	viper.SetDefault("server.host", "http://192.168.50.219")
	viper.SetDefault("server.port", 9090)
	viper.SetDefault("path.config", configPath)
	viper.SetDefault("log", "info")
	viper.SetDefault("ipfs.host", "http://192.168.50.219")
	// viper.SetDefault("ipfs.api", 5001)
	// viper.SetDefault("ipfs.gateway", 8080)
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
