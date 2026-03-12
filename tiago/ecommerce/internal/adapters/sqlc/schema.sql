create table if not exists products (
       id bigserial primary key,
       name text not null,
       price_in_cents integer not null check (price_in_cents >= 0),
       quantity integer default 0,
       created_at timestamptz default now()
);
