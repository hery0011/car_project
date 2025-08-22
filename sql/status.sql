CREATE TABLE status (
    id_status INT AUTO_INCREMENT PRIMARY KEY,
    nom_status VARCHAR(50) NOT NULL UNIQUE
);

-- Insérer les valeurs
INSERT INTO status (nom_status) VALUES
('commande ouvert'),
('commande en cours'),
('commande ferme'),
('panier ouvert'),
('panier en cours');