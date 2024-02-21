CREATE TABLE IF NOT EXISTS books (
         id uuid primary key,
         name text,
         author_name varchar(100),
         page_number integer,
         created_at timestamp default now(),
         updated_at timestamp default now(),
         deleted_at integer default 0
);
