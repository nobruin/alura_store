

-- Creation of product table
CREATE TABLE IF NOT EXISTS product_test (
  id INT NOT NULL ,
  name varchar(250) NOT NULL,
  price FLOAT NOT NULL,
  Quantity  int NOT NULL,
  description varchar(250) NOT NULL,
  PRIMARY KEY (id)
);
