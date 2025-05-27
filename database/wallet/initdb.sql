use wallet;

CREATE TABLE IF NOT EXISTS clients (id varchar(255), name varchar(255), email varchar(255), created_at date);
CREATE TABLE IF NOT EXISTS accounts (id varchar(255), client_id varchar(255), balance int, created_at date);
CREATE TABLE IF NOT EXISTS transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date);

INSERT INTO clients VALUES ("9633db35-2227-45c9-a1cd-fe1bc99c5c22", "John Doe", "j@j.com", now());
INSERT INTO clients VALUES ("c7ce4127-afea-47f6-9126-ed22a6dac0bd", "Jane Doe", "jane@j.com", now());

INSERT INTO accounts VALUES ("4ba94984-d733-41be-8a22-aae5052d6b1a", "9633db35-2227-45c9-a1cd-fe1bc99c5c22", 1000, now());
INSERT INTO accounts VALUES ("3506c046-f344-4eae-b18a-ca0b3b40666f", "c7ce4127-afea-47f6-9126-ed22a6dac0bd", 0, now());