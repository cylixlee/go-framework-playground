USE userdata;

CREATE TABLE user(
    id       bigint(20) NOT NULL AUTO_INCREMENT,
    username varchar(30),
    password varchar(64),
    PRIMARY KEY (id)
);