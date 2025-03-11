-- +goose Up
-- +goose StatementBegin
INSERT INTO users (username, password_hash, description, join_date) VALUES
    ('admin@ribbit.com', 'RibbitAdmin', 'Official Ribbit administrator account. Posts platform updates and community guidelines.', '2023-02-05 09:00:00'),
    ('ByteMaster', 'Techn0l0gy!', 'Software developer and AI enthusiast. Active in technology discussions.', '2023-03-15 14:30:00'),
    ('GreenGuru', 'Plant123!', 'Professional gardener and plant enthusiast.', '2023-07-22 11:45:00'),
    ('PixelPainter', 'Art1st!', 'Digital artist and graphic designer.', '2023-11-30 16:20:00'),
    ('BookWorm42', 'ReadMore!', 'Literature enthusiast and book reviewer.', '2024-02-14 10:15:00'),
    ('ChefCroak', 'Food1e!', 'Amateur chef and food photography enthusiast.', '2024-08-03 13:40:00'),
    ('LabRat', 'Sc1ence!', 'Science teacher and researcher.', '2025-01-20 15:55:00');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users;
-- +goose StatementEnd

-- Note: In a production environment, passwords should be properly hashed
-- This is just for development/testing purposes 