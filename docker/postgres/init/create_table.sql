\c application_server;

-- usersテーブル
CREATE TABLE users (
    id VARCHAR(26) PRIMARY KEY,
    name VARCHAR(50),
    mail VARCHAR(255),
    gender VARCHAR(10),
    age INT,
    occupation VARCHAR(50),
    password VARCHAR(50)
);

-- universitiesテーブル
CREATE TABLE universities　(
    id VARCHAR(26) PRIMARY KEY,
    -- 大学のID(既存)
    university_id VARCHAR(26),
    name VARCHAR(50)
);

-- undergraduatesテーブル
CREATE TABLE undergraduates (
    id VARCHAR(26) PRIMARY KEY,
    university_id VARCHAR(26) REFERENCES universities(id),
    name VARCHAR(50),
    department　VARCHAR(50),
    major VARCHAR(50)
);

--locationsテーブル
CREATE TABLE locations (
    id VARCHAR(26) PRIMARY KEY,
    building VARCHAR(50),
    room VARCHAR(50)
);

-- laboratoriesテーブル
CREATE TABLE laboratories (
    id VARCHAR(26) PRIMARY KEY,
    user_id VARCHAR(26) REFERENCES users(id),
    undergraduate_id VARCHAR REFERENCES undergraduates(id),
    location_id VARCHAR REFERENCES locations(id),
    name VARCHAR(50),
    homepage VARCHAR(50),

);

-- objectsテーブル
CREATE TABLE objects (
    id VARCHAR(26) PRIMARY KEY,
    lab_id VARCHAR(26) REFERENCES laboratories(id),
    aspect INT,
    height INT,
    size INT
);

