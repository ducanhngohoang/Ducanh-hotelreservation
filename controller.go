package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"booking/model"
)

func GetCities(c *gin.Context) {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	var cityInfo []model.Cities
	errGetCities := db.Find(&cityInfo).Error
	if errGetCities != nil {
		log.Println(errGetCities)
		return
	}

	// for _, v := range userInfo {
	// 	log.Println("Employee", v)
	// }

	c.JSON(200, cityInfo)

}

func GetRooms(c *gin.Context) {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	var roomInfo []model.Rooms
	errGetRooms := db.Find(&roomInfo).Error
	if errGetRooms != nil {
		log.Println(errGetRooms)
		return
	}

	// for _, v := range userInfo {
	// 	log.Println("Employee", v)
	// }

	c.JSON(200, roomInfo)

}

func GetRoomsByHotel(c *gin.Context) {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)
	var RoomsInfo []model.Rooms
	Hotel_id := c.Query("id")
	// Address := c.Query("address")
	// Lat := c.Query("lat")
	// Images := c.Query("images")
	// Restaurant := c.Query("restaurant")
	// Entertainment := c.Query("entertainment")
	// Phone := c.Query("phone")
	// Email := c.Query("email")
	// Hotel_detail := c.Query("hotel_detail")
	errGetRoomsInfo := db.Table("rooms").Select("room_id,hotel_id,room_type_id,room_status").
		Where("hotel_id = ?", Hotel_id).Scan(&RoomsInfo).Error
	// Khi code API thì chỗ này trả về HTTP status code 500
	if errGetRoomsInfo != nil {
		log.Println(errGetRoomsInfo)
		return
	}
	c.JSON(200, RoomsInfo)
}

func GetHotelByCity(c *gin.Context) {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)
	var HotelInfo []model.Hotels
	// Hotel_id := c.Query("id")
	// Hotel_name := c.Query("name")
	City_id := c.Query("id")
	// Address := c.Query("address")
	// Lat := c.Query("lat")
	// Images := c.Query("images")
	// Restaurant := c.Query("restaurant")
	// Entertainment := c.Query("entertainment")
	// Phone := c.Query("phone")
	// Email := c.Query("email")
	// Hotel_detail := c.Query("hotel_detail")
	errGetHotelInfo := db.Table("hotels").Select("hotel_id,hotel_name,city_id,address,lat,longi,images,restaurant,events,entertainment,phone,email,hotel_detail").
		Where("city_id = ?", City_id).Scan(&HotelInfo).Error
	// Khi code API thì chỗ này trả về HTTP status code 500
	if errGetHotelInfo != nil {
		log.Println(errGetHotelInfo)
		return
	}

	log.Println("Hotels", HotelInfo)
	c.JSON(200, HotelInfo)
}

func CreateCities(c *gin.Context) {

	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)
	// var headerInfo model.AuthorizationHeader

	// if err := c.ShouldBindHeader(&headerInfo); err != nil {
	// 	c.JSON(200, err)
	// }

	// tokenFromHeader := strings.Replace(headerInfo.Token, "Bearer ", "", -1)

	// claims := jwt_lib.MapClaims{}
	// tkn, err := jwt_lib.ParseWithClaims(tokenFromHeader, claims, func(token *jwt_lib.Token) (interface{}, error) {
	// 	return []byte(model.SecretKey), nil
	// })

	// if err != nil {
	// 	if err == jwt_lib.ErrSignatureInvalid {
	// 		log.Println("error 1")
	// 		c.JSON(http.StatusUnauthorized, model.ErrorMesssage{
	// 			Message: "Token không hợp lệ",
	// 		})
	// 		return
	// 	}
	// 	log.Println("error 2", err)
	// 	c.JSON(http.StatusBadRequest, model.ErrorMesssage{
	// 		Message: "Bad request",
	// 	})
	// 	return
	// }

	// if !tkn.Valid {
	// 	log.Println("error 3")
	// 	c.JSON(http.StatusUnauthorized, model.ErrorMesssage{
	// 		Message: "Token không hợp lệ",
	// 	})
	// 	return
	// }

	// var userId string
	// var roleFromToken int

	// for k, v := range claims {
	// 	if k == "userId" {
	// 		userId = v.(string)
	// 	}

	// 	if k == "Role" {
	// 		roleFromToken = int(v.(float64))
	// 	}
	// }

	// log.Println("--------", userId, roleFromToken)
	// if userId == "" {
	// 	c.JSON(http.StatusUnauthorized, model.ErrorMesssage{
	// 		Message: "Token không hợp lệ",
	// 	})
	// }

	var createCity model.CreateCityReq
	err := c.BindJSON(&createCity)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorMesssage{
			Message: "Yêu cầu tạo city không hợp lệ",
		})
		return
	}

	var CityInsert = model.Cities{
		City_id:   createCity.Id,
		City_name: createCity.Name,
	}

	errInsert := db.Table("cities").Create(&CityInsert).Error
	if errInsert != nil {
		log.Println(errInsert)
		c.JSON(http.StatusBadRequest, model.ErrorMesssage{
			Message: "Yêu cầu tạo thành phố không hợp lệ",
		})
		return
	}

	var rsp = model.CreateCityRsp{
		ResponseTime: time.Now().String(),
		Code:         0,
		Message:      "Tạo thành phố thành công",
		Data:         CityInsert,
	}

	c.JSON(http.StatusOK, rsp)
}
