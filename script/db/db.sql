create table body (
   id bigint not null auto_increment,
   type varchar(255) not null,
   name varchar(255) not null,
   description varchar(255) not null,
   moons int not null,
   primary key (id)
);

create table physical_data
(
    id bigint not null auto_increment,
    body_id bigint not null,
    density float not null,
    gravity float not null,
    mass_value float not null,
    mass_exponent int not null,
    volume_value float not null,
    volume_exponent float not null,
    primary key (id)
);

create table orbital_parameters
(
    id bigint not null auto_increment,
    body_id bigint not null,
    sideral_orbit float not null,
    sideral_rotation float not null,
    primary key (id)
)
