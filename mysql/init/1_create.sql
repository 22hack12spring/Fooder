DROP DATABASE IF EXISTS db_22spring;
CREATE DATABASE db_22spring;
USE db_22spring;

CREATE TABLE `searches` (
    `id` CHAR(36) PRIMARY KEY,
    `station` VARCHAR(100) DEFAULT NULL,
    `lat` DOUBLE(9, 6) DEFAULT NULL,
    `lng` DOUBLE(9, 6) DEFAULT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `genres` (
    `genre_code` CHAR(4) PRIMARY KEY,
    `name` VARCHAR(100) NOT NULL
);

CREATE TABLE `prices` (
    `price_code` CHAR(4) PRIMARY KEY,
    `name` VARCHAR(100) NOT NULL
);

CREATE TABLE `shops` (
    `shop_id` CHAR(10) PRIMARY KEY,
    `name` VARCHAR(200) NOT NULL,
    `image` VARCHAR(2048) NOT NULL,
    `genre_code` CHAR(4) NOT NULL,
    `subgenre_code` CHAR(4) DEFAULT NULL,
    `price_code` CHAR(4) NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`genre_code`) REFERENCES `genres`(`genre_code`),
    FOREIGN KEY (`subgenre_code`) REFERENCES `genres`(`genre_code`),
    FOREIGN KEY (`price_code`) REFERENCES `prices`(`price_code`)
);

CREATE TABLE `questions` (
    `id` INT(11) PRIMARY KEY AUTO_INCREMENT,
    `shop_id` CHAR(10) NOT NULL,
    `search_id` CHAR(36) NOT NULL,
    `number` INT NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`search_id`) REFERENCES `searches`(`id`),
    FOREIGN KEY (`shop_id`) REFERENCES `shops`(`shop_id`)
);

CREATE TABLE `gourmets` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `station` VARCHAR(100) DEFAULT NULL,
    `lat` DOUBLE(9, 6) DEFAULT NULL,
    `lng` DOUBLE(9, 6) DEFAULT NULL,
    `shops` TEXT NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
