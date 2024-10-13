CREATE TABLE links (
  id   INTEGER PRIMARY KEY,
  
  url text NOT NULL,
  name text,
  description text,

  folder INTEGER NOT NULL
);

CREATE TABLE folders (
  id   INTEGER PRIMARY KEY,
  
  path text NOT NULL,
  name text NOT NULL,
  description text,

  user INTEGER NOT NULL,
  public BOOLEAN
);

CREATE TABLE users (
  id   INTEGER PRIMARY KEY,

  name text NOT NULL,
  uuid string NOT NULL
);
