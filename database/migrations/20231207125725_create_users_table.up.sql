CREATE TABLE IF NOT EXISTS pessoas (
   id serial PRIMARY KEY,
   apelido VARCHAR(32) UNIQUE NOT NULL,
   nome VARCHAR(100) NOT NULL,
   nascimento DATE NOT NULL,
   stack VARCHAR(32)[]
);
