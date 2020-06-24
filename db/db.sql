CREATE DATABASE gomail;

CREATE TABLE business
(
  business_id serial PRIMARY KEY,
  name VARCHAR (64) NOT NULL
);

CREATE TABLE users (
    user_id bigserial PRIMARY KEY,
    email VARCHAR (128) NOT NULL,
    password bytea NOT NULL,
    business_id serial
);

CREATE TABLE campaign (
    campaign_id serial PRIMARY KEY,
    name VARCHAR (128) NOT NULL,
    mailing_list_id serial NOT NULL,
    business_id serial
);

CREATE TABLE mailing_list (
    mailing_list_id serial NOT NULL,
    name VARCHAR (128) NOT NULL,
    business_id serial NOT NULL
);

CREATE TABLE mailing_list_customer_assoc (
    mailing_list_id serial NOT NULL,
    customer_id serial NOT NULL
);

CREATE TABLE customer (
    customer_id serial NOT NULL,
    email VARCHAR (128) NOT NULL,
    name VARCHAR (128) NOT NULL,
    surname VARCHAR (128) NOT NULL,
    business_id serial NOT NULL
);