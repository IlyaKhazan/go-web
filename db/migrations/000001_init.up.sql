CREATE TABLE "users" (
                       id  UUID PRIMARY KEY,
                       "firstname" varchar(255) not null,
                       "lastname" varchar(255) not null,
                       "email" varchar(255) not null unique,
                       "age" bigint not null,
                       "created_at" timestamptz not null default 'now()'
);

