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
);

CREATE TABLE IF NOT EXISTS gx_state.machine_states (
  id serial primary key,
  machine_definition_id BIGINT REFERENCES gx_state.machine_definitions(id),
  name text not null,
  context jsonb null,
  final_state boolean not null default false
);

CREATE TABLE IF NOT EXISTS gx_state.action_definitions (
  id serial primary key,
  name text null,
  actor text not null,
  action text not null,
  args jsonb null,
  owner_id BIGINT not null
);

CREATE TABLE IF NOT EXISTS gx_state.message_actions (
  id serial primary key,
  message text not null,
  target text null,
  state_id BIGINT REFERENCES gx_state.machine_states(id),
  UNIQUE (message, target, state_id)
);

CREATE TABLE IF NOT EXISTS gx_state.entry_actions (
  id serial primary key,
  state_id BIGINT REFERENCES gx_state.machine_states(id),
  actor text not null,
  action text not null,
  args jsonb null
);

CREATE TABLE IF NOT EXISTS gx_state.exit_actions (
  id serial primary key,
  state_id BIGINT REFERENCES gx_state.machine_states(id),
  actor text not null,
  action text not null,
  args jsonb null
);

CREATE TABLE IF NOT EXISTS gx_state.activities (
  id serial primary key,
  name text not null,
  state_id BIGINT REFERENCES gx_state.machine_states(id),
  actor text not null,
  action text not null,
  args jsonb null
);

CREATE TABLE IF NOT EXISTS gx_state.machine_instances (
  id serial primary key,
  type_name text not null,
  definition jsonb not null,
  current_state_name text not null,
  created timestamptz default now(),
  updated timestamptz default now(),
  context jsonb null
);

CREATE TABLE IF NOT EXISTS gx_state.messages (
  id serial primary key,
  recipient BIGINT REFERENCES gx_state.machine_instances(id),
  sender text not null,
  message text not null,
  processed_state text not null default 'New',
  created timestamptz default now(),
  updated timestamptz default now(),
  context jsonb null
);

CREATE TABLE IF NOT EXISTS gx_state.remote_call (
  id serial primary key,
  machine_id BIGINT REFERENCES gx_state.machine_instances(id),
  actor text not null,
  action text not null,
  args jsonb null,
  result jsonb null,
  processed_state text not null default 'New',
  created timestamptz default now(),
  updated timestamptz default now()
);


