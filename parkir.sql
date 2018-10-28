CREATE TABLE gedung (
  id ID INT NOT NULL AUTO_INCREMENT,
  nama VARCHAR(50) NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO gedung (nama) VALUES
("Gedung Kuliah Bersama Fakultas Teknik"),
("Dekanat Fakultas Teknik"),
("Fakultas Hukum");

CREATE TABLE kendaraan (
  id INT NOT NULL AUTO_INCREMENT,
  plat_nomor VARCHAR(10) NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO kendaraan (plat_nomor) VALUES
("G100AAG"), ("H1212BDG"), ("G4545AG");

CREATE TABLE petugas (
  id INT NOT NULL AUTO_INCREMENT,
  username VARCHAR(50) NOT NULL,
  password VARCHAR(255) NOT NULL,
  nama VARCHAR(50) NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO petugas (username, password, nama) VALUES
("fannyhasbi", "hasbi12345", "Fanny Hasbi"),
("someone", "someone", "Someone Amazing");

CREATE TABLE kendaraan_masuk (
  id INT NOT NULL AUTO_INCREMENT,
  waktu TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  id_petugas INT NOT NULL,
  id_kendaraan INT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (id_petugas) REFERENCES petugas(id),
  FOREIGN KEY (id_kendaraan) REFERENCES kendaraan(id)
);

INSERT INTO kendaraan_masuk (id_petugas, id_kendaraan) VALUES
(1, 1), (1, 2);

CREATE TABLE kendaraan_keluar (
  id INT NOT NULL AUTO_INCREMENT,
  waktu TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  id_petugas INT NOT NULL,
  id_kendaraan INT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (id_petugas) REFERENCES petugas(id),
  FOREIGN KEY (id_kendaraan) REFERENCES kendaraan(id)
);

INSERT INTO kendaraan_keluar (id_petugas, id_kendaraan) VALUES
(2, 1), (2, 2);