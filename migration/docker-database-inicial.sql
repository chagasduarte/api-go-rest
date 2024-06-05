create table contas(
    id serial primary key,
    nome varchar,
    valor decimal,
    tipo int
)

insert into contas (nome, valor, tipo) values ("Santander", 350.00, 1)