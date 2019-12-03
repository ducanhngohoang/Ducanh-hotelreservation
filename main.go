package main

import (
	// "booking/controller"
	"booking/controller"
	"booking/model"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()

	r.GET("/getAllCities", controller.GetCities)
	r.GET("/getHotelByCity", controller.GetHotelByCity)
	r.GET("/getAllRooms", controller.GetRooms)
	r.GET("/getRoomsByHotel", controller.GetRoomsByHotel)
	r.POST("/create-city",  controller.CreateCities)

	r.Run(":8086")
}
