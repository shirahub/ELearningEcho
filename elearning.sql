CREATE DATABASE  IF NOT EXISTS `elearning` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `elearning`;
-- MySQL dump 10.13  Distrib 5.7.30, for Linux (x86_64)
--
-- Host: localhost    Database: elearning
-- ------------------------------------------------------
-- Server version	5.7.30

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
-- Table structure for table `paketsoal`
--

DROP TABLE IF EXISTS `paketsoal`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `paketsoal` (
  `id_paketsoal` int(11) NOT NULL AUTO_INCREMENT,
  `tingkat` varchar(45) DEFAULT NULL,
  `kelas` varchar(45) DEFAULT NULL,
  `mapel` varchar(45) DEFAULT NULL,
  `tema` varchar(45) DEFAULT NULL,
  `id_user` int(11) DEFAULT NULL,
  PRIMARY KEY (`id_paketsoal`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `paketsoal`
--

LOCK TABLES `paketsoal` WRITE;
/*!40000 ALTER TABLE `paketsoal` DISABLE KEYS */;
INSERT INTO `paketsoal` VALUES (1,'SD','1','Matematika','Bab 1',1),(2,'SD','1','Bahasa Inggris','Bab 1',1),(3,'SD','2','Matematika','Bab 2',1),(4,'SMA','12','Matematika','Easy',1),(6,'SMP','7','Matematika','Bab 1',1),(8,'SD','1','Matematika','Bab 3',2);
/*!40000 ALTER TABLE `paketsoal` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `soal`
--

DROP TABLE IF EXISTS `soal`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `soal` (
  `id_soal` int(11) NOT NULL AUTO_INCREMENT,
  `id_paketsoal` int(11) DEFAULT NULL,
  `pertanyaan` varchar(45) DEFAULT NULL,
  `jawaban` varchar(45) DEFAULT NULL,
  `pilihan1` varchar(45) DEFAULT NULL,
  `pilihan2` varchar(45) DEFAULT NULL,
  `pilihan3` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id_soal`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `soal`
--

LOCK TABLES `soal` WRITE;
/*!40000 ALTER TABLE `soal` DISABLE KEYS */;
INSERT INTO `soal` VALUES (7,2,'Good Morning','Selamat pagi','Selamat siang','Selamat sore','Selamat malam'),(8,2,'Good Night','Selamat Malam','Selamat pagi','Selamat siang','Selamat sore'),(9,3,'2x2=','4','6','8','10'),(10,3,'3x3=','9','12','8','15'),(11,3,'4x4=','16','20','10','15'),(14,6,'11x11=','121','111','1331','112'),(15,1,'1+10','11','3','4','5'),(16,1,'2+2=','4','5','6','10'),(17,1,'3+3=','6','7','8','9'),(18,1,'2+3=','5','6','7','8'),(19,1,'5+5=','10','9','7','8'),(20,8,'59+31=','90','100','190','80'),(21,8,'62+28','90','100','70','80'),(22,8,'12+88','100','90','80','70'),(23,4,'100x100=','10000','10','100','1000'),(24,4,'1000-999=','1','2','3','4');
/*!40000 ALTER TABLE `soal` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id_user` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(45) DEFAULT NULL,
  `password` varchar(45) DEFAULT NULL,
  `fullname` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id_user`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'shir','leen','Shirleen Adriana'),(2,'shirleen@gmail.com','@G2academy','Shirleen Sungkar');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `userhistory`
--

DROP TABLE IF EXISTS `userhistory`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `userhistory` (
  `id_history` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) DEFAULT NULL,
  `id_paketsoal` int(11) DEFAULT NULL,
  `nilai` float DEFAULT NULL,
  `waktu` datetime DEFAULT NULL,
  PRIMARY KEY (`id_history`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `userhistory`
--

LOCK TABLES `userhistory` WRITE;
/*!40000 ALTER TABLE `userhistory` DISABLE KEYS */;
INSERT INTO `userhistory` VALUES (1,2,1,60,'2020-08-17 22:49:24'),(2,2,2,100,'2020-08-17 22:51:50'),(3,2,4,0,'2020-08-17 23:12:11'),(4,2,8,33,'2020-08-18 00:13:35');
/*!40000 ALTER TABLE `userhistory` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-08-18  7:14:38
