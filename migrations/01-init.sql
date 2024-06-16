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
    description varchar(2048) null,
    how_to_get varchar(1024) null,
    what_to_take varchar(1024) null,
    in_emergency varchar(1024) null,
    recommendations varchar(1024) null,
    photo json null,
    gpx_data longtext null,

    length int null,
    duration int null,
    height int null,
    difficulty enum('easy', 'medium', 'hard') not null,
    load json null,
    max_load json null,
    group_distance float not null,
    average_time float not null,
    group_size int not null,
    days_on_route int not null,

    primary key (id),
    index (park_id)
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
    area float not null, -- площадь туристского объекта в кв. метрах
    area_per_visitor float not null, -- площадь, необходимая для одного посетителя в кв. метрах
    return_coefficient float not null, -- коэффициент возвращения
    days int not null, -- количество дней в рассматриваемую единицу времени
    used_in_calculations tinyint(1) not null default 1,

    primary key (id),
    foreign key (route_id) references route (id)
);

CREATE TABLE device (
    user_id int not null,
    device_id varchar(255) not null default '',

    primary key (device_id)
);

CREATE TABLE person ( 
    id int not null auto_increment,
    request_id varchar(255) not null,

    first_name varchar(255) not null default "",
    middle_name varchar(255) not null default "",
    last_name varchar(255) not null default "",
    sitizen varchar(255) not null default "",
    region varchar(255) not null default "",
    gender enum('man', 'woman') not null default 'man',
    passport varchar(255) not null default "",
    birthday date not null,
    email varchar(255) not null default "",
    phone varchar(255) not null default "",
    is_leader tinyint(1) not null default false,


    primary key (id),
    index (request_id)
);

CREATE TABLE org ( 
    id int not null auto_increment,
    request_id varchar(255) not null,

    name varchar(255) not null default "",
    inn varchar(255) not null default "",
    kpp varchar(255) not null default "",
    address varchar(255) not null default "",
    email varchar(255) not null default "",
    phone int null,

    primary key (id),
    unique (request_id)
);

CREATE TABLE request (
    request_id varchar(255) not null default "",
    user_id int not null,
    date_start date not null,
    quantity int not null,
    route_id int not null,

    primary key (request_id),
    index (user_id)
);    

CREATE TABLE message (
    id int not null auto_increment,
    status enum('pending', 'delivered', 'declined') not null default 'pending',

    user_id int not null,
    sent_at datetime not null,
    type enum('waste','waterPollution','other','air','animals','reserve','treatment','waterAccess') not null,
    message varchar(4096) not null default '',
    file_url varchar(255) not null default '',
    location json null,
    phone varchar(255),

    primary key (id),
    index (user_id)
);    

-- @JsonValue('waste') 
-- wasteDisposal('Незаконная утилизация отходов'), 

-- @JsonValue('waterPollution') 
-- waterPollution('Загрязнение воды'), 

-- @JsonValue('other') 
-- other('Иное'), 

-- @JsonValue('air') 
-- air('Загрязнение воздуха'), 

-- @JsonValue('animals') 
-- animals('Нарушение законодательства о животном мире'), 

-- @JsonValue('reserve') 
-- reserve('Нарушение в заповедной зоне'), 

-- @JsonValue('treatment') 
-- treatment('Нарушение при эксплуатации очистных сооружений'), 

-- @JsonValue('waterAccess') 
-- waterAccess('Незаконное ограничение доступа к воде');