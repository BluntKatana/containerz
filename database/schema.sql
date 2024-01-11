-- initialize the database
CREATE database IF NOT EXISTS chat;
use chat;

-- create the messages table
CREATE TABLE IF NOT EXISTS messages (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255),
    content VARCHAR(255),
    likes INTEGER,
    ys INTEGER,
    created_at timestamp default current_timestamp
);

-- seed the database
INSERT INTO messages (username, content, likes, ys)
VALUES
    ('admin', 'Welcome to the chat!', 0, 0),
    ('jack', 'just setting up my y', 12, 1),
    ('elon', 'why is no one using x anymore?', 212, 349)
;