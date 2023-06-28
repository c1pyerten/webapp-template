package config

import (
	"c1pherten/yet-webapp2/common/util"
	"embed"
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Dialect string ``
		Host    string
		Port    string
		DBName  string
		Username string
		Password string
		Migration bool
	}

	Redis struct {
		Host string
		Port string
		Url string
	}
	
	Extension struct {
		CorsEnabled bool
		SecurityEnabled bool
	}

	Log struct {
		RequestLogFormat string
	}

	Swagger struct {
		Enabled bool
		Path string
	}


	Security struct {
		AuthPath []string
		ExcludePath []string
		UserPath []string
		AdminPath []string
	}

}

const (
	DEV = "dev"
	PROD = "prod"

)

func LoadAppConfig(file embed.FS) (*Config, string) {
	var env *string
	if value := os.Getenv(ENV_NAME); value != "" {
		env = &value
	} else {
		env = flag.String("env", "dev", "switch configuration")
		flag.Parse()
	}
	b, err := file.ReadFile(fmt.Sprintf(AppConfigPath, *env))
	if err != nil {
		panic(err)
	}

	var config Config
	if err := yaml.Unmarshal(b, &config); err != nil {
		panic(err)
	}
	

	return &config, *env
}

func LoadMessagesConfig(file embed.FS) map[string]string {
	msgs, err := util.ReadPropertiesFile(file, MessagesConfigPath)
	if err != nil {
		panic(err)
	}
	
	return msgs
}

// // LoadAppConfig reads fiels
// func LoadAppConfig(file embed.FS) (*Config, string) {
// 	var env *string
// 	if value := os.Getenv(ENV_NAME); value != "" {
// 		env = &value
// 	} else {
// 		env = flag.String("env", "develope", "switch configuration")
// 		flag.Parse()
// 	}

// 	config := &Config{}
// 	f, err := file.ReadFile(fmt.Sprintf(AppConfigPath, *env))
// 	if err != nil {
// 		fmt.Printf("Failed to read application.%s.yml: %v", *env, err)
// 	}
// 	if err := yaml.Unmarshal(f, config); err != nil {
// 		fmt.Printf("failed to read %s, err: %v", *env, err)
// 		os.Exit(ErrExitStatus)
// 	}
	
// 	return config, *env
// }

// func LoadMessagesConfig(file embed.FS) map[string]string {
// 	messages := util.ReadFileProperties(file, MessagesConfigPath)
// 	if messages == nil {
// 		fmt.Println("Failed to load messages.properties")
// 		os.Exit(ErrExitStatus)
// 	}
// 	return messages
// }
