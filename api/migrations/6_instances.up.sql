CREATE TABLE IF NOT EXISTS gx_state.machine_instances (
  id serial primary key,
  type_name text not null,
  definition jsonb not null,
  current_state_name text not null,
  created timestamptz default now(),
  updated timestamptz default now(),
  context jsonb null
);

