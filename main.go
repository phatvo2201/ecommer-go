package main

import (
	"food-delivery-service/component"
	uploadprovider "food-delivery-service/component/upload"
	"food-delivery-service/middleware"
	restaurantgin "food-delivery-service/module/restaurant/transport/gin"
	uploadgin "food-delivery-service/module/upload/transport/gin"
	usertransport "food-delivery-service/module/user/usertransport/ginuser"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := "root:phat@123456@tcp(127.0.0.1:3306)/MANI?charset=utf8mb4&parseTime=True&loc=Local"

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	secretKey := os.Getenv("SECRET_KEY")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db, s3Provider, secretKey); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string) error {

	appCtx := component.NewAppContext(db, upProvider, secretKey)
	r := gin.Default()

	r.Use(middleware.Recover())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// CRUD

	v1 := r.Group("/v1")

	v1.POST("/upload", uploadgin.Upload(appCtx))

	v1.POST("/register", usertransport.Register(appCtx))
	v1.POST("/login", usertransport.Login(appCtx))
	//v1.GET("/profile", middleware.RequiredAuth(appCtx), ginuser.GetProfile(appCtx))

	restaurants := v1.Group("/restaurants")
	{
		restaurants.POST("", restaurantgin.CreateRestaurantHandler(appCtx))
		restaurants.GET("/:id", restaurantgin.GetRestaurantHandler(appCtx))
		restaurants.GET("", restaurantgin.ListRestaurant(appCtx))
		restaurants.PATCH("/:id", restaurantgin.UpdateRestaurantHandler(appCtx))
		restaurants.DELETE("/:id", restaurantgin.DeleteRestaurantHandler(appCtx))
	}

	return r.Run()
}
