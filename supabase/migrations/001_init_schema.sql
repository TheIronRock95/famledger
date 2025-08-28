-- Income table
create table if not exists public.income (
    id uuid primary key default gen_random_uuid(),
    user_id text not null,
    source text not null,
    amount numeric(12,2) not null,
    created_at timestamp with time zone default now()
);

-- Expenses table
create table if not exists public.expenses (
    id uuid primary key default gen_random_uuid(),
    user_id text not null,
    category text not null,
    amount numeric(12,2) not null,
    created_at timestamp with time zone default now()
);

-- Goals table
create table if not exists public.goals (
    id uuid primary key default gen_random_uuid(),
    user_id text not null,
    name text not null,
    target_amount numeric(12,2) not null,
    current_amount numeric(12,2) default 0,
    target_date date,
    created_at timestamp with time zone default now()
);

-- Grant permissions
grant select, insert on public.income to anon;
grant select, insert on public.expenses to anon;
grant select, insert on public.goals to anon;
