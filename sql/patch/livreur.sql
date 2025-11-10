ALTER TABLE livreur 
ADD COLUMN user_id INT(11) UNIQUE, 
ADD CONSTRAINT fk_livreur_user FOREIGN KEY (user_id) REFERENCES user(id)