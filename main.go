package main

import (
	"food-delivery-service/component"
	"food-delivery-service/middleware"
	restaurantgin "food-delivery-service/module/restaurant/transport/gin"
	usertransport "food-delivery-service/module/user/usertransport/ginuser"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	//dsn := os.Getenv("DB_CONN_STR")
	dsn := "root:phat@123456@tcp(127.0.0.1:3306)/MANI?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	log.Println("Connected:", db)

	router := gin.Default()
	router.Use(middleware.Recover())

	appCtx := component.NewAppContext(db)
	v1 := router.Group("/v1")

	v1.POST("/register", usertransport.Register(appCtx))

	{
		restaurants := v1.Group("/restaurants")
		{
			restaurants.POST("", restaurantgin.CreateRestaurantHandler(appCtx))
			restaurants.GET("", restaurantgin.ListRestaurant(appCtx))
			restaurants.GET("/:restaurant_id", restaurantgin.GetRestaurantHandler(appCtx))
			restaurants.PUT("/:restaurant_id", restaurantgin.UpdateRestaurantHandler(appCtx))
			restaurants.DELETE("/:restaurant_id", restaurantgin.DeleteRestaurantHandler(appCtx))
		}
	}

	//mainRoute(router, appCtx)

	router.Run(":3000")
}
