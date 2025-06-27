CREATE TABLE IF NOT EXISTS concerts
(
    id              UUID NOT NULL PRIMARY KEY,
    name            VARCHAR,
    slug            VARCHAR,
    description     TEXT,
    announced_at    TIMESTAMP with time zone,
    status              INT,

    created_at      TIMESTAMP with time zone default now(),
    updated_at      TIMESTAMP with time zone default now(),
    deleted_at      TIMESTAMP with time zone,
    UNIQUE (slug)
);

CREATE INDEX IF NOT EXISTS idx_concerts_deleted_at ON concerts (deleted_at);
CREATE INDEX IF NOT EXISTS idx_concerts_announced_at ON concerts (announced_at);
CREATE INDEX IF NOT EXISTS idx_concerts_status ON concerts (status);

CREATE TABLE IF NOT EXISTS concert_batches (
    id                  UUID PRIMARY KEY,
    concert_id          UUID REFERENCES concerts(id),
    name                VARCHAR, -- IDENTIFIER
    start_time          TIMESTAMPTZ,
    end_time            TIMESTAMPTZ,
    venue               VARCHAR,
    max_capacity        INT NOT NULL DEFAULT 0,
    available_tickets   INT,
    price               DOUBLE PRECISION,
    currency            VARCHAR(10),
    status              INT,    
    version             BIGINT default 0,

    created_at        TIMESTAMPTZ DEFAULT now(),
    updated_at        TIMESTAMPTZ DEFAULT now(),
    deleted_at        TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_concert_batches_created_at ON concert_batches (created_at);
CREATE INDEX IF NOT EXISTS idx_concert_batches_deleted_at ON concert_batches (deleted_at);
CREATE INDEX IF NOT EXISTS idx_concert_batches_concert_id ON concert_batches (concert_id);
CREATE INDEX IF NOT EXISTS idx_concert_batches_status ON concert_batches (status);

CREATE TABLE IF NOT EXISTS bookings (
    id              UUID PRIMARY KEY,
    concert_id      UUID REFERENCES concerts(id),
    batch_id         UUID REFERENCES concert_batches(id),
    number          VARCHAR,
    price           DOUBLE PRECISION,
    currency        VARCHAR(10),
    status          INT,
    reserved_at     TIMESTAMPTZ,
    paid_at         TIMESTAMPTZ,
    expired_at      TIMESTAMPTZ,
    failed_at       TIMESTAMPTZ,

    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now(),
    deleted_at      TIMESTAMPTZ,
    version         BIGINT DEFAULT 0,
);

CREATE INDEX IF NOT EXISTS idx_bookings_concert_id ON bookings (concert_id);
CREATE INDEX IF NOT EXISTS idx_bookings_batch_id ON bookings (batch_id);
CREATE INDEX IF NOT EXISTS idx_bookings_deleted_at ON bookings (deleted_at);
CREATE INDEX IF NOT EXISTS idx_bookings_expired_at ON bookings (expired_at);

ALTER TABLE bookings
  ADD COLUMN IF NOT EXISTS invoice_number VARCHAR,
  ADD COLUMN IF NOT EXISTS payment_type VARCHAR,
  ADD COLUMN IF NOT EXISTS cust_name VARCHAR NOT NULL DEFAULT '',
  ADD COLUMN IF NOT EXISTS cust_email VARCHAR NOT NULL DEFAULT '',
  ADD COLUMN IF NOT EXISTS cust_phone VARCHAR,
  ADD COLUMN IF NOT EXISTS ship_address TEXT,
  ADD COLUMN IF NOT EXISTS ship_apt TEXT,
  ADD COLUMN IF NOT EXISTS ship_city TEXT,
  ADD COLUMN IF NOT EXISTS ship_country TEXT,
  ADD COLUMN IF NOT EXISTS ship_zip TEXT,
  ADD COLUMN IF NOT EXISTS ship_state TEXT,
  ADD COLUMN IF NOT EXISTS bill_address TEXT,
  ADD COLUMN IF NOT EXISTS bill_apt TEXT,
  ADD COLUMN IF NOT EXISTS bill_city TEXT,
  ADD COLUMN IF NOT EXISTS bill_country TEXT,
  ADD COLUMN IF NOT EXISTS bill_zip TEXT,
  ADD COLUMN IF NOT EXISTS bill_state TEXT;
