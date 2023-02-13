create table if not exists users (
    id int not null auto_increment,
    username varchar(255) not null,
    password varchar(255) not null,
    date_created datetime not null,
    primary key(id)
);