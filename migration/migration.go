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

	migrate.Migration{
		ID: 1598207595,
		SQL: `CREATE TABLE IF NOT EXISTS workspaces (
			  id int(11) NOT NULL AUTO_INCREMENT,
			  name varchar(100) DEFAULT NULL,
			  client_id int(11) NOT NULL,
			  created_at datetime DEFAULT CURRENT_TIMESTAMP,
			  updated_at datetime DEFAULT NULL,
			  PRIMARY KEY (id),
              KEY name_idx (name),
              KEY client_id_idx (client_id));`,
	},

	migrate.Migration{
		ID: 1598207600,
		SQL: `CREATE TABLE IF NOT EXISTS datasources (
			  id int(11) NOT NULL AUTO_INCREMENT,
			  name varchar(100) DEFAULT NULL,
			  description varchar(100) DEFAULT NULL,
			  created_at datetime DEFAULT CURRENT_TIMESTAMP,
			  updated_at datetime DEFAULT NULL,
			  PRIMARY KEY (id),
              KEY name_idx (name));`,
	},

	migrate.Migration{
		ID: 1598207601,
		SQL: `INSERT INTO datasources (name, description) values ('HTTP API', 'HTTP api'),
				('JS SDK', 'JS SDK')`,
	},

	migrate.Migration{
		ID: 1598207602,
		SQL: `CREATE TABLE IF NOT EXISTS authentication (
			  id int(11) NOT NULL AUTO_INCREMENT,
			  client_id int(11) NOT NULL,
	          apikey varchar(100) DEFAULT NULL,
			  created_at datetime DEFAULT CURRENT_TIMESTAMP,
			  updated_at datetime DEFAULT NULL,
			  PRIMARY KEY (id),
              KEY apikey_idx (apikey),
              KEY client_id_idx (client_id));`,
	},

	migrate.Migration{
		ID: 1598207603,
		SQL: `CREATE TABLE IF NOT EXISTS properties (
			  id int(11) NOT NULL AUTO_INCREMENT,
			  client_id int(11) NOT NULL,
	          workspace_id int(11) NOT NULL,
			  property_name varchar(100) DEFAULT NULL, 	
			  created_at datetime DEFAULT CURRENT_TIMESTAMP,
			  updated_at datetime DEFAULT NULL,
			  PRIMARY KEY (id),
              KEY workspace_id_idx (workspace_id),
		      KEY property_name_idx (property_name),
              KEY client_id_idx (client_id));`,
	},

	migrate.Migration{
		ID: 1598207604,
		SQL: `CREATE TABLE IF NOT EXISTS properties_sources (
			  id int(11) NOT NULL AUTO_INCREMENT,
			  properties_id int(11) NOT NULL,
	          sources_id int(11) NOT NULL,
			  created_at datetime DEFAULT CURRENT_TIMESTAMP,
			  updated_at datetime DEFAULT NULL,
			  PRIMARY KEY (id),
              KEY properties_id_idx (properties_id),
              KEY sources_id_idx (sources_id));`,
	},
}
