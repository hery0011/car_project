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
);

CREATE TABLE `Commercant` (
  `commercant_id` int(11) NOT NULL,
  `nom` varchar(150) NOT NULL,
  `description` text DEFAULT NULL,
  `adresse` varchar(200) DEFAULT NULL,
  `telephone` varchar(50) DEFAULT NULL,
  `email` varchar(150) DEFAULT NULL,
  `latitude` varchar(45) DEFAULT NULL,
  `longitude` varchar(45) DEFAULT NULL
);

---------------  to execute ----------------
ALTER TABLE `user`
ADD COLUMN `commercant_id` INT(11) NULL AFTER `longitude`;

ALTER TABLE `user`
ADD CONSTRAINT fk_user_commercant
FOREIGN KEY (`commercant_id`) REFERENCES `Commercant`(`commercant_id`)
ON DELETE SET NULL
ON UPDATE CASCADE;