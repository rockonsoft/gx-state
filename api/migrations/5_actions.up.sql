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
  state_id BIGINT REFERENCES gx_state.machine_states(id)
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
