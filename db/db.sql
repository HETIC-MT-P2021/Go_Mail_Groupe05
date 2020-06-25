CREATE DATABASE gomail;

CREATE TABLE business
(
  business_id serial PRIMARY KEY,
  name VARCHAR (64) NOT NULL,
  UNIQUE (name)
);

CREATE TABLE users (
    user_id bigserial PRIMARY KEY,
    email VARCHAR (128) NOT NULL,
    password bytea NOT NULL,
    business_id serial references business(business_id)
);

CREATE TABLE campaign (
    campaign_id serial PRIMARY KEY NOT NULL,
    name VARCHAR (128) NOT NULL,
    mailing_list_id serial NOT NULL references mailing_list(mailing_list_id),
    business_id serial references business(business_id)
);

CREATE TABLE mailing_list (
    mailing_list_id serial PRIMARY KEY NOT NULL,
    name VARCHAR (128) NOT NULL,
    business_id serial NOT NULL references business(business_id)
);

CREATE TABLE mailing_list_customer_assoc (
    mailing_list_id serial NOT NULL references mailing_list(mailing_list_id),
    customer_id serial NOT NULL references customer(customer_id)
);

CREATE TABLE customer (
    customer_id serial PRIMARY KEY NOT NULL,
    email VARCHAR (128) NOT NULL,
    name VARCHAR (128) NOT NULL,
    surname VARCHAR (128) NOT NULL,
    business_id serial NOT NULL references business(business_id)
);