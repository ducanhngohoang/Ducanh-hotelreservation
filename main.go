package main

import (
	// "booking/controller"
	"booking/controller"
	"booking/model"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "booking/docs"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()

	url := ginSwagger.URL("http://localhost:8086/booking/doc.json")
	r.GET("/booking/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.GET("/getAllCities", controller.GetCities)
	r.GET("/getHotelByCity", controller.GetHotelByCity)
	// r.GET("/getAllRooms", controller.GetRooms)
	r.GET("/getRoomsByHotel", controller.GetRoomsByHotel)
	r.GET("/getallhotels", controller.GetAllHotels)
	r.GET("/getBookingByPhone", controller.GetBookingByPhone)
	r.GET("/getAllBooking", controller.GetAllBooking)
	r.POST("/create-city", controller.CreateCities)
	r.POST("/create-hotel", controller.CreateHotels)
	r.POST("/create-rooms", controller.CreateRooms)
	r.POST("/create-bookinginfo", controller.CreateBookingInfo)
	r.DELETE("/deletecity", controller.DeleteCity)
	r.POST("/availaberooms", controller.ProcessingCustomerRequest)
	r.GET("/getRoomsByBookingID", controller.GetRoomBookedByBookingID)
	r.Run(":8086")

}
