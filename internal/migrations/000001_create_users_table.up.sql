CREATE TABLE IF NOT EXISTS users(
  id varchar(50) PRIMARY KEY,
  name varchar(255),
  email varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  role varchar(100)
);
