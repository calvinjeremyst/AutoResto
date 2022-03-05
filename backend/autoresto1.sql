-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 05 Mar 2022 pada 08.40
-- Versi server: 10.4.21-MariaDB
-- Versi PHP: 8.0.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `autoresto`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `food`
--

CREATE TABLE `food` (
  `id` int(11) NOT NULL,
  `name` varchar(128) NOT NULL,
  `price` float NOT NULL,
  `idRecipeFK` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `food`
--

INSERT INTO `food` (`id`, `name`, `price`, `idRecipeFK`) VALUES
(1, 'Nasi Goreng Spesial', 20000, 1);

-- --------------------------------------------------------

--
-- Struktur dari tabel `inventory`
--

CREATE TABLE `inventory` (
  `id` int(11) NOT NULL,
  `capacity` int(11) NOT NULL,
  `location` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `inventory`
--

INSERT INTO `inventory` (`id`, `capacity`, `location`) VALUES
(1, 200, 1);

-- --------------------------------------------------------

--
-- Struktur dari tabel `location`
--

CREATE TABLE `location` (
  `id` int(11) NOT NULL,
  `loc_name` varchar(128) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `location`
--

INSERT INTO `location` (`id`, `loc_name`) VALUES
(1, 'WEST 1');

-- --------------------------------------------------------

--
-- Struktur dari tabel `material`
--

CREATE TABLE `material` (
  `id` int(11) NOT NULL,
  `name` varchar(128) NOT NULL,
  `quantity` float NOT NULL,
  `unit` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `material`
--

INSERT INTO `material` (`id`, `name`, `quantity`, `unit`) VALUES
(1, 'Beras', 10, 'KG'),
(2, 'Bawang Putih', 5, 'KG'),
(3, 'Bawang Merah', 5, 'KG'),
(4, 'Kecap', 2, 'KG'),
(5, 'Garam', 2, 'KG'),
(6, 'Micin', 2, 'KG'),
(7, 'Ayam', 10, 'KG');

-- --------------------------------------------------------

--
-- Struktur dari tabel `recipe`
--

CREATE TABLE `recipe` (
  `id` int(11) NOT NULL,
  `description` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `recipe`
--

INSERT INTO `recipe` (`id`, `description`) VALUES
(1, 'Nasi goreng rempah dengan ayam');

-- --------------------------------------------------------

--
-- Struktur dari tabel `recipedetail`
--

CREATE TABLE `recipedetail` (
  `id` int(11) NOT NULL,
  `quantity` int(11) DEFAULT NULL,
  `idMaterialFK` int(11) DEFAULT NULL,
  `idRecipeFK` int(11) DEFAULT NULL,
  `unit` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `recipedetail`
--

INSERT INTO `recipedetail` (`id`, `quantity`, `idMaterialFK`, `idRecipeFK`, `unit`) VALUES
(1, 300, 1, 1, 'Gr'),
(2, 3, 2, 1, 'Pcs'),
(3, 3, 3, 1, 'Pcs'),
(4, 5, 4, 1, 'Sendok'),
(5, 2, 5, 1, 'Sendok Kecil'),
(6, 1, 6, 1, 'Sendok Kecil'),
(7, 30, 7, 1, 'Gr');

-- --------------------------------------------------------

--
-- Struktur dari tabel `user`
--

CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `email` varchar(128) DEFAULT NULL,
  `password` varchar(128) DEFAULT NULL,
  `name` varchar(128) NOT NULL,
  `phoneNumber` varchar(50) NOT NULL,
  `position` varchar(10) NOT NULL,
  `idRole` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `user`
--

INSERT INTO `user` (`id`, `email`, `password`, `name`, `phoneNumber`, `position`, `idRole`) VALUES
(1, 'aji@gmail.com', 'aji', 'Aji Parama', '08154653', '1', 1),
(2, 'calvin@gmail.com', 'calvin', 'Calvin Jeremy', '081541534', '2', 2),
(3, 'christian@gmail.com', 'christian', 'Christian Satya', '081573682750', '2', 2),
(4, 'juniar@gmail.com', 'juniar`', 'William Juniar', '0812535435', '3', 3),
(5, 'kolis@gmail.com', 'kolis', 'William Kolis', '081534354', '3', 3);

-- --------------------------------------------------------

--
-- Struktur dari tabel `userrole`
--

CREATE TABLE `userrole` (
  `id` int(11) NOT NULL,
  `userType` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `userrole`
--

INSERT INTO `userrole` (`id`, `userType`) VALUES
(1, 'OWNER'),
(2, 'CHEF'),
(3, 'INVENTORY_PERSON');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `food`
--
ALTER TABLE `food`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idRecipeFK` (`idRecipeFK`) USING BTREE;

--
-- Indeks untuk tabel `inventory`
--
ALTER TABLE `inventory`
  ADD PRIMARY KEY (`id`),
  ADD KEY `location` (`location`);

--
-- Indeks untuk tabel `location`
--
ALTER TABLE `location`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `material`
--
ALTER TABLE `material`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `recipe`
--
ALTER TABLE `recipe`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `recipedetail`
--
ALTER TABLE `recipedetail`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idMaterialFK` (`idMaterialFK`),
  ADD KEY `idRecipeFK` (`idRecipeFK`);

--
-- Indeks untuk tabel `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idRole` (`idRole`);

--
-- Indeks untuk tabel `userrole`
--
ALTER TABLE `userrole`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `inventory`
--
ALTER TABLE `inventory`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT untuk tabel `recipedetail`
--
ALTER TABLE `recipedetail`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT untuk tabel `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `inventory`
--
ALTER TABLE `inventory`
  ADD CONSTRAINT `inventory_ibfk_1` FOREIGN KEY (`location`) REFERENCES `location` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Ketidakleluasaan untuk tabel `recipedetail`
--
ALTER TABLE `recipedetail`
  ADD CONSTRAINT `recipedetail_ibfk_1` FOREIGN KEY (`idMaterialFK`) REFERENCES `material` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `recipedetail_ibfk_2` FOREIGN KEY (`idRecipeFK`) REFERENCES `recipe` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Ketidakleluasaan untuk tabel `user`
--
ALTER TABLE `user`
  ADD CONSTRAINT `user_ibfk_1` FOREIGN KEY (`idRole`) REFERENCES `userrole` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
