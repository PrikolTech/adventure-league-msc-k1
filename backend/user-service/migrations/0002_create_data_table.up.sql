CREATE TABLE IF NOT EXISTS "user".data (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(100) UNIQUE NOT NULL,
    password CHAR(60) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    patronymic VARCHAR(100),
    birthdate DATE NOT NULL,
    phone VARCHAR(13),
    telegram VARCHAR(100)
);