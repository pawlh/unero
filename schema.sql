CREATE TABLE IF NOT EXISTS `transaction` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `date` DATE NOT NULL,
    `description` TEXT NOT NULL,
    `amount` NUMERIC(10,2) NOT NULL,
    `category` TEXT NOT NULL,
    `account` TEXT NOT NULL
);
