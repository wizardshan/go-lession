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
		"test")

	db, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	//if err := client.Schema.Create(context.Background()); err != nil {
	//	log.Fatalf("failed creating schema resources: %v", err)
	//}

	engine := gin.New()

	repoGoods := repository.NewGoods(db)
	repoGoodsCategory := repository.NewGoodsCategory(db)
	ctrGoods := controller.NewGoods(repoGoods, repoGoodsCategory)
	engine.GET("/goods/category", ctrGoods.Category)
	engine.GET("/v1/goods", ctrGoods.V1List)
	engine.GET("/v2/goods", ctrGoods.V2List)
	engine.GET("/v3/goods", ctrGoods.V3List)

	repoUser := repository.NewUser(db)
	ctrUser := controller.NewUser(repoUser)
	engine.GET("/user/register", ctrUser.Register)
	engine.GET("/users/:id", ctrUser.Get)
	engine.GET("/users", ctrUser.List)
	engine.Run()
}
