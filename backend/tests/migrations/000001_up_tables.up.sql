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
