INSERT INTO `inventory`(`capacity`, `location`) VALUES(200, 1);

INSERT INTO `location`(`loc_name`) VALUES('WEST 1');

INSERT INTO `material`(`name`, `quantity`, `unit`) VALUES
('Beras', 10, 'KG'),
('Bawang Putih', 5, 'KG'),
('Bawang Merah', 5, 'KG'),
('Kecap', 2, 'KG'),
('Garam', 2, 'KG'),
('Micin', 2, 'KG'),
('Ayam', 10, 'KG');

INSERT INTO `recipe`(`description`) VALUES
('Nasi goreng rempah dengan ayam');

INSERT INTO `menu`(`name`, `price`, `idRecipeFK`) VALUES
('Nasi Goreng Spesial', 20000, 1);

INSERT INTO `recipedetail`(`quantity`, `idMaterialFK`, `idRecipeFK`, `unit`) VALUES
(300, 1, 1, 'Gr'),
(3, 2, 1, 'Pcs'),
(3, 3, 1, 'Pcs'),
(5, 4, 1, 'Sendok'),
(2, 5, 1, 'Sendok Kecil'),
(1, 6, 1, 'Sendok Kecil'),
(30, 7, 1, 'Gr');

INSERT INTO `userrole`(`id`, `userType`) VALUES
(1, 'OWNER'),
(2, 'CHEF'),
(3, 'INVENTORY_PERSON');

INSERT INTO `user`(`email`, `password`, `name`, `phoneNumber`, `position`, `idRole`) VALUES
('aji@gmail.com', 'aji', 'Aji Parama', '08154653', '1', 1),
('calvin@gmail.com', 'calvin', 'Calvin Jeremy', '081541534', '2', 2),
('christian@gmail.com', 'christian', 'Christian Satya', '081573682750', '2', 2),
('juniar@gmail.com', 'juniar`', 'William Juniar', '0812535435', '3', 3),
('kolis@gmail.com', 'kolis', 'William Kolis', '081534354', '3', 3);