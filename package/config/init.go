package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../package/config")

	err := viper.ReadInConfig()
	return err
}

func InitForTests() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(`C:\Users\danby\GolandProjects\robbo_student_personal_account\package\config\`)
	fmt.Println()
	err := viper.ReadInConfig()
	return err
}
