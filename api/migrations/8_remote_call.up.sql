
CREATE TABLE IF NOT EXISTS gx_state.remote_call (
  id serial primary key,
  actor text not null,
  action text not null,
  args jsonb null,
  result jsonb null,
  processed_state text not null default 'New',
  created timestamptz default now(),
  updated timestamptz default now()
);
