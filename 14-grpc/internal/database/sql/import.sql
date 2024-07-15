CREATE TABLE categories (
    id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255),
    description TEXT
);
CREATE TABLE courses (
    id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255),
    description TEXT,
    category_id uuid,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);