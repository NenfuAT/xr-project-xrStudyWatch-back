package conf

import "github.com/spf13/viper"

var f, m *viper.Viper

func init() {
	f = viper.New()
	f.SetConfigType("yaml")
	f.SetConfigName("postgres")
	f.AddConfigPath("conf/environments/")

	m = viper.New()
	m.SetConfigType("yaml")
	m.SetConfigName("proxy")
	m.AddConfigPath("conf/environments/")
}

func GetPostgresConfig() *viper.Viper {
	if err := f.ReadInConfig(); err != nil {
		return nil
	}
	return f
}

func GetProxyConfig() *viper.Viper {
	if err := m.ReadInConfig(); err != nil {
		return nil
	}
	return m
}
