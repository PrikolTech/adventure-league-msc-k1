CREATE TABLE IF NOT EXISTS form.registration (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(100) UNIQUE NOT NULL,
    initiator_first_name VARCHAR(100) NOT NULL,
    initiator_last_name VARCHAR(100) NOT NULL,
    initiator_patronymic VARCHAR(100),
    birthdate DATE NOT NULL,
    supervisor_first_name VARCHAR(100) NOT NULL,
    supervisor_last_name VARCHAR(100) NOT NULL,
    supervisor_patronymic VARCHAR(100),
    department VARCHAR(150) NOT NULL,
    post VARCHAR(100) NOT NULL,
    history TEXT,
    achievements TEXT,
    motivation TEXT,
    phone VARCHAR(13),
    telegram VARCHAR(100)
);