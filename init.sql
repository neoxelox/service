--[Extra Databases]

--CREATE DATABASE extra;
--GRANT ALL PRIVILEGES ON DATABASE extra to postgres;

--[Extra Users]
-- CREATE USER metrics WITH PASSWORD 'your_password';
-- GRANT CONNECT ON DATABASE database_name TO username;
-- GRANT USAGE ON SCHEMA schema_name TO username;
-- GRANT SELECT ON ALL TABLES IN SCHEMA schema_name TO username;
-- ALTER DEFAULT PRIVILEGES IN SCHEMA public
--    GRANT SELECT ON TABLES TO xxx;