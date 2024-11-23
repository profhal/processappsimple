CREATE DATABASE PROCESSAPP;
USE PROCESSAPP;
-- SET GLOBAL local_infile=1;



CREATE TABLE customers (
  zipcode VARCHAR(100),
  customerid VARCHAR(36) NOT NULL PRIMARY KEY
);

LOAD DATA INFILE '/docker-entrypoint-initdb.d/init_data/customerIdsWithZips.txt' 
INTO TABLE customers
FIELDS TERMINATED BY ','
;

CREATE TABLE products (
  productid VARCHAR(36) NOT NULL PRIMARY KEY
);
LOAD DATA INFILE '/docker-entrypoint-initdb.d/init_data/productIds.txt' 
INTO TABLE products
;

CREATE TABLE purchase_history (
  purchaseid binary(16) default (uuid_to_bin(uuid())) not null primary key,
  customerid VARCHAR(36) REFERENCES customers(customerid),
  productid VARCHAR(36) REFERENCES products(productid)
);
LOAD DATA INFILE '/docker-entrypoint-initdb.d/init_data/purchaseHistory.txt' 
INTO TABLE purchase_history
FIELDS TERMINATED BY ','
(customerid, productid);
