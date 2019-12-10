-- phpMyAdmin SQL Dump
-- version 4.9.1
-- https://www.phpmyadmin.net/
--
-- Host: db
-- Generation Time: Dec 10, 2019 at 10:27 AM
-- Server version: 8.0.18
-- PHP Version: 7.2.23

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `vinpearl`
--

-- --------------------------------------------------------

--
-- Table structure for table `booking_infos`
--

CREATE TABLE `booking_infos` (
  `bookingid` int(11) NOT NULL,
  `hotelid` int(11) DEFAULT NULL,
  `startdate` date DEFAULT NULL,
  `enddate` date DEFAULT NULL,
  `roomquantity` int(11) DEFAULT NULL,
  `roomtype` int(11) DEFAULT NULL,
  `bookingname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `bookingphone` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `booking_infos`
--

INSERT INTO `booking_infos` (`bookingid`, `hotelid`, `startdate`, `enddate`, `roomquantity`, `roomtype`, `bookingname`, `bookingphone`) VALUES
(0, 7, '2019-12-01', '2019-12-10', 3, 1, 'Ducanh2', '0955059599'),
(1, 1, '2019-12-05', '2019-12-07', 3, 0, 'Ngô Hoàng Đức Anh', '0966095896'),
(1010, 7, '2019-12-01', '2019-12-10', 3, 1, 'Ducanh', '0955059599'),
(10101, 7, '2019-12-01', '2019-12-10', 3, 1, 'Ducanh', '0955059599'),
(107378, 1, '2019-12-12', '2019-12-10', 1, 0, 'Hvduc', '0962658415'),
(267894, 1, '2019-12-11', '2019-12-11', 1, 0, 'T', '0965625555'),
(431328, 1, '2019-12-11', '2019-12-15', 1, 1, 'Hvd', '09626666666'),
(862554, 1, '2019-12-11', '2019-12-11', 1, 0, 'T', '0962658444');

-- --------------------------------------------------------

--
-- Table structure for table `cities`
--

CREATE TABLE `cities` (
  `cityid` int(11) NOT NULL,
  `cityname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `cities`
--

INSERT INTO `cities` (`cityid`, `cityname`) VALUES
(1, 'Hồ Chí Minh'),
(2, 'Huế'),
(3, 'Nha Trang'),
(4, 'Quảng Bình'),
(5, 'Hà Tĩnh'),
(6, 'Thanh Hóa'),
(7, 'Phú Quốc'),
(8, 'Hải Phòng'),
(9, 'Vinh');

-- --------------------------------------------------------

--
-- Table structure for table `hotels`
--

CREATE TABLE `hotels` (
  `hotelid` int(11) NOT NULL,
  `hotelname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `cityid` int(11) DEFAULT NULL,
  `address` text,
  `lat` float DEFAULT NULL,
  `longi` float DEFAULT NULL,
  `images` varchar(10000) NOT NULL,
  `restaurant` text,
  `events` text,
  `entertainment` text,
  `phone` text,
  `email` text,
  `hoteldetail` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `vipprice` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `vvipprice` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `deluxevipprice` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `hotels`
--

INSERT INTO `hotels` (`hotelid`, `hotelname`, `cityid`, `address`, `lat`, `longi`, `images`, `restaurant`, `events`, `entertainment`, `phone`, `email`, `hoteldetail`, `vipprice`, `vvipprice`, `deluxevipprice`) VALUES
(1, 'Vinpearl Luxury Landmark 81', 1, 'Quận Bình Thạnh, Thành phố Hồ Chí Minh, Việt Nam', 10.8027, 106.697, 'https://vcdn-dulich.vnecdn.net/2019/04/25/24-4-201930-235787743-w500-4413-1556159282.png', 'Thưởng thức ẩm thực thượng hạng được chế biến và sáng tạo bởi những đầu bếp hàng đầu trong không gian nhà hàng & quán bar sang trọng, đẳng cấp bậc nhất.\r\nA.Oriental Pearl Restaurant\r\nĐịa điểm : Tầng 66\r\nSức chứa : 136-150.\r\nB.Blu Pearl Bar\r\nĐịa điểm : Tầng 47\r\nSức chứa : 15-20\r\nC.Pearl Club\r\nĐịa điểm : Tầng 71\r\nSức chứa : 56\r\nD.Far East Lounge\r\nĐịa điểm : Tầng 48\r\nSức chứa : 30', 'Với trang thiết bị tối tân, hiện đại, phòng họp hiện đại đáp ứng tối đa nhu cầu giao lưu, họp mặt, bàn bạc công việc và mang lại thành công trên cả sự mong đợi.', 'Tận hưởng sự thoải mái và tiện lợi của không gian nghỉ dưỡng hiện đại, sang trọng cùng các tiện ích giải trí đồng bộ 5 sao.', '(+84) 2839718888', 'res.VPLLM81@vinpearl.com', 'Nằm trên bờ sông Sài Gòn đẹp như tranh vẽ, tòa nhà Vinpearl Luxury Landmark 81 nằm gói gọn một cách sang trọng và nổi bật nhất trong khu vườn xanh của Vinhomes Central Park. Bắt đầu từ tầng 47 trở lên, Vinpearl Luxury Landmark 81 là khách sạn cao nhất Việt Nam và Đông Nam Á. Khách sạn đã đặt một chuẩn mực mới cho trải nghiệm “Luxury defined by Viet Nam” với 223 phòng nghỉ sang trọng, nội thất sang trọng bậc nhất với tầm nhìn toàn cảnh tuyệt đẹp của thành phố.', '5.000.000 VNĐ', '6.000.000 VNĐ', '7.000.000 VNĐ'),
(2, 'Vinpearl Hotel Huế', 2, '50A Hùng Vương, Phường Phú Nhuận, Thành Phố Huế, Tỉnh Thừa Thiên Huế, Việt Nam', 16.4619, 107.597, 'https://owa.bestprice.vn/images/combos/large/combo-4n3d-kham-pha-hue-tron-ven-tai-vinpearl-hue-lang-co-resort-ve-may-bay-5bc4124fea789-660x470.jpg', 'Các nhà hàng trong khách sạn là chốn tiếp đón riêng dành cho những thượng khách xứng tầm “hoàng đế”. Mọi bước đi của thượng khách đều được chăm sóc và phục vụ tận tâm nhất\r\nA.Nhà hàng Lotus\r\nĐịa điểm : Tầng 7\r\nSức chứa : 290\r\nB.Nhà hàng The Prime\r\nĐịa điểm : Tầng 33\r\nSức chứa : 142\r\nC.Lobby Bar\r\nĐịa điểm : Tầng 1\r\nSức chứa : 54\r\nD.Aqua Bar\r\nĐịa điểm : Tầng 8\r\nE.Sky Bar\r\nĐịa điểm : Tầng thượng\r\nSức chứa : 60\r\n\r\n', 'Cùng Vinpearl Hotel Huế tạo nên những khởi đầu viên mãn và chắp cánh mọi thành công với hệ thống phòng họp “Long-Phụng” với thiết kế hiện đại', 'Các tiện ích trong khách sạn như Vincharm Spa, trung tâm fitness & yoga, bể bơi bốn mùa hiện đại đem đến những trải nghiệm nghỉ dưỡng thú vị cho du khách.', '(+84) 1900 636 699 - (+84) 243 974 9999', 'callcenter@vinpearl.com', 'Tọa lạc giữa lòng xứ Huế cổ kính với vị trí đắc địa “hướng sông Hương, tựa núi Ngự “, tòa tháp Vinpearl Hotel Huế với 33 tầng nổi lên như một nút huyết mạch của thành phố. Thiết kế hình khối vững chắc thấm đượm hơi thở hiện đại, mạnh mẽ nhưng Vinpearl Hotel Huế vẫn ôm trọn trong mình nét tinh hoa của Đại Việt với họa tiết sông nước núi non, kim long phụng bảo ẩn hiện trong mọi không gian nghỉ ngơi của khách hàng.', '3.000.000VNĐ', '3.500.000VNĐ', '4.000.000VNĐ'),
(3, 'VINPEARL CONDOTEL BEACHFRONT NHA TRANG', 3, '78-80 Trần Phú, Phường Lộc Thọ, Thành phố Nha Trang, Tỉnh Khánh Hòa', 12.2147, 109.209, 'https://du-lich.chudu24.com/f/m/1905/31/vnr-overview-1803109.jpg?w=550&c=1', 'Với 4 nhà hàng và quán bar sang trọng Vinpearl Condotel Beach Front Nha Trang mang đến cho thực khách trải nghiệm những không gian ẩm thực tuyệt hảo, thưởng thức nghệ thuật ẩm thực Á – Âu với thực đơn hơn 100 món ăn mỗi ngày.', 'Hệ thống phòng họp và ballroom lên đến 430m2 cùng các tiện ích hỗ trợ hiện đại sẽ đáp ứng được mọi nhu cầu cho các buổi gặp mặt mang tính chất công việc của khách hàng. Vinpearl Condotel Beachfront Nha Trang là điểm đến cho tổ chức sự kiện mới mẻ khi đến với thành phố Nha Trang.', 'Vinpearl Condotel Empire Nha Trang đáp ứng hoàn hảo các nhu cầu của khách lưu trú với đa tiện ích như gym, bể bơi ngoài trời, Vincharm Spa và trung tâm thương mại Vincom Plaza tiện lợi với đủ dịch vụ nằm ngay dưới chân tòa nhà.', '(+84) 1900 636 699 - (+84) 243 974 9999', 'callcenter@vinpearl.com', 'Vinpearl Condotel Beachfront Nha Trang là những căn hộ khách sạn tiện nghi có tầm nhìn hướng biển, mang đến kỳ nghỉ trọn vẹn bên vịnh Nha Trang – một trong những vịnh biển đẹp nhất hành tinh. Đây là điểm đến mới đẳng cấp, nơi hội tụ “tất cả trong một” với trung tâm thương mại, nhà hàng sang trọng, bể bơi ngoài trời ngắm toàn cảnh vịnh biển, không gian hội họp đẳng cấp tiến bước thành công.', '4.000.000VNĐ', '5.000.000 VNĐ', '6.000.000 VNĐ'),
(4, 'Vinpearl Hotel Quảng Bình', 4, 'Đường Quách Xuân Kỳ, Thành phố Đồng Hới, Quảng Bình', 17.468, 106.625, 'https://pix10.agoda.net/hotelImages/5663774/-1/00f12730a33df7f9a31657bcdf749e93.jpg?s=1024x768', 'Không những thổi hồn tinh tế cho các món ăn đặc trưng xứ Quảng Bình các nhà hàng trong khách sạn còn mang tới du khách hành trình trải nghiệm ẩm thực miền Bắc – Trung – Nam và thực đơn Âu – Á dưới bàn tay tài hoa của những đầu bếp hàng đầu.', 'Vinpearl Hotel Quảng Bình trang bị sẵn hệ thống phòng họp với thiết kế cao cấp, trang thiết bị thông minh, hiện đại phù hợp cho những buổi họp/ ký kết hợp đồng quan trọng nhằm mang đến sự chuyên nghiệp, góp phần mang lại sự thành công cho chuyến công du của quý khách.', 'Các tiện ích trong khách sạn như Vincharm Spa, phòng gym, bể bơi trong nhà hiện đại đem đến những trải nghiệm nghỉ dưỡng thú vị cho du khách.', '(+84) 2323 900 888', 'res.VPCHQB@vinpearl.com', 'Vinpearl Hotel Quảng Bình, nằm ở trung tâm thành phố Đồng Hới bên dòng sông Nhật Lệ thanh bình. Đây là điểm đến lý tưởng để tận hưởng những khoảnh khắc thư giãn tuyệt vời bên gia đình, bạn bè và người thân của mình trên mảnh đất Quảng Bình đầy nắng và gió nhưng cũng đầy quyến rũ với những kỳ quan thế giới như vườn Quốc gia Phong Nha – Kẻ Bàng, Hang Sơn Đoòng, Đảo Yến – Vũng Chùa…', '4.000.000VNĐ', '5.000.000 VNĐ', '6.000.000 VNĐ'),
(5, 'Vinpearl Discovery Hà Tĩnh', 5, 'Phường Thịnh Lộc, Huyện Lộc Hà, tỉnh Hà Tĩnh', 18.4643, 105.876, 'https://q-cf.bstatic.com/images/hotel/max1024x768/194/194400346.jpg', 'Vinpearl Discovery Hà Tĩnh cung cấp thực đơn đa dạng gồm các món ăn tinh hoa của ẩm thực truyền thống Việt Nam và các món ăn nổi tiếng quốc tế trong không gian hiện đại và sang trọng đáp ứng mọi nhu cầu về ẩm thực của các du khách.', '', 'Vinpearl Discovery Hà Tĩnh đem đến nhiều trải nghiệm giải trí thú vị cho du khách với Vincharm Spa, bể bơi ngoài trời và đặc biệt Công viên nước Hà Tĩnh là công viên nước 5 sao lớn nhất Bắc Trung Bộ.', '(+84 239) 363 668 666', 'res.VPDSHT@vinpearl.com', 'Với tầm nhìn tuyệt đẹp hướng ra bãi biển Xuân Hải (Cửa Sót) hoang sơ, Vinpearl Discovery Hà Tĩnh được thiết kế với tất cả các cửa sổ và ban công nhìn ra biển. Đây cũng là biệt thự biển 5 sao đầu tiên có công viên nước trong tỉnh mang đến kỳ nghỉ thư giãn và giải trí cho khách hàng.', '3.000.000VNĐ', '3.500.000VNĐ', '4.000.000VNĐ'),
(6, 'Vinpearl Hotel Thanh Hóa', 6, 'Số 27 Trần Phú, phường Điện Biên, TP Thanh Hóa, tỉnh Thanh Hóa, Việt Nam', 19.8067, 105.785, 'https://pix10.agoda.net/hotelImages/5686721/-1/0c1b49233087f72121e2def06320c78f.jpg?s=1024x768', 'Khách hàng sẽ được thưởng thức trọn vẹn nét ẩm thực địa phương đặc trưng và các món Âu, Á tinh tế đến từ bàn tay các đầu bếp bậc thầy của chuỗi khách sạn 5 sao.', 'Vinpearl Hotel Thanh Hóa với hệ thống các phòng họp được trang bị đầy đủ thiết bị hiện đại là nơi khởi đầu thành công cho mọi sự kiện quan trọng', 'Các tiện ích trong khách sạn như Vincharm Spa, phòng gym, bể bơi bốn mùa hiện đại đem đến những trải nghiệm nghỉ dưỡng thú vị cho du khách.', '(+84) 1900 636 699 - (+84) 243 974 9999', 'callcenter@vinpearl.com', 'Nằm trên vùng đất xứ Thanh “địa linh nhân kiệt”, Vinpearl Hotel Thanh Hóa được xướng tên trên tuyến phố huyết mạch trong quần thể khu vực hành chính đầu não của thành phố, là một trong những điểm khí thiêng hội tụ. Đây là một trong những lựa chọn tuyệt vời cho chuyến công du, làm việc dài ngày với những tiện nghi phòng ở cao cấp cùng tổ hợp trung tâm thương mại Vincom Plaza, Shophouse hiện đại.', '2.000.000VNĐ', '2.200.000 VNĐ', '2.400.000 VNĐ'),
(7, 'Vinpearl Discovery 3 Phú Quốc', 7, 'Khu Bãi Dài, xã Gành Dầu, huyện Phú Quốc, tỉnh Kiên Giang', 10.3374, 103.946, 'https://d1nabgopwop1kh.cloudfront.net/hotel-asset/30000002100465703_wh_75', 'Với các nhà hàng và quán bar sang trọng, thực đơn đa dạng từ buffet đến A la carte Vinpearl Discovery 3 Phú Quốc đem đến cho du khách thiên đường ẩm thực phong phú đáp ứng mọi nhu cầu của du khách.', 'Vinpearl Discovery 3 Phú Quốc là địa điểm hoàn hảo để du lịch kết hợp cùng hội nghị, sự kiện đẳng cấp với trung tâm hội nghị đẳng cấp quốc tế Vinpearl Convention Center Phú Quốc và không gian teambuilding không giới hạn.', 'Bể bơi lớn ngoài trời, Akoya Spa, kết nối cùng các khu vui chơi khác của Vinpearl Phú Quốc như Vinpearl Land & Vinpearl Safari, Vinpearl Golf... đem đến những trải nghiệm vui chơi giải trí rất độc đáo và đáng nhớ.', '(+84) 1900 636 699 - (+84) 243 974 9999', 'callcenter@vinpearl.com', 'Vinpearl Discovery 3 Phú Quốc sở hữu các biệt thự nép mình dọc theo sườn núi với tầm nhìn khoáng đạt bao trọn đất trời. Không khí trong lành và cảnh quan độc đáo của khu nghỉ dưỡng thích hợp với những vị khách yêu thích khám phá cảnh sắc thiên nhiên xen kẽ giữa đại dương bao la và núi đồi trùng điệp.', '4.000.000VNĐ', '5.000.000VNĐ', '6.000.000VNĐ'),
(8, 'Vinpearl Hotel Rivera Hải Phòng', 8, 'Đường Manhattan 9, khu đô thị Vinhomes Imperia, đường Hà Nội, Hồng Bàng, Thành phố Hải Phòng', 20.8582, 106.684, 'http://cafefcdn.com/thumb_w/650/2017/http-channel-vcmedia-vn-prupload-270-2017-04-img-201704171501416268-1492477239137.jpg', 'Với bàn tay chế biến của những đầu bếp tài hoa và sự phục vụ chuyên nghiệp của đội ngũ nhân viên, du khách không chỉ được thưởng thức tinh hoa ẩm thực Việt Nam và quốc tế được bài trí đặc sắc mà còn được dịp ấn tượng bởi dịch vụ tận tâm và không gian sang trọng.', 'Với hệ thống phòng họp tối tân và đội ngũ chuyên gia đầy nhiệt huyết, Vinpearl Hotel Rivera Hải Phòng là lựa chọn hoàn hảo cho các chương trình hội nghị lớn nhỏ, các sự kiện đón tiếp, chiêu đãi khách VIP, gặp gỡ đối tác quan trọng.', 'Các dịch vụ khách sạn đẳng cấp như Vincharm Spa, trung tâm thể hình, bể bơi ngoài trời... cộng hưởng với các tiện ích nội khu hiện đại của Vinhomes Imperia Hải Phòng đem đến những trải nghiệm lưu trú trọn vẹn nhất cho khách hàng.', '(+84) 1900 636 699 - (+84) 243 974 9999', 'callcenter@vinpearl.com', 'Tọa lạc tại vị trí đắc địa ngay trong khuôn viên Vinhomes Imperia, Vinpearl Hotel Rivera Hải Phòng là tòa khách sạn 5 sao mang dáng vẻ thanh lịch và sang trọng. Hơi thở trữ tình và cổ điển của những công trình phong cách Pháp trên đất Cảng giao hòa với vẻ đẹp thời đại năng động và ấn tượng, được cảm nhận rõ nét qua lối kiến trúc tân cổ điển tinh mỹ.', '4.000.000VNĐ', '5.000.000VNĐ', '6.000.000VNĐ');

-- --------------------------------------------------------

--
-- Table structure for table `reservedrooms`
--

CREATE TABLE `reservedrooms` (
  `roombookedid` int(11) NOT NULL,
  `bookingid` int(11) DEFAULT NULL,
  `roomid` int(11) DEFAULT NULL,
  `startdate` date NOT NULL,
  `enddate` date NOT NULL,
  `roomtypeid` int(11) NOT NULL,
  `hotelid` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `reservedrooms`
--

INSERT INTO `reservedrooms` (`roombookedid`, `bookingid`, `roomid`, `startdate`, `enddate`, `roomtypeid`, `hotelid`) VALUES
(1, 1, 1, '2019-12-05', '2019-12-07', 0, 1),
(2, 1, 2, '2019-12-05', '2019-12-07', 0, 1),
(3, 1, 3, '2019-12-05', '2019-12-07', 0, 1);

-- --------------------------------------------------------

--
-- Table structure for table `rooms`
--

CREATE TABLE `rooms` (
  `roomid` int(11) NOT NULL,
  `hotelid` int(11) DEFAULT NULL,
  `roomtypeid` int(11) DEFAULT NULL,
  `roomstatus` int(11) DEFAULT NULL,
  `price` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `rooms`
--

INSERT INTO `rooms` (`roomid`, `hotelid`, `roomtypeid`, `roomstatus`, `price`) VALUES
(1, 1, 0, NULL, '5.000.000VNĐ'),
(2, 1, 0, NULL, '5.000.000VNĐ'),
(3, 1, 0, NULL, '5.000.000VNĐ'),
(4, 1, 0, NULL, '5.000.000VNĐ'),
(5, 1, 0, NULL, '5.000.000VNĐ'),
(6, 1, 1, NULL, '6.000.000VNĐ'),
(7, 1, 1, NULL, '6.000.000VNĐ'),
(8, 1, 1, NULL, '6.000.000VNĐ'),
(9, 1, 1, NULL, '6.000.000VNĐ'),
(10, 1, 1, NULL, '6.000.000VNĐ'),
(11, 1, 2, NULL, '7.000.000VNĐ'),
(12, 1, 2, NULL, '7.000.000VNĐ'),
(13, 1, 2, NULL, '7.000.000VNĐ'),
(14, 1, 2, NULL, '7.000.000VNĐ'),
(15, 1, 2, NULL, '7.000.000VNĐ'),
(16, 2, 0, NULL, '3.000.000VNĐ'),
(17, 2, 0, NULL, '3.000.000VNĐ'),
(18, 2, 0, NULL, '3.000.000VNĐ'),
(19, 2, 0, NULL, '3.000.000VNĐ'),
(20, 2, 1, NULL, '3.500.000VNĐ'),
(21, 2, 1, NULL, '3.500.000VNĐ'),
(22, 2, 1, NULL, '3.500.000VNĐ'),
(23, 2, 2, NULL, '4.000.000VNĐ'),
(24, 2, 2, NULL, '4.000.000VNĐ'),
(25, 2, 2, NULL, '4.000.000VNĐ'),
(26, 3, 0, NULL, '4.000.000VNĐ'),
(27, 3, 0, NULL, '4.000.000VNĐ'),
(28, 3, 0, NULL, '4.000.000VNĐ'),
(29, 3, 0, NULL, '4.000.000VNĐ'),
(30, 3, 1, NULL, '5.000.000VNĐ'),
(31, 3, 1, NULL, '5.000.000VNĐ'),
(32, 3, 1, NULL, '5.000.000VNĐ'),
(33, 3, 2, NULL, '6.000.000VNĐ'),
(34, 3, 2, NULL, '6.000.000VNĐ'),
(35, 3, 2, NULL, '6.000.000VNĐ'),
(36, 4, 0, NULL, '4.000.000VNĐ'),
(37, 4, 0, NULL, '4.000.000VNĐ'),
(38, 4, 0, NULL, '4.000.000VNĐ'),
(39, 4, 0, NULL, '4.000.000VNĐ'),
(40, 4, 1, NULL, '5.000.000VNĐ'),
(41, 4, 1, NULL, '5.000.000VNĐ'),
(42, 4, 1, NULL, '5.000.000VNĐ'),
(43, 4, 2, NULL, '6.000.000VNĐ'),
(44, 4, 2, NULL, '6.000.000VNĐ'),
(45, 4, 2, NULL, '6.000.000VNĐ'),
(46, 5, 0, NULL, '3.000.000VNĐ'),
(47, 5, 0, NULL, '3.000.000VNĐ'),
(48, 5, 0, NULL, '3.000.000VNĐ'),
(49, 5, 1, NULL, '3.500.000VNĐ'),
(50, 5, 1, NULL, '3.500.000VNĐ'),
(51, 5, 1, NULL, '3.500.000VNĐ'),
(52, 5, 2, NULL, '4.000.000VNĐ'),
(53, 5, 2, NULL, '4.000.000VNĐ'),
(54, 5, 2, NULL, '4.000.000VNĐ'),
(55, 5, 2, NULL, '4.000.000VNĐ'),
(56, 6, 0, NULL, '2.000.000VNĐ'),
(57, 6, 0, NULL, '2.000.000VNĐ'),
(58, 6, 0, NULL, '2.000.000VNĐ'),
(59, 6, 0, NULL, '2.000.000VNĐ'),
(60, 6, 1, NULL, '2.200.000VNĐ'),
(61, 6, 1, NULL, '2.200.000VNĐ'),
(62, 6, 1, NULL, '2.200.000VNĐ'),
(63, 6, 2, NULL, '2.400.000VNĐ'),
(64, 6, 2, NULL, '2.400.000VNĐ'),
(65, 6, 2, NULL, '2.400.000VNĐ'),
(66, 7, 0, NULL, '4.000.000VNĐ'),
(67, 7, 0, NULL, '4.000.000VNĐ'),
(68, 7, 0, NULL, '4.000.000VNĐ'),
(69, 7, 0, NULL, '4.000.000VNĐ'),
(70, 7, 1, NULL, '5.000.000VNĐ'),
(71, 7, 1, NULL, '5.000.000VNĐ'),
(72, 7, 1, NULL, '5.000.000VNĐ'),
(73, 7, 2, NULL, '6.000.000VNĐ'),
(74, 7, 2, NULL, '6.000.000VNĐ'),
(75, 7, 2, NULL, '6.000.000VNĐ'),
(76, 8, 0, NULL, '4.000.000VNĐ'),
(77, 8, 0, NULL, '4.000.000VNĐ'),
(79, 8, 0, NULL, '4.000.000VNĐ'),
(80, 8, 1, NULL, '5.000.000VNĐ'),
(81, 8, 1, NULL, '5.000.000VNĐ'),
(82, 8, 1, NULL, '5.000.000VNĐ'),
(83, 8, 2, NULL, '6.000.000VNĐ'),
(84, 8, 2, NULL, '6.000.000VNĐ'),
(85, 8, 2, NULL, '6.000.000VNĐ');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `booking_infos`
--
ALTER TABLE `booking_infos`
  ADD PRIMARY KEY (`bookingid`),
  ADD KEY `hotel_id` (`hotelid`);

--
-- Indexes for table `cities`
--
ALTER TABLE `cities`
  ADD PRIMARY KEY (`cityid`);

--
-- Indexes for table `hotels`
--
ALTER TABLE `hotels`
  ADD PRIMARY KEY (`hotelid`),
  ADD KEY `city_id` (`cityid`);

--
-- Indexes for table `reservedrooms`
--
ALTER TABLE `reservedrooms`
  ADD PRIMARY KEY (`roombookedid`),
  ADD KEY `booking_id` (`bookingid`),
  ADD KEY `room_id` (`roomid`);

--
-- Indexes for table `rooms`
--
ALTER TABLE `rooms`
  ADD PRIMARY KEY (`roomid`),
  ADD KEY `hotel_id` (`hotelid`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `hotels`
--
ALTER TABLE `hotels`
  MODIFY `hotelid` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=1112;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `booking_infos`
--
ALTER TABLE `booking_infos`
  ADD CONSTRAINT `booking_infos_ibfk_1` FOREIGN KEY (`hotelid`) REFERENCES `hotels` (`hotelid`),
  ADD CONSTRAINT `booking_infos_ibfk_2` FOREIGN KEY (`hotelid`) REFERENCES `hotels` (`hotelid`);

--
-- Constraints for table `hotels`
--
ALTER TABLE `hotels`
  ADD CONSTRAINT `hotels_ibfk_1` FOREIGN KEY (`cityid`) REFERENCES `cities` (`cityid`),
  ADD CONSTRAINT `hotels_ibfk_2` FOREIGN KEY (`cityid`) REFERENCES `cities` (`cityid`);

--
-- Constraints for table `reservedrooms`
--
ALTER TABLE `reservedrooms`
  ADD CONSTRAINT `reservedrooms_ibfk_1` FOREIGN KEY (`bookingid`) REFERENCES `booking_infos` (`bookingid`),
  ADD CONSTRAINT `reservedrooms_ibfk_2` FOREIGN KEY (`roomid`) REFERENCES `rooms` (`roomid`);

--
-- Constraints for table `rooms`
--
ALTER TABLE `rooms`
  ADD CONSTRAINT `rooms_ibfk_1` FOREIGN KEY (`hotelid`) REFERENCES `hotels` (`hotelid`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
