-- MySQL dump 10.13  Distrib 8.0.18, for macos10.14 (x86_64)
--
-- Host: localhost    Database: book
-- ------------------------------------------------------
-- Server version	8.0.17

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `md_book`
--

DROP TABLE IF EXISTS `md_book`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `md_book` (
  `book_id` int(11) NOT NULL AUTO_INCREMENT,
  `book_name` varchar(500) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `identify` varchar(100) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `order_index` int(11) NOT NULL DEFAULT '0',
  `description` varchar(1000) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `cover` varchar(1000) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `editor` varchar(50) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `status` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '0',
  `privately_owned` int(11) NOT NULL DEFAULT '0',
  `private_token` varchar(500) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `member_id` int(11) NOT NULL DEFAULT '0',
  `doc_count` int(11) NOT NULL DEFAULT '0',
  `comment_count` int(11) NOT NULL DEFAULT '0',
  `vcnt` int(11) NOT NULL DEFAULT '0',
  `star` int(11) NOT NULL DEFAULT '0',
  `score` int(11) NOT NULL DEFAULT '40',
  `cnt_score` int(11) NOT NULL DEFAULT '0',
  `cnt_comment` int(11) NOT NULL DEFAULT '0',
  `author` varchar(50) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `author_url` varchar(1000) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL,
  `modify_time` datetime NOT NULL,
  `release_time` datetime NOT NULL,
  PRIMARY KEY (`book_id`),
  UNIQUE KEY `identify` (`identify`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `md_book`
--

LOCK TABLES `md_book` WRITE;
/*!40000 ALTER TABLE `md_book` DISABLE KEYS */;
/*!40000 ALTER TABLE `md_book` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-01-20 15:06:23
