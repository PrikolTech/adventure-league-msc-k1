CREATE TABLE IF NOT EXISTS "user".role (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(20) UNIQUE NOT NULL,
    description VARCHAR(100) NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS "user".user_role (
    role_id UUID REFERENCES "user".role(id) ON DELETE CASCADE,
    data_id UUID REFERENCES "user".data(id) ON DELETE CASCADE,
    PRIMARY KEY (role_id, data_id)
);