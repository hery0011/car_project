-- ===========================
--  USERS (base)
-- ===========================
CREATE TABLE users (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  firstname VARCHAR(100),
  lastname VARCHAR(100),
  email VARCHAR(150) UNIQUE NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ===========================
--  PAYMENT METHOD (configurable)
-- ===========================
CREATE TABLE payment_method (
  id INT AUTO_INCREMENT PRIMARY KEY,
  code VARCHAR(50) UNIQUE NOT NULL,  -- 'stripe', 'paypal', 'cash', 'wallet'
  name VARCHAR(100) NOT NULL,
  description TEXT NULL,
  is_active BOOLEAN DEFAULT TRUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO payment_method (code, name) VALUES
('stripe', 'Paiement carte via Stripe'),
('paypal', 'Paiement PayPal'),
('cash', 'Paiement à la livraison'),
('wallet', 'Portefeuille interne');

-- ===========================
--  ORDER STATUS (configurable)
-- ===========================
CREATE TABLE order_status (
  id INT AUTO_INCREMENT PRIMARY KEY,
  code VARCHAR(50) UNIQUE NOT NULL,    -- 'pending_payment', 'paid', ...
  label VARCHAR(100) NOT NULL,
  is_final BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO order_status (code, label, is_final) VALUES
('cart', 'Panier', FALSE),
('pending_payment', 'En attente de paiement', FALSE),
('paid', 'Payée', FALSE),
('processing', 'En traitement', FALSE),
('shipped', 'Expédiée', FALSE),
('delivered', 'Livrée', TRUE),
('cancelled', 'Annulée', TRUE),
('refunded', 'Remboursée', TRUE);

-- ===========================
--  ORDERS
-- ===========================
CREATE TABLE orders (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  status_id INT NOT NULL,
  total_amount DECIMAL(12,2) NOT NULL DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES user(id),
  FOREIGN KEY (status_id) REFERENCES order_status(id)
);

-- ===========================
--  ORDER ITEMS
-- ===========================
CREATE TABLE order_items (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  order_id BIGINT NOT NULL,
  product_id BIGINT NOT NULL,
  product_name VARCHAR(255),
  quantity INT NOT NULL,
  unit_price DECIMAL(12,2) NOT NULL,
  total_price DECIMAL(12,2) NOT NULL,
  FOREIGN KEY (order_id) REFERENCES orders(id)
);

-- ===========================
--  PAYMENTS (paiement global d'une commande)
-- ===========================
CREATE TABLE payments (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  order_id BIGINT NOT NULL,
  method_id INT NOT NULL,
  amount DECIMAL(12,2) NOT NULL,
  status ENUM('initiated','completed','failed','refunded') DEFAULT 'initiated',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (order_id) REFERENCES orders(id),
  FOREIGN KEY (method_id) REFERENCES payment_method(id)
);

-- ===========================
--  PAYMENT TRANSACTION (détails financiers)
-- ===========================
CREATE TABLE payment_transaction (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  payment_id BIGINT NOT NULL,
  transaction_type ENUM('authorization', 'capture', 'refund', 'void') DEFAULT 'capture',
  transaction_reference VARCHAR(255),
  amount DECIMAL(12,2) NOT NULL,
  currency CHAR(3) DEFAULT 'EUR',
  status ENUM('pending', 'success', 'failed') DEFAULT 'pending',
  raw_response JSON NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (payment_id) REFERENCES payments(id)
);

-- ===========================
--  WALLET (solde interne utilisateur)
-- ===========================
CREATE TABLE wallet (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL UNIQUE,
  balance DECIMAL(12,2) NOT NULL DEFAULT 0,
  currency CHAR(3) DEFAULT 'EUR',
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES user(id)
);

-- ===========================
--  WALLET TRANSACTION (journal comptable interne)
-- ===========================
CREATE TABLE wallet_transaction (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  wallet_id BIGINT NOT NULL,
  transaction_type ENUM('credit', 'debit', 'refund', 'adjust') NOT NULL,
  amount DECIMAL(12,2) NOT NULL,
  reference VARCHAR(255) NULL,      -- ex: order_id, payment_id, note...
  description VARCHAR(255) NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (wallet_id) REFERENCES wallet(id)
);

CREATE TABLE IF NOT EXISTS `Adresse` (
  `adresse_id` INT(11) NOT NULL AUTO_INCREMENT,
  `client_id` INT(11) NOT NULL,
  `rue` VARCHAR(150) DEFAULT NULL,
  `ville` VARCHAR(100) DEFAULT NULL,
  `code_postal` VARCHAR(20) DEFAULT NULL,
  `pays` VARCHAR(50) DEFAULT NULL,
  `latitude` VARCHAR(255) DEFAULT NULL,
  `longitude` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`adresse_id`),
  CONSTRAINT `fk_adresse_client` FOREIGN KEY (`client_id`) 
      REFERENCES `user`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);
