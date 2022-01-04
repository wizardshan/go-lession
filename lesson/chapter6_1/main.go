package main

import (
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter6_1/controller"
	"go-web/lesson/chapter6_1/repository"
	"go-web/lesson/chapter6_1/repository/ent"
	"log"
	"time"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"",
		"127.0.0.1",
		"test")

	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	// 获取数据库驱动中的sql.DB对象。
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	client :=  ent.NewClient(ent.Driver(drv))

	engine := gin.New()

	repoGoods := repository.NewGoods(client)
	ctrGoods := controller.NewGoods(repoGoods)
	engine.GET("/goods/category", ctrGoods.Category)
	engine.GET("/v1/goods", ctrGoods.V1List)
	engine.GET("/v2/goods", ctrGoods.V2List)

	engine.GET("/goods/:id", ctrGoods.Get)
	engine.GET("/goods", ctrGoods.List)
	engine.POST("/goods", ctrGoods.Create)

	engine.Run()
}
