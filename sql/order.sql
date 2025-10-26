
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
  article_id BIGINT NOT NULL,
  article_name VARCHAR(255),
  quantity INT NOT NULL,
  unit_price DECIMAL(12,2) NOT NULL,
  total_price DECIMAL(12,2) NOT NULL,
  FOREIGN KEY (order_id) REFERENCES orders(id)
);