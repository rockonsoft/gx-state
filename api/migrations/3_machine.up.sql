CREATE SCHEMA IF NOT EXISTS gx_state;

CREATE TABLE IF NOT EXISTS gx_state.machine_definitions (
  id serial primary key,
  type_name text not null,
  documentation text not null,
  created timestamptz default now(),
  updated timestamptz default now(),
  initial_state text not null, -- how to make this 
  context jsonb null,
  current_state jsonb null,
  UNIQUE (type_name)
)

