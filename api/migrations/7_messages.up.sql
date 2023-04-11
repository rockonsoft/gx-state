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
