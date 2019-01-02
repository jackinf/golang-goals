DROP TABLE IF EXISTS goals;
DROP TABLE IF EXISTS goal;

CREATE TABLE goal
(
    id SERIAL PRIMARY KEY,
    title VARCHAR(160)  NOT NULL,
    description VARCHAR(500) NOT NULL,
    due DATE,
    motivation VARCHAR(500)
);

INSERT INTO goal (title, description, due, motivation)
VALUES ('6 pack', 'To have strong and beautiful body', current_date, 'Strong body makes one happier');