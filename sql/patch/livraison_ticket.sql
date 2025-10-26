CREATE TABLE delivery_ticket_status (
    id INT AUTO_INCREMENT PRIMARY KEY,
    code VARCHAR(50) UNIQUE NOT NULL,   -- 'pending', 'assigned', 'picked', 'delivered', 'cancelled'
    label VARCHAR(100) NOT NULL,
    is_final BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE delivery_tickets (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    order_id BIGINT NULL,                      -- optionnel si ticket vient d'une commande
    client_id INT NOT NULL,                    -- utilisateur qui a créé le ticket
    pickup_address_id INT NOT NULL,            -- FK vers adresse (récupération)
    dropoff_address_id INT NOT NULL,           -- FK vers adresse (livraison)
    delivery_price DECIMAL(10,2) NULL,        -- prix défini par client ou admin/livreur
    price_last_updated_by INT NULL,            -- utilisateur ayant mis à jour le prix
    status_id INT NOT NULL,                    -- FK vers delivery_ticket_status
    assigned_to INT NULL,                      -- livreur assigné
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (client_id) REFERENCES user(id),
    FOREIGN KEY (pickup_address_id) REFERENCES adresse(adresse_id),
    FOREIGN KEY (dropoff_address_id) REFERENCES adresse(adresse_id),
    FOREIGN KEY (price_last_updated_by) REFERENCES user(id),
    FOREIGN KEY (assigned_to) REFERENCES user(id),
    FOREIGN KEY (status_id) REFERENCES delivery_ticket_status(id)
);



INSERT INTO delivery_ticket_status (code, label, is_final) VALUES
('pending', 'En attente', FALSE),
('assigned', 'Assigné à un livreur', FALSE),
('picked', 'Colis récupéré', FALSE),
('in_transit', 'En cours de livraison', FALSE),
('delivered', 'Livré', TRUE),
('cancelled', 'Annulé', TRUE);


CREATE TABLE delivery_price_history (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    ticket_id BIGINT NOT NULL,
    old_price DECIMAL(10,2) NULL,
    new_price DECIMAL(10,2) NOT NULL,
    changed_by INT NOT NULL,
    changed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (ticket_id) REFERENCES delivery_tickets(id),
    FOREIGN KEY (changed_by) REFERENCES user(id)
);

CREATE TABLE order_addresses (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  order_id BIGINT NOT NULL,
  adresse_id INT NOT NULL,
  type ENUM('pickup','dropoff') NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (order_id) REFERENCES orders(id),
  FOREIGN KEY (adresse_id) REFERENCES adresse(adresse_id)
);

