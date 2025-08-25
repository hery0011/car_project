ALTER TABLE `livraisons`.`Commande` 
CHANGE COLUMN `status` `status` INT NULL ;

ALTER TABLE `livraisons`.`Commande` 
CHANGE COLUMN `status` `status_id` INT NULL DEFAULT NULL ;
