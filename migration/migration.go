package migration

import (
	"echo-jwt/migrate"
)

// LocalMigrations ...
var LocalMigrations = migrate.Migrations{
	migrate.Migration{
		ID: 1598207594,
		SQL: `CREATE TABLE IF NOT EXISTS clients (
			  id int(11) NOT NULL AUTO_INCREMENT,
			  email varchar(100) DEFAULT NULL,
			  user_name varchar(100) DEFAULT NULL,
			  password varchar(100) NOT NULL DEFAULT '',
			  created_at datetime DEFAULT CURRENT_TIMESTAMP,
			  updated_at datetime DEFAULT NULL,
			  PRIMARY KEY (id),
              KEY email_idx (email),
              KEY user_name_idx (user_name));`,
	},

}
