CREATE TABLE IF NOT EXISTS `users` ( 
 id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
 name VARCHAR(255) NOT NULL,
 email VARCHAR(255) NOT NULL UNIQUE,
 password VARCHAR(255) NOT NULL,
 phone VARCHAR(255) NOT NULL UNIQUE, 
 created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)