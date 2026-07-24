
CREATE TABLE establishments (
    id                UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name              VARCHAR(255) NOT NULL,
    enabled           BOOLEAN NOT NULL DEFAULT false,
    waba_id           VARCHAR(255),
    phone_number_id   VARCHAR(255),
    user_id           UUID NOT NULL REFERENCES users(id),
    description       TEXT,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at        TIMESTAMPTZ
);

CREATE INDEX idx_establishments_user_id ON establishments(user_id);
CREATE INDEX idx_establishments_enabled ON establishments(enabled);