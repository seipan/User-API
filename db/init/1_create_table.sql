-- DB切り替え
\c "hoge"

CREATE TABLE "users" (
  "id"                VARCHAR(255) NOT NULL PRIMARY KEY,
  "name"              VARCHAR(255) NOT NULL,
  "mail"              VARCHAR(255) NOT NULL
);

