
-- Create Admin User
CREATE USER xadmin;
ALTER USER xadmin WITH ENCRYPTED PASSWORD 'xadmin';
CREATE DATABASE people;
GRANT ALL PRIVILEGES ON DATABASE people TO xadmin;

-- Create Webuser
CREATE USER webuser;
ALTER USER webuser WITH ENCRYPTED PASSWORD 'webuser';
GRANT CONNECT ON DATABASE people to webuser;


