-- migrate:up
CREATE TABLE todos (
  id int not null AUTO_INCREMENT,
  name varchar(255) not null,
  description text not null,
  status smallint default 0,
  created_at timestamp null,
  updated_at timestamp null,
  deleted_at timestamp null,
  PRIMARY KEY (id)
);

-- migrate:down
DROP TABLE todos;

