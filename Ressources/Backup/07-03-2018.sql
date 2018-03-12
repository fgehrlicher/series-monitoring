-- MySQL dump 10.13  Distrib 5.7.21, for Linux (x86_64)
--
-- Host: localhost    Database: series
-- ------------------------------------------------------
-- Server version	5.7.21

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `BaseUrl`
--

DROP TABLE IF EXISTS `BaseUrl`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `BaseUrl` (
  `id` int(11) NOT NULL,
  `Url` varchar(255) DEFAULT NULL,
  `Series_id` int(11) NOT NULL,
  `Provider_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_BaseUrl_Series_idx` (`Series_id`),
  KEY `fk_BaseUrl_Provider_idx` (`Provider_id`),
  CONSTRAINT `fk_BaseUrl_Provider` FOREIGN KEY (`Provider_id`) REFERENCES `Provider` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_BaseUrl_Series` FOREIGN KEY (`Series_id`) REFERENCES `Series` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `BaseUrl`
--

LOCK TABLES `BaseUrl` WRITE;
/*!40000 ALTER TABLE `BaseUrl` DISABLE KEYS */;
/*!40000 ALTER TABLE `BaseUrl` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Credentials`
--

DROP TABLE IF EXISTS `Credentials`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Credentials` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Password` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Credentials`
--

LOCK TABLES `Credentials` WRITE;
/*!40000 ALTER TABLE `Credentials` DISABLE KEYS */;
/*!40000 ALTER TABLE `Credentials` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Episode`
--

DROP TABLE IF EXISTS `Episode`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Episode` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Series_id` int(11) NOT NULL,
  `Image_id` int(11) NOT NULL,
  `Episode` int(11) DEFAULT NULL,
  `Season` int(11) DEFAULT NULL,
  `Title` varchar(255) DEFAULT NULL,
  `Description` text,
  `ReleaseDate` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_Episode_Series_idx` (`Series_id`),
  KEY `fk_Episode_Image_idx` (`Image_id`),
  CONSTRAINT `fk_Episode_Image` FOREIGN KEY (`Image_id`) REFERENCES `Image` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_Episode_Series` FOREIGN KEY (`Series_id`) REFERENCES `Series` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Episode`
--

LOCK TABLES `Episode` WRITE;
/*!40000 ALTER TABLE `Episode` DISABLE KEYS */;
/*!40000 ALTER TABLE `Episode` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Image`
--

DROP TABLE IF EXISTS `Image`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Image` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Path` varchar(255) DEFAULT NULL,
  `OriginUrl` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Image`
--

LOCK TABLES `Image` WRITE;
/*!40000 ALTER TABLE `Image` DISABLE KEYS */;
/*!40000 ALTER TABLE `Image` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Log`
--

DROP TABLE IF EXISTS `Log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Type` int(11) DEFAULT NULL,
  `Message` varchar(1000) DEFAULT NULL,
  `Time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=104 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Log`
--

LOCK TABLES `Log` WRITE;
/*!40000 ALTER TABLE `Log` DISABLE KEYS */;
INSERT INTO `Log` VALUES (2,3,'Not Found | Caller Ip: 127.0.0.1:58162 | Request uri: /test/','2018-03-02 16:12:00'),(3,3,'Not Found | Caller Ip: 127.0.0.1:58162 | Request uri: /favicon.ico','2018-03-02 16:12:01'),(4,3,'Not Found | Caller Ip: 127.0.0.1:58162 | Request uri: /test/','2018-03-02 16:12:01'),(5,3,'Not Found | Caller Ip: 127.0.0.1:58174 | Request uri: /robots.txt','2018-03-02 16:12:01'),(6,3,'Not Found | Caller Ip: 127.0.0.1:58162 | Request uri: /favicon.ico','2018-03-02 16:12:02'),(7,3,'Not Found | Caller Ip: 127.0.0.1:58162 | Request uri: /favicon.ico','2018-03-02 16:12:06'),(8,3,'Not Found | Caller Ip: 127.0.0.1:58162 | Request uri: /favicon.ico','2018-03-02 16:12:07'),(9,3,'Not Found | Caller Ip: 127.0.0.1:58568 | Request uri: /robots.txt','2018-03-02 16:17:18'),(10,3,'Not Found | Caller Ip: 127.0.0.1:58566 | Request uri: /favicon.ico','2018-03-02 16:17:18'),(11,3,'Not Found | Caller Ip: 127.0.0.1:58566 | Request uri: /favicon.ico','2018-03-02 16:17:21'),(12,3,'Not Found | Caller Ip: 127.0.0.1:58566 | Request uri: /test/','2018-03-02 16:17:27'),(13,3,'Not Found | Caller Ip: 127.0.0.1:58566 | Request uri: /facion.ico/','2018-03-02 16:17:36'),(14,3,'Not Found | Caller Ip: 127.0.0.1:58566 | Request uri: /facion.ico','2018-03-02 16:17:38'),(15,3,'Not Found | Caller Ip: 127.0.0.1:58566 | Request uri: /facion.ico','2018-03-02 16:17:39'),(16,3,'Not Found | Caller Ip: 127.0.0.1:58566 | Request uri: /favicon.ico','2018-03-02 16:17:39'),(17,3,'Not Found | Caller Ip: 127.0.0.1:58566 | Request uri: /facion.ico','2018-03-02 16:17:40'),(18,3,'Not Found | Caller Ip: 127.0.0.1:58566 | Request uri: /favicon.ico','2018-03-02 16:17:40'),(19,3,'Not Found | Caller Ip: 127.0.0.1:58566 | Request uri: /facion.ico','2018-03-02 16:17:40'),(20,3,'Not Found | Caller Ip: 127.0.0.1:58601 | Request uri: /favicon.ico','2018-03-02 16:17:40'),(21,3,'Not Found | Caller Ip: 127.0.0.1:58566 | Request uri: /facion.ico','2018-03-02 16:17:40'),(22,3,'Not Found | Caller Ip: 127.0.0.1:58566 | Request uri: /facion.ico','2018-03-02 16:17:40'),(23,3,'Not Found | Caller Ip: 127.0.0.1:58566 | Request uri: /facion.ico','2018-03-02 16:17:41'),(24,3,'Not Found | Caller Ip: 127.0.0.1:58568 | Request uri: /robots.txt','2018-03-02 16:17:41'),(25,3,'Not Found | Caller Ip: 127.0.0.1:58566 | Request uri: /favicon.ico','2018-03-02 16:17:41'),(26,3,'Not Found | Caller Ip: 127.0.0.1:58566 | Request uri: /facion.ico','2018-03-02 16:17:41'),(27,3,'Not Found | Caller Ip: 127.0.0.1:58566 | Request uri: /favicon.ico','2018-03-02 16:17:41'),(28,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /favicon.ico','2018-03-02 16:23:04'),(29,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /favicon.ico','2018-03-02 16:23:07'),(30,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/','2018-03-02 16:23:10'),(31,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/','2018-03-02 16:23:10'),(32,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /favicon.ico','2018-03-02 16:23:11'),(33,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/','2018-03-02 16:23:11'),(34,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/','2018-03-02 16:23:11'),(35,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /favicon.ico','2018-03-02 16:23:11'),(36,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:13'),(37,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:13'),(38,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /favicon.ico','2018-03-02 16:23:14'),(39,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:14'),(40,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:14'),(41,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:14'),(42,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:14'),(43,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:15'),(44,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:15'),(45,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:15'),(46,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:15'),(47,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:15'),(48,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:15'),(49,3,'Not Found | Caller Ip: 127.0.0.1:58689 | Request uri: /favicon.ico','2018-03-02 16:23:15'),(50,3,'Not Found | Caller Ip: 127.0.0.1:58689 | Request uri: /test/asdasda','2018-03-02 16:23:16'),(51,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /favicon.ico','2018-03-02 16:23:16'),(52,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:16'),(53,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /favicon.ico','2018-03-02 16:23:16'),(54,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:16'),(55,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:16'),(56,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /favicon.ico','2018-03-02 16:23:16'),(57,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:16'),(58,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /favicon.ico','2018-03-02 16:23:17'),(59,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:17'),(60,3,'Not Found | Caller Ip: 127.0.0.1:58719 | Request uri: /robots.txt','2018-03-02 16:23:17'),(61,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /favicon.ico','2018-03-02 16:23:17'),(62,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:17'),(63,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:17'),(64,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /favicon.ico','2018-03-02 16:23:17'),(65,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:17'),(66,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /favicon.ico','2018-03-02 16:23:17'),(67,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:17'),(68,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /favicon.ico','2018-03-02 16:23:18'),(69,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:18'),(70,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /favicon.ico','2018-03-02 16:23:18'),(71,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:18'),(72,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /favicon.ico','2018-03-02 16:23:18'),(73,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:18'),(74,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /favicon.ico','2018-03-02 16:23:18'),(75,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:18'),(76,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /favicon.ico','2018-03-02 16:23:18'),(77,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:18'),(78,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /favicon.ico','2018-03-02 16:23:19'),(79,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:19'),(80,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /favicon.ico','2018-03-02 16:23:19'),(81,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:19'),(82,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /favicon.ico','2018-03-02 16:23:19'),(83,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:19'),(84,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /favicon.ico','2018-03-02 16:23:19'),(85,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:19'),(86,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /favicon.ico','2018-03-02 16:23:19'),(87,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:19'),(88,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /favicon.ico','2018-03-02 16:23:19'),(89,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /test/asdasda','2018-03-02 16:23:20'),(90,3,'Not Found | Caller Ip: 127.0.0.1:58648 | Request uri: /favicon.ico','2018-03-02 16:23:20'),(91,3,'Not Found | Caller Ip: 127.0.0.1:59054 | Request uri: /test/asdasda','2018-03-02 16:51:23'),(92,3,'Not Found | Caller Ip: 127.0.0.1:59058 | Request uri: /robots.txt','2018-03-02 16:51:23'),(93,3,'Not Found | Caller Ip: 127.0.0.1:59054 | Request uri: /favicon.ico','2018-03-02 16:51:24'),(94,3,'Not Found | Caller Ip: 127.0.0.1:59558 | Request uri: /t','2018-03-03 19:42:40'),(95,3,'Not Found | Caller Ip: 127.0.0.1:59558 | Request uri: /favicon.ico','2018-03-03 19:42:41'),(96,3,'Internal Server Error | Caller Ip: 127.0.0.1:59558 | Request uri: /logs/','2018-03-03 19:47:53'),(97,3,'Internal Server Error | Caller Ip: 127.0.0.1:59583 | Request uri: /logs/','2018-03-03 19:47:54'),(98,3,'Not Found | Caller Ip: 127.0.0.1:59728 | Request uri: /robots.txt','2018-03-03 20:05:01'),(99,3,'Not Found | Caller Ip: 127.0.0.1:59864 | Request uri: /robots.txt','2018-03-03 20:22:28'),(100,3,'Not Found | Caller Ip: 127.0.0.1:50408 | Request uri: /robots.txt','2018-03-03 20:29:37'),(101,3,'Not Found | Caller Ip: 127.0.0.1:50408 | Request uri: /robots.txt','2018-03-03 20:29:37'),(102,3,'Not Found | Caller Ip: 127.0.0.1:50406 | Request uri: /favicon.ico','2018-03-03 20:29:37'),(103,3,'Not Found | Caller Ip: 127.0.0.1:50406 | Request uri: /tesststs','2018-03-03 20:29:41');
/*!40000 ALTER TABLE `Log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ManualAction`
--

DROP TABLE IF EXISTS `ManualAction`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ManualAction` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `done` int(11) DEFAULT NULL,
  `Type` varchar(45) DEFAULT NULL,
  `Message` varchar(255) DEFAULT NULL,
  `Time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ManualAction`
--

LOCK TABLES `ManualAction` WRITE;
/*!40000 ALTER TABLE `ManualAction` DISABLE KEYS */;
/*!40000 ALTER TABLE `ManualAction` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Provider`
--

DROP TABLE IF EXISTS `Provider`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Provider` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) DEFAULT NULL,
  `Image_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_Provider_Image_idx` (`Image_id`),
  CONSTRAINT `fk_Provider_Image` FOREIGN KEY (`Image_id`) REFERENCES `Image` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Provider`
--

LOCK TABLES `Provider` WRITE;
/*!40000 ALTER TABLE `Provider` DISABLE KEYS */;
/*!40000 ALTER TABLE `Provider` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Series`
--

DROP TABLE IF EXISTS `Series`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Series` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Title` varchar(255) DEFAULT NULL,
  `Image_id` int(11) NOT NULL,
  `ProviderUrl` varchar(255) DEFAULT NULL,
  `ProviderType` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_Series_Image_idx` (`Image_id`),
  CONSTRAINT `fk_Series_Image` FOREIGN KEY (`Image_id`) REFERENCES `Image` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Series`
--

LOCK TABLES `Series` WRITE;
/*!40000 ALTER TABLE `Series` DISABLE KEYS */;
/*!40000 ALTER TABLE `Series` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ThirdPartyAccount`
--

DROP TABLE IF EXISTS `ThirdPartyAccount`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ThirdPartyAccount` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Provider_id` int(11) NOT NULL,
  `User_id` int(11) NOT NULL,
  `Credentials_id` int(11) NOT NULL,
  `Username` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_Account_Provider_idx` (`Provider_id`),
  KEY `fk_ThirdPartyAccount_Credentials_idx` (`Credentials_id`),
  CONSTRAINT `fk_Account_Provider` FOREIGN KEY (`Provider_id`) REFERENCES `Provider` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_ThirdPartyAccount_Credentials` FOREIGN KEY (`Credentials_id`) REFERENCES `Credentials` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ThirdPartyAccount`
--

LOCK TABLES `ThirdPartyAccount` WRITE;
/*!40000 ALTER TABLE `ThirdPartyAccount` DISABLE KEYS */;
/*!40000 ALTER TABLE `ThirdPartyAccount` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Url`
--

DROP TABLE IF EXISTS `Url`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Url` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Url` varchar(255) DEFAULT NULL,
  `Episode_id` int(11) NOT NULL,
  `Provider_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_EpisodeLinks_Episode_idx` (`Episode_id`),
  KEY `fk_Url_Provider_idx` (`Provider_id`),
  CONSTRAINT `fk_EpisodeLinks_Episode` FOREIGN KEY (`Episode_id`) REFERENCES `Episode` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_Url_Provider` FOREIGN KEY (`Provider_id`) REFERENCES `Provider` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Url`
--

LOCK TABLES `Url` WRITE;
/*!40000 ALTER TABLE `Url` DISABLE KEYS */;
/*!40000 ALTER TABLE `Url` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `WatchPointer`
--

DROP TABLE IF EXISTS `WatchPointer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `WatchPointer` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Episode_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_WatchPointer_Episode_idx` (`Episode_id`),
  CONSTRAINT `fk_WatchPointer_Episode` FOREIGN KEY (`Episode_id`) REFERENCES `Episode` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `WatchPointer`
--

LOCK TABLES `WatchPointer` WRITE;
/*!40000 ALTER TABLE `WatchPointer` DISABLE KEYS */;
/*!40000 ALTER TABLE `WatchPointer` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-03-07 14:28:41
