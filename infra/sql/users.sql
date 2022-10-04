create table if not exists syncs
(
    id uuid not null primary key
);

create type calendar_type as enum ('notion', 'google');

create table if not exists notion_users
(
    id        uuid not null primary key,
    sync_id   uuid not null references syncs (id),

    name text not null,
    token text not null
);

create table if not exists google_users
(
    id      uuid not null primary key,
    sync_id uuid not null references syncs (id),

    name text not null,
    token text not null
);

