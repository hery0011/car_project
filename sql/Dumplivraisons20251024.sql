-- phpMyAdmin SQL Dump
-- version 5.2.2
-- https://www.phpmyadmin.net/
--
-- Host: mariadb:3306
-- Generation Time: Oct 24, 2025 at 01:51 PM
-- Server version: 12.0.2-MariaDB-ubu2404
-- PHP Version: 8.2.27

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `livraisons`
--

-- --------------------------------------------------------

--
-- Table structure for table `Adresse`
--

CREATE TABLE `Adresse` (
  `adresse_id` int(11) NOT NULL,
  `client_id` int(11) NOT NULL,
  `rue` varchar(150) DEFAULT NULL,
  `ville` varchar(100) DEFAULT NULL,
  `code_postal` varchar(20) DEFAULT NULL,
  `pays` varchar(50) DEFAULT NULL,
  `latitude` varchar(255) DEFAULT NULL,
  `longitude` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `Article`
--

CREATE TABLE `Article` (
  `article_id` int(11) NOT NULL,
  `nom` varchar(150) NOT NULL,
  `description` text DEFAULT NULL,
  `prix` decimal(10,2) NOT NULL,
  `stock` int(11) DEFAULT 0,
  `commercant_id` int(11) NOT NULL,
  `categorie_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `Article`
--

INSERT INTO `Article` (`article_id`, `nom`, `description`, `prix`, `stock`, `commercant_id`, `categorie_id`) VALUES
(1, 'telephone', 'telephone', 15000.00, 50, 1, 1),
(2, 'telephone', 'telephone', 35000.00, 30, 1, 1),
(3, 'telephone', 'telephone', 45000.00, 20, 2, 2),
(4, 'telephone', 'telephone', 60000.00, 15, 1, 3),
(8, 'telephone', 'telephone', 123.00, 5, 1, 1),
(10, 'telephone', 'telephone', 123.00, 5, 1, 1),
(11, 'telephone', 'telephone', 123.00, 5, 1, 1),
(14, 'telephone', 'telephone', 323.00, 2, 1, 1),
(15, 'telephone', 'telephone', 323.00, 2, 1, 1);

-- --------------------------------------------------------

--
-- Table structure for table `Article_Image`
--

CREATE TABLE `Article_Image` (
  `image_id` int(11) NOT NULL,
  `article_id` int(11) NOT NULL,
  `url` varchar(255) NOT NULL,
  `largeur` int(11) DEFAULT NULL,
  `hauteur` int(11) DEFAULT NULL,
  `ordre` int(11) DEFAULT 0,
  `type` enum('main','gallery','thumbnail') DEFAULT 'gallery',
  `taille` varchar(45) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `Article_Image`
--

INSERT INTO `Article_Image` (`image_id`, `article_id`, `url`, `largeur`, `hauteur`, `ordre`, `type`, `taille`) VALUES
(1, 1, '/uploads/1.png', 800, 800, 1, 'main', 'L'),
(2, 1, '/uploads/2.png', 600, 600, 2, 'gallery', 'L'),
(3, 2, '/uploads/3.png', 800, 800, 1, 'main', 'M'),
(4, 3, '/uploads/4.png', 800, 800, 1, 'main', 'M'),
(5, 4, '/uploads/5.png', 800, 800, 1, 'main', 'M'),
(6, 8, '/uploads/6.png', 0, 0, 1, 'main', '69 KB'),
(7, 10, '/uploads/1.png', 500, 600, 0, 'main', 'L'),
(8, 11, '/uploads/2.png', 500, 600, 1, 'main', 'L'),
(9, 14, '/uploads/3.png', 333, 111, 1, 'main', '222'),
(10, 15, '/uploads/4.png', 333, 111, 1, 'main', '222');

-- --------------------------------------------------------

--
-- Table structure for table `Avis`
--

CREATE TABLE `Avis` (
  `avis_id` int(11) NOT NULL,
  `client_id` int(11) NOT NULL,
  `article_id` int(11) DEFAULT NULL,
  `commercant_id` int(11) DEFAULT NULL,
  `livreur_id` int(11) DEFAULT NULL,
  `note` int(11) DEFAULT NULL,
  `commentaire` text DEFAULT NULL,
  `date_avis` datetime DEFAULT current_timestamp()
) ;

--
-- Dumping data for table `Avis`
--

INSERT INTO `Avis` (`avis_id`, `client_id`, `article_id`, `commercant_id`, `livreur_id`, `note`, `commentaire`, `date_avis`) VALUES
(1, 1, 1, 1, 1, 5, 'Très bon produit, livraison rapide!', '2025-08-20 09:22:06'),
(2, 2, 3, 2, NULL, 4, 'Sac correct, qualité satisfaisante.', '2025-08-20 09:22:06');

-- --------------------------------------------------------

--
-- Table structure for table `Categorie`
--

CREATE TABLE `Categorie` (
  `categorie_id` int(11) NOT NULL,
  `nom` varchar(100) NOT NULL,
  `parent_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `Categorie`
--

INSERT INTO `Categorie` (`categorie_id`, `nom`, `parent_id`) VALUES
(1, 'Vêtements', NULL),
(2, 'Accessoires', NULL),
(3, 'Chaussures', 1);

-- --------------------------------------------------------

--
-- Table structure for table `Client`
--

CREATE TABLE `Client` (
  `client_id` int(11) NOT NULL,
  `nom` varchar(100) NOT NULL,
  `prenom` varchar(100) DEFAULT NULL,
  `email` varchar(150) NOT NULL,
  `telephone` varchar(50) DEFAULT NULL,
  `adresse` varchar(45) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `Client`
--

INSERT INTO `Client` (`client_id`, `nom`, `prenom`, `email`, `telephone`, `adresse`) VALUES
(1, 'Rakoto', 'Andry', 'andry@example.com', '0341234567', 'Ankaraobato'),
(2, 'Rasoa', 'Mialy', 'mialy@example.com', '0349876543', 'Ankaraobato');

-- --------------------------------------------------------

--
-- Table structure for table `Commande`
--

CREATE TABLE `Commande` (
  `commande_id` int(11) NOT NULL,
  `client_id` int(11) NOT NULL,
  `date_commande` datetime DEFAULT current_timestamp(),
  `montant_total` decimal(12,2) NOT NULL,
  `status_id` int(11) DEFAULT NULL,
  `livreur_assign` int(11) DEFAULT NULL,
  `lieux_livraison` varchar(100) DEFAULT NULL,
  `latitude` varchar(45) DEFAULT NULL,
  `longitude` varchar(45) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `Commande`
--

INSERT INTO `Commande` (`commande_id`, `client_id`, `date_commande`, `montant_total`, `status_id`, `livreur_assign`, `lieux_livraison`, `latitude`, `longitude`) VALUES
(1, 1, '2025-08-20 09:22:04', 75000.00, 1, NULL, NULL, NULL, NULL),
(2, 1, '2025-09-01 00:00:00', 14500.00, 1, NULL, NULL, NULL, NULL),
(3, 1, '2025-09-01 00:00:00', 14500.00, 1, NULL, NULL, NULL, NULL),
(4, 1, '2025-09-01 00:00:00', 14500.00, 1, NULL, NULL, NULL, NULL),
(5, 1, '2025-09-01 00:00:00', 14500.00, 1, NULL, NULL, NULL, NULL),
(6, 1, '2025-09-01 00:00:00', 14500.00, 1, NULL, NULL, NULL, NULL),
(7, 1, '2025-09-01 00:00:00', 14500.00, 1, NULL, NULL, NULL, NULL),
(8, 1, '2025-09-01 00:00:00', 14500.00, 1, NULL, NULL, NULL, NULL),
(9, 1, '2025-09-01 00:00:00', 14500.00, 1, NULL, NULL, NULL, NULL),
(10, 1, '2025-09-01 00:00:00', 14500.00, 1, NULL, NULL, NULL, NULL),
(11, 1, '2025-09-01 00:00:00', 14500.00, 1, NULL, NULL, NULL, NULL),
(12, 1, '2025-09-01 00:00:00', 14500.00, 1, NULL, NULL, NULL, NULL),
(13, 1, '2025-09-01 00:00:00', 14500.00, 1, NULL, NULL, NULL, NULL),
(14, 1, '2025-09-01 00:00:00', 14500.00, 1, NULL, NULL, NULL, NULL),
(15, 1, '2025-09-01 00:00:00', 14500.00, 1, NULL, NULL, NULL, NULL),
(16, 1, '2025-09-02 00:00:00', 14500.00, 2, 1, NULL, NULL, NULL),
(17, 1, '2025-09-10 00:00:00', 400.00, 1, 0, 'Andranomena', '-18.8792', '47.5079');

-- --------------------------------------------------------

--
-- Table structure for table `Commercant`
--

CREATE TABLE `Commercant` (
  `commercant_id` int(11) NOT NULL,
  `nom` varchar(150) NOT NULL,
  `description` text DEFAULT NULL,
  `adresse` varchar(200) DEFAULT NULL,
  `telephone` varchar(50) DEFAULT NULL,
  `email` varchar(150) DEFAULT NULL,
  `latitude` varchar(45) DEFAULT NULL,
  `longitude` varchar(45) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `Commercant`
--

INSERT INTO `Commercant` (`commercant_id`, `nom`, `description`, `adresse`, `telephone`, `email`, `latitude`, `longitude`) VALUES
(1, 'Boutique A', 'Vêtements', 'Antananarivo', '0321234567', 'contact@boutiquea.mg', NULL, NULL),
(2, 'Boutique B', 'Accessoires', 'Antsirabe', '0329876543', 'contact@boutiqueb.mg', NULL, NULL);

-- --------------------------------------------------------

--
-- Table structure for table `Livraison`
--

CREATE TABLE `Livraison` (
  `livraison_id` int(11) NOT NULL,
  `commande_id` int(11) NOT NULL,
  `livreur_id` int(11) NOT NULL,
  `article_id` int(11) NOT NULL,
  `client_id` int(11) NOT NULL,
  `commercant_id` int(11) NOT NULL,
  `date_prevue` datetime DEFAULT NULL,
  `date_effective` datetime DEFAULT NULL,
  `duree_estimee` int(11) DEFAULT NULL,
  `status` enum('en_attente','en_cours','livree','annulee','echouee') DEFAULT 'en_attente',
  `axe` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `Livraison`
--

INSERT INTO `Livraison` (`livraison_id`, `commande_id`, `livreur_id`, `article_id`, `client_id`, `commercant_id`, `date_prevue`, `date_effective`, `duree_estimee`, `status`, `axe`) VALUES
(1, 1, 1, 1, 1, 1, '2025-08-21 10:00:00', NULL, 60, 'en_attente', 'Antananarivo-Centre'),
(2, 1, 1, 3, 1, 2, '2025-08-21 10:30:00', NULL, 60, 'en_attente', 'Antananarivo-Centre');

-- --------------------------------------------------------

--
-- Table structure for table `Livreur`
--

CREATE TABLE `Livreur` (
  `livreur_id` int(11) NOT NULL,
  `nom` varchar(100) NOT NULL,
  `telephone` varchar(50) DEFAULT NULL,
  `vehicule` varchar(50) DEFAULT NULL,
  `zone_livraison` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `Livreur`
--

INSERT INTO `Livreur` (`livreur_id`, `nom`, `telephone`, `vehicule`, `zone_livraison`) VALUES
(1, 'Rajaonarison Tiana', '0331234567', 'Moto', 'Antananarivo');

-- --------------------------------------------------------

--
-- Table structure for table `orders`
--

CREATE TABLE `orders` (
  `id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `status_id` int(11) NOT NULL,
  `total_amount` decimal(12,2) NOT NULL DEFAULT 0.00,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `order_items`
--

CREATE TABLE `order_items` (
  `id` bigint(20) NOT NULL,
  `order_id` bigint(20) NOT NULL,
  `product_id` bigint(20) NOT NULL,
  `product_name` varchar(255) DEFAULT NULL,
  `quantity` int(11) NOT NULL,
  `unit_price` decimal(12,2) NOT NULL,
  `total_price` decimal(12,2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `order_status`
--

CREATE TABLE `order_status` (
  `id` int(11) NOT NULL,
  `code` varchar(50) NOT NULL,
  `label` varchar(100) NOT NULL,
  `is_final` tinyint(1) DEFAULT 0,
  `created_at` timestamp NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Dumping data for table `order_status`
--

INSERT INTO `order_status` (`id`, `code`, `label`, `is_final`, `created_at`) VALUES
(1, 'cart', 'Panier', 0, '2025-10-24 13:48:29'),
(2, 'pending_payment', 'En attente de paiement', 0, '2025-10-24 13:48:29'),
(3, 'paid', 'Payée', 0, '2025-10-24 13:48:29'),
(4, 'processing', 'En traitement', 0, '2025-10-24 13:48:29'),
(5, 'shipped', 'Expédiée', 0, '2025-10-24 13:48:29'),
(6, 'delivered', 'Livrée', 1, '2025-10-24 13:48:29'),
(7, 'cancelled', 'Annulée', 1, '2025-10-24 13:48:29'),
(8, 'refunded', 'Remboursée', 1, '2025-10-24 13:48:29');

-- --------------------------------------------------------

--
-- Table structure for table `Panier`
--

CREATE TABLE `Panier` (
  `panier_id` int(11) NOT NULL,
  `client_id` int(11) NOT NULL,
  `date_creation` datetime DEFAULT current_timestamp(),
  `status_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `Panier`
--

INSERT INTO `Panier` (`panier_id`, `client_id`, `date_creation`, `status_id`) VALUES
(1, 1, '2025-08-20 09:22:04', 1),
(5, 1, '2025-08-22 15:01:38', 4),
(6, 1, '2025-08-22 15:03:22', 4);

-- --------------------------------------------------------

--
-- Table structure for table `Panier_Article`
--

CREATE TABLE `Panier_Article` (
  `panier_id` int(11) NOT NULL,
  `article_id` int(11) NOT NULL,
  `quantite` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `Panier_Article`
--

INSERT INTO `Panier_Article` (`panier_id`, `article_id`, `quantite`) VALUES
(1, 1, 2),
(1, 3, 1),
(5, 1, 2),
(6, 1, 2);

-- --------------------------------------------------------

--
-- Table structure for table `payments`
--

CREATE TABLE `payments` (
  `id` bigint(20) NOT NULL,
  `order_id` bigint(20) NOT NULL,
  `method_id` int(11) NOT NULL,
  `amount` decimal(12,2) NOT NULL,
  `status` enum('initiated','completed','failed','refunded') DEFAULT 'initiated',
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `payment_method`
--

CREATE TABLE `payment_method` (
  `id` int(11) NOT NULL,
  `code` varchar(50) NOT NULL,
  `name` varchar(100) NOT NULL,
  `description` text DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT 1,
  `created_at` timestamp NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Dumping data for table `payment_method`
--

INSERT INTO `payment_method` (`id`, `code`, `name`, `description`, `is_active`, `created_at`) VALUES
(1, 'stripe', 'Paiement carte via Stripe', NULL, 1, '2025-10-24 13:48:28'),
(2, 'paypal', 'Paiement PayPal', NULL, 1, '2025-10-24 13:48:28'),
(3, 'cash', 'Paiement à la livraison', NULL, 1, '2025-10-24 13:48:28'),
(4, 'wallet', 'Portefeuille interne', NULL, 1, '2025-10-24 13:48:28');

-- --------------------------------------------------------

--
-- Table structure for table `payment_transaction`
--

CREATE TABLE `payment_transaction` (
  `id` bigint(20) NOT NULL,
  `payment_id` bigint(20) NOT NULL,
  `transaction_type` enum('authorization','capture','refund','void') DEFAULT 'capture',
  `transaction_reference` varchar(255) DEFAULT NULL,
  `amount` decimal(12,2) NOT NULL,
  `currency` char(3) DEFAULT 'EUR',
  `status` enum('pending','success','failed') DEFAULT 'pending',
  `raw_response` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`raw_response`)),
  `created_at` timestamp NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `profil`
--

CREATE TABLE `profil` (
  `idProfil` int(11) NOT NULL,
  `nomProfil` varchar(100) NOT NULL,
  `description` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `profil`
--

INSERT INTO `profil` (`idProfil`, `nomProfil`, `description`) VALUES
(1, 'client', 'Profil pour les clients'),
(2, 'commerçant', 'Profil pour les commerçants'),
(3, 'livreur', 'Profil pour les livreurs');

-- --------------------------------------------------------

--
-- Table structure for table `Promotion`
--

CREATE TABLE `Promotion` (
  `promotion_id` int(11) NOT NULL,
  `article_id` int(11) NOT NULL,
  `code` varchar(50) NOT NULL,
  `reduction` decimal(5,2) NOT NULL,
  `date_debut` datetime NOT NULL,
  `date_fin` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `Promotion`
--

INSERT INTO `Promotion` (`promotion_id`, `article_id`, `code`, `reduction`, `date_debut`, `date_fin`) VALUES
(1, 1, 'PROMO10', 10.00, '2025-08-20 00:00:00', '2025-08-31 23:59:59'),
(2, 3, 'PROMO15', 15.00, '2025-08-20 00:00:00', '2025-08-25 23:59:59');

-- --------------------------------------------------------

--
-- Table structure for table `status`
--

CREATE TABLE `status` (
  `id_status` int(11) NOT NULL,
  `nom_status` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `status`
--

INSERT INTO `status` (`id_status`, `nom_status`) VALUES
(2, 'commande en cours'),
(3, 'commande fermer'),
(1, 'commande ouvert'),
(5, 'panier fermer'),
(4, 'panier ouvert');

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `login` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `name` varchar(100) NOT NULL,
  `lastname` varchar(100) NOT NULL,
  `type` varchar(50) NOT NULL,
  `contact` varchar(20) NOT NULL,
  `mail` varchar(150) NOT NULL,
  `adresse` text NOT NULL,
  `latitude` varchar(45) DEFAULT NULL,
  `longitude` varchar(45) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `login`, `password`, `name`, `lastname`, `type`, `contact`, `mail`, `adresse`, `latitude`, `longitude`) VALUES
(1, 'hery', 'h', 'Rasolonjatovo', 'Hery', '', '12345', 'hery@gmail.mg', 'ffdqfqfq', '-18.8692', '47.5079'),
(3, 'kama', 'kely', 'kmajao', 'makaina', 'test', '12345', 'kama@gmail.com', 'anakrao', '-18.879', '47.508'),
(4, 'testdddddddddd', 'testddddddddd', 'testname', 'testlast', 'test', '1234', 'mail@gmail.com', 'adressetna', '-18.8792', '47.5180');

-- --------------------------------------------------------

--
-- Table structure for table `userProfil`
--

CREATE TABLE `userProfil` (
  `idUser` int(11) NOT NULL,
  `idProfil` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `userProfil`
--

INSERT INTO `userProfil` (`idUser`, `idProfil`) VALUES
(1, 2),
(3, 2),
(4, 2);

-- --------------------------------------------------------

--
-- Table structure for table `wallet`
--

CREATE TABLE `wallet` (
  `id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `balance` decimal(12,2) NOT NULL DEFAULT 0.00,
  `currency` char(3) DEFAULT 'EUR',
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `wallet_transaction`
--

CREATE TABLE `wallet_transaction` (
  `id` bigint(20) NOT NULL,
  `wallet_id` bigint(20) NOT NULL,
  `transaction_type` enum('credit','debit','refund','adjust') NOT NULL,
  `amount` decimal(12,2) NOT NULL,
  `reference` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `Adresse`
--
ALTER TABLE `Adresse`
  ADD PRIMARY KEY (`adresse_id`),
  ADD KEY `fk_adresse_client` (`client_id`);

--
-- Indexes for table `Article`
--
ALTER TABLE `Article`
  ADD PRIMARY KEY (`article_id`),
  ADD KEY `commercant_id` (`commercant_id`),
  ADD KEY `categorie_id` (`categorie_id`);

--
-- Indexes for table `Article_Image`
--
ALTER TABLE `Article_Image`
  ADD PRIMARY KEY (`image_id`),
  ADD KEY `article_id` (`article_id`);

--
-- Indexes for table `Avis`
--
ALTER TABLE `Avis`
  ADD PRIMARY KEY (`avis_id`),
  ADD KEY `client_id` (`client_id`),
  ADD KEY `article_id` (`article_id`),
  ADD KEY `commercant_id` (`commercant_id`),
  ADD KEY `livreur_id` (`livreur_id`);

--
-- Indexes for table `Categorie`
--
ALTER TABLE `Categorie`
  ADD PRIMARY KEY (`categorie_id`),
  ADD KEY `parent_id` (`parent_id`);

--
-- Indexes for table `Client`
--
ALTER TABLE `Client`
  ADD PRIMARY KEY (`client_id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Indexes for table `Commande`
--
ALTER TABLE `Commande`
  ADD PRIMARY KEY (`commande_id`),
  ADD KEY `client_id` (`client_id`);

--
-- Indexes for table `Commercant`
--
ALTER TABLE `Commercant`
  ADD PRIMARY KEY (`commercant_id`);

--
-- Indexes for table `Livraison`
--
ALTER TABLE `Livraison`
  ADD PRIMARY KEY (`livraison_id`),
  ADD KEY `commande_id` (`commande_id`),
  ADD KEY `livreur_id` (`livreur_id`),
  ADD KEY `article_id` (`article_id`),
  ADD KEY `client_id` (`client_id`),
  ADD KEY `commercant_id` (`commercant_id`);

--
-- Indexes for table `Livreur`
--
ALTER TABLE `Livreur`
  ADD PRIMARY KEY (`livreur_id`);

--
-- Indexes for table `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `status_id` (`status_id`);

--
-- Indexes for table `order_items`
--
ALTER TABLE `order_items`
  ADD PRIMARY KEY (`id`),
  ADD KEY `order_id` (`order_id`);

--
-- Indexes for table `order_status`
--
ALTER TABLE `order_status`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code` (`code`);

--
-- Indexes for table `Panier`
--
ALTER TABLE `Panier`
  ADD PRIMARY KEY (`panier_id`),
  ADD KEY `client_id` (`client_id`);

--
-- Indexes for table `Panier_Article`
--
ALTER TABLE `Panier_Article`
  ADD PRIMARY KEY (`panier_id`,`article_id`),
  ADD KEY `article_id` (`article_id`);

--
-- Indexes for table `payments`
--
ALTER TABLE `payments`
  ADD PRIMARY KEY (`id`),
  ADD KEY `order_id` (`order_id`),
  ADD KEY `method_id` (`method_id`);

--
-- Indexes for table `payment_method`
--
ALTER TABLE `payment_method`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code` (`code`);

--
-- Indexes for table `payment_transaction`
--
ALTER TABLE `payment_transaction`
  ADD PRIMARY KEY (`id`),
  ADD KEY `payment_id` (`payment_id`);

--
-- Indexes for table `profil`
--
ALTER TABLE `profil`
  ADD PRIMARY KEY (`idProfil`);

--
-- Indexes for table `Promotion`
--
ALTER TABLE `Promotion`
  ADD PRIMARY KEY (`promotion_id`),
  ADD KEY `article_id` (`article_id`);

--
-- Indexes for table `status`
--
ALTER TABLE `status`
  ADD PRIMARY KEY (`id_status`),
  ADD UNIQUE KEY `nom_status` (`nom_status`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `login` (`login`);

--
-- Indexes for table `userProfil`
--
ALTER TABLE `userProfil`
  ADD PRIMARY KEY (`idUser`,`idProfil`),
  ADD KEY `idProfil` (`idProfil`);

--
-- Indexes for table `wallet`
--
ALTER TABLE `wallet`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`);

--
-- Indexes for table `wallet_transaction`
--
ALTER TABLE `wallet_transaction`
  ADD PRIMARY KEY (`id`),
  ADD KEY `wallet_id` (`wallet_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `Adresse`
--
ALTER TABLE `Adresse`
  MODIFY `adresse_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `Article`
--
ALTER TABLE `Article`
  MODIFY `article_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT for table `Article_Image`
--
ALTER TABLE `Article_Image`
  MODIFY `image_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `Avis`
--
ALTER TABLE `Avis`
  MODIFY `avis_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `Categorie`
--
ALTER TABLE `Categorie`
  MODIFY `categorie_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `Client`
--
ALTER TABLE `Client`
  MODIFY `client_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `Commande`
--
ALTER TABLE `Commande`
  MODIFY `commande_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT for table `Commercant`
--
ALTER TABLE `Commercant`
  MODIFY `commercant_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `Livraison`
--
ALTER TABLE `Livraison`
  MODIFY `livraison_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `Livreur`
--
ALTER TABLE `Livreur`
  MODIFY `livreur_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `orders`
--
ALTER TABLE `orders`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `order_items`
--
ALTER TABLE `order_items`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `order_status`
--
ALTER TABLE `order_status`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `Panier`
--
ALTER TABLE `Panier`
  MODIFY `panier_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `payments`
--
ALTER TABLE `payments`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `payment_method`
--
ALTER TABLE `payment_method`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `payment_transaction`
--
ALTER TABLE `payment_transaction`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `profil`
--
ALTER TABLE `profil`
  MODIFY `idProfil` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `Promotion`
--
ALTER TABLE `Promotion`
  MODIFY `promotion_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `status`
--
ALTER TABLE `status`
  MODIFY `id_status` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `wallet`
--
ALTER TABLE `wallet`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `wallet_transaction`
--
ALTER TABLE `wallet_transaction`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `Adresse`
--
ALTER TABLE `Adresse`
  ADD CONSTRAINT `fk_adresse_client` FOREIGN KEY (`client_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `Article`
--
ALTER TABLE `Article`
  ADD CONSTRAINT `Article_ibfk_1` FOREIGN KEY (`commercant_id`) REFERENCES `Commercant` (`commercant_id`),
  ADD CONSTRAINT `Article_ibfk_2` FOREIGN KEY (`categorie_id`) REFERENCES `Categorie` (`categorie_id`);

--
-- Constraints for table `Article_Image`
--
ALTER TABLE `Article_Image`
  ADD CONSTRAINT `Article_Image_ibfk_1` FOREIGN KEY (`article_id`) REFERENCES `Article` (`article_id`);

--
-- Constraints for table `Avis`
--
ALTER TABLE `Avis`
  ADD CONSTRAINT `Avis_ibfk_1` FOREIGN KEY (`client_id`) REFERENCES `Client` (`client_id`),
  ADD CONSTRAINT `Avis_ibfk_2` FOREIGN KEY (`article_id`) REFERENCES `Article` (`article_id`),
  ADD CONSTRAINT `Avis_ibfk_3` FOREIGN KEY (`commercant_id`) REFERENCES `Commercant` (`commercant_id`),
  ADD CONSTRAINT `Avis_ibfk_4` FOREIGN KEY (`livreur_id`) REFERENCES `Livreur` (`livreur_id`);

--
-- Constraints for table `Categorie`
--
ALTER TABLE `Categorie`
  ADD CONSTRAINT `Categorie_ibfk_1` FOREIGN KEY (`parent_id`) REFERENCES `Categorie` (`categorie_id`);

--
-- Constraints for table `Commande`
--
ALTER TABLE `Commande`
  ADD CONSTRAINT `Commande_ibfk_1` FOREIGN KEY (`client_id`) REFERENCES `Client` (`client_id`);

--
-- Constraints for table `Livraison`
--
ALTER TABLE `Livraison`
  ADD CONSTRAINT `Livraison_ibfk_1` FOREIGN KEY (`commande_id`) REFERENCES `Commande` (`commande_id`),
  ADD CONSTRAINT `Livraison_ibfk_2` FOREIGN KEY (`livreur_id`) REFERENCES `Livreur` (`livreur_id`),
  ADD CONSTRAINT `Livraison_ibfk_3` FOREIGN KEY (`article_id`) REFERENCES `Article` (`article_id`),
  ADD CONSTRAINT `Livraison_ibfk_4` FOREIGN KEY (`client_id`) REFERENCES `Client` (`client_id`),
  ADD CONSTRAINT `Livraison_ibfk_5` FOREIGN KEY (`commercant_id`) REFERENCES `Commercant` (`commercant_id`);

--
-- Constraints for table `orders`
--
ALTER TABLE `orders`
  ADD CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  ADD CONSTRAINT `orders_ibfk_2` FOREIGN KEY (`status_id`) REFERENCES `order_status` (`id`);

--
-- Constraints for table `order_items`
--
ALTER TABLE `order_items`
  ADD CONSTRAINT `order_items_ibfk_1` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`);

--
-- Constraints for table `Panier`
--
ALTER TABLE `Panier`
  ADD CONSTRAINT `Panier_ibfk_1` FOREIGN KEY (`client_id`) REFERENCES `Client` (`client_id`);

--
-- Constraints for table `Panier_Article`
--
ALTER TABLE `Panier_Article`
  ADD CONSTRAINT `Panier_Article_ibfk_1` FOREIGN KEY (`panier_id`) REFERENCES `Panier` (`panier_id`),
  ADD CONSTRAINT `Panier_Article_ibfk_2` FOREIGN KEY (`article_id`) REFERENCES `Article` (`article_id`);

--
-- Constraints for table `payments`
--
ALTER TABLE `payments`
  ADD CONSTRAINT `payments_ibfk_1` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`),
  ADD CONSTRAINT `payments_ibfk_2` FOREIGN KEY (`method_id`) REFERENCES `payment_method` (`id`);

--
-- Constraints for table `payment_transaction`
--
ALTER TABLE `payment_transaction`
  ADD CONSTRAINT `payment_transaction_ibfk_1` FOREIGN KEY (`payment_id`) REFERENCES `payments` (`id`);

--
-- Constraints for table `Promotion`
--
ALTER TABLE `Promotion`
  ADD CONSTRAINT `Promotion_ibfk_1` FOREIGN KEY (`article_id`) REFERENCES `Article` (`article_id`);

--
-- Constraints for table `userProfil`
--
ALTER TABLE `userProfil`
  ADD CONSTRAINT `userProfil_ibfk_1` FOREIGN KEY (`idUser`) REFERENCES `user` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `userProfil_ibfk_2` FOREIGN KEY (`idProfil`) REFERENCES `profil` (`idProfil`) ON DELETE CASCADE;

--
-- Constraints for table `wallet`
--
ALTER TABLE `wallet`
  ADD CONSTRAINT `wallet_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

--
-- Constraints for table `wallet_transaction`
--
ALTER TABLE `wallet_transaction`
  ADD CONSTRAINT `wallet_transaction_ibfk_1` FOREIGN KEY (`wallet_id`) REFERENCES `wallet` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
