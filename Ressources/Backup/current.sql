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
  KEY `fk_Episode_Series1_idx` (`Series_id`),
  KEY `fk_Episode_Image1_idx` (`Image_id`),
  CONSTRAINT `fk_Episode_Image1` FOREIGN KEY (`Image_id`) REFERENCES `Image` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_Episode_Series1` FOREIGN KEY (`Series_id`) REFERENCES `Series` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
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
  `Message` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Log`
--

LOCK TABLES `Log` WRITE;
/*!40000 ALTER TABLE `Log` DISABLE KEYS */;
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
  `Type` varchar(255) DEFAULT NULL,
  `Parameter` varchar(500) DEFAULT NULL,
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
  KEY `fk_Provider_Image1_idx` (`Image_id`),
  CONSTRAINT `fk_Provider_Image1` FOREIGN KEY (`Image_id`) REFERENCES `Image` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
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
-- Table structure for table `ProviderUrl`
--

DROP TABLE IF EXISTS `ProviderUrl`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ProviderUrl` (
  `Episode_id` int(11) NOT NULL AUTO_INCREMENT,
  `Provider_id` int(11) NOT NULL,
  `Url` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`Episode_id`,`Provider_id`),
  KEY `fk_Episode_has_Provider_Provider1_idx` (`Provider_id`),
  KEY `fk_Episode_has_Provider_Episode1_idx` (`Episode_id`),
  CONSTRAINT `fk_Episode_has_Provider_Episode1` FOREIGN KEY (`Episode_id`) REFERENCES `Episode` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_Episode_has_Provider_Provider1` FOREIGN KEY (`Provider_id`) REFERENCES `Provider` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ProviderUrl`
--

LOCK TABLES `ProviderUrl` WRITE;
/*!40000 ALTER TABLE `ProviderUrl` DISABLE KEYS */;
/*!40000 ALTER TABLE `ProviderUrl` ENABLE KEYS */;
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
  `Seriescol` int(11) DEFAULT NULL,
  `DataProviderUrl` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_Series_Image1_idx` (`Image_id`),
  CONSTRAINT `fk_Series_Image1` FOREIGN KEY (`Image_id`) REFERENCES `Image` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
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
-- Table structure for table `SeriesProvider`
--

DROP TABLE IF EXISTS `SeriesProvider`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `SeriesProvider` (
  `Series_id` int(11) NOT NULL AUTO_INCREMENT,
  `Provider_id` int(11) NOT NULL,
  `BaseUrl` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`Series_id`,`Provider_id`),
  KEY `fk_Series_has_Provider_Provider1_idx` (`Provider_id`),
  KEY `fk_Series_has_Provider_Series1_idx` (`Series_id`),
  CONSTRAINT `fk_Series_has_Provider_Provider1` FOREIGN KEY (`Provider_id`) REFERENCES `Provider` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_Series_has_Provider_Series1` FOREIGN KEY (`Series_id`) REFERENCES `Series` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `SeriesProvider`
--

LOCK TABLES `SeriesProvider` WRITE;
/*!40000 ALTER TABLE `SeriesProvider` DISABLE KEYS */;
/*!40000 ALTER TABLE `SeriesProvider` ENABLE KEYS */;
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
  KEY `fk_Account_Provider1_idx` (`Provider_id`),
  KEY `fk_Account_User1_idx` (`User_id`),
  KEY `fk_ThirdPartyAccount_Credentials1_idx` (`Credentials_id`),
  CONSTRAINT `fk_Account_Provider1` FOREIGN KEY (`Provider_id`) REFERENCES `Provider` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_Account_User1` FOREIGN KEY (`User_id`) REFERENCES `User` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_ThirdPartyAccount_Credentials1` FOREIGN KEY (`Credentials_id`) REFERENCES `Credentials` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
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
-- Table structure for table `User`
--

DROP TABLE IF EXISTS `User`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `User` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) DEFAULT NULL,
  `Credentials_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_User_Credentials1_idx` (`Credentials_id`),
  CONSTRAINT `fk_User_Credentials1` FOREIGN KEY (`Credentials_id`) REFERENCES `Credentials` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `User`
--

LOCK TABLES `User` WRITE;
/*!40000 ALTER TABLE `User` DISABLE KEYS */;
/*!40000 ALTER TABLE `User` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `WatchPointer`
--

DROP TABLE IF EXISTS `WatchPointer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `WatchPointer` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `User_id` int(11) NOT NULL,
  `Episode_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_WatchPointer_User_idx` (`User_id`),
  KEY `fk_WatchPointer_Episode1_idx` (`Episode_id`),
  CONSTRAINT `fk_WatchPointer_Episode1` FOREIGN KEY (`Episode_id`) REFERENCES `Episode` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_WatchPointer_User` FOREIGN KEY (`User_id`) REFERENCES `User` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
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

-- Dump completed on 2018-01-17 12:40:02
