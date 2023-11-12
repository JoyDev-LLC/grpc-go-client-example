CREATE TABLE products
(
    id SERIAL,
    name varchar(255) not null,
    price float not null default 0.0,
    PRIMARY KEY (id)
);
