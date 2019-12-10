package model

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //lấy kiểu dữ liệu từ sql
	jsoniter "github.com/json-iterator/go"
)

//Cities là kiểu dữ liệu thành phố trong Db
type Cities struct {
	Cityid   int
	Cityname string
}

//Rooms là kiểu dữ liệu phòng trong DB
type Rooms struct {
	Roomid     int
	Hotelid    int
	Roomtypeid int
	Roomstatus int
	Price      string
}

//Hotels là kiểu dữ liệu khách sạn trong DB
type Hotels struct {
	Hotelid        int
	Hotelname      string
	Cityid         int
	Address        string
	Lat            float32
	Longi          float32
	Images         string
	Restaurant     string
	Events         string
	Entertainment  string
	Phone          string
	Email          string
	Hoteldetail    string
	Vipprice       string
	Vvipprice      string
	Deluxevipprice string
}

// CreateCityRsp là một struct thêm thành phố thành công
type CreateCityRsp struct {
	ResponseTime string `json:"responseTime"`
	Code         int    `json:"code"`
	Message      string `json:"message"`
	Data         Cities `json:"data"`
}

// CreateHotelRsp là một struct trả về thêm khách sạn thành công
type CreateHotelRsp struct {
	ResponseTime string `json:"responseTime"`
	Code         int    `json:"code"`
	Message      string `json:"message"`
	Data         Hotels `json:"data"`
}

// CreateReservedRoomsRsp trả về biến tạo phòng đã đặt

type CreateReservedRoomsRsp struct {
	ResponseTime string        `json:"responseTime"`
	Code         int           `json:"code"`
	Message      string        `json:"message"`
	Data         Reservedrooms `json:"data"`
}

// CreateRoomRsp là một struct trả về thêm phòng thành công
type CreateRoomRsp struct {
	ResponseTime string `json:"responseTime"`
	Code         int    `json:"code"`
	Message      string `json:"message"`
	Data         Rooms  `json:"data"`
}

//CreateBookingInfoRsp là một struct trả về thêm bookinginfo thành công
type CreateBookingInfoRsp struct {
	ResponseTime string      `json:"responseTime"`
	Code         int         `json:"code"`
	Message      string      `json:"message"`
	Data         BookingInfo `json:"data"`
}

// CreateCityReq là một struct thêm thành phố
type CreateCityReq struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// CreateReservedRoomReq là một struct thêm thông tin phòng đặt
type CreateReservedRoomReq struct {
	Roombookedid int       `json:"roombookedid"`
	Bookingid    int       `json:"bookingid"`
	Roomid       int       `json:"roomid"`
	Startdate    time.Time `json:"startdate"`
	Enddate      time.Time `json:"enddate"`
	Roomtypeid   int       `json:"roomtypeid"`
	Hotelid      int       `json:"hotelid"`
}

//CreateHotelReq là một struct thêm khách sạn từ yêu cầu json gửi lên
type CreateHotelReq struct {
	CityID         int     `json:"cityid"`
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Address        string  `json:"address"`
	Lat            float32 `json:"lat"`
	Longi          float32 `json:"long"`
	Images         string  `json:"images"`
	Restaurant     string  `json:"restaurant"`
	Events         string  `json:"events"`
	Entertainment  string  `json:"entertainment"`
	Phone          string  `json:"phone"`
	Email          string  `json:"email"`
	HotelDetail    string  `json:"detail"`
	Vipprice       string  `json:"vipprice"`
	Vvipprice      string  `json:"vvipprice"`
	Deluxevipprice string  `json:"deluxevipprice"`
}

//CreateRoomReq là request tạo phòng mới từ JSON
type CreateRoomReq struct {
	RoomID     int    `json:"roomID"`
	HotelID    int    `json:"hotelID"`
	RoomTypeID int    `json:"roomtypeID"`
	RoomStatus int    `json:"roomstatus"`
	Price      string `json:"price"`
}

//CreateBookingInfoReq là request tạo Booking mới thành công
type CreateBookingInfoReq struct {
	BookingID    int       `json:"bookingID"`
	HotelID      int       `json:"hotelID"`
	Startdate    time.Time `json:"startdate"`
	Enddate      time.Time `json:"enddate"`
	Roomquantity int       `json:"roomquantity"`
	Roomtype     int       `json:"roomtype"`
	Bookingname  string    `json:"bookingname"`
	Bookingphone string    `json:"bookingphone"`
}

//Reservedrooms khai báo kiểu dữ liệu thẻ phòng đã đặt trong DB
type Reservedrooms struct {
	Roombookedid int
	Bookingid    int
	Roomid       int
	Startdate    time.Time
	Enddate      time.Time
	Roomtypeid   int
	Hotelid      int
}

//BookingInfo khai báo kiểu bookinginfo
type BookingInfo struct {
	Bookingid    int
	Hotelid      int
	Startdate    time.Time
	Enddate      time.Time
	Roomquantity int
	Roomtype     int
	Bookingname  string
	Bookingphone string
}

//RequestBooking là struct gửi yêu cầu của khách hàng lên server
type RequestBooking struct {
	HotelID      int       `json:"hotelid"`
	RoomTypeID   int       `json:"roomtypeid"`
	Startdate    time.Time `json:"startdate"`
	Enddate      time.Time `json:"enddate"`
	Roomquantity int       `json:"roomquantity"`
}

//UnavailableRooms là struct các phòng không khả dụng khi đã có khách đặt từ trước
type UnavailableRooms struct {
	RoomID int
}

// type Message struct {
// 	Message int
// }

//AllRoomsCustomerWant là struct in ra tất cả các phòng phù hợp nhu cầu khách hàng về tên kahchs sạn và loại phòng
type AllRoomsCustomerWant struct {
	RoomID int
}

//Config là struct truy nhập vào DB
type Config struct {
	Database struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
		Address  string `json:"address"`
	} `json:"database"`
}

//ErrorMessage là biến mã lỗi trả về
type ErrorMessage struct {
	Message string `json:"message"`
}

//DecodeDataFromJsonFile đọc dữ liệu từ File JSON
func DecodeDataFromJsonFile(f *os.File, data interface{}) error {
	jsonParser := jsoniter.NewDecoder(f)
	err := jsonParser.Decode(&data)
	if err != nil {
		return err
	}

	return nil
}

//SetupConfig đọc dữ liệu từ config
func SetupConfig() Config {
	var conf Config

	// Đọc file config.dev.json
	configFile, err := os.Open("config.local.json")
	if err != nil {
		// Nếu không có file config.dev.json thì đọc file config.default.json
		configFile, err = os.Open("config.default.json")
		if err != nil {
			panic(err)
		}
		defer configFile.Close()
	}
	defer configFile.Close()

	// Parse dữ liệu JSON và bind vào conf
	err = DecodeDataFromJsonFile(configFile, &conf)
	if err != nil {
		log.Println("Không đọc được file config.")
		panic(err)
	}

	return conf
}

//ConnectDb kết nối với Database
func ConnectDb(user string, password string, database string, address string) *gorm.DB {
	connectionInfo := fmt.Sprintf(`%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local`, user, password, address, database)

	db, err := gorm.Open("mysql", connectionInfo)
	if err != nil {
		panic(err)
	}
	return db
}
