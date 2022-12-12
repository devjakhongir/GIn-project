

CREATE TABLE users (
    id UUID DEFAULT gen_random_uuid(),
    first_name varchar(24) NOT NULL,
    last_name varchar(36) NOT NULL
);

INSERT INTO users (first_name, last_name) VALUES
('Samandar', 'Foziljonov');

INSERT INTO users (first_name, last_name) VALUES
('Ravshanbek', 'Olimov');


-- UPDATE 
--     users 
-- SET 
--     first_name = 'Ravshanbek', 
--     last_name = 'Olimov' 
-- WHERE 
--     id = '32d98624-3aa2-4550-954f-c91798794dd7';

-- CREATE OR REPLACE FUNCTION updateById(user_id UUID,column VARCHAR) 
-- RETURNS VARCHAR LANGUAGE PLPGSQL
-- AS
-- $$
    
--     BEGIN

--        UPDATE 
-- 			users 
-- 		SET 
-- 			first_name = $2, 
-- 			last_name = $3
-- 		WHERE 
-- 			id = $1;

--         return movieName;
--     END;
-- $$;