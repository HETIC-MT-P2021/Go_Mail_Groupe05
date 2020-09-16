CREATE TABLE business (
  id serial PRIMARY KEY,
  name VARCHAR (64) NOT NULL
);

CREATE TABLE users (
    id bigserial PRIMARY KEY,
    email VARCHAR (128) NOT NULL,
    password bytea NOT NULL,

    business_id integer references business(id)
);

CREATE TABLE mailing_list (
    id serial PRIMARY KEY,
    name VARCHAR (128) NOT NULL,

    business_id integer references business(id)
);

CREATE TABLE campaign (
    id serial PRIMARY KEY,
    name VARCHAR (128) NOT NULL,

    mailing_list_id integer references mailing_list(id),
    business_id integer references business(id)
);

CREATE TABLE customer (
    id serial PRIMARY KEY,
    email VARCHAR (128) NOT NULL,
    name VARCHAR (128) NOT NULL,
    surname VARCHAR (128) NOT NULL,

    business_id integer references business(id)
);

CREATE TABLE mailing_list_customer_assoc (
    mailing_list_id integer references mailing_list(id),
    customer_id integer references customer(id)
);
