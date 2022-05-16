Create database auth_backend;

use auth_backend;

DROP TABLE IF EXISTS users;
CREATE TABLE users (
  id 			INT AUTO_INCREMENT NOT NULL,
  first_name 	VARCHAR(100),
  last_name  	VARCHAR(100),
  email     	VARCHAR(255) NOT NULL UNIQUE,
  password  	VARCHAR(50) NOT NULL,
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS sessions;
CREATE TABLE sessions (
  session_id 	VARCHAR(100) NOT NULL,
  user_id  	INT,
  last_active timestamp,
  PRIMARY KEY (`session_id`)
);