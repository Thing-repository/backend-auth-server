CREATE TABLE companies
(
    id           serial primary key,
    company_name varchar(255) not null,
    address      varchar(255) not null unique,
    image_url    varchar(255)
);

CREATE TABLE departments
(
    id              serial primary key,
    department_name varchar(255)                                    not null,
    company_id      int references companies (id) on delete cascade not null,
    image_url       varchar(255),
    UNIQUE (company_id, department_name)
);

CREATE TABLE users
(
    id                     serial primary key,
    first_name             varchar(255)          not null,
    last_name              varchar(255)          not null,
    email                  varchar(255)          not null unique,
    image_url              varchar(255),
    password_hash          varchar(255)          not null,
    company_id             int                   references companies (id) on delete set null,
    department_id          int                   references departments (id) on delete set null,
    email_is_validated     boolean default false not null,
    email_validation_token varchar(255)
);

CREATE TABLE vacations
(
    id                  serial primary key,
    user_id             int references users (id) on delete cascade not null,
    vacation_time_start timestamp,
    vacation_time_end   timestamp
);


CREATE TYPE credential_types as enum ('company_admin', 'company_user','department_admin', 'department_maintainer', 'department_user');
CREATE TABLE company_credentials
(
    id              serial primary key,
    credential_type credential_types                                not null,
    user_id         int references users (id) on delete cascade     not null,
    object_id       int references companies (id) on delete cascade not null,
    UNIQUE (credential_type, user_id, object_id)
);

CREATE TABLE department_credentials
(
    id              serial primary key,
    credential_type credential_types                                  not null,
    user_id         int references users (id) on delete cascade       not null,
    object_id       int references departments (id) on delete cascade not null,
    UNIQUE (credential_type, user_id, object_id)
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
    user_id    int references users (id) on delete cascade  not null,
    thing_id   int references things (id) on delete cascade not null,
    start_time timestamp                                    not null,
    end_time   timestamp,
    is_approve bool default false                           not null,
    is_taken   bool default false                           not null
);

CREATE TABLE blocking_things
(
    id         serial primary key,
    user_id    int references users (id) on delete cascade  not null,
    thing_id   int references things (id) on delete cascade not null,
    start_time timestamp                                    not null,
    end_time   timestamp,
    reason     varchar(255)
);