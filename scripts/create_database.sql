CREATE TABLE users_db.users (
	id BIGINT(20) AUTO_INCREMENT NOT NULL PRIMARY KEY,
	first_name varchar(100) NULL,
	last_name varchar(100) NULL,
	email varchar(100) NOT NULL,
	date_created varchar(100) NULL,
	CONSTRAINT email_UNIQUE UNIQUE (email)
)
ENGINE=InnoDB
DEFAULT CHARSET=latin1
COLLATE=latin1_swedish_ci;
