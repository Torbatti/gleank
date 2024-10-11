-- https://www.sqlite.org/datatype3.html
CREATE TABLE links (
  id   INTEGER PRIMARY KEY,
  
  url text NOT NULL,
  name text,
  description text,

  folder INTEGER NOT NULL
);

CREATE TABLE folders (
  id   INTEGER PRIMARY KEY,
  
  url text NOT NULL,
  name text NOT NULL,
  description text,

  user INTEGER NOT NULL,
  public BOOLEAN
);

CREATE TABLE users (
  id   INTEGER PRIMARY KEY,

  name text NOT NULL,
  uuid string
);

-- CREATE TABLE link_folders (
--   id   INTEGER PRIMARY KEY,
--   link_id   INTEGER NOT NULL,
--   folder_id   INTEGER NOT NULL
-- );

-- CREATE TABLE user_folders (
--   id   INTEGER PRIMARY KEY,
--   user_id   INTEGER NOT NULL,
--   folder_id   INTEGER NOT NULL
-- );