-- ✅ Merchant (propriétaire des articles)
CREATE TABLE merchants (
  merchant_id INT AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  shop_name VARCHAR(255) NOT NULL,
  status ENUM('pending','approved','suspended') DEFAULT 'pending',
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- ✅ Article
CREATE TABLE articles (
  article_id INT AUTO_INCREMENT PRIMARY KEY,
  merchant_id INT NOT NULL,
  sku VARCHAR(100) DEFAULT NULL, -- SKU global facultatif
  name VARCHAR(255) NOT NULL,
  slug VARCHAR(255) NOT NULL,
  short_description VARCHAR(512),
  description TEXT,
  status ENUM('draft','published','archived') DEFAULT 'draft',
  is_active TINYINT(1) DEFAULT 1,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (merchant_id) REFERENCES merchants(merchant_id) ON DELETE CASCADE,
  UNIQUE KEY merchant_slug_unique (merchant_id, slug), -- SEO unique par boutique
  INDEX (merchant_id)
);

-- ✅ Variantes d’article (taille/couleur…)
CREATE TABLE variants (
  variant_id INT AUTO_INCREMENT PRIMARY KEY,
  article_id INT NOT NULL,
  sku VARCHAR(120) NOT NULL, -- SKU unique
  attributes JSON DEFAULT NULL, -- {"color":"red","size":"M"}
  price_cents BIGINT NOT NULL,
  currency CHAR(3) DEFAULT 'USD',
  compare_at_price_cents BIGINT DEFAULT NULL,
  is_active TINYINT(1) DEFAULT 1,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (article_id) REFERENCES articles(article_id) ON DELETE CASCADE,
  INDEX (article_id),
  UNIQUE (sku)
);

-- ✅ Stock
CREATE TABLE inventories (
  inventory_id INT AUTO_INCREMENT PRIMARY KEY,
  variant_id INT NOT NULL,
  available INT DEFAULT 0,
  reserved INT DEFAULT 0,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (variant_id) REFERENCES variants(variant_id) ON DELETE CASCADE,
  INDEX (variant_id)
);

-- ✅ Images d’un article
CREATE TABLE images (
  image_id INT AUTO_INCREMENT PRIMARY KEY,
  article_id INT NOT NULL,
  url VARCHAR(1000) NOT NULL,
  alt_text VARCHAR(255),
  position INT DEFAULT 0, -- ordre d'affichage
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (article_id) REFERENCES articles(article_id) ON DELETE CASCADE,
  INDEX (article_id)
);

-- ✅ Catégories globales (hiérarchie possible)
CREATE TABLE categories (
  category_id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(200) NOT NULL,
  slug VARCHAR(200) UNIQUE,
  parent_id INT DEFAULT NULL,
  FOREIGN KEY (parent_id) REFERENCES categories(category_id) ON DELETE SET NULL
);

-- ✅ Liaison Article ↔ Category (N–N)
CREATE TABLE article_category (
  article_id INT NOT NULL,
  category_id INT NOT NULL,
  PRIMARY KEY (article_id, category_id),
  FOREIGN KEY (article_id) REFERENCES articles(article_id) ON DELETE CASCADE,
  FOREIGN KEY (category_id) REFERENCES categories(category_id) ON DELETE CASCADE
);

INSERT INTO category (name, description) VALUES
('Électronique', 'Produits high-tech et gadgets'),
('Vêtements', 'Mode et accessoires'),
('Maison', 'Objets pour la maison'),
('Alimentation', 'Produits alimentaires');

INSERT INTO merchant (user_id, shop_name, address) VALUES
(3, 'Tech Store', 'Lot IVC 123 Antananarivo'),
(4, 'Fashion Shop', 'Mahamasina Antananarivo');

INSERT INTO product (merchant_id, name, description, sku, price_cents, compare_at_price_cents, is_active) VALUES
(1, 'Smartphone Samsung A15', 'Smartphone Android 128GB', 'SKU-SAM-A15', 1200000, 1500000, 1),
(2, 'T-shirt Femme Bleu', 'T-shirt coton bleu', 'SKU-TSH-FEM-BL', 25000, 30000, 1);

INSERT INTO product_image (product_id, url, is_main) VALUES
(1, 'https://cdn.shop.com/products/samsung-a15.jpg', 1),
(1, 'https://cdn.shop.com/products/samsung-a15-2.jpg', 0),
(2, 'https://cdn.shop.com/products/tshirt-blue.jpg', 1);

INSERT INTO product_category (product_id, category_id) VALUES
(1, 1), -- Smartphone → Électronique
(2, 2); -- T-shirt → Vêtements


INSERT INTO product_variant 
(product_id, sku, name, price_cents, compare_at_price_cents, is_active) VALUES
(1, 'SKU-SAM-A15-128-BLK', '128Go - Noir', 1200000, 1500000, 1),
(2, 'SKU-TSH-FEM-BL-M', 'Bleu - M', 25000, 30000, 1),
(2, 'SKU-TSH-FEM-BL-L', 'Bleu - L', 25000, 30000, 1);

INSERT INTO product_inventory (variant_id, quantity) VALUES
(1, 50),  -- Samsung A15 noir
(2, 120), -- T-shirt M
(3, 80);  -- T-shirt L
