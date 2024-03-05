-- migrate:up
CREATE TABLE roominfo (ID SERIAL PRIMARY KEY, RooName VARCHAR(50) NOT NULL, CreatedBy INT64 NOT NULL REFERENCES users (ID), Members INT64[] NOT NULL);
CREATE TABLE roomchat (ID SERIAL PRIMARY KEY, SenderID INT64 NOT NULL REFERENCES users (ID), Message VARCHAR(50) NOT NULL, Time TIMESTAMP DEFAULT CURRENT_TIMESTAMP, RoomID INT64 NOT NULL REFERENCES roominfo (ID));

-- migrate:down
DROP TABLE IF EXISTS roominfo
DROP TABLE IF EXISTS roomchat
