CREATE DATABASE gomail;

CREATE TABLE business
(
  business_id serial PRIMARY KEY,
  name VARCHAR (64) NOT NULL
);

CREATE TABLE users
(
  user_id bigserial PRIMARY KEY,
  email VARCHAR (128) NOT NULL,
  password bytea NOT NULL,
  enterprise_id serial
);

CREATE TABLE campaign
(
  campaign_id serial PRIMARY KEY,
  broadcastListId serial NOT NULL,
  title VARCHAR (128) NOT NULL,
  body TEXT NOT NULL,
);