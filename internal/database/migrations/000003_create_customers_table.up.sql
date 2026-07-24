CREATE TABLE customers (
    id                UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name        VARCHAR(255) NOT NULL,
    last_name         VARCHAR(255) NOT NULL,
    phone             VARCHAR(20) NOT NULL,
    email             VARCHAR(255),
    birthday          DATE,
    establishment_id  UUID NOT NULL REFERENCES establishments(id),
    created_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at        TIMESTAMPTZ
);

CREATE INDEX idx_customers_establishment_id ON customers(establishment_id);
CREATE INDEX idx_customers_phone ON customers(phone);