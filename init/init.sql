CREATE TABLE IF NOT EXISTS telemetry (
  id SERIAL PRIMARY KEY,
  latency INT,
  cpu INT,
  players INT,
  timestamp TIMESTAMP
);
