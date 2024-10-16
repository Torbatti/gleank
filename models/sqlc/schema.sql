CREATE TABLE if not exists links (
  id   INTEGER PRIMARY KEY,
  
  url text NOT NULL,
  name text,
  description text,

  folder INTEGER NOT NULL
);

CREATE TABLE if not exists folders (
  id   INTEGER PRIMARY KEY,
  
  path text NOT NULL,
  name text NOT NULL,
  description text,

  user INTEGER NOT NULL,
  public BOOLEAN
);

CREATE TABLE if not exists users (
  id   INTEGER PRIMARY KEY,

  email text NOT NULL,
  name text NOT NULL,
  uuid string NOT NULL
);
