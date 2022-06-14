DROP DATABASE IF EXISTS db_22spring;
CREATE DATABASE db_22spring;
USE db_22spring;

CREATE TABLE `searchs` (
    `id` CHAR(36) NOT NULL,
    `station` VARCHAR(100) DEFAULT NULL,
    `lat` DOUBLE(9, 6) DEFAULT NULL,
    `lng` DOUBLE(9, 6) DEFAULT NULL,
    `created_at` DATETIME(6) DEFAULT NULL
);

CREATE TABLE `questions` (
    `id` INT(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `shop_id` CHAR(10) NOT NULL,
    `search_id` CHAR(36) NOT NULL,
    `number` INT NOT NULL,
    `created_at` DATETIME(6) DEFAULT NULL
);

CREATE TABLE `shops` (
    `shop_id` CHAR(10) NOT NULL,
    `name` VARCHAR(200) NOT NULL,
    `image` VARCHAR(2048) NOT NULL,
    `genre_code` CHAR(4) NOT NULL,
    `subgenre_code` CHAR(4) DEFAULT NULL,
    `price_code` CHAR(4) NOT NULL,
    `created_at` DATETIME(6) DEFAULT NULL
);

CREATE TABLE `genres` (
    `genre_code` CHAR(4) NOT NULL,
    `name` VARCHAR(100) NOT NULL
);

CREATE TABLE `prices` (
    `price_code` CHAR(4) NOT NULL,
    `name` VARCHAR(100) NOT NULL
);

CREATE TABLE `gourmets` (
    `station` VARCHAR(100) DEFAULT NULL,
    `lat` DOUBLE(9, 6) DEFAULT NULL,
    `lng` DOUBLE(9, 6) DEFAULT NULL,
    `shops` TEXT NOT NULL,
    `created_at` DATETIME(6) DEFAULT NULL
);