CREATE USER postgres SUPERUSER;

create extension if not exists pgcrypto;

create schema if not exists saladRecipes;
create schema if not exists keywords;

create table if not exists keywords.word (
    id uuid default gen_random_uuid() primary key,
    word varchar(32)
);

create table if not exists saladRecipes.user (
    id uuid default gen_random_uuid() primary key,
    name varchar(64) not null,
    email text not null check ( email like '%@%.%' ) unique,
    login varchar(64) not null unique,
    password varchar(256) not null,
    role varchar(25) not null default 'user'
);

create table if not exists saladRecipes.saladType (
    id uuid default gen_random_uuid() primary key,
    name varchar(25) not null,
    description varchar(256) default ''
);

create table if not exists saladRecipes.salad (
    id uuid default gen_random_uuid() primary key,
    name varchar(32) not null,
    authorId uuid references saladRecipes.user(id),
    description varchar(64) not null default ''
);

create table if not exists saladRecipes.modStatus (
    id serial primary key,
    name varchar(32),
    description varchar(256)
);

create table if not exists saladRecipes.recipe (
    id uuid default gen_random_uuid() primary key,
    saladId uuid not null references saladRecipes.salad(id) on delete cascade,
    status int not null references saladRecipes.modStatus(id),
    numberOfServings int not null, check ( numberOfServings > 0 ),
    timeToCook int not null, check ( timeToCook > 0 ),
    rating decimal(3, 1) default 0.0, check ( rating >= 0 ), check ( rating <= 5 )
);

create table if not exists saladRecipes.ingredientType (
    id uuid default gen_random_uuid() primary key,
    name varchar(32) not null,
    description varchar(256) default ''
);

create table if not exists saladRecipes.ingredient (
    id uuid default gen_random_uuid() primary key,
    name varchar(32) not null,
    calories int not null, check ( calories >= 0 ),
    type uuid not null references saladRecipes.ingredientType(id)
);

create table if not exists saladRecipes.comment (
    id uuid default gen_random_uuid() primary key,
    author uuid not null references saladRecipes.user(id),
    salad uuid not null references saladRecipes.salad(id) on delete cascade,
    text text default '',
    rating int not null, check ( rating >= 1 ), check ( rating <= 5 ),
    unique (author, salad)
);

create table if not exists saladRecipes.measurement (
    id uuid default gen_random_uuid() primary key,
    name varchar(32) not null,
    grams int, check ( grams > 0 )
);

create table if not exists saladRecipes.recipeStep (
    id uuid default gen_random_uuid() primary key,
    name varchar(32) not null,
    description text not null,
    recipeId uuid not null references saladRecipes.recipe(id) on delete cascade,
    stepNum int not null, check ( stepNum > 0 )
);

-- Links between tables
create table if not exists saladRecipes.recipeIngredient (
    id uuid default gen_random_uuid() primary key,
    recipeId uuid not null references saladRecipes.recipe(id) on delete cascade,
    ingredientId uuid not null references saladRecipes.ingredient(id),
    measurement uuid not null references saladRecipes.measurement(id) default '01000000-0000-0000-0000-000000000000',
    amount int not null default 1, check ( amount > 0 ),
    unique (recipeId, ingredientId)
);

create table if not exists saladRecipes.typesOfSalads (
    id uuid default gen_random_uuid() primary key,
    saladId uuid not null references saladRecipes.salad(id) on delete cascade,
    typeId uuid not null references saladRecipes.saladType(id)
);

-- TRIGGERS
create or replace function calcRate(saladId uuid)
    returns float
as $$
begin
    return (select case
                when avg(rating) is null then 0.0
                else avg(rating)
                end as avgRate
            from saladrecipes.comment
            where salad = saladId);
end;
$$ language plpgsql;

create or replace function updateRate()
    returns trigger
as $$
begin
    update saladRecipes.recipe
    set rating = calcRate(new.salad)
    where saladId = new.salad;
    return new;
end;
$$ language plpgsql;

create or replace function updateRateOnDelete()
    returns trigger
as $$
begin
    update saladRecipes.recipe
    set rating = calcRate(old.salad)
    where saladId = old.salad;
    return old;
end;
$$ language plpgsql;

create trigger updateRateTrigger after insert on saladrecipes.comment
    for each row execute procedure updateRate();
create trigger updateRateUpdate after update on saladrecipes.comment
    for each row execute procedure updateRate();
create trigger updateRateDelete after delete on saladrecipes.comment
    for each row execute procedure updateRateOnDelete();

-- DEFAULT USERS TO USE DUE DEMONSTRATION
insert into keywords.word(word) values
    ('banned');

insert into saladRecipes.user(id, name, email, login, password, role) values
    ('02000000-0000-0000-0000-000000000000', 'user', 'user@mail.ru', 'user', '$2a$10$AZ7UvXvrrmMl3VUZ1Sq/4uk6TRCB1d3/zzAm0j.A2tN5beX/il3ye', 'user'),
    ('01000000-0000-0000-0000-000000000000', 'admin', 'admin@mail.ru', 'admin', '$2a$10$53A6AVUmQFr3nx3taVYnjOKQpHh2JTeBxEaIX1NYEh2I9nnepfVPC', 'admin');

insert into saladrecipes.ingredienttype(name)
values
    ('фрукт'),
    ('овощ'),
    ('мясо'),
    ('рыба'),
    ('молоко');

insert into saladrecipes.ingredient(id, name, calories, type)
values ('f1fc4bfc-799c-4471-a971-1bb00f7dd30a', 'яблоко', 1, (select id from saladrecipes.ingredienttype where name = 'фрукт'));

insert into saladrecipes.ingredient(name, calories, type)
values
    ('морковь', 2, (select id from saladrecipes.ingredienttype where name = 'овощ')),
    ('говядина', 3, (select id from saladrecipes.ingredienttype where name = 'мясо')),
    ('лосось', 4,  (select id from saladrecipes.ingredienttype where name = 'рыба')),
    ('молоко', 5, (select id from saladrecipes.ingredienttype where name = 'молоко'));

insert into saladrecipes.salad(id, name, authorId, description)
values ('fbabc2aa-cd4a-42b0-b68d-d3cf67fba06f', 'цезарь', '02000000-0000-0000-0000-000000000000', 'описание цезаря');

insert into saladrecipes.salad(name, authorId, description)
values
    ('овощной', '02000000-0000-0000-0000-000000000000', 'описание овощного салата'),
    ('сезонный', '01000000-0000-0000-0000-000000000000', 'описание сезонного салата'),
    ('сельдь под шубой', '01000000-0000-0000-0000-000000000000', 'описание'),
    ('греческий', '01000000-0000-0000-0000-000000000000', 'описание греческого салата салата');

insert into saladrecipes.saladtype(id, name)
values ('7e17866b-2b97-4d2b-b399-42ceeebd5480', 'зима');

insert into saladrecipes.saladtype(name)
values
    ('лето'),
    ('осень'),
    ('весна'),
    ('мясной');

insert into saladrecipes.typesofsalads(saladid, typeid)
values
    ((select id from saladrecipes.salad where name = 'цезарь'),
     (select id from saladrecipes.saladtype where name = 'зима')),

    ((select id from saladrecipes.salad where name = 'овощной'),
     (select id from saladrecipes.saladtype where name = 'лето')),
    ((select id from saladrecipes.salad where name = 'овощной'),
     (select id from saladrecipes.saladtype where name = 'зима')),

    ((select id from saladrecipes.salad where name = 'сезонный'),
     (select id from saladrecipes.saladtype where name = 'лето')),
    ((select id from saladrecipes.salad where name = 'сезонный'),
     (select id from saladrecipes.saladtype where name = 'зима')),
    ((select id from saladrecipes.salad where name = 'сезонный'),
     (select id from saladrecipes.saladtype where name = 'весна')),
    ((select id from saladrecipes.salad where name = 'сезонный'),
     (select id from saladrecipes.saladtype where name = 'осень')),

    ((select id from saladrecipes.salad where name = 'сельдь под шубой'),
     (select id from saladrecipes.saladtype where name = 'зима')),
    ((select id from saladrecipes.salad where name = 'сельдь под шубой'),
     (select id from saladrecipes.saladtype where name = 'мясной')),

    ((select id from saladrecipes.salad where name = 'греческий'),
     (select id from saladrecipes.saladtype where name = 'зима'));

insert into saladrecipes.modstatus(name)
values
    ('редактирование'),
    ('на модерации'),
    ('отклонено'),
    ('опубликовано'),
    ('снято с публикации');

insert into saladrecipes.recipe(id, saladid, status, numberofservings, timetocook, rating)
values
    ('01000000-0000-0000-0000-000000000000', (select id from saladrecipes.salad where name = 'цезарь'),
     (select id from saladrecipes.modstatus where name = 'опубликовано'),
     1, 1, 0.0),
    ('02000000-0000-0000-0000-000000000000', (select id from saladrecipes.salad where name = 'овощной'),
     (select id from saladrecipes.modstatus where name = 'опубликовано'),
     2, 2, 0.0),
    ('03000000-0000-0000-0000-000000000000', (select id from saladrecipes.salad where name = 'сельдь под шубой'),
     (select id from saladrecipes.modstatus where name = 'опубликовано'),
     3, 3, 0.0),
    ('04000000-0000-0000-0000-000000000000', (select id from saladrecipes.salad where name = 'сезонный'),
     (select id from saladrecipes.modstatus where name = 'опубликовано'),
     4, 4, 0.0),
    ('05000000-0000-0000-0000-000000000000', (select id from saladrecipes.salad where name = 'греческий'),
     (select id from saladrecipes.modstatus where name = 'опубликовано'),
     5, 5, 0.0);

insert into saladRecipes.measurement(id, name, grams)
values ('01000000-0000-0000-0000-000000000000', 'граммов', 1);

insert into saladrecipes.measurement(name, grams)
values
    ('чайная ложка', 1),
    ('штук', 1),
    ('килограмм', 1000);

insert into saladrecipes.recipeingredient(recipeid, ingredientid, measurement, amount)
values
    ((select id from saladrecipes.recipe where numberofservings = 1),
     (select id from saladrecipes.ingredient where name = 'яблоко'),
     (select id from saladrecipes.measurement where name = 'граммов'),
     1),

    ((select id from saladrecipes.recipe where numberofservings = 2),
     (select id from saladrecipes.ingredient where name = 'морковь'),
     (select id from saladrecipes.measurement where name = 'граммов'),
     2),
    ((select id from saladrecipes.recipe where numberofservings = 2),
     (select id from saladrecipes.ingredient where name = 'яблоко'),
     (select id from saladrecipes.measurement where name = 'граммов'),
     3),

    ((select id from saladrecipes.recipe where numberofservings = 3),
     (select id from saladrecipes.ingredient where name = 'говядина'),
     (select id from saladrecipes.measurement where name = 'граммов'),
     4),
    ((select id from saladrecipes.recipe where numberofservings = 3),
     (select id from saladrecipes.ingredient where name = 'яблоко'),
     (select id from saladrecipes.measurement where name = 'граммов'),
     5),

    ((select id from saladrecipes.recipe where numberofservings = 4),
     (select id from saladrecipes.ingredient where name = 'говядина'),
     (select id from saladrecipes.measurement where name = 'граммов'),
     6),
    ((select id from saladrecipes.recipe where numberofservings = 4),
     (select id from saladrecipes.ingredient where name = 'яблоко'),
     (select id from saladrecipes.measurement where name = 'граммов'),
     7),
    ((select id from saladrecipes.recipe where numberofservings = 4),
     (select id from saladrecipes.ingredient where name = 'морковь'),
     (select id from saladrecipes.measurement where name = 'граммов'),
     8),
    ((select id from saladrecipes.recipe where numberofservings = 4),
     (select id from saladrecipes.ingredient where name = 'лосось'),
     (select id from saladrecipes.measurement where name = 'граммов'),
     9),

    ((select id from saladrecipes.recipe where numberofservings = 5),
     (select id from saladrecipes.ingredient where name = 'говядина'),
     (select id from saladrecipes.measurement where name = 'граммов'),
     10),
    ((select id from saladrecipes.recipe where numberofservings = 5),
     (select id from saladrecipes.ingredient where name = 'яблоко'),
     (select id from saladrecipes.measurement where name = 'граммов'),
     11),
    ((select id from saladrecipes.recipe where numberofservings = 5),
     (select id from saladrecipes.ingredient where name = 'морковь'),
     (select id from saladrecipes.measurement where name = 'граммов'),
     12),
    ((select id from saladrecipes.recipe where numberofservings = 5),
     (select id from saladrecipes.ingredient where name = 'лосось'),
     (select id from saladrecipes.measurement where name = 'граммов'),
     13),
    ((select id from saladrecipes.recipe where numberofservings = 5),
     (select id from saladrecipes.ingredient where name = 'молоко'),
     (select id from saladrecipes.measurement where name = 'граммов'),
     14);

insert into saladrecipes.recipestep(id, name, description, recipeid, stepnum)
values ('01000000-0000-0000-0000-000000000000', 'step', 'description', '02000000-0000-0000-0000-000000000000', 1),
       ('07000000-0000-0000-0000-000000000000', 'step', 'description', '03000000-0000-0000-0000-000000000000', 1),

       ('02000000-0000-0000-0000-000000000000', 'first', 'first', '01000000-0000-0000-0000-000000000000', 1),
       ('03000000-0000-0000-0000-000000000000', 'second', 'second', '01000000-0000-0000-0000-000000000000', 2),
       ('04000000-0000-0000-0000-000000000000', 'third', 'third', '01000000-0000-0000-0000-000000000000', 3),
       ('05000000-0000-0000-0000-000000000000', 'fourth', 'fourth', '01000000-0000-0000-0000-000000000000', 4),
       ('06000000-0000-0000-0000-000000000000', 'fifth', 'fifth', '01000000-0000-0000-0000-000000000000', 5),

       ('08000000-0000-0000-0000-000000000000', 'first', 'first', '04000000-0000-0000-0000-000000000000', 1),
       ('09000000-0000-0000-0000-000000000000', 'second', 'second', '04000000-0000-0000-0000-000000000000', 2),
       ('0a000000-0000-0000-0000-000000000000', 'third', 'third', '04000000-0000-0000-0000-000000000000', 3);
