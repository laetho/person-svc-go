package persons

const PersonMigration = `CREATE TABLE IF NOT EXISTS person (
id serial primary key,
name varchar(150) unique not null,
phone int not null,
address varchar(150) not null
);`

type Person struct {
	Id      int32  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address,omitempty"`
	Phone   int32  `json:"phone,omitempty"`
}

type CreatePerson struct {
	Name    string `json:"name"`
	Address string `json:"address,omitempty"`
	Phone   int32  `json:"phone,omitempty"`
}
