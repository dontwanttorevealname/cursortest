-- +goose Up
-- +goose StatementBegin

-- Debug: Show all users
SELECT id, username FROM users;

-- Debug: Show all ponds
SELECT id, name FROM ponds;

-- ByteMaster's ponds
WITH user_data AS (
    SELECT id FROM users WHERE username = 'ByteMaster'
)
INSERT INTO user_ponds (user_id, pond_id)
SELECT user_data.id, p.id
FROM user_data, ponds p
WHERE p.name IN ('TechTalk', 'CodingPond', 'ScienceLab');

-- GreenGuru's ponds
WITH user_data AS (
    SELECT id FROM users WHERE username = 'GreenGuru'
)
INSERT INTO user_ponds (user_id, pond_id)
SELECT user_data.id, p.id
FROM user_data, ponds p
WHERE p.name IN ('GreenThumb', 'ScienceLab');

-- PixelPainter's ponds
WITH user_data AS (
    SELECT id FROM users WHERE username = 'PixelPainter'
)
INSERT INTO user_ponds (user_id, pond_id)
SELECT user_data.id, p.id
FROM user_data, ponds p
WHERE p.name IN ('ArtistsCorner', 'TechTalk');

-- BookWorm42's ponds
WITH user_data AS (
    SELECT id FROM users WHERE username = 'BookWorm42'
)
INSERT INTO user_ponds (user_id, pond_id)
SELECT user_data.id, p.id
FROM user_data, ponds p
WHERE p.name IN ('BookClub', 'ArtistsCorner');

-- ChefCroak's ponds
WITH user_data AS (
    SELECT id FROM users WHERE username = 'ChefCroak'
)
INSERT INTO user_ponds (user_id, pond_id)
SELECT user_data.id, p.id
FROM user_data, ponds p
WHERE p.name IN ('FoodiesUnite', 'GreenThumb');

-- LabRat's ponds
WITH user_data AS (
    SELECT id FROM users WHERE username = 'LabRat'
)
INSERT INTO user_ponds (user_id, pond_id)
SELECT user_data.id, p.id
FROM user_data, ponds p
WHERE p.name IN ('ScienceLab', 'TechTalk', 'BookClub');

-- Admin's ponds
WITH user_data AS (
    SELECT id FROM users WHERE username = 'admin@ribbit.com'
)
INSERT INTO user_ponds (user_id, pond_id)
SELECT user_data.id, p.id
FROM user_data, ponds p;


-- Debug: Show final results
SELECT 
    u.username,
    p.name as pond_name
FROM user_ponds up
JOIN users u ON u.id = up.user_id
JOIN ponds p ON p.id = up.pond_id
ORDER BY u.username, p.name;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM user_ponds;
UPDATE ponds SET member_count = 0;
-- +goose StatementEnd 