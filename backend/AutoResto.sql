CREATE TABLE `role`(
    id int NOT NULL PRIMARY KEY,
    userType varchar(20)
);

CREATE TABLE `user`(
    id int PRIMARY KEY NOT NULL,
    email varchar(50),
    password varchar(50),
    name varchar(50) NOT NULL,
    phoneNumber varchar(10) NOT NULL,
    position varchar(10) NOT NULL,
    idRole int,
    FOREIGN KEY (idRole) REFERENCES role(id)
);

CREATE TABLE `material`(
    id int NOT NULL PRIMARY KEY,
    name varchar(20) NOT NULL,
    weight int NOT NULL
);

CREATE TABLE `recipe`(
    id int NOT NULL PRIMARY KEY,
    description varchar(255)
);

CREATE TABLE `food`(
    id int PRIMARY KEY NOT NULL,
    name varchar(20) NOT NULL,
    price float(8) NOT NULL,
    idMaterialFK int,
    idRecipeFK int,
    FOREIGN KEY(idMaterialFK) REFERENCES material(id),
    FOREIGN KEY(idRecipeFK) REFERENCES recipe(id)
);

CREATE TABLE `recipedetail`(
    id int NOT NULL PRIMARY KEY,
    quantity int,
    idMaterialFK int,
    idRecipeFK int,
    FOREIGN KEY (idMaterialFK) REFERENCES material(id),    
    FOREIGN KEY (idRecipeFK) REFERENCES recipe(id)
);

CREATE TABLE `inventory`(
     id int NOT NULL PRIMARY KEY,
     capacity int NOT NULL,
     loc varchar(20)
);
