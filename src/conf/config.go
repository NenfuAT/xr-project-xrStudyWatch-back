package conf

import "github.com/spf13/viper"

var f, m *viper.Viper

func init() {
	f = viper.New()
	f.SetConfigType("yaml")
	f.SetConfigName("postgres")
	f.AddConfigPath("conf/environments/")
}

func GetPostgresConfig() *viper.Viper {
	if err := f.ReadInConfig(); err != nil {
		return nil
	}
	return f
}
