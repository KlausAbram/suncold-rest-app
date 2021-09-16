CREATE TABLE agents
(
    id            serial       not null unique,
    name          varchar(255) not null,
    agent_name    varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE requests
(
    id          serial       not null unique,
    author_name varchar(255) not null unique,
    date        date         not null default now(),
    mod         int          not null default 0
);

CREATE TABLE states
(
    id          serial       not null unique,
    location    varchar(255) not null,
    temperature int          not null,
    pressure    int          not null,
    rain        varchar(255) not null,
    clouds      varchar(255),
    wind        int
);

CREATE TABLE links
(
    id         serial                                         not null unique,
    request_id int references requests (id) on delete cascade not null,
    state_id   int references states (id) on delete cascade   not null,
    agent_id   int references agents (id) on delete cascade   not null
);