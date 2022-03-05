CREATE TABLE `userrole`(
    id int PRIMARY KEY,
    userType varchar(20) NOT NULL
);

CREATE TABLE `user`(
    id int PRIMARY KEY,
    email varchar(128),
    password varchar(128),
    name varchar(128) NOT NULL,
    phoneNumber varchar(50) NOT NULL,
    position varchar(10) NOT NULL,
    idRole int,
    FOREIGN KEY (idRole) REFERENCES userrole(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE `material`(
    id int PRIMARY KEY,
    name varchar(128) NOT NULL,
    quantity int(10) NOT NULL,
    unit varchar(10)
);

CREATE TABLE `recipe`(
    id int PRIMARY KEY,
    description varchar(255)
);

CREATE TABLE `recipedetail`(
    id int PRIMARY KEY,
    quantity int,
    idMaterialFK int,
    idRecipeFK int,
    FOREIGN KEY (idMaterialFK) REFERENCES material(id) ON UPDATE CASCADE ON DELETE CASCADE,    
    FOREIGN KEY (idRecipeFK) REFERENCES recipe(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE `food`(
    id int PRIMARY KEY,
    name varchar(128) NOT NULL,
    price float(8) NOT NULL,
    idMaterialFK int,
    idRecipeFK int,
    FOREIGN KEY(idMaterialFK) REFERENCES material(id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY(idRecipeFK) REFERENCES recipe(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE `location`(
	id int PRIMARY KEY,
	loc_name varchar(128) NOT NULL
);

CREATE TABLE `inventory`(
     id int PRIMARY KEY,
     capacity int NOT NULL,
     location int,
	 FOREIGN KEY(location) REFERENCES location(id) ON UPDATE CASCADE ON DELETE CASCADE
);