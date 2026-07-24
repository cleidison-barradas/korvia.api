CREATE TABLE services (
  id                UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name              VARCHAR(255) NOT NULL,
  enabled           BOOLEAN NOT NULL DEFAULT true,
  description       TEXT,
  price             BIGINT NOT NULL,
  establishment_id  UUID NOT NULL REFERENCES establishments(id),
  duration          INTEGER NOT NULL,
  created_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at        TIMESTAMPTZ
);

create index idx_services_establishment_id on services(establishment_id);
