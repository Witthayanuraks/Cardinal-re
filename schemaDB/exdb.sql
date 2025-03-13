CREATE DATABASE ctf_competition;

USE ctf_competition;

-- Tabel untuk tim yang berpartisipasi dalam kompetisi
CREATE TABLE teams (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    token VARCHAR(100) UNIQUE NOT NULL,
    score INT DEFAULT 0
);

-- Tabel untuk daftar soal (challenges)
CREATE TABLE challenges (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    flag VARCHAR(255) NOT NULL,
    points INT NOT NULL
);

-- Tabel untuk menyimpan leaderboard (hasil kompetisi)
CREATE TABLE leaderboard (
    id SERIAL PRIMARY KEY,
    team_id INT NOT NULL,
    score INT DEFAULT 0,
    FOREIGN KEY (team_id) REFERENCES teams(id) ON DELETE CASCADE
);

-- Tabel untuk menyimpan admin
CREATE TABLE admins (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- Tambah admin
INSERT INTO admins (username, password) VALUES ('admin', 'kominfoo');

-- Tambah Tim
INSERT INTO teams (name, token, score) VALUES 
('Team Alpha', 'token123', 100), 
('Team Bravo', 'token456', 50), 
('Team Charlie', 'token789', 0);

-- Tambah Soal (Challenges)
INSERT INTO challenges (title, flag, points) VALUES 
('Reverse Engineering Challenge', 'flag{rev123}', 100), 
('Web Exploit', 'flag{webXploit}', 200), 
('Forensic Analysis', 'flag{forensics_2024}', 150);

-- Tambah Skor ke Leaderboard
INSERT INTO leaderboard (team_id, score) VALUES 
(1, 100), 
(2, 50), 
(3, 0);
