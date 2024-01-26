package test

import (
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	_ "gorm.io/gorm/logger"
	"testing"
)

func TestDir01(t *testing.T) {
	dsn := "tcp://localhost:9000?database=default&username=root&password=123456&parseTime=True&loc=Local"
	//dsn := "root:123456@tcp(127.0.0.1:9000)/trade?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	println(dsn)
}
