-- THIS IS SCRIPT FOR CREATING USERS TABLE
CREATE TABLE users (
   uuid UUID PRIMARY KEY NOT NULL,
   name VARCHAR(255) NOT NULL,
   email VARCHAR(255) NOT NULL,
   password TEXT NOT NULL,
   role VARCHAR(10) NOT NULL,
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP,
   created_by VARCHAR(255),
   updated_by VARCHAR(255),
   deleted_by VARCHAR(255),
   CONSTRAINT unique_email UNIQUE (email)
);

-- THIS IS SCRIPT FOR CREATING TOKENS TABLE
CREATE TABLE tokens (
   uuid UUID PRIMARY KEY NOT NULL,
   user_uuid UUID NOT NULL,
   token TEXT NOT NULL,
   is_active BOOLEAN NOT NULL,
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   created_by VARCHAR(255),
   updated_by VARCHAR(255),
   FOREIGN KEY (user_uuid) REFERENCES users(uuid)
);