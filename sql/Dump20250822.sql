-- MySQL dump 10.13  Distrib 8.0.36, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: livraisons
-- ------------------------------------------------------
-- Server version	8.0.43-0ubuntu0.24.04.1

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
-- Table structure for table `Adresse`
--

DROP TABLE IF EXISTS `Adresse`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Adresse` (
  `adresse_id` int NOT NULL AUTO_INCREMENT,
  `client_id` int NOT NULL,
  `rue` varchar(150) DEFAULT NULL,
  `ville` varchar(100) DEFAULT NULL,
  `code_postal` varchar(20) DEFAULT NULL,
  `pays` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`adresse_id`),
  KEY `client_id` (`client_id`),
  CONSTRAINT `Adresse_ibfk_1` FOREIGN KEY (`client_id`) REFERENCES `Client` (`client_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Adresse`
--

LOCK TABLES `Adresse` WRITE;
/*!40000 ALTER TABLE `Adresse` DISABLE KEYS */;
INSERT INTO `Adresse` VALUES (1,1,'Rue 1','Antananarivo','101','Madagascar'),(2,2,'Rue 2','Antsirabe','110','Madagascar');
/*!40000 ALTER TABLE `Adresse` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Article`
--

DROP TABLE IF EXISTS `Article`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Article` (
  `article_id` int NOT NULL AUTO_INCREMENT,
  `nom` varchar(150) NOT NULL,
  `description` text,
  `prix` decimal(10,2) NOT NULL,
  `stock` int DEFAULT '0',
  `commercant_id` int NOT NULL,
  `categorie_id` int DEFAULT NULL,
  PRIMARY KEY (`article_id`),
  KEY `commercant_id` (`commercant_id`),
  KEY `categorie_id` (`categorie_id`),
  CONSTRAINT `Article_ibfk_1` FOREIGN KEY (`commercant_id`) REFERENCES `Commercant` (`commercant_id`),
  CONSTRAINT `Article_ibfk_2` FOREIGN KEY (`categorie_id`) REFERENCES `Categorie` (`categorie_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Article`
--

LOCK TABLES `Article` WRITE;
/*!40000 ALTER TABLE `Article` DISABLE KEYS */;
INSERT INTO `Article` VALUES (1,'T-shirt rouge','T-shirt coton rouge',15000.00,50,1,1),(2,'Pantalon bleu','Pantalon jean bleu',35000.00,30,1,1),(3,'Sac à main','Sac en cuir noir',45000.00,20,2,2),(4,'Chaussures sport','Chaussures running',60000.00,15,1,3),(8,'\"testa\"','\"test description\"',123.00,5,1,1),(10,'\"testa\"','\"test description\"',123.00,5,1,1),(11,'\"testa\"','\"test description\"',123.00,5,1,1);
/*!40000 ALTER TABLE `Article` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Article_Image`
--

DROP TABLE IF EXISTS `Article_Image`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Article_Image` (
  `image_id` int NOT NULL AUTO_INCREMENT,
  `article_id` int NOT NULL,
  `url` varchar(255) NOT NULL,
  `largeur` int DEFAULT NULL,
  `hauteur` int DEFAULT NULL,
  `ordre` int DEFAULT '0',
  `type` enum('main','gallery','thumbnail') DEFAULT 'gallery',
  `taille` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`image_id`),
  KEY `article_id` (`article_id`),
  CONSTRAINT `Article_Image_ibfk_1` FOREIGN KEY (`article_id`) REFERENCES `Article` (`article_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Article_Image`
--

LOCK TABLES `Article_Image` WRITE;
/*!40000 ALTER TABLE `Article_Image` DISABLE KEYS */;
INSERT INTO `Article_Image` VALUES (1,1,'/uploads/articles/1/tshirt_red_main.jpg',800,800,1,'main','L'),(2,1,'/uploads/articles/1/tshirt_red_2.jpg',600,600,2,'gallery','L'),(3,2,'/uploads/articles/2/pantalon_bleu_main.jpg',800,800,1,'main','M'),(4,3,'/uploads/articles/3/sac_main.jpg',800,800,1,'main','M'),(5,4,'/uploads/articles/4/chaussures_main.jpg',800,800,1,'main','M'),(6,8,'uploads/8_devant.jpg',0,0,1,'main','69 KB'),(7,10,'uploads/10_devant.jpg',500,600,0,'main','L'),(8,11,'uploads/11_devant.jpg',500,600,1,'main','L');
/*!40000 ALTER TABLE `Article_Image` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Avis`
--

DROP TABLE IF EXISTS `Avis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Avis` (
  `avis_id` int NOT NULL AUTO_INCREMENT,
  `client_id` int NOT NULL,
  `article_id` int DEFAULT NULL,
  `commercant_id` int DEFAULT NULL,
  `livreur_id` int DEFAULT NULL,
  `note` int DEFAULT NULL,
  `commentaire` text,
  `date_avis` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`avis_id`),
  KEY `client_id` (`client_id`),
  KEY `article_id` (`article_id`),
  KEY `commercant_id` (`commercant_id`),
  KEY `livreur_id` (`livreur_id`),
  CONSTRAINT `Avis_ibfk_1` FOREIGN KEY (`client_id`) REFERENCES `Client` (`client_id`),
  CONSTRAINT `Avis_ibfk_2` FOREIGN KEY (`article_id`) REFERENCES `Article` (`article_id`),
  CONSTRAINT `Avis_ibfk_3` FOREIGN KEY (`commercant_id`) REFERENCES `Commercant` (`commercant_id`),
  CONSTRAINT `Avis_ibfk_4` FOREIGN KEY (`livreur_id`) REFERENCES `Livreur` (`livreur_id`),
  CONSTRAINT `Avis_chk_1` CHECK ((`note` between 1 and 5))
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Avis`
--

LOCK TABLES `Avis` WRITE;
/*!40000 ALTER TABLE `Avis` DISABLE KEYS */;
INSERT INTO `Avis` VALUES (1,1,1,1,1,5,'Très bon produit, livraison rapide!','2025-08-20 09:22:06'),(2,2,3,2,NULL,4,'Sac correct, qualité satisfaisante.','2025-08-20 09:22:06');
/*!40000 ALTER TABLE `Avis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Categorie`
--

DROP TABLE IF EXISTS `Categorie`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Categorie` (
  `categorie_id` int NOT NULL AUTO_INCREMENT,
  `nom` varchar(100) NOT NULL,
  `parent_id` int DEFAULT NULL,
  PRIMARY KEY (`categorie_id`),
  KEY `parent_id` (`parent_id`),
  CONSTRAINT `Categorie_ibfk_1` FOREIGN KEY (`parent_id`) REFERENCES `Categorie` (`categorie_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Categorie`
--

LOCK TABLES `Categorie` WRITE;
/*!40000 ALTER TABLE `Categorie` DISABLE KEYS */;
INSERT INTO `Categorie` VALUES (1,'Vêtements',NULL),(2,'Accessoires',NULL),(3,'Chaussures',1);
/*!40000 ALTER TABLE `Categorie` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Client`
--

DROP TABLE IF EXISTS `Client`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Client` (
  `client_id` int NOT NULL AUTO_INCREMENT,
  `nom` varchar(100) NOT NULL,
  `prenom` varchar(100) DEFAULT NULL,
  `email` varchar(150) NOT NULL,
  `telephone` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`client_id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Client`
--

LOCK TABLES `Client` WRITE;
/*!40000 ALTER TABLE `Client` DISABLE KEYS */;
INSERT INTO `Client` VALUES (1,'Rakoto','Andry','andry@example.com','0341234567'),(2,'Rasoa','Mialy','mialy@example.com','0349876543');
/*!40000 ALTER TABLE `Client` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Commande`
--

DROP TABLE IF EXISTS `Commande`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Commande` (
  `commande_id` int NOT NULL AUTO_INCREMENT,
  `client_id` int NOT NULL,
  `date_commande` datetime DEFAULT CURRENT_TIMESTAMP,
  `montant_total` decimal(12,2) NOT NULL,
  `status` enum('en_attente','payee','annulee') DEFAULT 'en_attente',
  PRIMARY KEY (`commande_id`),
  KEY `client_id` (`client_id`),
  CONSTRAINT `Commande_ibfk_1` FOREIGN KEY (`client_id`) REFERENCES `Client` (`client_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Commande`
--

LOCK TABLES `Commande` WRITE;
/*!40000 ALTER TABLE `Commande` DISABLE KEYS */;
INSERT INTO `Commande` VALUES (1,1,'2025-08-20 09:22:04',75000.00,'en_attente');
/*!40000 ALTER TABLE `Commande` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Commande_Article`
--

DROP TABLE IF EXISTS `Commande_Article`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Commande_Article` (
  `commande_id` int NOT NULL,
  `article_id` int NOT NULL,
  `quantite` int NOT NULL,
  `prix_unitaire` decimal(10,2) NOT NULL,
  PRIMARY KEY (`commande_id`,`article_id`),
  KEY `article_id` (`article_id`),
  CONSTRAINT `Commande_Article_ibfk_1` FOREIGN KEY (`commande_id`) REFERENCES `Commande` (`commande_id`),
  CONSTRAINT `Commande_Article_ibfk_2` FOREIGN KEY (`article_id`) REFERENCES `Article` (`article_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Commande_Article`
--

LOCK TABLES `Commande_Article` WRITE;
/*!40000 ALTER TABLE `Commande_Article` DISABLE KEYS */;
INSERT INTO `Commande_Article` VALUES (1,1,2,15000.00),(1,3,1,45000.00);
/*!40000 ALTER TABLE `Commande_Article` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Commercant`
--

DROP TABLE IF EXISTS `Commercant`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Commercant` (
  `commercant_id` int NOT NULL AUTO_INCREMENT,
  `nom` varchar(150) NOT NULL,
  `description` text,
  `adresse` varchar(200) DEFAULT NULL,
  `telephone` varchar(50) DEFAULT NULL,
  `email` varchar(150) DEFAULT NULL,
  PRIMARY KEY (`commercant_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Commercant`
--

LOCK TABLES `Commercant` WRITE;
/*!40000 ALTER TABLE `Commercant` DISABLE KEYS */;
INSERT INTO `Commercant` VALUES (1,'Boutique A','Vêtements','Antananarivo','0321234567','contact@boutiquea.mg'),(2,'Boutique B','Accessoires','Antsirabe','0329876543','contact@boutiqueb.mg');
/*!40000 ALTER TABLE `Commercant` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Livraison`
--

DROP TABLE IF EXISTS `Livraison`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Livraison` (
  `livraison_id` int NOT NULL AUTO_INCREMENT,
  `commande_id` int NOT NULL,
  `livreur_id` int NOT NULL,
  `article_id` int NOT NULL,
  `client_id` int NOT NULL,
  `commercant_id` int NOT NULL,
  `date_prevue` datetime DEFAULT NULL,
  `date_effective` datetime DEFAULT NULL,
  `duree_estimee` int DEFAULT NULL,
  `status` enum('en_attente','en_cours','livree','annulee','echouee') DEFAULT 'en_attente',
  `axe` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`livraison_id`),
  KEY `commande_id` (`commande_id`),
  KEY `livreur_id` (`livreur_id`),
  KEY `article_id` (`article_id`),
  KEY `client_id` (`client_id`),
  KEY `commercant_id` (`commercant_id`),
  CONSTRAINT `Livraison_ibfk_1` FOREIGN KEY (`commande_id`) REFERENCES `Commande` (`commande_id`),
  CONSTRAINT `Livraison_ibfk_2` FOREIGN KEY (`livreur_id`) REFERENCES `Livreur` (`livreur_id`),
  CONSTRAINT `Livraison_ibfk_3` FOREIGN KEY (`article_id`) REFERENCES `Article` (`article_id`),
  CONSTRAINT `Livraison_ibfk_4` FOREIGN KEY (`client_id`) REFERENCES `Client` (`client_id`),
  CONSTRAINT `Livraison_ibfk_5` FOREIGN KEY (`commercant_id`) REFERENCES `Commercant` (`commercant_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Livraison`
--

LOCK TABLES `Livraison` WRITE;
/*!40000 ALTER TABLE `Livraison` DISABLE KEYS */;
INSERT INTO `Livraison` VALUES (1,1,1,1,1,1,'2025-08-21 10:00:00',NULL,60,'en_attente','Antananarivo-Centre'),(2,1,1,3,1,2,'2025-08-21 10:30:00',NULL,60,'en_attente','Antananarivo-Centre');
/*!40000 ALTER TABLE `Livraison` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Livreur`
--

DROP TABLE IF EXISTS `Livreur`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Livreur` (
  `livreur_id` int NOT NULL AUTO_INCREMENT,
  `nom` varchar(100) NOT NULL,
  `telephone` varchar(50) DEFAULT NULL,
  `vehicule` varchar(50) DEFAULT NULL,
  `zone_livraison` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`livreur_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Livreur`
--

LOCK TABLES `Livreur` WRITE;
/*!40000 ALTER TABLE `Livreur` DISABLE KEYS */;
INSERT INTO `Livreur` VALUES (1,'Rajaonarison Tiana','0331234567','Moto','Antananarivo');
/*!40000 ALTER TABLE `Livreur` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ModePaiement`
--

DROP TABLE IF EXISTS `ModePaiement`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ModePaiement` (
  `mode_id` int NOT NULL AUTO_INCREMENT,
  `type` enum('MobileMoney','Cash','Wallet') NOT NULL,
  `operateur` enum('Airtel','Mvola','OrangeMoney') DEFAULT NULL,
  PRIMARY KEY (`mode_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ModePaiement`
--

LOCK TABLES `ModePaiement` WRITE;
/*!40000 ALTER TABLE `ModePaiement` DISABLE KEYS */;
INSERT INTO `ModePaiement` VALUES (1,'MobileMoney','Mvola'),(2,'Wallet',NULL);
/*!40000 ALTER TABLE `ModePaiement` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Paiement`
--

DROP TABLE IF EXISTS `Paiement`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Paiement` (
  `paiement_id` int NOT NULL AUTO_INCREMENT,
  `commande_id` int NOT NULL,
  `mode_id` int NOT NULL,
  `montant` decimal(12,2) NOT NULL,
  `reference` varchar(100) DEFAULT NULL,
  `date_paiement` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`paiement_id`),
  KEY `commande_id` (`commande_id`),
  KEY `mode_id` (`mode_id`),
  CONSTRAINT `Paiement_ibfk_1` FOREIGN KEY (`commande_id`) REFERENCES `Commande` (`commande_id`),
  CONSTRAINT `Paiement_ibfk_2` FOREIGN KEY (`mode_id`) REFERENCES `ModePaiement` (`mode_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Paiement`
--

LOCK TABLES `Paiement` WRITE;
/*!40000 ALTER TABLE `Paiement` DISABLE KEYS */;
INSERT INTO `Paiement` VALUES (1,1,1,45000.00,'MM12345','2025-08-20 09:22:04'),(2,1,2,45000.00,'WLT56789','2025-08-20 09:22:04');
/*!40000 ALTER TABLE `Paiement` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Panier`
--

DROP TABLE IF EXISTS `Panier`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Panier` (
  `panier_id` int NOT NULL AUTO_INCREMENT,
  `client_id` int NOT NULL,
  `date_creation` datetime DEFAULT CURRENT_TIMESTAMP,
  `status` enum('en_cours','valide','abandonne') DEFAULT 'en_cours',
  PRIMARY KEY (`panier_id`),
  KEY `client_id` (`client_id`),
  CONSTRAINT `Panier_ibfk_1` FOREIGN KEY (`client_id`) REFERENCES `Client` (`client_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Panier`
--

LOCK TABLES `Panier` WRITE;
/*!40000 ALTER TABLE `Panier` DISABLE KEYS */;
INSERT INTO `Panier` VALUES (1,1,'2025-08-20 09:22:04','en_cours');
/*!40000 ALTER TABLE `Panier` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Panier_Article`
--

DROP TABLE IF EXISTS `Panier_Article`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Panier_Article` (
  `panier_id` int NOT NULL,
  `article_id` int NOT NULL,
  `quantite` int NOT NULL,
  PRIMARY KEY (`panier_id`,`article_id`),
  KEY `article_id` (`article_id`),
  CONSTRAINT `Panier_Article_ibfk_1` FOREIGN KEY (`panier_id`) REFERENCES `Panier` (`panier_id`),
  CONSTRAINT `Panier_Article_ibfk_2` FOREIGN KEY (`article_id`) REFERENCES `Article` (`article_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Panier_Article`
--

LOCK TABLES `Panier_Article` WRITE;
/*!40000 ALTER TABLE `Panier_Article` DISABLE KEYS */;
INSERT INTO `Panier_Article` VALUES (1,1,2),(1,3,1);
/*!40000 ALTER TABLE `Panier_Article` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Promotion`
--

DROP TABLE IF EXISTS `Promotion`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Promotion` (
  `promotion_id` int NOT NULL AUTO_INCREMENT,
  `article_id` int NOT NULL,
  `code` varchar(50) NOT NULL,
  `reduction` decimal(5,2) NOT NULL,
  `date_debut` datetime NOT NULL,
  `date_fin` datetime NOT NULL,
  PRIMARY KEY (`promotion_id`),
  KEY `article_id` (`article_id`),
  CONSTRAINT `Promotion_ibfk_1` FOREIGN KEY (`article_id`) REFERENCES `Article` (`article_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Promotion`
--

LOCK TABLES `Promotion` WRITE;
/*!40000 ALTER TABLE `Promotion` DISABLE KEYS */;
INSERT INTO `Promotion` VALUES (1,1,'PROMO10',10.00,'2025-08-20 00:00:00','2025-08-31 23:59:59'),(2,3,'PROMO15',15.00,'2025-08-20 00:00:00','2025-08-25 23:59:59');
/*!40000 ALTER TABLE `Promotion` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Wallet`
--

DROP TABLE IF EXISTS `Wallet`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Wallet` (
  `wallet_id` int NOT NULL AUTO_INCREMENT,
  `client_id` int NOT NULL,
  `solde` decimal(12,2) DEFAULT '0.00',
  `date_creation` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`wallet_id`),
  KEY `client_id` (`client_id`),
  CONSTRAINT `Wallet_ibfk_1` FOREIGN KEY (`client_id`) REFERENCES `Client` (`client_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Wallet`
--

LOCK TABLES `Wallet` WRITE;
/*!40000 ALTER TABLE `Wallet` DISABLE KEYS */;
INSERT INTO `Wallet` VALUES (1,1,10000.00,'2025-08-20 09:22:04'),(2,2,5000.00,'2025-08-20 09:22:04');
/*!40000 ALTER TABLE `Wallet` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `profil`
--

DROP TABLE IF EXISTS `profil`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `profil` (
  `idProfil` int NOT NULL AUTO_INCREMENT,
  `nomProfil` varchar(100) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`idProfil`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `profil`
--

LOCK TABLES `profil` WRITE;
/*!40000 ALTER TABLE `profil` DISABLE KEYS */;
INSERT INTO `profil` VALUES (1,'client','Profil pour les clients'),(2,'commerçant','Profil pour les commerçants'),(3,'livreur','Profil pour les livreurs');
/*!40000 ALTER TABLE `profil` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `login` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `name` varchar(100) NOT NULL,
  `lastname` varchar(100) NOT NULL,
  `type` varchar(50) NOT NULL,
  `contact` varchar(20) NOT NULL,
  `mail` varchar(150) NOT NULL,
  `adresse` text NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `login` (`login`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'hery','hery1234','Rasolonjatovo','Hery','','12345','hery@test.mg','ffdqfqfq'),(4,'testdddddddddd','testddddddddd','testname','testlast','test','1234','mail@gmail.com','adressetna');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `userProfil`
--

DROP TABLE IF EXISTS `userProfil`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `userProfil` (
  `idUser` int NOT NULL,
  `idProfil` int NOT NULL,
  PRIMARY KEY (`idUser`,`idProfil`),
  KEY `idProfil` (`idProfil`),
  CONSTRAINT `userProfil_ibfk_1` FOREIGN KEY (`idUser`) REFERENCES `user` (`id`) ON DELETE CASCADE,
  CONSTRAINT `userProfil_ibfk_2` FOREIGN KEY (`idProfil`) REFERENCES `profil` (`idProfil`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `userProfil`
--

LOCK TABLES `userProfil` WRITE;
/*!40000 ALTER TABLE `userProfil` DISABLE KEYS */;
INSERT INTO `userProfil` VALUES (1,1);
/*!40000 ALTER TABLE `userProfil` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-08-22  9:16:49
