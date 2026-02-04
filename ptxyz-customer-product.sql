-- MySQL dump 10.13  Distrib 8.0.45, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: ptxyz-customer-product
-- ------------------------------------------------------
-- Server version	8.0.44

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `konsumen_tenor_limit`
--

DROP TABLE IF EXISTS `konsumen_tenor_limit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `konsumen_tenor_limit` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `public_id` varchar(255) NOT NULL,
  `user_id` bigint unsigned NOT NULL,
  `tenor_ref_id` bigint unsigned NOT NULL,
  `limit` decimal(15,0) NOT NULL,
  `used_amount` decimal(15,0) DEFAULT NULL,
  `balance` decimal(15,0) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `created_by` int NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `updated_by` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `konsumen_tenor_limit_users_FK` (`user_id`),
  KEY `konsumen_tenor_limit_tenor_ref_FK` (`tenor_ref_id`),
  CONSTRAINT `konsumen_tenor_limit_tenor_ref_FK` FOREIGN KEY (`tenor_ref_id`) REFERENCES `tenor_ref` (`id`),
  CONSTRAINT `konsumen_tenor_limit_users_FK` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `konsumen_tenor_limit`
--

LOCK TABLES `konsumen_tenor_limit` WRITE;
/*!40000 ALTER TABLE `konsumen_tenor_limit` DISABLE KEYS */;
INSERT INTO `konsumen_tenor_limit` VALUES (1,'01J4YQ6A7M9K2Z8W3F5R1D0H8C',1,1,100000,NULL,NULL,'2026-02-03 10:57:24',1,NULL,NULL),(2,'01J6B9ABCD9F5R2W7H1C4XPZKM',1,2,200000,NULL,NULL,'2026-02-03 10:57:26',1,NULL,NULL),(3,'01J6B9ABCDA1Z3K9M5R2W7H4XP',1,3,500000,NULL,NULL,'2026-02-03 10:57:28',1,NULL,NULL),(4,'01J6B9ABCDQ8Z5M3R2W7H1C4XP',1,6,700000,12000,1487000,'2026-02-03 10:57:31',1,'2026-02-04 06:52:50',NULL),(5,'01J6B9ABCD7KZ9M5F3R2W1H8CX',2,1,1000000,321,123,'2026-02-03 10:57:24',1,'2026-02-04 06:05:11',NULL),(6,'01J6B9ABCD5MZ9F3R2W7H1C4XP',2,2,1200000,111,222,'2026-02-03 10:57:26',1,'2026-02-04 05:55:28',NULL),(7,'01J6B9ABCDZ9M5F3R2W7H1C4XP',1,3,1500000,NULL,NULL,'2026-02-03 10:57:28',1,NULL,NULL),(8,'01J6B9ABCDKMZ9F3R2W7H1C4XP',2,6,2000000,NULL,NULL,'2026-02-03 10:57:31',1,NULL,NULL);
/*!40000 ALTER TABLE `konsumen_tenor_limit` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `public_id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `category` enum('White Goods','Motor','Mobil') NOT NULL,
  `harga_otr` decimal(15,0) NOT NULL,
  `created_at` datetime NOT NULL,
  `created_by` int NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `updated_by` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1,'01J7ZQ8W6XK3M4A9F2P7C5R8LQ','Kulkas 2 Pintu Polytron Belleza','White Goods',3750000,'2026-01-31 23:55:44',1,NULL,NULL),(2,'01J7ZQB1N8E6H9M5K2W3F7RPLA','Mesin Cuci Front Load Sharp','White Goods',4250000,'2026-01-31 23:55:46',1,NULL,NULL),(3,'01J7ZQG9M2R5A8KXH4W6F3NPEB','Honda Beat CBS','Motor',18500000,'2026-01-31 23:55:48',1,NULL,NULL),(4,'01J7ZQKDW8A9P6R2H5M4FEXN3C','Yamaha NMAX 155','Motor',32000000,'2026-01-31 23:55:50',1,NULL,NULL),(5,'01J7ZQQ4R5W2M6P8K9A3HXFNE7','Toyota Avanza','Mobil',310000000,'2026-01-31 23:55:52',1,NULL,NULL),(6,'01J7ZQT8F3A5KXW9M6H2P4RNEC','Daihatsu Xenia','Mobil',295000000,'2026-01-31 23:55:54',1,NULL,NULL);
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schema_migrations`
--

DROP TABLE IF EXISTS `schema_migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `schema_migrations` (
  `id` varchar(191) NOT NULL,
  `applied_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schema_migrations`
--

LOCK TABLES `schema_migrations` WRITE;
/*!40000 ALTER TABLE `schema_migrations` DISABLE KEYS */;
INSERT INTO `schema_migrations` VALUES ('20260131221942_create_users_table','2026-01-31 15:39:15.707');
/*!40000 ALTER TABLE `schema_migrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tenor_ref`
--

DROP TABLE IF EXISTS `tenor_ref`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tenor_ref` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `desc` text,
  `created_at` datetime NOT NULL,
  `created_by` int NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `updated_by` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tenor_ref`
--

LOCK TABLES `tenor_ref` WRITE;
/*!40000 ALTER TABLE `tenor_ref` DISABLE KEYS */;
INSERT INTO `tenor_ref` VALUES (1,'1 bulan','2026-01-31 23:50:22',1,NULL,NULL),(2,'2 bulan','2026-01-31 23:50:22',1,NULL,NULL),(3,'3 bulan','2026-01-31 23:50:22',1,NULL,NULL),(6,'6 bulan','2026-01-31 23:50:22',1,NULL,NULL);
/*!40000 ALTER TABLE `tenor_ref` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `public_id` varchar(255) NOT NULL,
  `nik` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `full_name` varchar(255) NOT NULL,
  `legal_name` varchar(255) NOT NULL,
  `tempat_lahir` varchar(50) NOT NULL,
  `tanggal_lahir` date NOT NULL,
  `gaji` decimal(15,2) NOT NULL,
  `foto_ktp` varchar(255) NOT NULL,
  `foto_selfie` varchar(255) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `created_by` bigint NOT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `updated_by` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_users_id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'01HZQ9YF6Z8A3W2C4B1M7N5RKP','3201010101010001','$2a$10$CwTycUXWue0Thq9StjUM0uJ8YzY5rG6K8R9GZJkQ2d9OQeFJ5jZHy','Andi Pratama','Andi Pratama','Jakarta','1995-03-15',8500000.00,'ktp_andi.jpg','selfie_andi.jpg','2026-02-01 09:35:10.799',1,'2026-02-01 09:35:10.799',NULL),(2,'01HZQ9Z2M4P7C9XW6B5A8F1KRN','3202020202020002','$2a$10$CwTycUXWue0Thq9StjUM0uJ8YzY5rG6K8R9GZJkQ2d9OQeFJ5jZHy','Budi Santoso','Budi Santoso','Bandung','1992-07-21',12000000.00,'ktp_budi.jpg','selfie_budi.jpg','2026-02-01 09:35:10.799',1,'2026-02-01 09:35:10.799',NULL),(3,'01HZQA01D9R3F6M7KZB8XW2A5C','3203030303030003','$2a$10$CwTycUXWue0Thq9StjUM0uJ8YzY5rG6K8R9GZJkQ2d9OQeFJ5jZHy','Citra Lestari','Citra Lestari','Surabaya','1998-11-05',9500000.00,'ktp_citra.jpg','selfie_citra.jpg','2026-02-01 09:35:10.799',1,'2026-02-01 09:35:10.799',NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'ptxyz-customer-product'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-02-04 16:03:24
