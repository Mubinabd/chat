CREATE TYPE role_type AS enum ('user','admin');
-- USER TABLE
CREATE TABLE users (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(100),
    date_of_birth DATE,
    role role_type NOT NULL DEFAULT 'user',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

-- SETTING TABLE
CREATE TABLE settings (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    privacy_level VARCHAR(50) NOT NULL DEFAULT 'private',
    notification VARCHAR(30) NOT NULL DEFAULT 'on',
    language VARCHAR(255) NOT NULL DEFAULT 'en',
    theme VARCHAR(255) NOT NULL DEFAULT 'light',
    updated_at TIMESTAMP DEFAULT NOW()
);

-- TOKEN
CREATE TABLE tokens (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    token VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

-- Groups table
CREATE TABLE groups (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    members JSONB, 
    messages JSONB,
    files TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

-- Group Members table 
CREATE TABLE group_members (
    group_id UUID REFERENCES groups(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (group_id, user_id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

-- Messages table
CREATE TABLE messages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sender_id UUID REFERENCES users(id) ON DELETE CASCADE,
    group_id UUID REFERENCES groups(id) ON DELETE CASCADE,
    content TEXT,
    file_url TEXT,
    timestamp BIGINT NOT NULL
);