
-- -----------------------------------------------------
-- Table `serieswatcher`.`FileFormat`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `serieswatcher`.`FileFormat` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `Type` VARCHAR(45) NULL,
  PRIMARY KEY (`id`))
  ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `serieswatcher`.`Image`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `serieswatcher`.`Image` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `Data` BLOB NULL,
  `FileFormat_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_Image_FileFormat1_idx` (`FileFormat_id` ASC),
  CONSTRAINT `fk_Image_FileFormat1`
  FOREIGN KEY (`FileFormat_id`)
  REFERENCES `serieswatcher`.`FileFormat` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
  ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `serieswatcher`.`Series`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `serieswatcher`.`Series` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `Title` VARCHAR(255) NULL,
  `Image_id` INT NOT NULL,
  `Seriescol` INT NULL,
  `DataProviderUrl` VARCHAR(45) NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_Series_Image1_idx` (`Image_id` ASC),
  CONSTRAINT `fk_Series_Image1`
  FOREIGN KEY (`Image_id`)
  REFERENCES `serieswatcher`.`Image` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
  ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `serieswatcher`.`Provider`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `serieswatcher`.`Provider` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `Name` VARCHAR(255) NULL,
  `Scraper` VARCHAR(255) NULL,
  `Image_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_Provider_Image1_idx` (`Image_id` ASC),
  CONSTRAINT `fk_Provider_Image1`
  FOREIGN KEY (`Image_id`)
  REFERENCES `serieswatcher`.`Image` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
  ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `serieswatcher`.`Episode`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `serieswatcher`.`Episode` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `Series_id` INT NOT NULL,
  `Image_id` INT NOT NULL,
  `Episode` INT NULL,
  `Season` INT NULL,
  `Title` VARCHAR(255) NULL,
  `Description` TEXT NULL,
  `ReleaseDate` DATETIME NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_Episode_Series1_idx` (`Series_id` ASC),
  INDEX `fk_Episode_Image1_idx` (`Image_id` ASC),
  CONSTRAINT `fk_Episode_Series1`
  FOREIGN KEY (`Series_id`)
  REFERENCES `serieswatcher`.`Series` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Episode_Image1`
  FOREIGN KEY (`Image_id`)
  REFERENCES `serieswatcher`.`Image` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
  ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `serieswatcher`.`Credentials`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `serieswatcher`.`Credentials` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `Password` VARCHAR(255) NULL,
  PRIMARY KEY (`id`))
  ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `serieswatcher`.`User`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `serieswatcher`.`User` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `Name` VARCHAR(255) NULL,
  `Credentials_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_User_Credentials1_idx` (`Credentials_id` ASC),
  CONSTRAINT `fk_User_Credentials1`
  FOREIGN KEY (`Credentials_id`)
  REFERENCES `serieswatcher`.`Credentials` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
  ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `serieswatcher`.`WatchPointer`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `serieswatcher`.`WatchPointer` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `User_id` INT NOT NULL,
  `Episode_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_WatchPointer_User_idx` (`User_id` ASC),
  INDEX `fk_WatchPointer_Episode1_idx` (`Episode_id` ASC),
  CONSTRAINT `fk_WatchPointer_User`
  FOREIGN KEY (`User_id`)
  REFERENCES `serieswatcher`.`User` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_WatchPointer_Episode1`
  FOREIGN KEY (`Episode_id`)
  REFERENCES `serieswatcher`.`Episode` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
  ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `serieswatcher`.`ProviderUrl`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `serieswatcher`.`ProviderUrl` (
  `Episode_id` INT NOT NULL,
  `Provider_id` INT NOT NULL,
  `Url` VARCHAR(100) NULL,
  PRIMARY KEY (`Episode_id`, `Provider_id`),
  INDEX `fk_Episode_has_Provider_Provider1_idx` (`Provider_id` ASC),
  INDEX `fk_Episode_has_Provider_Episode1_idx` (`Episode_id` ASC),
  CONSTRAINT `fk_Episode_has_Provider_Episode1`
  FOREIGN KEY (`Episode_id`)
  REFERENCES `serieswatcher`.`Episode` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Episode_has_Provider_Provider1`
  FOREIGN KEY (`Provider_id`)
  REFERENCES `serieswatcher`.`Provider` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
  ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `serieswatcher`.`ThirdPartyAccount`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `serieswatcher`.`ThirdPartyAccount` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `Provider_id` INT NOT NULL,
  `User_id` INT NOT NULL,
  `Credentials_id` INT NOT NULL,
  `Username` VARCHAR(45) NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_Account_Provider1_idx` (`Provider_id` ASC),
  INDEX `fk_Account_User1_idx` (`User_id` ASC),
  INDEX `fk_ThirdPartyAccount_Credentials1_idx` (`Credentials_id` ASC),
  CONSTRAINT `fk_Account_Provider1`
  FOREIGN KEY (`Provider_id`)
  REFERENCES `serieswatcher`.`Provider` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Account_User1`
  FOREIGN KEY (`User_id`)
  REFERENCES `serieswatcher`.`User` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_ThirdPartyAccount_Credentials1`
  FOREIGN KEY (`Credentials_id`)
  REFERENCES `serieswatcher`.`Credentials` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
  ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `serieswatcher`.`ProviderUrl`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `serieswatcher`.`ProviderUrl` (
  `Episode_id` INT NOT NULL AUTO_INCREMENT,
  `Provider_id` INT NOT NULL,
  `Url` VARCHAR(100) NULL,
  PRIMARY KEY (`Episode_id`, `Provider_id`),
  INDEX `fk_Episode_has_Provider_Provider1_idx` (`Provider_id` ASC),
  INDEX `fk_Episode_has_Provider_Episode1_idx` (`Episode_id` ASC),
  CONSTRAINT `fk_Episode_has_Provider_Episode1`
  FOREIGN KEY (`Episode_id`)
  REFERENCES `serieswatcher`.`Episode` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Episode_has_Provider_Provider1`
  FOREIGN KEY (`Provider_id`)
  REFERENCES `serieswatcher`.`Provider` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
  ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `serieswatcher`.`SeriesProvider`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `serieswatcher`.`SeriesProvider` (
  `Series_id` INT NOT NULL AUTO_INCREMENT,
  `Provider_id` INT NOT NULL,
  PRIMARY KEY (`Series_id`, `Provider_id`),
  INDEX `fk_Series_has_Provider_Provider1_idx` (`Provider_id` ASC),
  INDEX `fk_Series_has_Provider_Series1_idx` (`Series_id` ASC),
  CONSTRAINT `fk_Series_has_Provider_Series1`
  FOREIGN KEY (`Series_id`)
  REFERENCES `serieswatcher`.`Series` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Series_has_Provider_Provider1`
  FOREIGN KEY (`Provider_id`)
  REFERENCES `serieswatcher`.`Provider` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
  ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `serieswatcher`.`errorLog`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `serieswatcher`.`errorLog` (
  `id` INT NOT NULL,
  `message` VARCHAR(500) NULL,
  `created` DATETIME NULL,
  PRIMARY KEY (`id`))
  ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `serieswatcher`.`workQueue`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `serieswatcher`.`workQueue` (
  `id` INT NOT NULL,
  `jobType` VARCHAR(45) NULL,
  `parameter` VARCHAR(100) NULL,
  `time` DATETIME NULL,
  PRIMARY KEY (`id`))
  ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;