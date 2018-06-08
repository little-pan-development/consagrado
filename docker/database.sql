CREATE SCHEMA IF NOT EXISTS `palmirinha_data` DEFAULT CHARACTER SET utf8 ;
USE `palmirinha_data` ;

-- -----------------------------------------------------
-- Table `palmirinha_data`.`cart`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `palmirinha_data`.`cart` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `description` VARCHAR(255) NULL,
  `status` TINYINT(1) NOT NULL DEFAULT 1,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `palmirinha_data`.`item`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `palmirinha_data`.`item` (
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
    REFERENCES `palmirinha_data`.`cart` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
