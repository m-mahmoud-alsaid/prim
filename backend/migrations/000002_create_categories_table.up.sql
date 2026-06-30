CREATE TABLE categories (
  id UUID PRIMARY KEY NOT NULL,
  name TEXT UNIQUE NOT NULL,
  slug TEXT UNIQUE NOT NULL,
  parent_id UUID,
  deleted_at TIMESTAMPTZ NULL,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL,
  FOREIGN KEY (parent_id) REFERENCES categories(id)
);

CREATE INDEX IF NOT EXISTS categories_slug_idx ON categories(slug)
    WHERE deleted_at IS NULL;
