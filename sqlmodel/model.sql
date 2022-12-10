create table citizen (
     id UUID DEFAULT gen_random_uuid(),
     firts_name varchar not null,
     second_name varchar not null,
     middle_name varchar not null,
     addres varchar not null,
     passport_seria  varchar not null,
     passport_number int not null
     
);

insert into citizen (firts_name, second_name, middle_name,addres, passport_seria, passport_number) values
                    ('Jahongir','Muhiddinov','Fazliddin','Namangan viloyati','AC',3072956),
                    ('Ravshanbek', 'Olimov', 'Akramjon', 'Namangan viloyati', 'AB',4565678);