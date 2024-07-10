CREATE TABLE quotes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    code TEXT,
    codein TEXT,
    name TEXT,
    high REAL,
    low REAL,
    varBid REAL,
    pctChange TEXT,
    bid REAL,
    ask REAL,
    timestamp TEXT,
    createDate TEXT
);