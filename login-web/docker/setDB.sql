CREATE DATABASE gura DEFAULT CHARACTER SET UTF8;

CREATE USER 'admin'@'%' IDENTIFIED BY 'admin';

GRANT ALL PRIVILEGES ON gura.* TO 'admin'@'%';

FLUSH PRIVILEGES;

USE gura;

CREATE TABLE users (
    code INT PRIMARY KEY NOT NULL,
    id VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    pwd VARCHAR(255) NOT NULL
);

INSERT INTO users (code,id,name,pwd) VALUES (1,'karina','yujimin','000411');
INSERT INTO users (code,id,name,pwd) VALUES (2,'winter','kimminjeong','010101');
INSERT INTO users (code,id,name,pwd) VALUES (3,'ningning','ning yizhuo','021023');
INSERT INTO users (code,id,name,pwd) VALUES (4,'giselle','aeriuchinaga','001030');