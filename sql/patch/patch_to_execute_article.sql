-- 1️⃣ Ajouter les colonnes
ALTER TABLE article
  ADD COLUMN slug VARCHAR(255) NOT NULL AFTER nom,
  ADD COLUMN short_description VARCHAR(512) DEFAULT NULL AFTER slug,
  ADD COLUMN status ENUM('draft','published','archived') DEFAULT 'draft' AFTER description,
  ADD COLUMN is_active TINYINT(1) DEFAULT 1 AFTER status,
  ADD COLUMN created_at DATETIME DEFAULT CURRENT_TIMESTAMP AFTER is_active,
  ADD COLUMN updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER created_at;

-- 2️⃣ Modifier les colonnes existantes si nécessaire
ALTER TABLE article
  CHANGE COLUMN nom nom VARCHAR(255) NOT NULL,
  CHANGE COLUMN prix prix_cents BIGINT NOT NULL,
  CHANGE COLUMN commercant_id merchant_id INT NOT NULL;

-- 3️⃣ Ajouter les clés étrangères
ALTER TABLE article ADD CONSTRAINT fk_article_merchant FOREIGN KEY (commercant_id) REFERENCES commercant(commercant_id) ON DELETE CASCADE

-- 4️⃣ Ajouter index et clé unique
ALTER TABLE article ADD INDEX idx_commercant_id (commercant_id)

ALTER TABLE article
DROP FOREIGN KEY Article_ibfk_2;

ALTER TABLE article
DROP INDEX categorie_id;

ALTER TABLE article
DROP COLUMN categorie_id;

UPDATE categorie
SET slug = LOWER(REPLACE(nom, ' ', '-'))
WHERE slug IS NULL OR slug = '';

------

ALTER TABLE categorie
  ADD COLUMN slug VARCHAR(255) NOT NULL AFTER nom;

UPDATE categorie
SET slug = LOWER(REPLACE(nom, ' ', '-'))
WHERE slug IS NULL OR slug = '';

ALTER TABLE categorie
ADD UNIQUE KEY slug_unique (slug);
-----------------------------------
CREATE TABLE article_category (
  article_id INT NOT NULL,
  categorie_id INT NOT NULL,
  PRIMARY KEY (article_id, categorie_id),
  FOREIGN KEY (article_id) REFERENCES article(article_id) ON DELETE CASCADE,
  FOREIGN KEY (categorie_id) REFERENCES categorie(categorie_id) ON DELETE CASCADE
)






