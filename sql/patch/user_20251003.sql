-- 2. Ajouter les colonnes d'amélioration
ALTER TABLE `userprofil`
ADD COLUMN `date_assigned` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP AFTER `idProfil`,
ADD COLUMN `status` ENUM('ACTIVE', 'INACTIVE') NOT NULL DEFAULT 'ACTIVE' AFTER `date_assigned`;

-- 3. Ajouter la contrainte d'unicité (supprime les doublons avant)
ALTER TABLE `userprofil`
ADD CONSTRAINT unique_user_profil UNIQUE (`idUser`, `idProfil`);

-- 4. Ajouter les clés étrangères
ALTER TABLE `userprofil`
ADD CONSTRAINT fk_userprofil_user
    FOREIGN KEY (`idUser`) REFERENCES `user`(`id`) ON DELETE CASCADE,
ADD CONSTRAINT fk_userprofil_profil
    FOREIGN KEY (`idProfil`) REFERENCES `profil`(`idProfil`) ON DELETE CASCADE;

CREATE TABLE menu (
    id INT AUTO_INCREMENT PRIMARY KEY,
    label VARCHAR(100) NOT NULL,
    icon VARCHAR(50) NOT NULL,
    link VARCHAR(255) NOT NULL
);

CREATE TABLE menu_roles (
    menu_id INT NOT NULL,
    role VARCHAR(50) NOT NULL,
    PRIMARY KEY(menu_id, role),
    FOREIGN KEY (menu_id) REFERENCES menu(id)
);

INSERT INTO menu (id, label, icon, link) VALUES
(1, 'Mes commandes', 'cart-outline', '/profile/orders'),
(2, 'Tickets', 'ticket-outline', '/profile/tickets'),
(3, 'Mes produits', 'list-outline', '/profile/products'),
(4, 'Créer produit', 'add-circle-outline', '/profile/create-product');

INSERT INTO menu_roles (menu_id, role) VALUES
(1, 'client'),
(2, 'client'),
(2, 'livreur'),
(3, 'merchant'),
(4, 'merchant');


