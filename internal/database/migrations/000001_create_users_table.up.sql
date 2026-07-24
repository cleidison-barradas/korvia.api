DO $$ BEGIN
    CREATE TYPE user_role AS ENUM ('admin', 'owner', 'establishment', 'professional');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

DO $$ BEGIN
    CREATE TYPE document_type AS ENUM ('CPF', 'CNPJ');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  first_name varchar(255),
  last_name varchar(255),
  email varchar(255) NOT NULL UNIQUE,
  phone varchar(20) NOT NULL UNIQUE,
  enabled boolean DEFAULT false,
  role user_role DEFAULT 'owner',
  password_hash varchar(255) NOT NULL,
  document varchar(20) NOT NULL UNIQUE,
  document_type document_type DEFAULT 'CPF',
  verified_at TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP
);