-- +goose Up
-- +goose StatementBegin
-- Insert Ponds
INSERT INTO ponds (name, description, member_count) VALUES
    ('TechTalk', 'All things technology', 15000),
    ('GreenThumb', 'Gardening enthusiasts', 8500),
    ('CodingPond', 'Programming discussions', 12000),
    ('ArtistsCorner', 'Share your creations', 9800),
    ('ScienceLab', 'Scientific discoveries and discussions', 11200),
    ('BookClub', 'For literature lovers', 7300),
    ('FoodiesUnite', 'Cooking and recipes', 13400),
    ('MusicLounge', 'For music lovers and creators', 12300),
    ('GamerHaven', 'Gaming discussions and news', 18700),
    ('CinemaSpot', 'Movie reviews and discussions', 14200),
    ('PetPals', 'All about our furry friends', 10100),
    ('FitnessZone', 'Health and workout tips', 13400);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM user_ponds;
DELETE FROM ponds;
-- +goose StatementEnd