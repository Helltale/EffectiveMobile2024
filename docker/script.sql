
CREATE TABLE IF NOT EXISTS songs (
    id SERIAL PRIMARY KEY,              
    group_name VARCHAR(255) NOT NULL,   
    title VARCHAR(255) NOT NULL,        
    release_date DATE,                  
    text TEXT,                          
    link VARCHAR(255)                   
);

-- INSERT INTO songs (group_name, title, release_date, text, link) VALUES
-- ('Muse', 'Supermassive Black Hole', '2006-07-16', 'Ooh baby, don''t you know I suffer? Ooh baby, can you hear me moan? You caught me under false pretenses How long before you let me go? Ooh You set my soul alight Ooh You set my soul alight', 'https://www.youtube.com/watch?v=Xsp3_a-PMTw');
