CREATE TABLE user(
    email varchar(50),
    idUser int PRIMARY KEY NOT NULL,
    name_ varchar(50) NOT NULL,
    phoneNumber varchar(10) NOT NULL,
    position varchar(10) NOT NULL,
);

CREATE TABLE food(
    idFood varchar PRIMARY KEY NOT NULL,
    foodname varchar(20) NOT NULL,
    price float NOT NULL,
    idMaterialFK varchar,
    idRecipeFK varchar,
    PRIMARY KEY(idFood),
    FOREIGN KEY(idMaterialFK) REFERENCES material(idMaterial),
    FOREIGN KEY(idRecipeFK) REFERENCES recipe(idRecipeFK)
);

CREATE TABLE recipe(
    idRecipe varchar(5) NOT NULL PRIMARY KEY,
    descript varchar(255)

);

CREATE TABLE material(
    idMaterial varchar(5) NOT NULL PRIMARY KEY,
    names varchar(20) NOT NULL,
    weight int NOT NULL
);

CREATE TABLE recipeDetail(
    quantity int,
    idDetilRecipe varchar NOT NULL,
    idMaterialFK varchar,
    idRecipeFK varchar,
    PRIMARY KEY(idDetilRecipe),
    FOREIGN KEY (idMaterialFK) REFERENCES material(idMaterial),    
    FOREIGN KEY (idRecipeFK) REFERENCES recipe(idRecipe)

);

CREATE TABLE inventory(
     capacity int NOT NULL,
     loc varchar(20)

)

CREATE TABLE userRole(
    userType int PRIMARY KEY NOT NULL,
    idUserFK int,
    PRIMARY KEY(userType),
    FOREIGN KEY (idUserFK) REFERENCES user(idUser)

);
