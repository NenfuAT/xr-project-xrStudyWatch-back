\c application_server;

-- usersテーブルにVARCHAR(26)をPrimary Keyとしてサンプルデータを挿入
CREATE TABLE users (
    id VARCHAR(26) PRIMARY KEY,
    name VARCHAR(50),
    mail VARCHAR(255),
    gender VARCHAR(10),
    age INT,
    height DECIMAL(5, 2),
    weight DECIMAL(5, 2),
    address VARCHAR(255),
    occupation VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

