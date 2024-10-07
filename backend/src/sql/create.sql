create table if not exists "user"
(
    id serial primary key not null ,
    login text not null ,
    password text not null ,
    role text not null ,
    first_name text not null ,
    second_name text not null ,
    third_name text not null
);

create table if not exists studio
(
    id serial primary key not null ,
    name text not null
);

create table if not exists Room
(
    id serial primary key not null,
    name text not null ,
    studio_id int references studio(id) on delete cascade not null,
    start_hour int not null ,
    end_hour int not null
);

create table if not exists equipment
(
    id serial primary key not null ,
    name text not null ,
    type int not null ,
    studio_id int references studio(id) on delete cascade not null
);

create table if not exists producer
(
    id serial primary key not null ,
    name text not null ,
    studio_id int references studio(id) on delete cascade not null ,
    start_hour int not null,
    end_hour int not null
);

create table if not exists instrumentalist
(
    id serial primary key not null ,
    name text not null ,
    studio_id int references studio(id) on delete cascade not null ,
    start_hour int not null ,
    end_hour int not null
);

create table if not exists reserve
(
    id serial primary key not null ,
    user_id int not null,
    room_id int references room(id) on delete cascade not null,
    producer_id int not null,
    instrumentalist_id int not null,
    start_time timestamp not null ,
    end_time timestamp not null
);

create table if not exists reserved_equipments
(
    id serial primary key not null ,
    reserve_id int references reserve(id) on delete cascade not null ,
    equipment_id int references equipment(id) on delete cascade not null
);

-- insert into "user"(login, password, role, first_name, second_name, third_name) values ('admin', 'admin', 'admin', 'admin','admin', 'admin')


create or replace procedure update_reserved_equipment(reserve_id integer, equipment_id integer) as '
        insert into reserved_equipments(reserve_id, equipment_id) values (reserve_id, equipment_id);
    ' language sql;

create or replace procedure delete_reserved_equipment(reserveId integer, equipmentId integer) as '
    delete from reserved_equipments where reserved_equipments.reserve_id = reserveId and reserved_equipments.equipment_id = equipmentId;
' language sql;

create or replace function is_intersect(reserveStartTime timestamp, 
					reserveEndTime timestamp,
					choosenStartTime timestamp,
					choosenEndTime timestamp) returns boolean language plpgsql as $$
    declare
        ans boolean;
    begin
    ans = false;
    if ((choosenStartTime >= reserveStartTime and choosenStartTime < reserveEndTime) or
		(choosenEndTime <= reserveEndTime and choosenEndTime > reserveStartTime) or
		(choosenStartTime <= reserveStartTime and choosenEndTime >= reserveEndTime)) then
		ans = true;
		end if;
    return ans;
	end;
    $$;

create or replace function is_reserve(userId integer,
                                      roomId integer,
                                      producerId integer,
                                      instrumentalistId integer,
                                      startTime timestamp,
                                      endTime timestamp) returns integer language plpgsql as $$
    declare
        count integer;
    begin
        select count(*) into count
        from reserve
        where (user_id = userId or
               room_id = roomId or
               producer_id = producerId or
               instrumentalist_id = instrumentalistId)
          and is_intersect(reserve.start_time,
                           reserve.end_time,
                           startTime,
                           endTime);

        return count;
    end;
    $$;

-- create trigger update_reserved_equipment after insert
--     on reserve execute update_reserved_equipment()

-- select is_reserve(1,1,1,1,'2024-05-03 15:00:00.000000', '2024-05-03 17:00:00.000000')
-- select is_intersect('2024-05-14 13:00:00.000000', '2024-05-14 15:00:00.000000', '2024-05-15 13:00:00.000000', '2024-05-15 15:00:00.000000')
-- call update_reserved_equipment(8, 3)

