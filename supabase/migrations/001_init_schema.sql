-- Income table
create table if not exists income (
    id uuid primary key default gen_random_uuid(),
    user_id uuid not null,
    source text not null,
    amount numeric(12,2) not null,
    created_at timestamp default now()
);

-- Expenses table
create table if not exists expenses (
    id uuid primary key default gen_random_uuid(),
    user_id uuid not null,
    category text not null,
    amount numeric(12,2) not null,
    created_at timestamp default now()
);

-- Goals table
create table if not exists goals (
    id uuid primary key default gen_random_uuid(),
    user_id uuid not null,
    name text not null,
    target_amount numeric(12,2) not null,
    current_amount numeric(12,2) default 0,
    target_date date,
    created_at timestamp default now()
);

