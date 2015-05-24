package core

import (
	"time"

	"github.com/spf13/viper"
)

// ConfigStore is the storage for the configuration singleton
type ConfigStore struct {
	Binding        string
	Port           string
	RedisIP        string
	RedisPort      string
	RethinkDbIP    string
	RethinkDbPort  string
	DockerHost     string
	DockerCertPath string
	TaskTimeout    time.Duration
}

var (
	// Config is the configuration singleton
	Config ConfigStore
)

func init() {
	viper.BindEnv("binding", "BINDING")
	viper.SetDefault("binding", "127.0.0.1")

	viper.BindEnv("port", "PORT")
	viper.SetDefault("port", "8080")

	viper.BindEnv("redisIp", "REDIS_PORT_6379_TCP_ADDR")

	viper.BindEnv("redisPort", "REDIS_PORT_6379_TCP_PORT")
	viper.SetDefault("redisPort", "6379")

	viper.BindEnv("rethinkdbIp", "RETHINKDB_PORT_28015_TCP_ADDR")

	viper.BindEnv("rethinkdbPort", "RETHINKDB_PORT_28015_TCP_PORT")
	viper.SetDefault("rethinkdbPort", "28015")

	viper.BindEnv("dockerHost", "DOCKER_HOST")
	viper.SetDefault("dockerHost", "unix:///var/run/docker.sock")

	viper.BindEnv("dockerCertPath", "DOCKER_CERT_PATH")
	viper.SetDefault("dockerCertPath", "")

	viper.SetDefault("taskTimeout", 30*time.Second)

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	Config = ConfigStore{
		Binding:        viper.GetString("binding"),
		Port:           viper.GetString("port"),
		RedisIP:        viper.GetString("redisIp"),
		RedisPort:      viper.GetString("redisPort"),
		RethinkDbIP:    viper.GetString("rethinkdbIp"),
		RethinkDbPort:  viper.GetString("rethinkdbPort"),
		DockerHost:     viper.GetString("dockerHost"),
		DockerCertPath: viper.GetString("dockerCertPath"),
		TaskTimeout:    viper.GetDuration("taskTimeout"),
	}
}
