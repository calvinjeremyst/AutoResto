CREATE TABLE `inventory`(
  id int(11) AUTO_INCREMENT,
  capacity int(11) NOT NULL,
  location int(11) NOT NULL,
  PRIMARY KEY (id)
); 

CREATE TABLE `location`(
  id int(11) AUTO_INCREMENT,
  loc_name varchar(128) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE `material`(
  id int(11) AUTO_INCREMENT,
  name varchar(128) NOT NULL,
  quantity float NOT NULL,
  unit varchar(50) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE `recipe`(
  id int(11) AUTO_INCREMENT,
  description varchar(255) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE `recipedetail`(
  id int(11) AUTO_INCREMENT,
  quantity int(11) NOT NULL,
  idMaterialFK int(11) DEFAULT NULL,
  idRecipeFK int(11) DEFAULT NULL,
  unit varchar(50) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (idMaterialFK) REFERENCES material(id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (idRecipeFK) REFERENCES recipe(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE `menu`(
  id int(11) AUTO_INCREMENT,
  name varchar(128) NOT NULL,
  price float NOT NULL,
  idRecipeFK int(11) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (idRecipeFK) REFERENCES recipe(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE `userrole`(
  id int(11) AUTO_INCREMENT,
  userType varchar(20) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE `user`(
  id int(11) AUTO_INCREMENT,
  email varchar(128) NOT NULL,
  password varchar(128) NOT NULL,
  name varchar(128) NOT NULL,
  phoneNumber varchar(50) NOT NULL,
  position varchar(10) NOT NULL,
  idRole int(11) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (idRole) REFERENCES userrole(id) ON DELETE CASCADE ON UPDATE CASCADE
);
