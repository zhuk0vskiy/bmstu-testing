-- studio fill

insert into studio(name) values ('first_studio');
insert into studio(name) values ('second_studio');
insert into studio(name) values ('third_studio');

-- room fill

insert into room(name, studio_id, start_hour, end_hour) values ('red', 1, 9, 21);
insert into room(name, studio_id, start_hour, end_hour) values ('blue', 1, 9, 21);
insert into room(name, studio_id, start_hour, end_hour) values ('green', 2, 9, 21);

-- producer fill

insert into producer(name, studio_id, start_hour, end_hour) values ('nik', 1, 9, 21);
insert into producer(name, studio_id, start_hour, end_hour) values ('ivan', 1, 9, 21);
insert into producer(name, studio_id, start_hour, end_hour) values ('masha', 2, 9, 21);

-- instrumentalist fill

insert into instrumentalist(name, studio_id, start_hour, end_hour) values ('slash', 1, 9, 21);
insert into instrumentalist(name, studio_id, start_hour, end_hour) values ('page', 1, 9, 21);
insert into instrumentalist(name, studio_id, start_hour, end_hour) values ('jack', 1, 9, 21);

--equipment fill

insert into equipment(name, studio_id, type) values ('les paul', 1, 2);
insert into equipment(name, studio_id, type) values ('telecaster', 1, 2);
insert into equipment(name, studio_id, type) values ('strat', 1, 2);

