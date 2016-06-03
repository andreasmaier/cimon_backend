CREATE TABLE servers
(
  id INT(11) PRIMARY KEY NOT NULL,
  url VARCHAR(512)
);

CREATE TABLE jobs
(
  id INT(11) PRIMARY KEY NOT NULL,
  path VARCHAR(255),
  status VARCHAR(255)
);
CREATE INDEX jobs_path_index ON jobs (path);