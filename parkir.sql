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
("someone", "someone", "Someone Amazin");