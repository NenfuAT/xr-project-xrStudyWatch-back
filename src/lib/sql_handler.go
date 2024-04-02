package lib

import (
	"fmt"
	"time"

	"github.com/NenfuAT/xr-project-xrStudyWatch-back/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SqlConnect() (database *gorm.DB) {
	var db *gorm.DB
	var err error
	c := conf.GetPostgresConfig()
	dsn := "host=" + c.GetString("postgres.host") + " user=" + c.GetString("postgres.user") + " password=" + c.GetString("postgres.password") + " dbname=" + c.GetString("postgres.dbname") + " port=" + c.GetString("postgres.port") + " sslmode=disable TimeZone=Asia/Tokyo"
	dialector := postgres.Open(dsn)

	fmt.Println(dsn)
	if db, err = gorm.Open(dialector); err != nil {
		db = connect(dialector, 10)
	}
	fmt.Println("db connected!!")

	return db
}

func connect(dialector gorm.Dialector, count uint) *gorm.DB {
	var err error
	var db *gorm.DB
	if db, err = gorm.Open(dialector); err != nil {
		if count > 1 {
			time.Sleep(time.Second * 2)
			count--
			fmt.Printf("retry... count:%v\n", count)
			connect(dialector, count)
		}
		panic(err.Error())
	}
	return (db)
}
