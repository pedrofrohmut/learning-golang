create table if not exists products (
       id bigserial,
       name text not null,
       price_in_cents integer not null check (price_in_cents >= 0),
       quantity integer default 0,
       created_at timestamptz default now(),
       primary key (id)
);

create table if not exists orders (
       id bigserial,
       customer_id bigint not null,
       created_at timestamptz default now(),
       primary key (id)
);

create table if not exists order_items (
       id bigserial,
       product_id bigint not null,
       quantity integer default 0 check (quantity >= 0),
       price_in_cents integer not null check (price_in_cents >= 0),

       order_id bigint not null,
       constraint fk_order foreign key (order_id) references orders (id),
       primary key (id)
);
