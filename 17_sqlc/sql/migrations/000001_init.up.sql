CREATE TABLE categories (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT
);

CREATE TABLE courses (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    category_id VARCHAR(36),
    price DECIMAL(10,2),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);
