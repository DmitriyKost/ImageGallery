CREATE TABLE IF NOT EXISTS images (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    path TEXT NOT NULL,
    description TEXT
);
query_separator
CREATE TABLE IF NOT EXISTS videos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    path TEXT NOT NULL,
    description TEXT
);
