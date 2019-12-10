package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"booking/model"
)

//Get All Cities
// @Description Get tất cả dữ liệu từ bảng thành phố trả về JSON
// @Accept json
// @Success 200 {object} model.Cities
// @Router /GetAllCities [get]

//GetCities là lấy dữ liệu tất cả thành phố về in ra JSON
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

	c.JSON(200, cityInfo)

}

// GetRooms là Lấy dữ liệu tất cả phòng về
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

	c.JSON(200, roomInfo)

}

//GetRoomsByHotel là lấy dữ liệu tất cả các phòng của một khách sạn
func GetRoomsByHotel(c *gin.Context) {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)
	var RoomsInfo []model.Rooms
	Hotelid := c.Query("id")
	errGetRoomsInfo := db.Table("rooms").Select("roomid,hotelid,roomtypeid,roomstatus,price").
		Where("hotelid = ?", Hotelid).Scan(&RoomsInfo).Error
	if errGetRoomsInfo != nil {
		log.Println(errGetRoomsInfo)
		return
	}
	c.JSON(200, RoomsInfo)
}

//GetHotelByCity là in ra các khách sạn của một thành phố
func GetHotelByCity(c *gin.Context) {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)
	var HotelInfo []model.Hotels
	Cityid := c.Query("id")
	errGetHotelInfo := db.Table("hotels").Select("hotelid,hotelname,cityid,address,lat,longi,images,restaurant,events,entertainment,phone,email,hoteldetail,vipprice,vvipprice,deluxevipprice").
		Where("cityid = ?", Cityid).Scan(&HotelInfo).Error

	if errGetHotelInfo != nil {
		log.Println(errGetHotelInfo)
		return
	}

	log.Println("Hotels", HotelInfo)
	c.JSON(200, HotelInfo)
}

// GetAllHotels là lấy dữ liệu tất cả các khách sạn
func GetAllHotels(c *gin.Context) {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	var AllHotelsInfo []model.Hotels
	errGetAllHotels := db.Find(&AllHotelsInfo).Error
	if errGetAllHotels != nil {
		log.Println(errGetAllHotels)
		return
	}

	// for _, v := range userInfo {
	// 	log.Println("Employee", v)
	// }

	c.JSON(200, AllHotelsInfo)

}

//GetAllBooking là hàm get về tất cả các thông tin Booking hiện có
func GetAllBooking(c *gin.Context) {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	var AllBookingInfo []model.BookingInfo
	errGetAllBookings := db.Find(&AllBookingInfo).Error
	if errGetAllBookings != nil {
		log.Println(errGetAllBookings)
		return
	}

	c.JSON(200, AllBookingInfo)

}

//GetBookingByPhone là in ra các khách sạn của một thành phố
func GetBookingByPhone(c *gin.Context) {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	var BookingInfo []model.BookingInfo
	Bookingphone := c.Query("bookingphone")

	errGetBookingInfo := db.Table("booking_infos").Select("bookingid,hotelid,startdate,enddate,roomquantity,roomtype,bookingname,bookingphone").
		Where("bookingphone = ?", Bookingphone).Scan(&BookingInfo).Error

	if errGetBookingInfo != nil {
		log.Println(errGetBookingInfo)
		return
	}

	log.Println("Booking", BookingInfo)
	c.JSON(200, BookingInfo)
}

//GetRoomBookedByBookingID trả về các phòng đã đặt theo ID.
func GetRoomBookedByBookingID(c *gin.Context) {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	var ReservedRoomsInfo []model.Reservedrooms
	Bookingid := c.Query("bookingid")

	errGetReservedRoomsInfo := db.Table("reservedrooms").Select("roombookedid,bookingid,roomid,startdate,enddate,roomtypeid,hotelid").
		Where("bookingid = ?", Bookingid).Scan(&ReservedRoomsInfo).Error

	if errGetReservedRoomsInfo != nil {
		log.Println(errGetReservedRoomsInfo)
		return
	}

	log.Println("Reserved", ReservedRoomsInfo)
	c.JSON(200, ReservedRoomsInfo)
}

//CreateCities là tạo thêm thành phố mới trong cơ sở dữ liệu
func CreateCities(c *gin.Context) {

	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	var createCity model.CreateCityReq
	err := c.BindJSON(&createCity)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorMessage{
			Message: "Yêu cầu tạo city không hợp lệ",
		})
		return
	}

	var CityInsert = model.Cities{
		Cityid:   createCity.ID,
		Cityname: createCity.Name,
	}

	errInsert := db.Table("cities").Create(&CityInsert).Error
	if errInsert != nil {
		log.Println(errInsert)
		c.JSON(http.StatusBadRequest, model.ErrorMessage{
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

//CreateBookingInfo tạo ra thông tin booking lên database
func CreateBookingInfo(c *gin.Context) {

	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	var createBookingInfo model.CreateBookingInfoReq
	err := c.BindJSON(&createBookingInfo)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorMessage{
			Message: "Yêu cầu tạo city không hợp lệ",
		})
		return
	}

	var BookingInfoInsert = model.BookingInfo{
		Bookingid:    createBookingInfo.BookingID,
		Hotelid:      createBookingInfo.HotelID,
		Startdate:    createBookingInfo.Startdate,
		Enddate:      createBookingInfo.Enddate,
		Roomquantity: createBookingInfo.Roomquantity,
		Roomtype:     createBookingInfo.Roomtype,
		Bookingname:  createBookingInfo.Bookingname,
		Bookingphone: createBookingInfo.Bookingphone,
	}

	errInsert := db.Table("booking_infos").Create(&BookingInfoInsert).Error
	if errInsert != nil {
		log.Println(errInsert)
		c.JSON(http.StatusBadRequest, model.ErrorMessage{
			Message: "Yêu cầu tạo booking không hợp lệ",
		})
		return
	}

	var rsp = model.CreateBookingInfoRsp{
		ResponseTime: time.Now().String(),
		Code:         0,
		Message:      "Tạo booking thành công",
		Data:         BookingInfoInsert,
	}

	c.JSON(http.StatusOK, rsp)
}

//CreateHotels là tạo thêm khách sạn mới trong cơ sở dữ liệu
func CreateHotels(c *gin.Context) {

	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	var createHotel model.CreateHotelReq
	err := c.BindJSON(&createHotel)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorMessage{
			Message: "Yêu cầu tạo hotel không hợp lệ",
		})
		return
	}

	var HotelInsert = model.Hotels{
		Hotelid:        createHotel.ID,
		Hotelname:      createHotel.Name,
		Cityid:         createHotel.CityID,
		Address:        createHotel.Address,
		Lat:            createHotel.Lat,
		Longi:          createHotel.Longi,
		Images:         createHotel.Images,
		Restaurant:     createHotel.Restaurant,
		Events:         createHotel.Events,
		Entertainment:  createHotel.Entertainment,
		Phone:          createHotel.Phone,
		Email:          createHotel.Email,
		Hoteldetail:    createHotel.HotelDetail,
		Vipprice:       createHotel.Vipprice,
		Vvipprice:      createHotel.Vvipprice,
		Deluxevipprice: createHotel.Deluxevipprice,
	}

	errInsert := db.Table("hotels").Create(&HotelInsert).Error
	if errInsert != nil {
		log.Println(errInsert)
		c.JSON(http.StatusBadRequest, model.ErrorMessage{
			Message: "Yêu cầu tạo khách sạn không hợp lệ",
		})
		return
	}

	var rsp = model.CreateHotelRsp{
		ResponseTime: time.Now().String(),
		Code:         0,
		Message:      "Tạo khách sạn thành công",
		Data:         HotelInsert,
	}

	c.JSON(http.StatusOK, rsp)
}

//CreateRooms là tạo thêm phòng mới trong cơ sở dữ liệu
func CreateRooms(c *gin.Context) {

	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	var createRoom model.CreateRoomReq
	err := c.BindJSON(&createRoom)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorMessage{
			Message: "Yêu cầu tạo Room không hợp lệ",
		})
		return
	}

	var RoomInsert = model.Rooms{
		Roomid:     createRoom.RoomID,
		Hotelid:    createRoom.HotelID,
		Roomtypeid: createRoom.RoomTypeID,
		Roomstatus: createRoom.RoomStatus,
		Price:      createRoom.Price,
	}

	errInsert := db.Table("rooms").Create(&RoomInsert).Error
	if errInsert != nil {
		log.Println(errInsert)
		c.JSON(http.StatusBadRequest, model.ErrorMessage{
			Message: "Yêu cầu tạo phòng không hợp lệ",
		})
		return
	}

	var rsp = model.CreateRoomRsp{
		ResponseTime: time.Now().String(),
		Code:         0,
		Message:      "Tạo phòng thành công",
		Data:         RoomInsert,
	}

	c.JSON(http.StatusOK, rsp)
}

//CreateReservedRoom thêm danh sách phòng đã có khách đặt
func CreateReservedRoom(c *gin.Context) {

	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	var createReservedRoomsInfo model.CreateReservedRoomReq
	err := c.BindJSON(&createReservedRoomsInfo)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorMessage{
			Message: "Yêu cầu tạo phòng đặt không hợp lệ",
		})
		return
	}

	var ReservedRoomsInsert = model.Reservedrooms{
		Roombookedid: createReservedRoomsInfo.Roombookedid,
		Bookingid:    createReservedRoomsInfo.Bookingid,
		Roomid:       createReservedRoomsInfo.Roomid,
		Startdate:    createReservedRoomsInfo.Startdate,
		Enddate:      createReservedRoomsInfo.Enddate,
		Roomtypeid:   createReservedRoomsInfo.Roomtypeid,
		Hotelid:      createReservedRoomsInfo.Hotelid,
	}

	errInsert := db.Table("reservedrooms").Create(&ReservedRoomsInsert).Error
	if errInsert != nil {
		log.Println(errInsert)
		c.JSON(http.StatusBadRequest, model.ErrorMessage{
			Message: "Yêu cầu tạo phòng đặt không hợp lệ",
		})
		return
	}

	var rsp = model.CreateReservedRoomsRsp{
		ResponseTime: time.Now().String(),
		Code:         0,
		Message:      "Tạo các phòng đã đặt thành công",
		Data:         ReservedRoomsInsert,
	}

	c.JSON(http.StatusOK, rsp)
}

//DeleteCity là thêm tinh nang delete
func DeleteCity(c *gin.Context) {
	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	// Trước khi delete
	var InfoCity model.Cities

	cityid := c.Query("id")

	errDeletePost := db.Table("cities").Delete("cityid,cityname").
		// Khi code API thì chỗ này trả về HTTP status code 500
		Where("cityid = ?", cityid).Scan(&InfoCity).Error
	if errDeletePost != nil {
		log.Println(errDeletePost)
		return
	}

	// Kiểm tra xem đã delete chưa
	var InfoCityAfter model.Cities
	// Khi code API thì ID của user sẽ lấy từ đường dẫn (GET request) hoặc form-data/JSON (POST)
	errGetPost := db.Where("id = ?", c.Query("id")).Find(&InfoCityAfter).Error
	// Khi code API thì chỗ này trả về HTTP status code 500
	if errGetPost != nil {
		log.Println(errGetPost)
		return
	}
	log.Println("After update", InfoCityAfter)

	defer db.Close()
}



//ProcessingCustomerRequest là tính năng trả về số phòng khách có thể đặt được theo yêu cầu tên khách sạn, loại phòng, và ngày giờ đặt.
func ProcessingCustomerRequest(c *gin.Context) {

	config := model.SetupConfig()
	db := model.ConnectDb(config.Database.User, config.Database.Password, config.Database.Database, config.Database.Address)
	defer db.Close()
	db.LogMode(true)

	var processrequest model.RequestBooking
	err := c.BindJSON(&processrequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorMessage{
			Message: "Yêu cầu tạo request không hợp lệ",
		})
		log.Println(err)
		return
	}

	var UnavailableRooms1 []model.UnavailableRooms
	db.Raw(`
		SELECT roomid
		FROM reservedrooms
		WHERE hotelid = ? and roomtypeid = ? and Enddate > ? and Enddate <= ?`, processrequest.HotelID, processrequest.RoomTypeID, processrequest.Startdate, processrequest.Enddate).Scan(&UnavailableRooms1)

	var UnavailableRooms2 []model.UnavailableRooms
	db.Raw(`
		SELECT roomid
		FROM reservedrooms
		WHERE hotelid = ? and roomtypeid = ? and Startdate >= ? and Startdate < ?`, processrequest.HotelID, processrequest.RoomTypeID, processrequest.Startdate, processrequest.Enddate).Scan(&UnavailableRooms2)

	var AllRoomsCustomerWant []model.AllRoomsCustomerWant
	db.Raw(`
		SELECT roomid
		FROM rooms
		WHERE hotelid = ? and roomtypeid = ?`, processrequest.HotelID, processrequest.RoomTypeID).Scan(&AllRoomsCustomerWant)
	var a = len(AllRoomsCustomerWant) - len(UnavailableRooms1)
	var b = len(AllRoomsCustomerWant) - len(UnavailableRooms2)
	var NumberOfRoomsAvailable int
	if a >= b {
		NumberOfRoomsAvailable = b
	} else {
		NumberOfRoomsAvailable = a
	}

	//Result trả về kết quả đặt thành công hay không
	var Result int
	if NumberOfRoomsAvailable >= processrequest.Roomquantity {
		Result = 1
	} else {
		Result = 0
	}
	// c.JSON(200, UnavailableRooms1)
	// c.JSON(200, UnavailableRooms2)
	c.JSON(200, Result)
}
