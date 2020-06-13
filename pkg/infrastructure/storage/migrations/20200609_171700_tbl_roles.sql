CREATE TABLE IF NOT EXISTS role (
      id SERIAL NOT NULL,
      name VARCHAR NOT NULL,
      id_company VARCHAR NOT NULL,
      created_at TIMESTAMP DEFAULT NOW() NOT NULL,
      updated_at TIMESTAMP,
      CONSTRAINT role_unique UNIQUE (name,id_company)
);


CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = now(); 
   RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_module_change_timestamp BEFORE UPDATE
    ON role FOR EACH ROW EXECUTE PROCEDURE 
    update_timestamp();
