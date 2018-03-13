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
  `OriginUrl` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Image`
--

LOCK TABLES `Image` WRITE;
/*!40000 ALTER TABLE `Image` DISABLE KEYS */;
INSERT INTO `Image` VALUES (4,'Provider/burningseries.png','https://pbs.twimg.com/profile_images/60964220'),(5,'Provider/netflix.png','http://o.aolcdn.com/hss/storage/midas/160cc7dc83bd656bfd15fcd8a13f06b0/203978869/13502130_10153696123838870_522765110773053650_n.png'),(6,'Provider/primevideo.png','https://is4-ssl.mzstatic.com/image/thumb/Purple118/v4/f6/00/72/f600720a-b799-4bc4-ad70-2026419ec1fe/AppIcon-1x_U007emarketing-85-220-0-9.png/600x600bf.jpg'),(7,'Series/the-flash','https://image.tmdb.org/t/p/w300_and_h450_bestv2/lUFK7ElGCk9kVEryDJHICeNdmd1.jpg'),(8,'Series/bojack-horseman','https://image.tmdb.org/t/p/w300_and_h450_bestv2/1bnrSsJNejoQq8lGBDECQroGjPz.jpg'),(9,'Series/marvel\'s-agents-of-s.h.i.e.l.d.','https://image.tmdb.org/t/p/w300_and_h450_bestv2/xjm6uVktPuKXNILwjLXwVG5d5BU.jpg'),(10,'Series/american-gods','https://image.tmdb.org/t/p/w300_and_h450_bestv2/gevw5nZRYz2kWj1PqW9pz4sgeeZ.jpg'),(11,'Series/archer','https://image.tmdb.org/t/p/w300_and_h450_bestv2/k10h9wJW06YcCA2aOKWaJWfVjHQ.jpg'),(12,'Series/arrow','https://image.tmdb.org/t/p/w300_and_h450_bestv2/mo0FP1GxOFZT4UDde7RFDz5APXF.jpg'),(13,'Series/marvel\'s-daredevil','https://image.tmdb.org/t/p/w300_and_h450_bestv2/cidOqJL8tayqvv3TpfTQCsgeITu.jpg'),(14,'Series/marvel\'s-jessica-jones','https://image.tmdb.org/t/p/w300_and_h450_bestv2/8a7e2GNpMnjI2hgRZH3jq2c7ffv.jpg'),(15,'Series/marvel\'s-iron-fist','https://image.tmdb.org/t/p/w300_and_h450_bestv2/nv4nLXbDhcISPP8C1mgaxKU50KO.jpg'),(16,'Series/marvel\'s-the-punisher','https://image.tmdb.org/t/p/w300_and_h450_bestv2/s2YM9zHF3tf2coi8t0UEzYrOHg8.jpg'),(17,'Series/marvel\'s-the-defenders','https://image.tmdb.org/t/p/w300_and_h450_bestv2/49XzINhH4LFsgz7cx6TOPcHUJUL.jpg'),(18,'Series/marvel\'s-luke-cage','https://image.tmdb.org/t/p/w300_and_h450_bestv2/9nWZZ1ghE0LuXEWJi7QjCymHygi.jpg'),(19,'Series/dexter','https://image.tmdb.org/t/p/w300_and_h450_bestv2/ydmfheI5cJ4NrgcupDEwk8I8y5q.jpg'),(20,'Series/the-walking-dead','https://image.tmdb.org/t/p/w300_and_h450_bestv2/yn7psGTZsHumHOkLUmYpyrIcA2G.jpg'),(21,'Series/futurama','https://image.tmdb.org/t/p/w300_and_h450_bestv2/iN0LOeE2JnJpIy4jF7imUjO6jwn.jpg'),(22,'Series/family-guy','https://image.tmdb.org/t/p/w300_and_h450_bestv2/gBGUL1UTUNmdRQT8gA1LUV4yg39.jpg'),(23,'Series/the-simpsons','https://image.tmdb.org/t/p/w300_and_h450_bestv2/yTZQkSsxUFJZJe67IenRM0AEklc.jpg'),(24,'Series/game-of-thrones','https://image.tmdb.org/t/p/w300_and_h450_bestv2/gwPSoYUHAKmdyVywgLpKKA4BjRr.jpg'),(25,'Series/hannibal','https://image.tmdb.org/t/p/w300_and_h450_bestv2/kArcGHvTb7m1v6RFHdU40jpiGZD.jpg'),(26,'Series/limitless','https://image.tmdb.org/t/p/w300_and_h450_bestv2/ieXc6ryp53DGOv9X3PFhTvTtLau.jpg'),(27,'Series/lucifer','https://image.tmdb.org/t/p/w300_and_h450_bestv2/wmOuYyqVaczBLXxQNFSaRfAUgPz.jpg'),(28,'Series/mr.-robot','https://image.tmdb.org/t/p/w300_and_h450_bestv2/qE0t9rlClIReax0d5tr3j300wUt.jpg'),(29,'Series/ncis','https://image.tmdb.org/t/p/w300_and_h450_bestv2/1ubAPydzsb9VzhqeUGGDA7DZCUy.jpg'),(30,'Series/new-girl','https://image.tmdb.org/t/p/w300_and_h450_bestv2/tvqi2y2FjiWTL5TRL1Pa6UX8Jef.jpg'),(31,'Series/rick-and-morty','https://image.tmdb.org/t/p/w300_and_h450_bestv2/qJdfO3ahgAMf2rcmhoqngjBBZW1.jpg'),(32,'Series/south-park','https://image.tmdb.org/t/p/w300_and_h450_bestv2/v9zc0cZpy5aPSfAy6Tgb6I1zWgV.jpg'),(33,'Series/star-trek:-discovery','https://image.tmdb.org/t/p/w300_and_h450_bestv2/ihvG9dCEnVU3gmMUftTkRICNdJf.jpg'),(34,'Series/westworld','https://image.tmdb.org/t/p/w300_and_h450_bestv2/x2WKIbiwhLoWgLFbT2I0Gwq8U1J.jpg'),(35,'Series/your-pretty-face-is-going-to-hell','https://image.tmdb.org/t/p/w300_and_h450_bestv2/xmls5Vzh4tRxLpjeN04xvwJ4l8s.jpg'),(36,'Series/better-call-saul','https://image.tmdb.org/t/p/w300_and_h450_bestv2/s73aqo43YLTkhFGWc3FmCv751iY.jpg'),(37,'Series/breaking-bad','https://image.tmdb.org/t/p/w300_and_h450_bestv2/1yeVJox3rjo2jBKrrihIMj7uoS9.jpg'),(38,'Series/altered-carbon','https://image.tmdb.org/t/p/w300_and_h450_bestv2/pZg2NUDPJA54AmDs1Y1ZLizBrpd.jpg'),(39,'Series/stranger-things','https://image.tmdb.org/t/p/w300_and_h450_bestv2/lXS60geme1LlEob5Wgvj3KilClA.jpg'),(40,'Series/stargate-universe','https://image.tmdb.org/t/p/w300_and_h450_bestv2/rsVo9mMayCXTXP4X266zcKkf4dz.jpg'),(41,'Series/stargate-atlantis','https://image.tmdb.org/t/p/w300_and_h450_bestv2/dcofw4ByyS0o2CZrYjzhgV2Pgd0.jpg'),(42,'Series/brooklyn-nine-nine','https://image.tmdb.org/t/p/w300_and_h450_bestv2/7tTvREykun6WeGJ4hy7sCG9TOeP.jpg'),(43,'Series/the-office','https://image.tmdb.org/t/p/w300_and_h450_bestv2/iP0uuzWOR5uornNpkkZiemVonMi.jpg'),(44,'Series/black-lightning','https://image.tmdb.org/t/p/w300_and_h450_bestv2/95nmr01SjhjCs0Aa3ZWEfSpZ9oX.jpg'),(45,'Series/house-of-cards','https://image.tmdb.org/t/p/w300_and_h450_bestv2/6AquMx9MoJZTaZIR2AokSDATFCt.jpg'),(46,'Series/lost','https://image.tmdb.org/t/p/w300_and_h450_bestv2/jyGspygDXJMydTOJj7iWNx9Elyd.jpg'),(47,'Series/suits','https://image.tmdb.org/t/p/w300_and_h450_bestv2/i6Iu6pTzfL6iRWhXuYkNs8cPdJF.jpg'),(48,'Series/parks-and-recreation','https://image.tmdb.org/t/p/w300_and_h450_bestv2/9kWSJ9c8NlBY2WgfvOZZ71kafSx.jpg');
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
  `StatusCode` int(11) DEFAULT NULL,
  `RequestUri` varchar(255) DEFAULT NULL,
  `Caller` varchar(255) DEFAULT NULL,
  `Message` varchar(1000) DEFAULT NULL,
  `Time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=453 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Log`
--

LOCK TABLES `Log` WRITE;
/*!40000 ALTER TABLE `Log` DISABLE KEYS */;
INSERT INTO `Log` VALUES (3,1,200,'/provider/','127.0.0.1:56033','Authorized Call','2018-03-09 13:43:11'),(4,1,200,'/provider/','127.0.0.1:56905','Authorized Call','2018-03-09 13:46:26'),(5,1,200,'/provider/','127.0.0.1:56905','Authorized Call','2018-03-09 13:46:29'),(6,1,200,'/provider/','127.0.0.1:56905','Authorized Call','2018-03-09 13:46:29'),(7,1,200,'/provider/','127.0.0.1:56905','Authorized Call','2018-03-09 13:46:30'),(8,1,200,'/provider/','127.0.0.1:56905','Authorized Call','2018-03-09 13:46:30'),(9,1,200,'/provider/','127.0.0.1:56905','Authorized Call','2018-03-09 13:46:30'),(10,1,200,'/provider/','127.0.0.1:56905','Authorized Call','2018-03-09 13:46:30'),(11,1,200,'/provider/','127.0.0.1:56905','Authorized Call','2018-03-09 13:46:30'),(12,1,200,'/robots.txt','127.0.0.1:56977','Not Found','2018-03-09 13:46:39'),(13,1,200,'/sdasd','127.0.0.1:56905','Not Found','2018-03-09 13:46:56'),(14,1,404,'/sdasd','127.0.0.1:57447','Not Found','2018-03-09 13:48:46'),(15,1,404,'/sdasd','127.0.0.1:57447','Not Found','2018-03-09 13:48:46'),(16,1,404,'/sdasd','127.0.0.1:57447','Not Found','2018-03-09 13:48:47'),(17,1,404,'/sdasd','127.0.0.1:57447','Not Found','2018-03-09 13:48:47'),(18,1,404,'/sdasd','127.0.0.1:57447','Not Found','2018-03-09 13:48:47'),(19,1,404,'/sdasd','127.0.0.1:57447','Not Found','2018-03-09 13:48:47'),(20,1,404,'/sdasd','127.0.0.1:57447','Not Found','2018-03-09 13:48:48'),(21,1,404,'/test/','127.0.0.1:57447','Not Found','2018-03-09 13:48:50'),(22,1,404,'/log/','127.0.0.1:57447','Not Found','2018-03-09 13:48:54'),(23,1,200,'/logs/','127.0.0.1:57447','Authorized Call','2018-03-09 13:48:58'),(24,1,200,'/logs/','127.0.0.1:57447','Authorized Call','2018-03-09 13:49:07'),(25,1,200,'/logs/','127.0.0.1:57447','Authorized Call','2018-03-09 13:49:07'),(26,1,200,'/logs/','127.0.0.1:57447','Authorized Call','2018-03-09 13:49:07'),(27,1,200,'/logs/','127.0.0.1:57447','Authorized Call','2018-03-09 13:49:08'),(28,1,200,'/logs/','127.0.0.1:57447','Authorized Call','2018-03-09 13:49:08'),(29,1,200,'/logs/','127.0.0.1:57447','Authorized Call','2018-03-09 13:49:08'),(30,1,200,'/logs/','127.0.0.1:57447','Authorized Call','2018-03-09 13:49:25'),(31,1,200,'/logs/','127.0.0.1:57447','Authorized Call','2018-03-09 13:49:27'),(32,1,200,'/logs/','127.0.0.1:57447','Authorized Call','2018-03-09 13:49:28'),(33,1,200,'/logs/','127.0.0.1:57447','Authorized Call','2018-03-09 13:49:28'),(34,1,200,'/logs/','127.0.0.1:57447','Authorized Call','2018-03-09 13:49:29'),(35,1,200,'/logs/','127.0.0.1:57447','Authorized Call','2018-03-09 13:49:29'),(36,1,404,'/logs/adasd','127.0.0.1:57447','Not Found','2018-03-09 13:49:31'),(37,1,404,'/logs/adasd','127.0.0.1:57447','Not Found','2018-03-09 13:49:33'),(38,2,404,'/logs/adasd','127.0.0.1:58052','Not Found','2018-03-09 13:51:01'),(39,2,404,'/logs/adasd','127.0.0.1:58052','Not Found','2018-03-09 13:51:26'),(40,2,404,'/robots.txt','127.0.0.1:58162','Not Found','2018-03-09 13:51:26'),(41,2,404,'/logs/adasd','127.0.0.1:58052','Not Found','2018-03-09 13:51:28'),(42,1,200,'/logs/','127.0.0.1:58052','Authorized Call','2018-03-09 13:51:33'),(43,1,200,'/logs/','127.0.0.1:52681','Authorized Call','2018-03-09 15:45:11'),(44,1,200,'/logs/','127.0.0.1:54861','Authorized Call','2018-03-09 15:54:47'),(45,2,404,'/logs/1','127.0.0.1:54861','Not Found','2018-03-09 15:54:50'),(46,2,404,'/logs/12','127.0.0.1:54861','Not Found','2018-03-09 15:54:52'),(47,2,404,'/robots.txt','127.0.0.1:54891','Not Found','2018-03-09 15:54:52'),(48,2,404,'/logs/12+','127.0.0.1:54861','Not Found','2018-03-09 15:55:07'),(49,2,404,'/logs/12+','127.0.0.1:55063','Not Found','2018-03-09 15:55:33'),(50,1,200,'/logs/12','127.0.0.1:55063','Authorized Call','2018-03-09 15:55:35'),(51,1,200,'/logs/12','127.0.0.1:55063','Authorized Call','2018-03-09 15:55:36'),(52,1,200,'/logs/12','127.0.0.1:55063','Authorized Call','2018-03-09 15:55:37'),(53,1,200,'/logs/12','127.0.0.1:55063','Authorized Call','2018-03-09 15:55:38'),(54,1,200,'/logs/','127.0.0.1:55063','Authorized Call','2018-03-09 15:55:40'),(55,1,200,'/logs/5','127.0.0.1:55063','Authorized Call','2018-03-09 15:55:41'),(56,1,200,'/logs/5','127.0.0.1:55063','Authorized Call','2018-03-09 15:55:41'),(57,1,200,'/logs/5','127.0.0.1:55063','Authorized Call','2018-03-09 15:55:42'),(58,1,200,'/logs/5','127.0.0.1:55063','Authorized Call','2018-03-09 15:55:42'),(59,1,200,'/logs/5','127.0.0.1:55063','Authorized Call','2018-03-09 15:55:42'),(60,1,200,'/logs/5','127.0.0.1:55063','Authorized Call','2018-03-09 15:55:42'),(61,1,200,'/logs/5','127.0.0.1:55063','Authorized Call','2018-03-09 15:55:43'),(62,1,200,'/logs/5','127.0.0.1:55063','Authorized Call','2018-03-09 15:55:44'),(63,1,200,'/logs/5','127.0.0.1:55063','Authorized Call','2018-03-09 15:55:45'),(64,1,200,'/logs/5','127.0.0.1:55063','Authorized Call','2018-03-09 15:55:45'),(65,1,200,'/logs/5','127.0.0.1:55063','Authorized Call','2018-03-09 15:55:47'),(66,1,200,'/logs/5','127.0.0.1:55063','Authorized Call','2018-03-09 15:55:48'),(67,1,200,'/logs/5','127.0.0.1:55063','Authorized Call','2018-03-09 15:55:49'),(68,1,200,'/logs/5','127.0.0.1:55063','Authorized Call','2018-03-09 15:55:50'),(69,1,200,'/logs/5','127.0.0.1:55063','Authorized Call','2018-03-09 15:55:51'),(70,1,200,'/logs/5','127.0.0.1:55063','Authorized Call','2018-03-09 15:55:51'),(71,1,200,'/logs/5','127.0.0.1:55063','Authorized Call','2018-03-09 15:55:52'),(72,1,200,'/logs/5','127.0.0.1:55063','Authorized Call','2018-03-09 15:55:52'),(73,1,200,'/logs/5','127.0.0.1:55063','Authorized Call','2018-03-09 15:55:52'),(74,2,404,'/robots.txt','127.0.0.1:55364','Not Found','2018-03-09 16:00:52'),(75,1,200,'/logs/','127.0.0.1:55362','Authorized Call','2018-03-09 16:00:56'),(76,1,200,'/logs/0','127.0.0.1:55362','Authorized Call','2018-03-09 16:00:58'),(77,1,200,'/logs/1','127.0.0.1:55362','Authorized Call','2018-03-09 16:01:01'),(78,1,200,'/logs/2','127.0.0.1:55362','Authorized Call','2018-03-09 16:01:03'),(79,1,200,'/logs/1','127.0.0.1:55362','Authorized Call','2018-03-09 16:01:06'),(80,1,200,'/logs/1','127.0.0.1:55362','Authorized Call','2018-03-09 16:01:08'),(81,1,200,'/logs/1','127.0.0.1:55362','Authorized Call','2018-03-09 16:01:10'),(82,1,200,'/logs/1','127.0.0.1:55362','Authorized Call','2018-03-09 16:01:15'),(83,1,200,'/logs/1','127.0.0.1:55362','Authorized Call','2018-03-09 16:01:16'),(84,1,200,'/logs/1','127.0.0.1:55362','Authorized Call','2018-03-09 16:01:17'),(85,1,200,'/logs/1','127.0.0.1:55362','Authorized Call','2018-03-09 16:01:17'),(86,1,200,'/logs/1','127.0.0.1:55362','Authorized Call','2018-03-09 16:01:17'),(87,1,200,'/logs/1','127.0.0.1:55362','Authorized Call','2018-03-09 16:01:17'),(88,1,200,'/logs/1','127.0.0.1:55362','Authorized Call','2018-03-09 16:01:18'),(89,1,200,'/logs/1','127.0.0.1:55362','Authorized Call','2018-03-09 16:01:18'),(90,1,200,'/logs/0','127.0.0.1:55362','Authorized Call','2018-03-09 16:01:20'),(91,2,404,'/robots.txt','127.0.0.1:55364','Not Found','2018-03-09 16:01:20'),(92,1,200,'/logs/1','127.0.0.1:55362','Authorized Call','2018-03-09 16:01:23'),(93,1,200,'/logs/1','127.0.0.1:55362','Authorized Call','2018-03-09 16:01:35'),(94,1,200,'/logs/10','127.0.0.1:55362','Authorized Call','2018-03-09 16:01:40'),(95,1,200,'/logs/10','127.0.0.1:55362','Authorized Call','2018-03-09 16:01:49'),(96,1,200,'/logs/10','127.0.0.1:55362','Authorized Call','2018-03-09 16:01:54'),(97,1,200,'/logs/10','127.0.0.1:55362','Authorized Call','2018-03-09 16:01:55'),(98,1,200,'/logs/10','127.0.0.1:55362','Authorized Call','2018-03-09 16:01:55'),(99,2,404,'/asd','127.0.0.1:55476','Not Found','2018-03-09 16:08:22'),(100,2,404,'/asd','127.0.0.1:56162','Not Found','2018-03-10 09:03:26'),(101,2,404,'/robots.txt','127.0.0.1:56166','Not Found','2018-03-10 09:03:26'),(102,2,404,'/favicon.ico','127.0.0.1:56162','Not Found','2018-03-10 09:03:26'),(103,1,200,'/logs/','127.0.0.1:56162','Authorized Call','2018-03-10 09:03:39'),(104,1,200,'/logs/','127.0.0.1:56171','Authorized Call','2018-03-10 09:03:39'),(105,1,200,'/logs/6400','127.0.0.1:56192','Authorized Call','2018-03-10 09:08:23'),(106,1,200,'/logs/6400','127.0.0.1:56462','Authorized Call','2018-03-10 10:06:28'),(107,2,404,'/favicon.ico','127.0.0.1:56462','Not Found','2018-03-10 10:06:29'),(108,1,200,'/logs/6400','127.0.0.1:56462','Authorized Call','2018-03-10 10:06:29'),(109,2,404,'/favicon.ico','127.0.0.1:56462','Not Found','2018-03-10 10:06:29'),(110,2,404,'/logs/warnings','127.0.0.1:56462','Not Found','2018-03-10 10:06:33'),(111,2,404,'/logs/warnings','127.0.0.1:56462','Not Found','2018-03-10 10:06:37'),(112,1,200,'/logs/warning/','127.0.0.1:56462','Authorized Call','2018-03-10 10:06:39'),(113,1,200,'/logs/warning/1','127.0.0.1:56462','Authorized Call','2018-03-10 10:06:46'),(114,1,200,'/logs/warning/300','127.0.0.1:56462','Authorized Call','2018-03-10 10:06:54'),(115,2,404,'/robots.txt','127.0.0.1:56487','Not Found','2018-03-10 10:06:55'),(116,1,200,'/logs/warning/300','127.0.0.1:56462','Authorized Call','2018-03-10 10:06:55'),(117,2,404,'/favicon.ico','127.0.0.1:56462','Not Found','2018-03-10 10:06:55'),(118,1,200,'/logs/warning/300','127.0.0.1:56462','Authorized Call','2018-03-10 10:06:56'),(119,2,404,'/favicon.ico','127.0.0.1:56462','Not Found','2018-03-10 10:06:56'),(120,1,200,'/logs/warning/300','127.0.0.1:56462','Authorized Call','2018-03-10 10:06:56'),(121,2,404,'/favicon.ico','127.0.0.1:56462','Not Found','2018-03-10 10:06:56'),(122,1,200,'/logs/warning/300','127.0.0.1:56462','Authorized Call','2018-03-10 10:06:56'),(123,1,200,'/logs/warning/300','127.0.0.1:56462','Authorized Call','2018-03-10 10:06:56'),(124,2,404,'/favicon.ico','127.0.0.1:56462','Not Found','2018-03-10 10:06:57'),(125,2,404,'/logs/errors/300','127.0.0.1:56462','Not Found','2018-03-10 10:06:59'),(126,1,200,'/logs/error/300','127.0.0.1:56462','Authorized Call','2018-03-10 10:07:03'),(127,1,200,'/logs/error/','127.0.0.1:56462','Authorized Call','2018-03-10 10:07:04'),(128,1,200,'/logs/error/','127.0.0.1:56462','Authorized Call','2018-03-10 10:07:06'),(129,1,200,'/logs/message/','127.0.0.1:56462','Authorized Call','2018-03-10 10:07:11'),(130,1,200,'/logs/warning/300','127.0.0.1:56462','Authorized Call','2018-03-10 10:07:38'),(131,1,200,'/logs/warning/300','127.0.0.1:56462','Authorized Call','2018-03-10 10:07:48'),(132,1,200,'/logs/warning/300','127.0.0.1:56555','Authorized Call','2018-03-10 10:29:30'),(133,2,404,'/robots.txt','127.0.0.1:62521','Not Found','2018-03-10 11:26:06'),(134,1,200,'/logs/warning/300','127.0.0.1:62518','Authorized Call','2018-03-10 11:26:06'),(135,2,404,'/favicon.ico','127.0.0.1:62518','Not Found','2018-03-10 11:26:06'),(136,2,404,'/favicon.ico','127.0.0.1:62518','Not Found','2018-03-10 11:26:10'),(137,2,404,'/favicon.ico','127.0.0.1:62536','Not Found','2018-03-10 11:28:01'),(138,2,404,'/favicon.ico','127.0.0.1:62544','Not Found','2018-03-10 11:28:31'),(139,2,404,'/robots.txt','127.0.0.1:62849','Not Found','2018-03-10 14:10:43'),(140,2,404,'/favicon.ico','127.0.0.1:62846','Not Found','2018-03-10 14:10:43'),(141,2,404,'/favicon.ico','127.0.0.1:62846','Not Found','2018-03-10 14:10:45'),(142,2,404,'/favicon.ico','127.0.0.1:62846','Not Found','2018-03-10 14:10:45'),(143,1,200,'/logs/warning/300','127.0.0.1:62934','Authorized Call','2018-03-10 15:13:12'),(144,2,404,'/favicon.ico','127.0.0.1:62934','Not Found','2018-03-10 15:13:12'),(145,1,200,'/logs/warning/300','127.0.0.1:62934','Authorized Call','2018-03-10 15:13:12'),(146,2,404,'/favicon.ico','127.0.0.1:62934','Not Found','2018-03-10 15:13:13'),(147,1,200,'/logs/warning/300','127.0.0.1:62934','Authorized Call','2018-03-10 15:13:15'),(148,2,404,'/favicon.ico','127.0.0.1:62934','Not Found','2018-03-10 15:13:15'),(149,2,404,'/robots.txt','127.0.0.1:63684','Not Found','2018-03-11 17:35:13'),(150,2,404,'/favicon.ico','127.0.0.1:63683','Not Found','2018-03-11 17:35:14'),(151,1,200,'/provider/','127.0.0.1:63683','Authorized Call','2018-03-11 17:35:17'),(152,1,200,'/provider/netflix','127.0.0.1:63683','Authorized Call','2018-03-11 17:35:32'),(153,1,200,'/logs/','127.0.0.1:64691','Authorized Call','2018-03-11 17:40:33'),(154,1,200,'/logs/10','127.0.0.1:64691','Authorized Call','2018-03-11 17:40:38'),(155,1,200,'/logs/10','127.0.0.1:64691','Authorized Call','2018-03-11 17:40:39'),(156,1,200,'/logs/10','127.0.0.1:64691','Authorized Call','2018-03-11 17:40:40'),(157,1,200,'/logs/10','127.0.0.1:64691','Authorized Call','2018-03-11 17:40:40'),(158,1,200,'/logs/10','127.0.0.1:64691','Authorized Call','2018-03-11 17:40:41'),(159,2,404,'/robots.txt','127.0.0.1:64737','Not Found','2018-03-11 17:40:41'),(160,1,200,'/logs/10','127.0.0.1:64691','Authorized Call','2018-03-11 17:40:41'),(161,1,200,'/logs/10','127.0.0.1:64691','Authorized Call','2018-03-11 17:40:41'),(162,1,200,'/logs/10','127.0.0.1:64691','Authorized Call','2018-03-11 17:40:41'),(163,1,200,'/logs/10','127.0.0.1:64691','Authorized Call','2018-03-11 17:40:42'),(164,1,200,'/logs/10','127.0.0.1:64691','Authorized Call','2018-03-11 17:40:42'),(165,1,200,'/logs/10','127.0.0.1:64691','Authorized Call','2018-03-11 17:40:42'),(166,1,200,'/logs/10','127.0.0.1:64691','Authorized Call','2018-03-11 17:40:42'),(167,1,200,'/logs/10','127.0.0.1:64691','Authorized Call','2018-03-11 17:40:42'),(168,1,200,'/logs/warning/','127.0.0.1:64691','Authorized Call','2018-03-11 17:40:49'),(169,2,404,'/logs/info/','127.0.0.1:64691','Not Found','2018-03-11 17:40:53'),(170,2,404,'/logs/messages/','127.0.0.1:64691','Not Found','2018-03-11 17:40:57'),(171,1,200,'/logs/message/','127.0.0.1:64691','Authorized Call','2018-03-11 17:41:02'),(172,2,404,'/logs/messages/','127.0.0.1:64691','Not Found','2018-03-11 17:41:10'),(173,1,200,'/provider/','127.0.0.1:64691','Authorized Call','2018-03-11 17:42:04'),(174,1,200,'/provider/Burning%20Series','127.0.0.1:64691','Authorized Call','2018-03-11 17:42:34'),(175,2,404,'/provider/Burning%20Series/imge','127.0.0.1:64691','Not Found','2018-03-11 17:42:53'),(176,1,200,'/provider/','127.0.0.1:64691','Authorized Call','2018-03-11 17:43:02'),(177,1,200,'/provider/','127.0.0.1:50587','Authorized Call','2018-03-11 18:08:31'),(178,2,404,'/robots.txt','127.0.0.1:50590','Not Found','2018-03-11 18:08:31'),(179,2,404,'/favicon.ico','127.0.0.1:50587','Not Found','2018-03-11 18:08:31'),(180,1,200,'/provider/netflix','127.0.0.1:50587','Authorized Call','2018-03-11 18:08:35'),(181,2,404,'/provider/primevideo/image','127.0.0.1:50587','Not Found','2018-03-11 18:08:47'),(182,1,200,'/provider/netflix','127.0.0.1:50587','Authorized Call','2018-03-11 18:09:29'),(183,1,200,'/provider/netflix/image','127.0.0.1:50629','Authorized Call','2018-03-11 18:13:29'),(184,2,404,'/favicon.ico','127.0.0.1:50629','Not Found','2018-03-11 18:13:29'),(185,1,200,'/provider/netflix/image','127.0.0.1:50629','Authorized Call','2018-03-11 18:13:29'),(186,2,404,'/favicon.ico','127.0.0.1:50629','Not Found','2018-03-11 18:13:30'),(187,1,200,'/provider/netflix/image','127.0.0.1:50629','Authorized Call','2018-03-11 18:13:30'),(188,1,200,'/provider/netflix/image','127.0.0.1:50629','Authorized Call','2018-03-11 18:13:30'),(189,2,404,'/favicon.ico','127.0.0.1:50630','Not Found','2018-03-11 18:13:30'),(190,1,200,'/provider/netflix/image','127.0.0.1:50630','Authorized Call','2018-03-11 18:13:31'),(191,2,404,'/favicon.ico','127.0.0.1:50630','Not Found','2018-03-11 18:13:31'),(192,1,200,'/provider/burning%20series/image','127.0.0.1:50630','Authorized Call','2018-03-11 18:13:32'),(193,2,404,'/robots.txt','127.0.0.1:50655','Not Found','2018-03-11 18:13:32'),(194,2,404,'/favicon.ico','127.0.0.1:50630','Not Found','2018-03-11 18:13:33'),(195,1,200,'/provider/','127.0.0.1:50630','Authorized Call','2018-03-11 18:13:33'),(196,2,404,'/favicon.ico','127.0.0.1:50630','Not Found','2018-03-11 18:13:33'),(197,1,200,'/provider/burning%20series/image','127.0.0.1:50630','Authorized Call','2018-03-11 18:13:35'),(198,2,404,'/favicon.ico','127.0.0.1:50630','Not Found','2018-03-11 18:13:35'),(199,1,200,'/provider/burning%20series/image','127.0.0.1:50630','Authorized Call','2018-03-11 18:13:39'),(200,1,200,'/provider/','127.0.0.1:51260','Authorized Call','2018-03-11 18:23:10'),(201,1,200,'/provider/netflix','127.0.0.1:51260','Authorized Call','2018-03-11 18:23:19'),(202,1,200,'/provider/netflix','127.0.0.1:51260','Authorized Call','2018-03-11 18:23:33'),(203,1,200,'/provider/netflix/image','127.0.0.1:51260','Authorized Call','2018-03-11 18:23:52'),(204,1,200,'/provider/netflix/image','127.0.0.1:51260','Authorized Call','2018-03-11 18:23:53'),(205,1,200,'/provider/netflix/image','127.0.0.1:51642','Authorized Call','2018-03-11 18:26:01'),(206,1,200,'/provider/netflix/image','127.0.0.1:51786','Authorized Call','2018-03-11 18:26:36'),(207,1,200,'/provider/netflix/image','127.0.0.1:52906','Authorized Call','2018-03-11 18:30:40'),(208,2,404,'/robots.txt','127.0.0.1:53420','Not Found','2018-03-12 07:24:18'),(209,1,200,'/logs/','127.0.0.1:53419','Authorized Call','2018-03-12 07:24:21'),(210,1,200,'/logs/10','127.0.0.1:53419','Authorized Call','2018-03-12 07:24:29'),(211,2,404,'/logs/warnings','127.0.0.1:53447','Not Found','2018-03-12 07:25:47'),(212,1,200,'/logs/warning/','127.0.0.1:53447','Authorized Call','2018-03-12 07:25:50'),(213,2,404,'/errors/warning/','127.0.0.1:53487','Not Found','2018-03-12 07:27:55'),(214,2,404,'/error/warning/','127.0.0.1:53487','Not Found','2018-03-12 07:27:58'),(215,2,404,'/error/warning/','127.0.0.1:53527','Not Found','2018-03-12 07:31:03'),(216,2,404,'/robots.txt','127.0.0.1:53529','Not Found','2018-03-12 07:31:03'),(217,2,404,'/favicon.ico','127.0.0.1:53527','Not Found','2018-03-12 07:31:03'),(218,1,200,'/provider/','127.0.0.1:53561','Authorized Call','2018-03-12 07:32:44'),(219,1,200,'/provider/netflix','127.0.0.1:53561','Authorized Call','2018-03-12 07:33:24'),(220,3,500,'/provider/netflix','127.0.0.1:54021','Internal Server Error','2018-03-12 07:57:28'),(221,2,404,'/favicon.ico','127.0.0.1:54021','Not Found','2018-03-12 07:57:29'),(222,1,200,'/provider/netflix','127.0.0.1:54048','Authorized Call','2018-03-12 07:59:30'),(223,2,404,'/favicon.ico','127.0.0.1:54048','Not Found','2018-03-12 07:59:30'),(224,1,200,'/provider/netflix/image','127.0.0.1:54048','Authorized Call','2018-03-12 07:59:34'),(225,1,200,'/provider/netflix/image','127.0.0.1:54048','Authorized Call','2018-03-12 08:00:14'),(226,2,404,'/favicon.ico','127.0.0.1:54048','Not Found','2018-03-12 08:00:14'),(227,1,200,'/provider/netflix/image','127.0.0.1:54048','Authorized Call','2018-03-12 08:00:15'),(228,2,404,'/favicon.ico','127.0.0.1:54048','Not Found','2018-03-12 08:00:15'),(229,1,200,'/provider/netflix/image','127.0.0.1:54048','Authorized Call','2018-03-12 08:00:16'),(230,2,404,'/favicon.ico','127.0.0.1:54048','Not Found','2018-03-12 08:00:16'),(231,1,200,'/provider/netflix/image','127.0.0.1:54048','Authorized Call','2018-03-12 08:00:16'),(232,2,404,'/favicon.ico','127.0.0.1:54048','Not Found','2018-03-12 08:00:16'),(233,1,200,'/provider/netflix/image','127.0.0.1:54176','Authorized Call','2018-03-12 08:06:11'),(234,2,404,'/favicon.ico','127.0.0.1:54176','Not Found','2018-03-12 08:06:12'),(235,1,200,'/provider/netflix/image','127.0.0.1:54176','Authorized Call','2018-03-12 08:06:12'),(236,2,404,'/favicon.ico','127.0.0.1:54176','Not Found','2018-03-12 08:06:13'),(237,2,404,'/provider/netflix/image','127.0.0.1:54234','Not Found','2018-03-12 08:10:48'),(238,2,404,'/favicon.ico','127.0.0.1:54234','Not Found','2018-03-12 08:10:48'),(239,2,404,'/robots.txt','127.0.0.1:54256','Not Found','2018-03-12 08:11:15'),(240,2,404,'/provider/netflix/image','127.0.0.1:54254','Not Found','2018-03-12 08:11:15'),(241,2,404,'/favicon.ico','127.0.0.1:54254','Not Found','2018-03-12 08:11:15'),(242,2,404,'/provider/netflix/image','127.0.0.1:54254','Not Found','2018-03-12 08:11:16'),(243,2,404,'/favicon.ico','127.0.0.1:54254','Not Found','2018-03-12 08:11:16'),(244,2,404,'/provider/netflix/image','127.0.0.1:54254','Not Found','2018-03-12 08:11:16'),(245,2,404,'/favicon.ico','127.0.0.1:54254','Not Found','2018-03-12 08:11:17'),(246,2,404,'/provider/netflix/image','127.0.0.1:54254','Not Found','2018-03-12 08:11:17'),(247,2,404,'/favicon.ico','127.0.0.1:54254','Not Found','2018-03-12 08:11:17'),(248,2,404,'/provider/netflix/image','127.0.0.1:54276','Not Found','2018-03-12 08:11:23'),(249,2,404,'/favicon.ico','127.0.0.1:54284','Not Found','2018-03-12 08:11:23'),(250,1,200,'/provider/netflix/image','127.0.0.1:54386','Authorized Call','2018-03-12 08:17:02'),(251,2,404,'/favicon.ico','127.0.0.1:54386','Not Found','2018-03-12 08:17:03'),(252,1,200,'/provider/netflix/image','127.0.0.1:54386','Authorized Call','2018-03-12 08:17:10'),(253,2,404,'/favicon.ico','127.0.0.1:54386','Not Found','2018-03-12 08:17:10'),(254,1,200,'/provider/burning%20series/image','127.0.0.1:54386','Authorized Call','2018-03-12 08:17:28'),(255,2,404,'/provider/prime%20video/image','127.0.0.1:54386','Not Found','2018-03-12 08:17:38'),(256,1,200,'/provider/amazon%20prime/image','127.0.0.1:54386','Authorized Call','2018-03-12 08:17:54'),(257,2,404,'/robots.txt','127.0.0.1:57328','Not Found','2018-03-12 10:46:44'),(258,1,200,'/logs/','127.0.0.1:57327','Authorized Call','2018-03-12 10:46:44'),(259,1,200,'/logs/?asd=1','127.0.0.1:57327','Authorized Call','2018-03-12 10:46:52'),(260,1,200,'/logs/?endDate=0','127.0.0.1:57327','Authorized Call','2018-03-12 10:47:00'),(261,1,200,'/logs/?endDate=12','127.0.0.1:57327','Authorized Call','2018-03-12 10:47:03'),(262,1,200,'/logs/?endDate=asac','127.0.0.1:57327','Authorized Call','2018-03-12 10:47:08'),(263,1,200,'/logs/?endDate=sad4','127.0.0.1:57327','Authorized Call','2018-03-12 10:47:12'),(264,1,200,'/logs/?endDate=sad400','127.0.0.1:57327','Authorized Call','2018-03-12 10:47:15'),(265,1,200,'/logs/?endDate=400adsasd','127.0.0.1:57327','Authorized Call','2018-03-12 10:47:19'),(266,1,200,'/logs/?endDate=400adsasd','127.0.0.1:57327','Authorized Call','2018-03-12 10:47:21'),(267,2,404,'/favicon.ico','127.0.0.1:57327','Not Found','2018-03-12 10:47:21'),(268,1,200,'/logs/?endDate=400','127.0.0.1:57327','Authorized Call','2018-03-12 10:47:23'),(269,2,404,'/robots.txt','127.0.0.1:57402','Not Found','2018-03-12 10:49:15'),(270,1,200,'/logs/?endDate=400','127.0.0.1:57397','Authorized Call','2018-03-12 10:49:15'),(271,2,404,'/favicon.ico','127.0.0.1:57397','Not Found','2018-03-12 10:49:15'),(272,1,200,'/logs/?endDate=400','127.0.0.1:57397','Authorized Call','2018-03-12 10:49:16'),(273,2,404,'/favicon.ico','127.0.0.1:57397','Not Found','2018-03-12 10:49:17'),(274,1,200,'/logs/?since=400','127.0.0.1:57397','Authorized Call','2018-03-12 10:49:19'),(275,1,200,'/logs/?since=400','127.0.0.1:57397','Authorized Call','2018-03-12 10:49:20'),(276,1,200,'/logs/warning/?since=400','127.0.0.1:57397','Authorized Call','2018-03-12 10:49:27'),(277,1,200,'/logs/error/?since=400','127.0.0.1:57397','Authorized Call','2018-03-12 10:49:34'),(278,1,200,'/logs/error/?since=40000','127.0.0.1:57397','Authorized Call','2018-03-12 10:49:38'),(279,1,200,'/logs/message/?since=40000','127.0.0.1:57397','Authorized Call','2018-03-12 10:49:44'),(280,1,200,'/logs/warning/?since=40000','127.0.0.1:57397','Authorized Call','2018-03-12 10:49:49'),(281,1,200,'/logs/message/?since=40000','127.0.0.1:57397','Authorized Call','2018-03-12 10:49:55'),(282,2,404,'/robots.txt','127.0.0.1:60337','Not Found','2018-03-12 14:10:42'),(283,3,500,'/series/','127.0.0.1:60335','Internal Server Error','2018-03-12 14:10:42'),(284,3,500,'/provider/','127.0.0.1:60335','Internal Server Error','2018-03-12 14:11:21'),(285,1,200,'/provider/','127.0.0.1:60357','Authorized Call','2018-03-12 14:11:43'),(286,2,404,'/favicon.ico','127.0.0.1:60357','Not Found','2018-03-12 14:11:43'),(287,1,200,'/provider/','127.0.0.1:60357','Authorized Call','2018-03-12 14:12:01'),(288,2,404,'/favicon.ico','127.0.0.1:60357','Not Found','2018-03-12 14:12:01'),(289,1,200,'/series/','127.0.0.1:60357','Authorized Call','2018-03-12 14:12:04'),(290,1,200,'/series/','127.0.0.1:60357','Authorized Call','2018-03-12 14:12:05'),(291,2,404,'/favicon.ico','127.0.0.1:60357','Not Found','2018-03-12 14:12:05'),(292,2,404,'/robots.txt','127.0.0.1:60670','Not Found','2018-03-12 14:37:45'),(293,1,200,'/series/','127.0.0.1:60668','Authorized Call','2018-03-12 14:37:45'),(294,2,404,'/favicon.ico','127.0.0.1:60668','Not Found','2018-03-12 14:37:45'),(295,3,500,'/series/Test/','127.0.0.1:60668','Internal Server Error','2018-03-12 14:37:51'),(296,3,500,'/series/Test/','127.0.0.1:60684','Internal Server Error','2018-03-12 14:38:06'),(297,2,404,'/favicon.ico','127.0.0.1:60684','Not Found','2018-03-12 14:38:06'),(298,3,500,'/series/Test/','127.0.0.1:60684','Internal Server Error','2018-03-12 14:38:07'),(299,2,404,'/favicon.ico','127.0.0.1:60684','Not Found','2018-03-12 14:38:07'),(300,1,200,'/series/Test/','127.0.0.1:60707','Authorized Call','2018-03-12 14:38:53'),(301,2,404,'/favicon.ico','127.0.0.1:60707','Not Found','2018-03-12 14:38:53'),(302,1,200,'/series/','127.0.0.1:60707','Authorized Call','2018-03-12 14:38:57'),(303,1,200,'/series/Test/','127.0.0.1:60784','Authorized Call','2018-03-12 14:43:14'),(304,2,404,'/favicon.ico','127.0.0.1:60784','Not Found','2018-03-12 14:43:14'),(305,1,200,'/series/Test/','127.0.0.1:60784','Authorized Call','2018-03-12 14:43:14'),(306,2,404,'/favicon.ico','127.0.0.1:60784','Not Found','2018-03-12 14:43:15'),(307,1,200,'/series/','127.0.0.1:60812','Authorized Call','2018-03-12 14:44:23'),(308,1,200,'/series/Test/','127.0.0.1:61041','Authorized Call','2018-03-12 14:56:30'),(309,1,200,'/series/Test/image','127.0.0.1:61041','Authorized Call','2018-03-12 14:56:33'),(310,2,404,'/robots.txt','127.0.0.1:51590','Not Found','2018-03-12 15:45:59'),(311,1,200,'/series/','127.0.0.1:51588','Authorized Call','2018-03-12 15:45:59'),(312,2,404,'/series/testseries/','127.0.0.1:52048','Not Found','2018-03-12 15:47:46'),(313,2,404,'/series/testseries/','127.0.0.1:52638','Not Found','2018-03-12 15:50:30'),(314,2,404,'/series/testseries/','127.0.0.1:52673','Not Found','2018-03-12 15:50:37'),(315,1,200,'/series/','127.0.0.1:52698','Authorized Call','2018-03-12 15:50:42'),(316,2,405,'/series','127.0.0.1:52698','Method Not Allowed','2018-03-12 15:50:47'),(317,2,404,'/series/a/','127.0.0.1:52760','Not Found','2018-03-12 15:50:54'),(318,2,400,'/series/asdasdasd/','127.0.0.1:52760','Bad Request','2018-03-12 15:51:01'),(319,2,400,'/series/asdasdasd/','127.0.0.1:52792','Bad Request','2018-03-12 15:51:01'),(320,2,400,'/series/asdasdasd/','127.0.0.1:53052','Bad Request','2018-03-12 15:52:34'),(321,2,400,'/series/asdasdasd/','127.0.0.1:53428','Bad Request','2018-03-12 15:54:07'),(322,2,400,'/series/asdasdasd/','127.0.0.1:53964','Bad Request','2018-03-12 15:56:07'),(323,2,400,'/series/asdasdasd/','127.0.0.1:53976','Bad Request','2018-03-12 15:56:09'),(324,2,400,'/series/asdasdasd/','127.0.0.1:54002','Bad Request','2018-03-12 15:56:14'),(325,2,400,'/series/asdasdasd/','127.0.0.1:54139','Bad Request','2018-03-12 15:57:06'),(326,1,200,'/series/asdasdasd/','127.0.0.1:54293','Authorized Call','2018-03-12 15:57:47'),(327,1,200,'/series/asdasdasd/','127.0.0.1:57017','Authorized Call','2018-03-12 16:07:45'),(328,1,200,'/series/asdasdasd/','127.0.0.1:58565','Authorized Call','2018-03-12 16:18:09'),(329,1,200,'/series/asdasdasd/','127.0.0.1:61433','Authorized Call','2018-03-12 16:25:41'),(330,1,200,'/series/asdasdasd/','127.0.0.1:62370','Authorized Call','2018-03-12 16:29:08'),(331,1,200,'/series/asdasdasd/','127.0.0.1:63137','Authorized Call','2018-03-13 09:23:43'),(332,3,500,'/series/asdasdasd/','127.0.0.1:63740','Internal Server Error','2018-03-13 09:50:28'),(333,1,200,'/series/asdasdasd/','127.0.0.1:63740','Authorized Call','2018-03-13 09:50:28'),(334,1,200,'/series/asdasdasd/','127.0.0.1:63760','Authorized Call','2018-03-13 09:52:06'),(335,1,200,'/series/flash/','127.0.0.1:63760','Authorized Call','2018-03-13 09:52:15'),(336,1,200,'/series/flash/','127.0.0.1:64179','Authorized Call','2018-03-13 10:26:03'),(337,1,200,'/series/flash/','127.0.0.1:49276','Authorized Call','2018-03-13 10:42:53'),(338,1,200,'/series/flash/','127.0.0.1:49276','Authorized Call','2018-03-13 10:43:08'),(339,1,200,'/series/flash/','127.0.0.1:49504','Authorized Call','2018-03-13 10:43:46'),(340,1,200,'/series/flash/','127.0.0.1:61213','Authorized Call','2018-03-13 13:12:10'),(341,1,200,'/series/flash/','127.0.0.1:61213','Authorized Call','2018-03-13 13:12:15'),(342,1,200,'/series/flash/','127.0.0.1:61213','Authorized Call','2018-03-13 13:12:16'),(343,1,200,'/series/flash/','127.0.0.1:61529','Authorized Call','2018-03-13 13:13:25'),(344,1,200,'/series/flash/','127.0.0.1:63761','Authorized Call','2018-03-13 13:22:26'),(345,1,200,'/series/flash/','127.0.0.1:63994','Authorized Call','2018-03-13 13:23:07'),(346,1,200,'/series/flash/','127.0.0.1:64666','Authorized Call','2018-03-13 13:25:38'),(347,1,200,'/series/flash/','127.0.0.1:64666','Authorized Call','2018-03-13 13:26:00'),(348,1,200,'/series/flash/','127.0.0.1:64666','Authorized Call','2018-03-13 13:26:07'),(349,1,200,'/series/flash/','127.0.0.1:64976','Authorized Call','2018-03-13 13:26:56'),(350,3,500,'/series/flash/','127.0.0.1:65249','Internal Server Error','2018-03-13 13:28:03'),(351,1,200,'/series/flash/','127.0.0.1:62441','Authorized Call','2018-03-13 14:34:22'),(352,2,404,'/robots.txt','127.0.0.1:62788','Not Found','2018-03-13 14:36:02'),(353,1,200,'/series/the%20flash/','127.0.0.1:62786','Authorized Call','2018-03-13 14:36:02'),(354,1,200,'/series/the%20flash/image','127.0.0.1:62786','Authorized Call','2018-03-13 14:36:06'),(355,1,200,'/series/flash/','127.0.0.1:62962','Authorized Call','2018-03-13 14:36:41'),(356,2,404,'/robots.txt','127.0.0.1:63109','Not Found','2018-03-13 14:38:43'),(357,1,200,'/series/','127.0.0.1:63108','Authorized Call','2018-03-13 14:38:43'),(358,1,200,'/provider/','127.0.0.1:63122','Authorized Call','2018-03-13 14:39:45'),(359,2,400,'/series/flash/','127.0.0.1:63451','Bad Request','2018-03-13 14:59:07'),(360,2,400,'/series/flash/','127.0.0.1:63451','Bad Request','2018-03-13 14:59:14'),(361,2,400,'/series/flash/','127.0.0.1:63451','Bad Request','2018-03-13 14:59:16'),(362,2,400,'/series/flash/','127.0.0.1:63451','Bad Request','2018-03-13 14:59:19'),(363,2,404,'/robots.txt','127.0.0.1:63543','Not Found','2018-03-13 15:03:16'),(364,1,200,'/series/','127.0.0.1:63539','Authorized Call','2018-03-13 15:03:16'),(365,2,404,'/favicon.ico','127.0.0.1:63539','Not Found','2018-03-13 15:03:16'),(366,1,200,'/series/','127.0.0.1:63539','Authorized Call','2018-03-13 15:03:16'),(367,2,404,'/favicon.ico','127.0.0.1:63539','Not Found','2018-03-13 15:03:17'),(368,1,200,'/series/','127.0.0.1:63539','Authorized Call','2018-03-13 15:03:17'),(369,2,404,'/favicon.ico','127.0.0.1:63539','Not Found','2018-03-13 15:03:17'),(370,1,200,'/series/','127.0.0.1:63539','Authorized Call','2018-03-13 15:03:17'),(371,2,404,'/favicon.ico','127.0.0.1:63539','Not Found','2018-03-13 15:03:17'),(372,1,200,'/series/','127.0.0.1:63539','Authorized Call','2018-03-13 15:03:17'),(373,1,200,'/series/','127.0.0.1:63539','Authorized Call','2018-03-13 15:03:18'),(374,1,200,'/series/','127.0.0.1:63539','Authorized Call','2018-03-13 15:03:18'),(375,2,404,'/favicon.ico','127.0.0.1:63539','Not Found','2018-03-13 15:03:18'),(376,1,200,'/series/flash/','127.0.0.1:63570','Authorized Call','2018-03-13 15:03:23'),(377,2,400,'/series/flash/','127.0.0.1:63570','Bad Request','2018-03-13 15:03:25'),(378,1,200,'/series/flash/','127.0.0.1:63570','Authorized Call','2018-03-13 15:03:52'),(379,1,200,'/series/BoJack%20Horseman/image','127.0.0.1:63620','Authorized Call','2018-03-13 15:07:09'),(380,1,200,'/series/','127.0.0.1:63686','Authorized Call','2018-03-13 15:10:16'),(381,2,404,'/series/flash/','127.0.0.1:63710','Not Found','2018-03-13 15:12:52'),(382,2,404,'/series/flash/','127.0.0.1:63710','Not Found','2018-03-13 15:12:57'),(383,2,400,'/series/flash/','127.0.0.1:63710','Bad Request','2018-03-13 15:13:07'),(384,2,400,'/series/flash/','127.0.0.1:63761','Bad Request','2018-03-13 15:17:33'),(385,2,400,'/series/flash/','127.0.0.1:63761','Bad Request','2018-03-13 15:17:35'),(386,2,404,'/series/flash/','127.0.0.1:63761','parser error: Invalid Series Url passed','2018-03-13 15:18:00'),(387,1,200,'/series/','127.0.0.1:63852','Authorized Call','2018-03-13 15:20:52'),(388,2,400,'/series/flash/','127.0.0.1:63857','Bad Request','2018-03-13 15:21:01'),(389,1,200,'/series/','127.0.0.1:63866','Authorized Call','2018-03-13 15:21:44'),(390,1,200,'/series/','127.0.0.1:63866','Authorized Call','2018-03-13 15:21:45'),(391,1,200,'/series/flash/','127.0.0.1:63871','Authorized Call','2018-03-13 15:21:48'),(392,1,200,'/series/','127.0.0.1:63866','Authorized Call','2018-03-13 15:22:00'),(393,1,200,'/series/BoJack%20Horseman/','127.0.0.1:63866','Authorized Call','2018-03-13 15:22:06'),(394,1,200,'/series/BoJack%20Horseman/image','127.0.0.1:63866','Authorized Call','2018-03-13 15:22:11'),(395,2,404,'/robots.txt','127.0.0.1:63888','Not Found','2018-03-13 15:22:11'),(396,1,200,'/series/','127.0.0.1:63914','Authorized Call','2018-03-13 15:25:21'),(397,1,200,'/series/','127.0.0.1:64642','Authorized Call','2018-03-13 15:28:26'),(398,1,200,'/series/x/','127.0.0.1:49858','Authorized Call','2018-03-13 15:35:16'),(399,1,200,'/series/','127.0.0.1:49936','Authorized Call','2018-03-13 15:35:33'),(400,1,200,'/series/Marvel%27s%20Agents%20of%20S.H.I.E.L.D./','127.0.0.1:49936','Authorized Call','2018-03-13 15:35:47'),(401,1,200,'/series/Marvel%27s%20Agents%20of%20S.H.I.E.L.D./image','127.0.0.1:49936','Authorized Call','2018-03-13 15:35:50'),(402,1,200,'/series/x/','127.0.0.1:50166','Authorized Call','2018-03-13 15:36:28'),(403,1,200,'/series/x/','127.0.0.1:50580','Authorized Call','2018-03-13 15:38:06'),(404,1,200,'/series/x/','127.0.0.1:50580','Authorized Call','2018-03-13 15:38:48'),(405,1,200,'/series/x/','127.0.0.1:51105','Authorized Call','2018-03-13 15:40:08'),(406,1,200,'/series/x/','127.0.0.1:51105','Authorized Call','2018-03-13 15:40:17'),(407,1,200,'/series/x/','127.0.0.1:51105','Authorized Call','2018-03-13 15:40:24'),(408,1,200,'/series/x/','127.0.0.1:51105','Authorized Call','2018-03-13 15:40:30'),(409,1,200,'/series/x/','127.0.0.1:51105','Authorized Call','2018-03-13 15:40:38'),(410,1,200,'/series/x/','127.0.0.1:51105','Authorized Call','2018-03-13 15:40:46'),(411,1,200,'/series/x/','127.0.0.1:51105','Authorized Call','2018-03-13 15:41:04'),(412,1,200,'/series/','127.0.0.1:51391','Authorized Call','2018-03-13 15:41:12'),(413,1,200,'/series/x/','127.0.0.1:51105','Authorized Call','2018-03-13 15:41:38'),(414,1,200,'/series/','127.0.0.1:51391','Authorized Call','2018-03-13 15:41:43'),(415,1,200,'/series/x/','127.0.0.1:51958','Authorized Call','2018-03-13 15:43:24'),(416,1,200,'/series/x/','127.0.0.1:51958','Authorized Call','2018-03-13 15:43:45'),(417,1,200,'/series/x/','127.0.0.1:51958','Authorized Call','2018-03-13 15:43:57'),(418,1,200,'/series/x/','127.0.0.1:51958','Authorized Call','2018-03-13 15:44:14'),(419,1,200,'/series/x/','127.0.0.1:51958','Authorized Call','2018-03-13 15:44:41'),(420,1,200,'/series/x/','127.0.0.1:51958','Authorized Call','2018-03-13 15:45:15'),(421,1,200,'/series/x/','127.0.0.1:51958','Authorized Call','2018-03-13 15:45:41'),(422,1,200,'/series/x/','127.0.0.1:51958','Authorized Call','2018-03-13 15:46:09'),(423,1,200,'/series/x/','127.0.0.1:51958','Authorized Call','2018-03-13 15:46:29'),(424,1,200,'/series/x/','127.0.0.1:51958','Authorized Call','2018-03-13 15:46:43'),(425,1,200,'/series/x/','127.0.0.1:51958','Authorized Call','2018-03-13 15:47:06'),(426,1,200,'/series/x/','127.0.0.1:51958','Authorized Call','2018-03-13 15:47:21'),(427,1,200,'/series/x/','127.0.0.1:51958','Authorized Call','2018-03-13 15:47:53'),(428,1,200,'/series/x/','127.0.0.1:51958','Authorized Call','2018-03-13 15:48:16'),(429,1,200,'/series/x/','127.0.0.1:51958','Authorized Call','2018-03-13 15:48:38'),(430,1,200,'/series/x/','127.0.0.1:51958','Authorized Call','2018-03-13 15:49:25'),(431,1,200,'/series/x/','127.0.0.1:51958','Authorized Call','2018-03-13 15:49:31'),(432,1,200,'/series/x/','127.0.0.1:53888','Authorized Call','2018-03-13 15:50:54'),(433,1,200,'/series/x/','127.0.0.1:53888','Authorized Call','2018-03-13 15:51:02'),(434,1,200,'/series/x/','127.0.0.1:53888','Authorized Call','2018-03-13 15:51:39'),(435,1,200,'/series/','127.0.0.1:55021','Authorized Call','2018-03-13 15:54:51'),(436,2,404,'/series/Stargate%20Universe/images','127.0.0.1:55021','Not Found','2018-03-13 15:55:07'),(437,1,200,'/series/Stargate%20Universe/image','127.0.0.1:55021','Authorized Call','2018-03-13 15:55:10'),(438,1,200,'/series/x/','127.0.0.1:55222','Authorized Call','2018-03-13 15:55:39'),(439,2,400,'/series/x/','127.0.0.1:55222','Bad Request','2018-03-13 15:55:40'),(440,1,200,'/series/x/','127.0.0.1:56360','Authorized Call','2018-03-13 16:00:04'),(441,1,200,'/series/x/','127.0.0.1:56360','Authorized Call','2018-03-13 16:00:15'),(442,1,200,'/series/x/','127.0.0.1:56360','Authorized Call','2018-03-13 16:00:25'),(443,1,200,'/series/x/','127.0.0.1:56360','Authorized Call','2018-03-13 16:00:37'),(444,1,200,'/series/x/','127.0.0.1:56360','Authorized Call','2018-03-13 16:00:44'),(445,1,200,'/series/x/','127.0.0.1:56360','Authorized Call','2018-03-13 16:00:52'),(446,2,400,'/series/x/','127.0.0.1:56360','Bad Request','2018-03-13 16:01:52'),(447,1,200,'/series/x/','127.0.0.1:57120','Authorized Call','2018-03-13 16:03:02'),(448,1,200,'/series/','127.0.0.1:57429','Authorized Call','2018-03-13 16:04:14'),(449,1,200,'/series/','127.0.0.1:57429','Authorized Call','2018-03-13 16:04:15'),(450,2,404,'/robots.txt','127.0.0.1:58290','Not Found','2018-03-13 16:07:10'),(451,2,405,'/series/','127.0.0.1:58503','Method Not Allowed','2018-03-13 16:17:17'),(452,1,200,'/series/','127.0.0.1:58517','Authorized Call','2018-03-13 16:18:22');
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Provider`
--

LOCK TABLES `Provider` WRITE;
/*!40000 ALTER TABLE `Provider` DISABLE KEYS */;
INSERT INTO `Provider` VALUES (1,'Netflix',5),(2,'Burning Series',4),(3,'Amazon Prime',6);
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
  `Image_id` int(11) NOT NULL,
  `Title` varchar(255) DEFAULT NULL,
  `ProviderUrl` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_Series_Image_idx` (`Image_id`),
  CONSTRAINT `fk_Series_Image` FOREIGN KEY (`Image_id`) REFERENCES `Image` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Series`
--

LOCK TABLES `Series` WRITE;
/*!40000 ALTER TABLE `Series` DISABLE KEYS */;
INSERT INTO `Series` VALUES (5,8,'BoJack Horseman','https://www.themoviedb.org/tv/61222-bojack-horseman'),(6,9,'Marvel\'s Agents of S.H.I.E.L.D.','https://www.themoviedb.org/tv/1403-marvel-s-agents-of-s-h-i-e-l-d'),(7,10,'American Gods','https://www.themoviedb.org/tv/46639-american-gods'),(8,11,'Archer','https://www.themoviedb.org/tv/10283-archer'),(9,12,'Arrow','https://www.themoviedb.org/tv/1412-arrow'),(10,13,'Marvel\'s Daredevil','https://www.themoviedb.org/tv/61889-daredevil'),(11,14,'Marvel\'s Jessica Jones','https://www.themoviedb.org/tv/38472-marvel-s-jessica-jones'),(12,15,'Marvel\'s Iron Fist','https://www.themoviedb.org/tv/62127-marvel-s-iron-fist'),(13,16,'Marvel\'s The Punisher','https://www.themoviedb.org/tv/67178-marvel-s-the-punisher'),(14,17,'Marvel\'s The Defenders','https://www.themoviedb.org/tv/62285-marvel-s-the-defenders'),(15,18,'Marvel\'s Luke Cage','https://www.themoviedb.org/tv/62126-marvel-s-luke-cage'),(16,19,'Dexter','https://www.themoviedb.org/tv/1405-dexter'),(17,20,'The Walking Dead','https://www.themoviedb.org/tv/1402-the-walking-dead'),(18,21,'Futurama','https://www.themoviedb.org/tv/615-futurama'),(19,22,'Family Guy','https://www.themoviedb.org/tv/1434-family-guy'),(20,23,'The Simpsons','https://www.themoviedb.org/tv/456-the-simpsons'),(21,24,'Game of Thrones','https://www.themoviedb.org/tv/1399-game-of-thrones'),(22,25,'Hannibal','https://www.themoviedb.org/tv/40008-hannibal'),(23,26,'Limitless','https://www.themoviedb.org/tv/62687-limitless'),(24,27,'Lucifer','https://www.themoviedb.org/tv/63174-lucifer'),(25,28,'Mr. Robot','https://www.themoviedb.org/tv/62560-mr-robot'),(26,29,'NCIS','https://www.themoviedb.org/tv/4614-ncis'),(27,30,'New Girl','https://www.themoviedb.org/tv/1420-new-girl'),(28,31,'Rick and Morty','https://www.themoviedb.org/tv/60625-rick-and-morty'),(29,32,'South Park','https://www.themoviedb.org/tv/2190-south-park'),(30,33,'Star Trek: Discovery','https://www.themoviedb.org/tv/67198-star-trek-discovery'),(31,34,'Westworld','https://www.themoviedb.org/tv/63247-westworld'),(32,35,'Your Pretty Face Is Going to Hell','https://www.themoviedb.org/tv/48000-your-pretty-face-is-going-to-hell'),(33,36,'Better Call Saul','https://www.themoviedb.org/tv/60059-better-call-saul'),(34,37,'Breaking Bad','https://www.themoviedb.org/tv/1396-breaking-bad'),(35,38,'Altered Carbon','https://www.themoviedb.org/tv/68421-altered-carbon'),(36,39,'Stranger Things','https://www.themoviedb.org/tv/66732-stranger-things'),(37,40,'Stargate Universe','https://www.themoviedb.org/tv/5148-stargate-universe'),(38,41,'Stargate Atlantis','https://www.themoviedb.org/tv/2290-stargate-atlantis'),(39,42,'Brooklyn Nine-Nine','https://www.themoviedb.org/tv/48891-brooklyn-nine-nine'),(40,43,'The Office','https://www.themoviedb.org/tv/2316-the-office'),(41,44,'Black Lightning','https://www.themoviedb.org/tv/71663-black-lightning'),(42,45,'House of Cards','https://www.themoviedb.org/tv/1425-house-of-cards'),(43,46,'Lost','https://www.themoviedb.org/tv/4607-lost'),(44,47,'Suits','https://www.themoviedb.org/tv/37680-suits'),(45,48,'Parks and Recreation','https://www.themoviedb.org/tv/8592-parks-and-recreation');
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

-- Dump completed on 2018-03-13 16:21:28
