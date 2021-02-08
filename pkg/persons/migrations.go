package persons

const PersonMigration = `CREATE TABLE IF NOT EXISTS person (
id serial primary key,
name varchar(150) unique not null,
phone int not null,
address varchar(150) not null
);`