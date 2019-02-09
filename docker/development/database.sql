CREATE SCHEMA IF NOT EXISTS `palmirinha` DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci;
USE `palmirinha` ;

-- -----------------------------------------------------
-- Table `palmirinha`.`cart`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `palmirinha`.`cart` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `channel_id` VARCHAR(255) NULL,
  `description` VARCHAR(255) NULL,
  `status` TINYINT(1) NOT NULL DEFAULT 1,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`))
ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;

-- -----------------------------------------------------
-- Table `palmirinha`.`item`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `palmirinha`.`item` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `cart_id` INT(11) NOT NULL,
  `discord_user_id` VARCHAR(45) NOT NULL,
  `description` VARCHAR(255) NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `fk_item_cart_idx` (`cart_id` ASC),
  CONSTRAINT `fk_item_cart`
    FOREIGN KEY (`cart_id`)
    REFERENCES `palmirinha`.`cart` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;
