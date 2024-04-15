CREATE TABLE IF NOT EXISTS athletes (
    slug TEXT NOT NULL PRIMARY KEY,
    name TEXT COLLATE NOCASE,
    gender TEXT,
    age INTEGER,
    sport_id INTEGER,
    famous_for TEXT
);

CREATE TABLE IF NOT EXISTS sports (
    id INTEGER NOT NULL PRIMARY KEY autoincrement,
    name TEXT COLLATE NOCASE
);


INSERT OR IGNORE INTO sports (id, name) VALUES (1, 'Pugilato');
INSERT OR IGNORE INTO sports (id, name) VALUES (2, 'Corsa');
INSERT OR IGNORE INTO sports (id, name) VALUES (3, 'Lotta');
INSERT OR IGNORE INTO sports (id, name) VALUES (4, 'Danza');
INSERT OR IGNORE INTO sports (id, name) VALUES (5, 'Lancio del disco');
INSERT OR IGNORE INTO sports (id, name) VALUES (6, 'Salto in lungo');
INSERT OR IGNORE INTO sports (id, name) VALUES (7, 'Lancio del giavellotto');

INSERT OR IGNORE INTO
    athletes (slug, name, gender, age, sport_id, famous_for)
VALUES (
    'thalia-di-mileto',
    'Thalia di Mileto',
    'F',
    25,
    6,
    'Famosa per la sua destrezza nel salto in lungo.'
);
INSERT OR IGNORE INTO
    athletes (slug, name, gender, age, sport_id, famous_for)
VALUES (
    'callias-di-atene',
    'Callias di Atene',
    'M',
    24,
    4,
    'Rinomato per la sua destrezza nella danza.'
);
INSERT OR IGNORE INTO
    athletes (slug, name, gender, age, sport_id, famous_for)
VALUES (
    'astioco-di-taranto',
    'Astioco di Taranto',
    'M',
    25,
    3,
    'Conosciuto per la sua forza e abilità nella lotta.'
);
INSERT OR IGNORE INTO
    athletes (slug, name, gender, age, sport_id, famous_for)
VALUES (
    'chryseis-di-larissa',
    'Chryseis di Larissa',
    'F',
    24,
    7,
    'Famosa per la sua agilità e velocità nel lancio del giavellotto.'
);
INSERT OR IGNORE INTO
    athletes (slug, name, gender, age, sport_id, famous_for)
VALUES (
    'arrhichione-di-fliunte',
    'Arrhichione di Fliunte',
    'M',
    26,
    3,
    'Conosciuto per il suo successo nella lotta e il suo sacrificio durante una competizione.'
);
INSERT OR IGNORE INTO
    athletes (slug, name, gender, age, sport_id, famous_for)
VALUES (
    'diagora-di-rodi',
    'Diagora di Rodi',
    'M',
    32,
    1,
    'Celebre per il suo successo nel pugilato.'
);
INSERT OR IGNORE INTO
    athletes (slug, name, gender, age, sport_id, famous_for)
VALUES (
    'eurydice-di-corinto',
    'Eurydice di Corinto',
    'F',
    25,
    2,
    'Rinomata per la sua abilità nella corsa.'
);
INSERT OR IGNORE INTO
    athletes (slug, name, gender, age, sport_id, famous_for)
VALUES (
    'ismene-di-corinto',
    'Ismene di Corinto',
    'F',
    23,
    4,
    'Celebre per la sua agilità e precisione nella danza rituale.'
);
INSERT OR IGNORE INTO
    athletes (slug, name, gender, age, sport_id, famous_for)
VALUES (
    'leonida-di-rodi',
    'Leonida di Rodi',
    'M',
    28,
    2,
    'Celebrato per le sue abilità nella corsa.'
);
INSERT OR IGNORE INTO
    athletes (slug, name, gender, age, sport_id, famous_for)
VALUES (
    'milone-di-crotone',
    'Milone di Crotone',
    'M',
    30,
    1,
    'Famoso per le sue vittorie ripetute nel pugilato.'
);
INSERT OR IGNORE INTO
    athletes (slug, name, gender, age, sport_id, famous_for)
VALUES (
    'lysander-di-tebe',
    'Lysander di Tebe',
    'M',
    26,
    5,
    'Conosciuto per la sua abilità nel lancio del disco.'
);
INSERT OR IGNORE INTO
    athletes (slug, name, gender, age, sport_id, famous_for)
VALUES (
    'melankomas-di-caria',
    'Melankomas di Caria',
    'M',
    29,
    1,
    'Rinomato per le sue abilità nel pugilato.'
);
INSERT OR IGNORE INTO
    athletes (slug, name, gender, age, sport_id, famous_for)
VALUES (
    'myrto-di-salamina',
    'Myrto di Salamina',
    'F',
    22,
    2,
    'Celebre per la sua velocità e agilità nella corsa.'
);
INSERT OR IGNORE INTO
    athletes (slug, name, gender, age, sport_id, famous_for)
VALUES (
    'theodora-di-creta',
    'Theodora di Creta',
    'F',
    26,
    5,
    'Famosa per la sua abilità nel lancio del disco.'
);
INSERT OR IGNORE INTO
    athletes (slug, name, gender, age, sport_id, famous_for)
VALUES (
    'polidamo-di-skotoussa',
    'Polidamo di Skotoussa',
    'M',
    27,
    2,
    'Famoso per la sua resistenza e velocità nella corsa.'
);
