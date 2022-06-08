CREATE TABLE companies
(
    id           serial primary key,
    company_name varchar(255) not null,
    address      varchar(255) not null,
    image_url    varchar(255)
);

CREATE TABLE departments
(
    id              serial primary key,
    department_name varchar(255) not null unique,
    image_url       varchar(255)
);

CREATE TABLE users
(
    id                     serial primary key,
    first_name             varchar(255)          not null,
    last_name              varchar(255)          not null,
    email                  varchar(255)          not null unique,
    image_url              varchar(255),
    password_hash          varchar(255)          not null,
    company_id             int references companies (id) on delete cascade,
    department_id          int references departments (id) on delete cascade,
    vacation_time_start    timestamp,
    vacation_time_end      timestamp,
    email_is_validated     boolean default false not null,
    email_validation_token varchar(255)
);

CREATE TABLE companies_admins
(
    id         serial primary key,
    user_id    int references users (id) on delete cascade     not null,
    company_id int references companies (id) on delete cascade not null
);

CREATE TABLE departments_admins
(
    id            serial primary key,
    user_id       int references users (id) on delete cascade       not null,
    department_id int references departments (id) on delete cascade not null
);

CREATE TABLE department_maintainers
(
    id            serial primary key,
    user_id       int references users (id) on delete cascade       not null,
    department_id int references departments (id) on delete cascade not null
);

CREATE TYPE thing_types as enum ('equipment', 'consumables');
CREATE TYPE remainder_types as enum ('pcs', 'capacity', 'percents');

CREATE TABLE things
(
    id                  serial primary key,
    thing_name          varchar(255)                                      not null,
    company_id          int references companies (id) on delete cascade   not null,
    department_id       int references departments (id) on delete cascade not null,
    image_url           varchar(255),
    thing_type          thing_types                                       not null,
    thing_remainder     real,
    remainder_type      remainder_types,
    is_blocked          boolean                                           not null,
    blocked_time_start  timestamp,
    blocked_time_end    timestamp,
    need_admin_approval boolean                                           not null
);

CREATE TABLE using_things
(
    id         serial primary key,
    user_id    int references users (id)  not null,
    thing_id   int references things (id) not null,
    start_time timestamp                  not null,
    end_time   timestamp,
    is_approve bool default false         not null,
    is_taken  bool default false         not null
);

CREATE TABLE blocking_things
(
    id         serial primary key,
    user_id    int references users (id)  not null,
    thing_id   int references things (id) not null,
    start_time timestamp                  not null,
    end_time   timestamp,
    reason     varchar(255)
);