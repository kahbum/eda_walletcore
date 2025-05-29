use balance;

CREATE TABLE IF NOT EXISTS accounts (id varchar(255), balance int, updated_at date);

INSERT INTO accounts VALUES ("4ba94984-d733-41be-8a22-aae5052d6b1a", 1000, now());
INSERT INTO accounts VALUES ("3506c046-f344-4eae-b18a-ca0b3b40666f", 0, now());