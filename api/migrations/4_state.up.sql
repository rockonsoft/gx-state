CREATE TABLE IF NOT EXISTS gx_state.machine_states (
  id serial primary key,
  machine_definition_id BIGINT REFERENCES gx_state.machine_definitions(id),
  name text not null,
  context jsonb null,
  final_state boolean not null default false
)
