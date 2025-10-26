-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1
-- Généré le :  Dim 26 oct. 2025 à 12:52
-- Version du serveur :  10.1.38-MariaDB
-- Version de PHP :  5.6.40

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données :  `livraison`
--

-- --------------------------------------------------------

--
-- Structure de la table `adresse`
--

CREATE TABLE `adresse` (
  `adresse_id` int(11) NOT NULL,
  `client_id` int(11) NOT NULL,
  `rue` varchar(150) DEFAULT NULL,
  `ville` varchar(100) DEFAULT NULL,
  `code_postal` varchar(20) DEFAULT NULL,
  `pays` varchar(50) DEFAULT NULL,
  `latitude` varchar(255) DEFAULT NULL,
  `longitude` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `adresse`
--

INSERT INTO `adresse` (`adresse_id`, `client_id`, `rue`, `ville`, `code_postal`, `pays`, `latitude`, `longitude`) VALUES
(21, 1, '', '', '', '', '-18.891345331705605', '47.5275502505808'),
(22, 1, '', '', '', '', '-18.88208719207098', '47.51999400965967'),
(23, 1, '', '', '', '', '-18.890208119563212', '47.52942766293057');

-- --------------------------------------------------------

--
-- Structure de la table `article`
--

CREATE TABLE `article` (
  `article_id` int(11) NOT NULL,
  `nom` varchar(150) NOT NULL,
  `description` text,
  `prix` decimal(10,2) NOT NULL,
  `stock` int(11) DEFAULT '0',
  `commercant_id` int(11) NOT NULL,
  `categorie_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `article`
--

INSERT INTO `article` (`article_id`, `nom`, `description`, `prix`, `stock`, `commercant_id`, `categorie_id`) VALUES
(1, 'telephone', 'telephone', '15000.00', 50, 1, 1),
(2, 'telephone', 'telephone', '35000.00', 30, 1, 1),
(3, 'telephone', 'telephone', '45000.00', 20, 2, 2),
(4, 'telephone', 'telephone', '60000.00', 15, 1, 3),
(8, 'telephone', 'telephone', '123.00', 5, 1, 1),
(10, 'telephone', 'telephone', '123.00', 5, 1, 1),
(11, 'telephone', 'telephone', '123.00', 5, 1, 1),
(14, 'telephone', 'telephone', '323.00', 2, 1, 1),
(15, 'telephone', 'telephone', '323.00', 2, 1, 1);

-- --------------------------------------------------------

--
-- Structure de la table `article_image`
--

CREATE TABLE `article_image` (
  `image_id` int(11) NOT NULL,
  `article_id` int(11) NOT NULL,
  `url` varchar(255) NOT NULL,
  `largeur` int(11) DEFAULT NULL,
  `hauteur` int(11) DEFAULT NULL,
  `ordre` int(11) DEFAULT '0',
  `type` enum('main','gallery','thumbnail') DEFAULT 'gallery',
  `taille` varchar(45) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `article_image`
--

INSERT INTO `article_image` (`image_id`, `article_id`, `url`, `largeur`, `hauteur`, `ordre`, `type`, `taille`) VALUES
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
-- Structure de la table `avis`
--

CREATE TABLE `avis` (
  `avis_id` int(11) NOT NULL,
  `client_id` int(11) NOT NULL,
  `article_id` int(11) DEFAULT NULL,
  `commercant_id` int(11) DEFAULT NULL,
  `livreur_id` int(11) DEFAULT NULL,
  `note` int(11) DEFAULT NULL,
  `commentaire` text,
  `date_avis` datetime DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `avis`
--

INSERT INTO `avis` (`avis_id`, `client_id`, `article_id`, `commercant_id`, `livreur_id`, `note`, `commentaire`, `date_avis`) VALUES
(1, 1, 1, 1, 1, 5, 'Très bon produit, livraison rapide!', '2025-08-20 09:22:06'),
(2, 2, 3, 2, NULL, 4, 'Sac correct, qualité satisfaisante.', '2025-08-20 09:22:06');

-- --------------------------------------------------------

--
-- Structure de la table `categorie`
--

CREATE TABLE `categorie` (
  `categorie_id` int(11) NOT NULL,
  `nom` varchar(100) NOT NULL,
  `parent_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `categorie`
--

INSERT INTO `categorie` (`categorie_id`, `nom`, `parent_id`) VALUES
(1, 'Vêtements', NULL),
(2, 'Accessoires', NULL),
(3, 'Chaussures', 1);

-- --------------------------------------------------------

--
-- Structure de la table `client`
--

CREATE TABLE `client` (
  `client_id` int(11) NOT NULL,
  `nom` varchar(100) NOT NULL,
  `prenom` varchar(100) DEFAULT NULL,
  `email` varchar(150) NOT NULL,
  `telephone` varchar(50) DEFAULT NULL,
  `adresse` varchar(45) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `client`
--

INSERT INTO `client` (`client_id`, `nom`, `prenom`, `email`, `telephone`, `adresse`) VALUES
(1, 'Rakoto', 'Andry', 'andry@example.com', '0341234567', 'Ankaraobato'),
(2, 'Rasoa', 'Mialy', 'mialy@example.com', '0349876543', 'Ankaraobato');

-- --------------------------------------------------------

--
-- Structure de la table `commande`
--

CREATE TABLE `commande` (
  `commande_id` int(11) NOT NULL,
  `client_id` int(11) NOT NULL,
  `date_commande` datetime DEFAULT CURRENT_TIMESTAMP,
  `montant_total` decimal(12,2) NOT NULL,
  `status_id` int(11) DEFAULT NULL,
  `livreur_assign` int(11) DEFAULT NULL,
  `lieux_livraison` varchar(100) DEFAULT NULL,
  `latitude` varchar(45) DEFAULT NULL,
  `longitude` varchar(45) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `commande`
--

INSERT INTO `commande` (`commande_id`, `client_id`, `date_commande`, `montant_total`, `status_id`, `livreur_assign`, `lieux_livraison`, `latitude`, `longitude`) VALUES
(1, 1, '2025-08-20 09:22:04', '75000.00', 1, NULL, NULL, NULL, NULL),
(2, 1, '2025-09-01 00:00:00', '14500.00', 1, NULL, NULL, NULL, NULL),
(3, 1, '2025-09-01 00:00:00', '14500.00', 1, NULL, NULL, NULL, NULL),
(4, 1, '2025-09-01 00:00:00', '14500.00', 1, NULL, NULL, NULL, NULL),
(5, 1, '2025-09-01 00:00:00', '14500.00', 1, NULL, NULL, NULL, NULL),
(6, 1, '2025-09-01 00:00:00', '14500.00', 1, NULL, NULL, NULL, NULL),
(7, 1, '2025-09-01 00:00:00', '14500.00', 1, NULL, NULL, NULL, NULL),
(8, 1, '2025-09-01 00:00:00', '14500.00', 1, NULL, NULL, NULL, NULL),
(9, 1, '2025-09-01 00:00:00', '14500.00', 1, NULL, NULL, NULL, NULL),
(10, 1, '2025-09-01 00:00:00', '14500.00', 1, NULL, NULL, NULL, NULL),
(11, 1, '2025-09-01 00:00:00', '14500.00', 1, NULL, NULL, NULL, NULL),
(12, 1, '2025-09-01 00:00:00', '14500.00', 1, NULL, NULL, NULL, NULL),
(13, 1, '2025-09-01 00:00:00', '14500.00', 1, NULL, NULL, NULL, NULL),
(14, 1, '2025-09-01 00:00:00', '14500.00', 1, NULL, NULL, NULL, NULL),
(15, 1, '2025-09-01 00:00:00', '14500.00', 1, NULL, NULL, NULL, NULL),
(16, 1, '2025-09-02 00:00:00', '14500.00', 2, 1, NULL, NULL, NULL),
(17, 1, '2025-09-10 00:00:00', '400.00', 1, 0, 'Andranomena', '-18.8792', '47.5079');

-- --------------------------------------------------------

--
-- Structure de la table `commercant`
--

CREATE TABLE `commercant` (
  `commercant_id` int(11) NOT NULL,
  `nom` varchar(150) NOT NULL,
  `description` text,
  `adresse` varchar(200) DEFAULT NULL,
  `telephone` varchar(50) DEFAULT NULL,
  `email` varchar(150) DEFAULT NULL,
  `latitude` varchar(45) DEFAULT NULL,
  `longitude` varchar(45) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `commercant`
--

INSERT INTO `commercant` (`commercant_id`, `nom`, `description`, `adresse`, `telephone`, `email`, `latitude`, `longitude`) VALUES
(1, 'Boutique A', 'Vêtements', 'Antananarivo', '0321234567', 'contact@boutiquea.mg', NULL, NULL),
(2, 'Boutique B', 'Accessoires', 'Antsirabe', '0329876543', 'contact@boutiqueb.mg', NULL, NULL);

-- --------------------------------------------------------

--
-- Structure de la table `livraison`
--

CREATE TABLE `livraison` (
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `livraison`
--

INSERT INTO `livraison` (`livraison_id`, `commande_id`, `livreur_id`, `article_id`, `client_id`, `commercant_id`, `date_prevue`, `date_effective`, `duree_estimee`, `status`, `axe`) VALUES
(1, 1, 1, 1, 1, 1, '2025-08-21 10:00:00', NULL, 60, 'en_attente', 'Antananarivo-Centre'),
(2, 1, 1, 3, 1, 2, '2025-08-21 10:30:00', NULL, 60, 'en_attente', 'Antananarivo-Centre');

-- --------------------------------------------------------

--
-- Structure de la table `livreur`
--

CREATE TABLE `livreur` (
  `livreur_id` int(11) NOT NULL,
  `nom` varchar(100) NOT NULL,
  `telephone` varchar(50) DEFAULT NULL,
  `vehicule` varchar(50) DEFAULT NULL,
  `zone_livraison` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `livreur`
--

INSERT INTO `livreur` (`livreur_id`, `nom`, `telephone`, `vehicule`, `zone_livraison`) VALUES
(1, 'Rajaonarison Tiana', '0331234567', 'Moto', 'Antananarivo');

-- --------------------------------------------------------

--
-- Structure de la table `orders`
--

CREATE TABLE `orders` (
  `id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `status_id` int(11) NOT NULL,
  `total_amount` decimal(12,2) NOT NULL DEFAULT '0.00',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `orders`
--

INSERT INTO `orders` (`id`, `user_id`, `status_id`, `total_amount`, `created_at`, `updated_at`) VALUES
(17, 1, 3, '165000.00', '2025-10-26 06:39:34', '2025-10-26 06:39:34'),
(18, 1, 3, '646.00', '2025-10-26 06:49:00', '2025-10-26 06:49:01'),
(19, 1, 3, '646.00', '2025-10-26 11:43:51', '2025-10-26 11:43:52');

-- --------------------------------------------------------

--
-- Structure de la table `order_items`
--

CREATE TABLE `order_items` (
  `id` bigint(20) NOT NULL,
  `order_id` bigint(20) NOT NULL,
  `article_id` bigint(20) NOT NULL,
  `product_name` varchar(255) DEFAULT NULL,
  `quantity` int(11) NOT NULL,
  `unit_price` decimal(12,2) NOT NULL,
  `total_price` decimal(12,2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `order_items`
--

INSERT INTO `order_items` (`id`, `order_id`, `article_id`, `product_name`, `quantity`, `unit_price`, `total_price`) VALUES
(29, 17, 3, NULL, 1, '45000.00', '45000.00'),
(30, 17, 4, NULL, 2, '60000.00', '120000.00'),
(31, 18, 14, NULL, 1, '323.00', '323.00'),
(32, 18, 15, NULL, 1, '323.00', '323.00'),
(33, 19, 14, NULL, 1, '323.00', '323.00'),
(34, 19, 15, NULL, 1, '323.00', '323.00');

-- --------------------------------------------------------

--
-- Structure de la table `order_status`
--

CREATE TABLE `order_status` (
  `id` int(11) NOT NULL,
  `code` varchar(50) NOT NULL,
  `label` varchar(100) NOT NULL,
  `is_final` tinyint(1) DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `order_status`
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
-- Structure de la table `panier`
--

CREATE TABLE `panier` (
  `panier_id` int(11) NOT NULL,
  `client_id` int(11) NOT NULL,
  `date_creation` datetime DEFAULT CURRENT_TIMESTAMP,
  `status_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `panier`
--

INSERT INTO `panier` (`panier_id`, `client_id`, `date_creation`, `status_id`) VALUES
(1, 1, '2025-08-20 09:22:04', 1),
(5, 1, '2025-08-22 15:01:38', 4),
(6, 1, '2025-08-22 15:03:22', 4);

-- --------------------------------------------------------

--
-- Structure de la table `panier_article`
--

CREATE TABLE `panier_article` (
  `panier_id` int(11) NOT NULL,
  `article_id` int(11) NOT NULL,
  `quantite` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `panier_article`
--

INSERT INTO `panier_article` (`panier_id`, `article_id`, `quantite`) VALUES
(1, 1, 2),
(1, 3, 1),
(5, 1, 2),
(6, 1, 2);

-- --------------------------------------------------------

--
-- Structure de la table `payments`
--

CREATE TABLE `payments` (
  `id` bigint(20) NOT NULL,
  `order_id` bigint(20) NOT NULL,
  `method_id` int(11) NOT NULL,
  `amount` decimal(12,2) NOT NULL,
  `status` enum('initiated','completed','failed','refunded') DEFAULT 'initiated',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `payments`
--

INSERT INTO `payments` (`id`, `order_id`, `method_id`, `amount`, `status`, `created_at`, `updated_at`) VALUES
(4, 17, 4, '165000.00', 'completed', '2025-10-26 06:39:34', '2025-10-26 06:39:34'),
(5, 18, 4, '646.00', 'completed', '2025-10-26 06:49:00', '2025-10-26 06:49:00'),
(6, 19, 4, '646.00', 'completed', '2025-10-26 08:43:51', '2025-10-26 08:43:51');

-- --------------------------------------------------------

--
-- Structure de la table `payment_method`
--

CREATE TABLE `payment_method` (
  `id` int(11) NOT NULL,
  `code` varchar(50) NOT NULL,
  `name` varchar(100) NOT NULL,
  `description` text,
  `is_active` tinyint(1) DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `payment_method`
--

INSERT INTO `payment_method` (`id`, `code`, `name`, `description`, `is_active`, `created_at`) VALUES
(1, 'stripe', 'Paiement carte via Stripe', NULL, 1, '2025-10-24 16:48:28'),
(2, 'paypal', 'Paiement PayPal', NULL, 1, '2025-10-24 16:48:28'),
(3, 'cash', 'Paiement à la livraison', NULL, 1, '2025-10-24 16:48:28'),
(4, 'wallet', 'Portefeuille interne', NULL, 1, '2025-10-24 16:48:28');

-- --------------------------------------------------------

--
-- Structure de la table `payment_transaction`
--

CREATE TABLE `payment_transaction` (
  `id` bigint(20) NOT NULL,
  `payment_id` bigint(20) NOT NULL,
  `transaction_type` enum('authorization','capture','refund','void') DEFAULT 'capture',
  `transaction_reference` varchar(255) DEFAULT NULL,
  `amount` decimal(12,2) NOT NULL,
  `currency` char(3) DEFAULT 'EUR',
  `status` enum('pending','success','failed') DEFAULT 'pending',
  `raw_response` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `payment_transaction`
--

INSERT INTO `payment_transaction` (`id`, `payment_id`, `transaction_type`, `transaction_reference`, `amount`, `currency`, `status`, `raw_response`, `created_at`) VALUES
(4, 4, 'capture', '', '165000.00', 'Ar', 'success', '', '2025-10-26 06:39:34'),
(5, 5, 'capture', '', '646.00', 'Ar', 'success', '', '2025-10-26 06:49:00'),
(6, 6, 'capture', '', '646.00', 'Ar', 'success', '', '2025-10-26 08:43:52');

-- --------------------------------------------------------

--
-- Structure de la table `profil`
--

CREATE TABLE `profil` (
  `idProfil` int(11) NOT NULL,
  `nomProfil` varchar(100) NOT NULL,
  `description` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `profil`
--

INSERT INTO `profil` (`idProfil`, `nomProfil`, `description`) VALUES
(1, 'client', 'Profil pour les clients'),
(2, 'commerçant', 'Profil pour les commerçants'),
(3, 'livreur', 'Profil pour les livreurs');

-- --------------------------------------------------------

--
-- Structure de la table `promotion`
--

CREATE TABLE `promotion` (
  `promotion_id` int(11) NOT NULL,
  `article_id` int(11) NOT NULL,
  `code` varchar(50) NOT NULL,
  `reduction` decimal(5,2) NOT NULL,
  `date_debut` datetime NOT NULL,
  `date_fin` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `promotion`
--

INSERT INTO `promotion` (`promotion_id`, `article_id`, `code`, `reduction`, `date_debut`, `date_fin`) VALUES
(1, 1, 'PROMO10', '10.00', '2025-08-20 00:00:00', '2025-08-31 23:59:59'),
(2, 3, 'PROMO15', '15.00', '2025-08-20 00:00:00', '2025-08-25 23:59:59');

-- --------------------------------------------------------

--
-- Structure de la table `status`
--

CREATE TABLE `status` (
  `id_status` int(11) NOT NULL,
  `nom_status` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `status`
--

INSERT INTO `status` (`id_status`, `nom_status`) VALUES
(2, 'commande en cours'),
(3, 'commande fermer'),
(1, 'commande ouvert'),
(5, 'panier fermer'),
(4, 'panier ouvert');

-- --------------------------------------------------------

--
-- Structure de la table `user`
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `user`
--

INSERT INTO `user` (`id`, `login`, `password`, `name`, `lastname`, `type`, `contact`, `mail`, `adresse`, `latitude`, `longitude`) VALUES
(1, 'hery', 'h', 'Rasolonjatovo', 'Hery', '', '12345', 'hery@gmail.mg', 'ffdqfqfq', '-18.8692', '47.5079'),
(3, 'kama', 'kely', 'kmajao', 'makaina', 'test', '12345', 'kama@gmail.com', 'anakrao', '-18.879', '47.508'),
(4, 'testdddddddddd', 'testddddddddd', 'testname', 'testlast', 'test', '1234', 'mail@gmail.com', 'adressetna', '-18.8792', '47.5180');

-- --------------------------------------------------------

--
-- Structure de la table `userprofil`
--

CREATE TABLE `userprofil` (
  `idUser` int(11) NOT NULL,
  `idProfil` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `userprofil`
--

INSERT INTO `userprofil` (`idUser`, `idProfil`) VALUES
(1, 2),
(3, 2),
(4, 2);

-- --------------------------------------------------------

--
-- Structure de la table `wallet`
--

CREATE TABLE `wallet` (
  `id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `balance` decimal(12,2) NOT NULL DEFAULT '0.00',
  `currency` char(3) DEFAULT 'EUR',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `wallet`
--

INSERT INTO `wallet` (`id`, `user_id`, `balance`, `currency`, `updated_at`) VALUES
(1, 1, '667416.00', 'Ar', '2025-10-26 11:43:52');

-- --------------------------------------------------------

--
-- Structure de la table `wallet_transaction`
--

CREATE TABLE `wallet_transaction` (
  `id` bigint(20) NOT NULL,
  `wallet_id` bigint(20) NOT NULL,
  `transaction_type` enum('credit','debit','refund','adjust') NOT NULL,
  `amount` decimal(12,2) NOT NULL,
  `reference` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `wallet_transaction`
--

INSERT INTO `wallet_transaction` (`id`, `wallet_id`, `transaction_type`, `amount`, `reference`, `description`, `created_at`) VALUES
(7, 1, 'debit', '165000.00', 'order:', 'Paiement commande via wallet', '2025-10-26 06:39:34'),
(8, 1, 'debit', '165000.00', 'order_', 'Paiement commande', '2025-10-26 06:39:34'),
(9, 1, 'debit', '646.00', 'order:', 'Paiement commande via wallet', '2025-10-26 06:49:01'),
(10, 1, 'debit', '646.00', 'order_', 'Paiement commande', '2025-10-26 06:49:01'),
(11, 1, 'debit', '646.00', 'order:', 'Paiement commande via wallet', '2025-10-26 08:43:52'),
(12, 1, 'debit', '646.00', 'order_', 'Paiement commande', '2025-10-26 08:43:52');

--
-- Index pour les tables déchargées
--

--
-- Index pour la table `adresse`
--
ALTER TABLE `adresse`
  ADD PRIMARY KEY (`adresse_id`),
  ADD KEY `fk_adresse_client` (`client_id`);

--
-- Index pour la table `article`
--
ALTER TABLE `article`
  ADD PRIMARY KEY (`article_id`),
  ADD KEY `commercant_id` (`commercant_id`),
  ADD KEY `categorie_id` (`categorie_id`);

--
-- Index pour la table `article_image`
--
ALTER TABLE `article_image`
  ADD PRIMARY KEY (`image_id`),
  ADD KEY `article_id` (`article_id`);

--
-- Index pour la table `avis`
--
ALTER TABLE `avis`
  ADD PRIMARY KEY (`avis_id`),
  ADD KEY `client_id` (`client_id`),
  ADD KEY `article_id` (`article_id`),
  ADD KEY `commercant_id` (`commercant_id`),
  ADD KEY `livreur_id` (`livreur_id`);

--
-- Index pour la table `categorie`
--
ALTER TABLE `categorie`
  ADD PRIMARY KEY (`categorie_id`),
  ADD KEY `parent_id` (`parent_id`);

--
-- Index pour la table `client`
--
ALTER TABLE `client`
  ADD PRIMARY KEY (`client_id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Index pour la table `commande`
--
ALTER TABLE `commande`
  ADD PRIMARY KEY (`commande_id`),
  ADD KEY `client_id` (`client_id`);

--
-- Index pour la table `commercant`
--
ALTER TABLE `commercant`
  ADD PRIMARY KEY (`commercant_id`);

--
-- Index pour la table `livraison`
--
ALTER TABLE `livraison`
  ADD PRIMARY KEY (`livraison_id`),
  ADD KEY `commande_id` (`commande_id`),
  ADD KEY `livreur_id` (`livreur_id`),
  ADD KEY `article_id` (`article_id`),
  ADD KEY `client_id` (`client_id`),
  ADD KEY `commercant_id` (`commercant_id`);

--
-- Index pour la table `livreur`
--
ALTER TABLE `livreur`
  ADD PRIMARY KEY (`livreur_id`);

--
-- Index pour la table `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `status_id` (`status_id`);

--
-- Index pour la table `order_items`
--
ALTER TABLE `order_items`
  ADD PRIMARY KEY (`id`),
  ADD KEY `order_id` (`order_id`);

--
-- Index pour la table `order_status`
--
ALTER TABLE `order_status`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code` (`code`);

--
-- Index pour la table `panier`
--
ALTER TABLE `panier`
  ADD PRIMARY KEY (`panier_id`),
  ADD KEY `client_id` (`client_id`);

--
-- Index pour la table `panier_article`
--
ALTER TABLE `panier_article`
  ADD PRIMARY KEY (`panier_id`,`article_id`),
  ADD KEY `article_id` (`article_id`);

--
-- Index pour la table `payments`
--
ALTER TABLE `payments`
  ADD PRIMARY KEY (`id`),
  ADD KEY `order_id` (`order_id`),
  ADD KEY `method_id` (`method_id`);

--
-- Index pour la table `payment_method`
--
ALTER TABLE `payment_method`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code` (`code`);

--
-- Index pour la table `payment_transaction`
--
ALTER TABLE `payment_transaction`
  ADD PRIMARY KEY (`id`),
  ADD KEY `payment_id` (`payment_id`);

--
-- Index pour la table `profil`
--
ALTER TABLE `profil`
  ADD PRIMARY KEY (`idProfil`);

--
-- Index pour la table `promotion`
--
ALTER TABLE `promotion`
  ADD PRIMARY KEY (`promotion_id`),
  ADD KEY `article_id` (`article_id`);

--
-- Index pour la table `status`
--
ALTER TABLE `status`
  ADD PRIMARY KEY (`id_status`),
  ADD UNIQUE KEY `nom_status` (`nom_status`);

--
-- Index pour la table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `login` (`login`);

--
-- Index pour la table `userprofil`
--
ALTER TABLE `userprofil`
  ADD PRIMARY KEY (`idUser`,`idProfil`),
  ADD KEY `idProfil` (`idProfil`);

--
-- Index pour la table `wallet`
--
ALTER TABLE `wallet`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`);

--
-- Index pour la table `wallet_transaction`
--
ALTER TABLE `wallet_transaction`
  ADD PRIMARY KEY (`id`),
  ADD KEY `wallet_id` (`wallet_id`);

--
-- AUTO_INCREMENT pour les tables déchargées
--

--
-- AUTO_INCREMENT pour la table `adresse`
--
ALTER TABLE `adresse`
  MODIFY `adresse_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=24;

--
-- AUTO_INCREMENT pour la table `article`
--
ALTER TABLE `article`
  MODIFY `article_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT pour la table `article_image`
--
ALTER TABLE `article_image`
  MODIFY `image_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT pour la table `avis`
--
ALTER TABLE `avis`
  MODIFY `avis_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT pour la table `categorie`
--
ALTER TABLE `categorie`
  MODIFY `categorie_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT pour la table `client`
--
ALTER TABLE `client`
  MODIFY `client_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT pour la table `commande`
--
ALTER TABLE `commande`
  MODIFY `commande_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT pour la table `commercant`
--
ALTER TABLE `commercant`
  MODIFY `commercant_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT pour la table `livraison`
--
ALTER TABLE `livraison`
  MODIFY `livraison_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT pour la table `livreur`
--
ALTER TABLE `livreur`
  MODIFY `livreur_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT pour la table `orders`
--
ALTER TABLE `orders`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=20;

--
-- AUTO_INCREMENT pour la table `order_items`
--
ALTER TABLE `order_items`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=35;

--
-- AUTO_INCREMENT pour la table `order_status`
--
ALTER TABLE `order_status`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT pour la table `panier`
--
ALTER TABLE `panier`
  MODIFY `panier_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT pour la table `payments`
--
ALTER TABLE `payments`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT pour la table `payment_method`
--
ALTER TABLE `payment_method`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT pour la table `payment_transaction`
--
ALTER TABLE `payment_transaction`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT pour la table `profil`
--
ALTER TABLE `profil`
  MODIFY `idProfil` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT pour la table `promotion`
--
ALTER TABLE `promotion`
  MODIFY `promotion_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT pour la table `status`
--
ALTER TABLE `status`
  MODIFY `id_status` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT pour la table `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT pour la table `wallet`
--
ALTER TABLE `wallet`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT pour la table `wallet_transaction`
--
ALTER TABLE `wallet_transaction`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- Contraintes pour les tables déchargées
--

--
-- Contraintes pour la table `adresse`
--
ALTER TABLE `adresse`
  ADD CONSTRAINT `fk_adresse_client` FOREIGN KEY (`client_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `article`
--
ALTER TABLE `article`
  ADD CONSTRAINT `Article_ibfk_1` FOREIGN KEY (`commercant_id`) REFERENCES `commercant` (`commercant_id`),
  ADD CONSTRAINT `Article_ibfk_2` FOREIGN KEY (`categorie_id`) REFERENCES `categorie` (`categorie_id`);

--
-- Contraintes pour la table `article_image`
--
ALTER TABLE `article_image`
  ADD CONSTRAINT `Article_Image_ibfk_1` FOREIGN KEY (`article_id`) REFERENCES `article` (`article_id`);

--
-- Contraintes pour la table `avis`
--
ALTER TABLE `avis`
  ADD CONSTRAINT `Avis_ibfk_1` FOREIGN KEY (`client_id`) REFERENCES `client` (`client_id`),
  ADD CONSTRAINT `Avis_ibfk_2` FOREIGN KEY (`article_id`) REFERENCES `article` (`article_id`),
  ADD CONSTRAINT `Avis_ibfk_3` FOREIGN KEY (`commercant_id`) REFERENCES `commercant` (`commercant_id`),
  ADD CONSTRAINT `Avis_ibfk_4` FOREIGN KEY (`livreur_id`) REFERENCES `livreur` (`livreur_id`);

--
-- Contraintes pour la table `categorie`
--
ALTER TABLE `categorie`
  ADD CONSTRAINT `Categorie_ibfk_1` FOREIGN KEY (`parent_id`) REFERENCES `categorie` (`categorie_id`);

--
-- Contraintes pour la table `commande`
--
ALTER TABLE `commande`
  ADD CONSTRAINT `Commande_ibfk_1` FOREIGN KEY (`client_id`) REFERENCES `client` (`client_id`);

--
-- Contraintes pour la table `livraison`
--
ALTER TABLE `livraison`
  ADD CONSTRAINT `Livraison_ibfk_1` FOREIGN KEY (`commande_id`) REFERENCES `commande` (`commande_id`),
  ADD CONSTRAINT `Livraison_ibfk_2` FOREIGN KEY (`livreur_id`) REFERENCES `livreur` (`livreur_id`),
  ADD CONSTRAINT `Livraison_ibfk_3` FOREIGN KEY (`article_id`) REFERENCES `article` (`article_id`),
  ADD CONSTRAINT `Livraison_ibfk_4` FOREIGN KEY (`client_id`) REFERENCES `client` (`client_id`),
  ADD CONSTRAINT `Livraison_ibfk_5` FOREIGN KEY (`commercant_id`) REFERENCES `commercant` (`commercant_id`);

--
-- Contraintes pour la table `orders`
--
ALTER TABLE `orders`
  ADD CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  ADD CONSTRAINT `orders_ibfk_2` FOREIGN KEY (`status_id`) REFERENCES `order_status` (`id`);

--
-- Contraintes pour la table `order_items`
--
ALTER TABLE `order_items`
  ADD CONSTRAINT `order_items_ibfk_1` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`);

--
-- Contraintes pour la table `panier`
--
ALTER TABLE `panier`
  ADD CONSTRAINT `Panier_ibfk_1` FOREIGN KEY (`client_id`) REFERENCES `client` (`client_id`);

--
-- Contraintes pour la table `panier_article`
--
ALTER TABLE `panier_article`
  ADD CONSTRAINT `Panier_Article_ibfk_1` FOREIGN KEY (`panier_id`) REFERENCES `panier` (`panier_id`),
  ADD CONSTRAINT `Panier_Article_ibfk_2` FOREIGN KEY (`article_id`) REFERENCES `article` (`article_id`);

--
-- Contraintes pour la table `payments`
--
ALTER TABLE `payments`
  ADD CONSTRAINT `payments_ibfk_1` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`),
  ADD CONSTRAINT `payments_ibfk_2` FOREIGN KEY (`method_id`) REFERENCES `payment_method` (`id`);

--
-- Contraintes pour la table `payment_transaction`
--
ALTER TABLE `payment_transaction`
  ADD CONSTRAINT `payment_transaction_ibfk_1` FOREIGN KEY (`payment_id`) REFERENCES `payments` (`id`);

--
-- Contraintes pour la table `promotion`
--
ALTER TABLE `promotion`
  ADD CONSTRAINT `Promotion_ibfk_1` FOREIGN KEY (`article_id`) REFERENCES `article` (`article_id`);

--
-- Contraintes pour la table `userprofil`
--
ALTER TABLE `userprofil`
  ADD CONSTRAINT `userProfil_ibfk_1` FOREIGN KEY (`idUser`) REFERENCES `user` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `userProfil_ibfk_2` FOREIGN KEY (`idProfil`) REFERENCES `profil` (`idProfil`) ON DELETE CASCADE;

--
-- Contraintes pour la table `wallet`
--
ALTER TABLE `wallet`
  ADD CONSTRAINT `wallet_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

--
-- Contraintes pour la table `wallet_transaction`
--
ALTER TABLE `wallet_transaction`
  ADD CONSTRAINT `wallet_transaction_ibfk_1` FOREIGN KEY (`wallet_id`) REFERENCES `wallet` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
