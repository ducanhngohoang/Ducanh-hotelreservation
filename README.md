# BookingProject
Hotel Reservation Project mô phỏng một phần mềm đặt phòng của hệ thống khách sạn Vinpearl.
## Contents
- [Giới thiệu](#Giới-thiệu)
- [Thành viên](#Thành-viên)
- [Công cụ](#Công-cụ)
- [Cấu trúc Database](#Cấu-trúc-Database)
- [Kiến trúc](#Kiến-trúc)
- [Cài đặt](#Cài-đặt)


## Giới thiệu
Server API của project mô phỏng một App Mobile đặt phòng khách sạn cao cấp Vinpearl. Mục tiêu của project nhằm giúp cho khách hàng đặt phòng một cách thuận lợi hơn, giúp cho nhà quản lý có thể chủ động hơn trong việc sắp xếp lượng khách đặt giúp tối ưu hiệu quả việc kinh doanh của mình.

## Thành viên
Thành viên của project bao gồm 2 thành viên :

-Ngô Hoàng Đức Anh : phụ trách BackEnd, bao gồm thiết kế database, thiết kế các REST API để liên kết database với bên FrontEnd.

-Hoàng Văn Đức : phụ trách phần Ios, tạo giao diện đẹp mắt liên kết client với server.

## Công cụ
Cấu trúc và các công cụ được sử dụng của project:
- Virtualization: docker
- Version control: git
- Database: mysql, phpmyadmin
- BackEnd: Golang
  - Gin-gonic : https://github.com/gin-gonic
  - Gorm : https://github.com/jinzhu/gorm
  - IDE : Visual Studio Code
- FrontEnd:Ios swift
- Api documentation: GoSwagger: https://github.com/go-swagger/go-swagger
- Proxy: Localxpose: https://localxpose.io/

## Cấu trúc Database
Cấu trúc các bảng trong database của project

https://imgur.com/a/jRLNF5V


## Cài đặt
Repositories này chứa phần BackEnd của project bao gồm các Rest API, swagger documentation:

main.go:

API declaration

Database initialization

Set listening port (default 8086)

config.local.json: Thông tin của server mysql

controller/

API definition, mỗi file tương ứng với một chức năng chính của webiste

internalfunction.go : Function sử dụng nội bộ trong package controller

model/

Chứa các struct model cần thiết cho gorm và gửi JSON về FrontEnd
1. Kéo thư mục về máy
```sh
$ git clone https://github.com/ducanhngohoang/Ducanh-hotelreservation
```

2. Cài đặt các thư viện cần thiết
```sh
$ go get -u github.com/gin-gonic/gin
$ go get -u github.com/gin-gonic/contrib/jwt
$ go get -u github.com/swaggo/files
$ go get -u github.com/swaggo/gin-swagger
$ go get -u github.com/jinzhu/gorm
$ go get -u github.com/jinzhu/gorm/dialects/
$ go get -u github.com/json-iterator/go
```
3. Vào thư mục chứa Booking Project
```sh
$ cd Booking
```

  
4. Khởi động server
```sh
$ go run main.go
```
