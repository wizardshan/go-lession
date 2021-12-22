package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go-web/lesson/chapter5_1/controller"
	"go-web/lesson/chapter5_1/repository"
	"go-web/lesson/chapter5_1/repository/ent"
	"log"
)

func main() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"",
		"127.0.0.1",
		"test",
	)

	db, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	//
	////if err := client.Schema.Create(context.Background()); err != nil {
	////	log.Fatalf("failed creating schema resources: %v", err)
	////}
	//
	//orders, err := client.Debug().Order.Query().WithGoods().All(context.Background())
	//fmt.Println(orders)
	//fmt.Println(goods[0].Edges.Goods)


	engine := gin.New()

	repoGoods := repository.NewGoods(db)
	ctrGoods := controller.NewGoods(repoGoods)
	engine.GET("/v1/goods", ctrGoods.V1List)
	engine.GET("/v1/goods/category", ctrGoods.V1Category)
	engine.GET("/v2/goods", ctrGoods.V2List)
	engine.GET("/v3/goods", ctrGoods.V3List)
	//engine.GET("/users/:id", ctrUser.Info)
	//engine.POST("/user/update", ctrUser.Update)
	engine.Run()
}
