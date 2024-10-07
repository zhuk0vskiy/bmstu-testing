create table if not exists "user"
(
    id serial primary key not null ,
    login text not null ,
    password text not null ,
    role int not null ,
    first_name text not null ,
    second_name text not null ,
    third_name text not null
);

create table if not exists studio
(
    id serial primary key not null ,
    name text not null
);

create table if not exists room
(
    id serial primary key not null ,
    name text not null ,
    studio_id int not null,
    start_hour int not null ,
    end_hour int not null
);

create table if not exists equipment
(
    id serial primary key not null ,
    name text not null ,
    type int not null ,
    studio_id int not null
);

create table if not exists producer
(
    id serial primary key not null ,
    name text not null ,
    studio_id int not null ,
    start_hour int not null ,
    end_hour int not null
);

create table if not exists instrumentalist
(
    id serial primary key not null ,
    name text not null ,
    studio_id int not null ,
    start_hour int not null ,
    end_hour int not null
);


create table if not exists reserve
(
    id serial primary key not null ,
    user_id int not null ,
    room_id int not null ,
    producer_id int,
    instrumentalist_id int,
    start_time timestamp not null ,
    end_time timestamp not null
);

create table if not exists reserved_equipments
(
    id serial primary key not null ,
    reserve_id int not null ,
    equipment_id int not null
);
