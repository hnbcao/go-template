-- name: create-table-user_namespace

CREATE TABLE user_namespace (
  id int(11) NOT NULL AUTO_INCREMENT,
  uid int(11) NOT NULL,
  namespace varchar(255) NOT NULL,
  status tinyint(1) NOT NULL DEFAULT '1',
  create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
