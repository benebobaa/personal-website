-- Admin table
CREATE TABLE IF NOT EXISTS admins (
                                      id SERIAL PRIMARY KEY,
                                      username VARCHAR(255) NOT NULL UNIQUE,
                                        password VARCHAR(255) NOT NULL
    );

-- Hero table
CREATE TABLE IF NOT EXISTS heroes (
                                      id SERIAL PRIMARY KEY,
                                      first_name VARCHAR(255) NOT NULL,
                                        last_name VARCHAR(255) NOT NULL,
                                        description TEXT NOT NULL
    );

-- Socmed table
CREATE TABLE IF NOT EXISTS socmeds (
                                       id SERIAL PRIMARY KEY,
                                       name VARCHAR(255) NOT NULL,
                                        link VARCHAR(255) NOT NULL
    );