CREATE TABLE `tb_rekening` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `kode_rekening` varchar(255) UNIQUE NOT NULL,
  `nama_rekening` varchar(255) NOT NULL,
  `tahun` varchar(30) NOT NULL,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP
);