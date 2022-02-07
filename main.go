package main

import (
	"context"
	"food-delivery/component/appctx"
	"food-delivery/component/uploadprovider"
	"food-delivery/middleware"
	ginrestaurant "food-delivery/module/restaurant/transport/gin"
	ginupload "food-delivery/module/upload/transport/gin"
	userstorage "food-delivery/module/user/storage"
	"food-delivery/module/user/transport/ginuser"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//testData := restaurantmodel.Restaurant{Name: "Test", Address: "Somewhere"}
	//byteData, err := json.Marshal(testData)
	//
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//log.Println(string(byteData))
	//os.Exit(0)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("CONN_STRING")
	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	secretKey := os.Getenv("SYSTEM_SECRET")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Error mysql:", err)
	}

	db = db.Debug()
	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	url, _ := s3Provider.GetPresignedURL(context.Background(), "img/test.jpg")
	log.Println(url)

	appCtx := appctx.NewAppContext(db, s3Provider, secretKey)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	r.Static("/static", "./static")

	userStore := userstorage.NewSQLStore(appCtx.GetMaiDBConnection())
	midAuthorize := middleware.RequiredAuth(appCtx, userStore)

	v1 := r.Group("/v1")
	{
		v1.POST("/upload", ginupload.Upload(appCtx))

		v1.POST("/register", ginuser.Register(appCtx))
		v1.POST("/authenticate", ginuser.Login(appCtx))
		v1.GET("/profile", midAuthorize, ginuser.Profile(appCtx))

		restaurants := v1.Group("/restaurants", midAuthorize)
		{
			// CRUD
			restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
			restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
			restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
			restaurants.PUT("/:id", ginrestaurant.UpdateRestaurant(appCtx))
			restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()

}
