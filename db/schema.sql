CREATE TABLE posts(
  id UUID PRIMARY KEY,
  title TEXT NOT NULL,
  slug TEXT NOT NULL,
  body TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  published_at TIMESTAMP
);

CREATE TABLE post_tags(
    id UUID PRIMARY KEY,
    post_id UUID REFERENCES posts(id),
    label varchar(64) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE photos(
    id UUID PRIMARY KEY,
    alt_text TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);