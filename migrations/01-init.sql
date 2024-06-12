create database lct2024;

CREATE TABLE park (
    id int not null auto_increment,

    name varchar(255) not null,
    description varchar(255) not null,
    area json null,

    primary key (id)
);    

CREATE TABLE route ( 
    id int not null auto_increment,
    
    park_id int null,
    name varchar(255) not null,
    description varchar(255) not null,
    how_to_get varchar(255) not null,
    what_to_take varchar(255) not null,
    in_emergency varchar(255) not null,
    recommendations varchar(255) not null,

    length int null,
    duration int null,
    height int null,
    difficulty enum('easy', 'medium', 'hard') not null,
    load json null,

    primary key (id),
    index (park_id)
);

CREATE TABLE route_to_month ( 
    route_id int not null,
    month int not null,
    workload_cur int not null,
    workload_avg int not null,
    capacity int not null,

    primary key (route_id, month)
);

CREATE TABLE review ( 
    id int not null auto_increment,
    route_id int not null,
    user_id int not null,

    message varchar(4096) not null,

    primary key (id),
    index (route_id)
);

CREATE TABLE place ( 
    id int not null auto_increment, 

    route_id int not null,
    name varchar(255) not null,
    description varchar(255) not null,
    icon varchar(255) null,
    location json null,

    primary key (id)
);

CREATE TABLE user (
    id int not null auto_increment,
    status enum('active','blocked') not null default 'active',
    device_id varchar(255) not null default '',

    flags set('gosuslugi_connected') default '',

    primary key (id)
);

CREATE TABLE person ( 
    id int not null auto_increment,
    request_id int not null,

    first_name varchar(255) not null default "",
    middle_name varchar(255) not null default "",
    last_name varchar(255) not null default "",
    sitizen varchar(255) not null default "",
    region varchar(255) not null default "",
    gender enum('man', 'woman') not null,
    passport varchar(255) not null default "",
    birthday date not null,
    email varchar(255) not null default "",
    phone int null,


    primary key (id),
    index (request_id)
);

CREATE TABLE org ( 
    id int not null auto_increment,
    request_id int not null,

    name varchar(255) not null default "",
    inn varchar(255) not null default "",
    kpp varchar(255) not null default "",
    address varchar(255) not null default "",
    email varchar(255) not null default "",
    phone int null,

    primary key (id),
    index (request_id)
);

CREATE TABLE request (
    id int not null auto_increment,
    status enum('pending', 'approved', 'declined') not null default 'pending',

    user_id int not null,
    date_start date not null,

    primary key (id),
    index (user_id)
);    

CREATE TABLE message (
    id int not null auto_increment,
    status enum('pending', 'delivered', 'declined') not null default 'pending',

    user_id int not null,
    sent_at datetime not null,
    type enum('trash', 'fire', 'damage', 'other') not null,
    title varchar(255) not null default '',
    message varchar(4096) not null default '',
    file varchar(255) not null default '',

    primary key (id),
    index (user_id)
);    

