-- phpMyAdmin SQL Dump
-- version 4.6.6deb5
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: 23 Jul 2018 pada 08.17
-- Versi Server: 10.1.29-MariaDB-6
-- PHP Version: 5.6.36-1+ubuntu18.04.1+deb.sury.org+1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_shopee_exchange`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `sp_exchange_rates`
--

CREATE TABLE `sp_exchange_rates` (
  `id` int(10) UNSIGNED NOT NULL,
  `exchange_date` date NOT NULL,
  `currency_from` char(3) COLLATE utf8mb4_unicode_ci NOT NULL,
  `currency_to` char(3) COLLATE utf8mb4_unicode_ci NOT NULL,
  `rate` decimal(12,6) UNSIGNED NOT NULL DEFAULT '0.000000',
  `created_by_ip` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL,
  `updated_by_ip` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data untuk tabel `sp_exchange_rates`
--

INSERT INTO `sp_exchange_rates` (`id`, `exchange_date`, `currency_from`, `currency_to`, `rate`, `created_by_ip`, `updated_by_ip`, `created_at`, `updated_at`) VALUES
(1, '2018-07-01', 'USD', 'GBP', '0.100000', '172.19.0.1', '172.19.0.1', '2018-07-20 10:05:03', '2018-07-20 10:05:03'),
(2, '2018-07-01', 'USD', 'GBP', '0.650000', '172.19.0.1', '172.19.0.1', '2018-07-20 10:06:40', '2018-07-20 10:06:40'),
(3, '2018-07-01', 'SGP', 'GBP', '0.757090', '172.19.0.1', '172.19.0.1', '2018-07-20 10:07:20', '2018-07-20 10:07:20'),
(5, '2018-07-01', 'USD', 'GBP', '0.100000', '172.8.0.1', '172.8.0.1', '2018-07-23 00:52:51', '2018-07-23 00:52:51'),
(6, '2018-07-02', 'SGP', 'IDR', '14300.000000', '172.8.0.1', '172.8.0.1', '2018-07-23 01:05:26', '2018-07-23 01:05:26'),
(7, '2018-07-01', 'GBP', 'USD', '1.314233', '172.8.0.1', '172.8.0.1', '2018-07-23 01:07:02', '2018-07-23 01:07:02'),
(8, '2018-07-02', 'GBP', 'USD', '1.214233', '172.8.0.1', '172.8.0.1', '2018-07-23 01:07:32', '2018-07-23 01:07:32'),
(10, '2018-07-02', 'USD', 'IDR', '14200.400000', '172.8.0.1', '172.8.0.1', '2018-07-23 01:12:11', '2018-07-23 01:12:11'),
(11, '2018-07-02', 'GBP', 'IDR', '8200.400000', '172.8.0.1', '172.8.0.1', '2018-07-23 01:12:12', '2018-07-23 01:12:12');

-- --------------------------------------------------------

--
-- Struktur dari tabel `sp_users`
--

CREATE TABLE `sp_users` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `username` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status_active` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '1:inactive; 2:active;3:banned;',
  `created_by_ip` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL,
  `updated_by_ip` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data untuk tabel `sp_users`
--

INSERT INTO `sp_users` (`id`, `name`, `email`, `username`, `password`, `status_active`, `created_by_ip`, `updated_by_ip`, `created_at`, `updated_at`) VALUES
(1, '', 'maulanateguh87@gmail.com', 'teguh', '$2a$14$zQfAMepBDuR1J3oHpUSJnepYaDVatr9X1ojSOHfpKuGF6nxwGJUii', 2, '172.8.0.1', '172.8.0.1', '2018-07-23 00:00:55', '2018-07-23 00:00:55');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `sp_exchange_rates`
--
ALTER TABLE `sp_exchange_rates`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `sp_users`
--
ALTER TABLE `sp_users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `users_email_unique` (`email`),
  ADD UNIQUE KEY `users_username_unique` (`username`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `sp_exchange_rates`
--
ALTER TABLE `sp_exchange_rates`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;
--
-- AUTO_INCREMENT for table `sp_users`
--
ALTER TABLE `sp_users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
