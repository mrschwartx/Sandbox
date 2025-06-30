-- courses table
CREATE TABLE IF NOT EXISTS courses (
    id            UUID PRIMARY KEY,
    name          VARCHAR(255),
    slug          VARCHAR(255) UNIQUE,
    description   TEXT,
    status        INT,
    published_at  TIMESTAMPTZ,
    created_at    TIMESTAMPTZ DEFAULT now(),
    updated_at    TIMESTAMPTZ DEFAULT now(),
    deleted_at    TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_courses_deleted_at     ON courses (deleted_at);
CREATE INDEX IF NOT EXISTS idx_courses_published_at   ON courses (published_at);
CREATE INDEX IF NOT EXISTS idx_courses_status         ON courses (status);

-- course_batches table
CREATE TABLE IF NOT EXISTS course_batches (
    id              UUID PRIMARY KEY,
    course_id       UUID REFERENCES courses(id),
    name            VARCHAR(255),
    max_seats       INT NOT NULL DEFAULT 0,
    available_seats INT,
    price           DOUBLE PRECISION,
    currency        VARCHAR(10),
    status          INT,
    start_date      TIMESTAMPTZ,
    end_date        TIMESTAMPTZ,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now(),
    deleted_at      TIMESTAMPTZ,
    version         BIGINT DEFAULT 0
);

CREATE INDEX IF NOT EXISTS idx_course_batches_created_at   ON course_batches (created_at);
CREATE INDEX IF NOT EXISTS idx_course_batches_deleted_at   ON course_batches (deleted_at);
CREATE INDEX IF NOT EXISTS idx_course_batches_course_id    ON course_batches (course_id);
CREATE INDEX IF NOT EXISTS idx_course_batches_status       ON course_batches (status);

-- bookings table
CREATE TABLE IF NOT EXISTS bookings (
    id              UUID PRIMARY KEY,
    course_id       UUID REFERENCES courses(id),
    course_batch_id UUID REFERENCES course_batches(id),
    price           DOUBLE PRECISION,
    currency        VARCHAR(10),
    status          INT,
    reserved_at     TIMESTAMPTZ,
    paid_at         TIMESTAMPTZ,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now(),
    deleted_at      TIMESTAMPTZ,
    version         BIGINT DEFAULT 0
);

CREATE INDEX IF NOT EXISTS idx_bookings_deleted_at        ON bookings (deleted_at);
CREATE INDEX IF NOT EXISTS idx_bookings_course_id         ON bookings (course_id);
CREATE INDEX IF NOT EXISTS idx_bookings_course_batch_id   ON bookings (course_batch_id);

ALTER TABLE bookings
    ADD COLUMN IF NOT EXISTS invoice_number   VARCHAR,
    ADD COLUMN IF NOT EXISTS payment_type     VARCHAR,
    ADD COLUMN IF NOT EXISTS expired_at       TIMESTAMPTZ,
    ADD COLUMN IF NOT EXISTS cust_name        VARCHAR NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS cust_email       VARCHAR NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS cust_phone       VARCHAR;