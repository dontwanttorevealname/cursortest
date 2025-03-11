-- +goose Up
-- +goose StatementBegin

-- ByteMaster's posts (joined March 2023)
INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('The Future of AI Development',
    'Discussing the latest breakthroughs in artificial intelligence and what they mean for developers.',
    432, 156, 'ByteMaster', 'TechTalk',
    '2023-03-20 14:23:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Anyone else feel like frameworks are getting out of hand?',
    'Just venting here but I swear every time I start a new project there''s ANOTHER framework I need to learn...',
    892, 287, 'ByteMaster', 'TechTalk',
    '2023-05-15 09:45:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('My first open source contribution!',
    'After months of lurking, I finally submitted a PR to a major project. Here''s what I learned...',
    567, 156, 'ByteMaster', 'CodingPond',
    '2023-08-28 16:30:00');

-- GreenGuru's posts (joined July 2023)
INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Spring Gardening Tips',
    'Essential tips for preparing your garden for the spring season. Share your experiences!',
    245, 89, 'GreenGuru', 'GreenThumb',
    '2023-07-25 10:15:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Secret weapon for pest control: LADYBUGS!',
    'Y''all. I released 1500 ladybugs in my garden last week and the aphids are GONE. Best $20 I ever spent.',
    556, 167, 'GreenGuru', 'GreenThumb',
    '2023-09-12 15:45:00');

-- PixelPainter's posts (joined November 2023)
INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Digital Art Techniques Workshop',
    'Join our weekly digital art workshop! This week we''re focusing on character design.',
    389, 167, 'PixelPainter', 'ArtistsCorner',
    '2023-12-05 11:30:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Can we talk about art theft on Twitter?',
    'Found my work being sold as NFTs without permission. Here''s what I did to get them taken down...',
    1234, 445, 'PixelPainter', 'ArtistsCorner',
    '2024-01-20 17:20:00');

-- BookWorm42's posts (joined February 2024)
INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Monthly Book Discussion: Sci-Fi Classics',
    'Join us as we explore the foundational works of science fiction literature.',
    276, 142, 'BookWorm42', 'BookClub',
    '2024-02-16 13:45:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Just finished Project Hail Mary and I''m EMOTIONAL',
    'No spoilers but that ending?? I need to talk about this with someone who understands...',
    678, 234, 'BookWorm42', 'BookClub',
    '2024-02-28 22:15:00');

-- ChefCroak's posts (joined August 2024)
INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Seasonal Cooking with Home-Grown Herbs',
    'Make the most of your garden''s bounty with these delicious recipes!',
    312, 98, 'ChefCroak', 'FoodiesUnite',
    '2024-08-05 16:30:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Hot take: Air fryers are overrated',
    'Fight me in the comments but it''s just a tiny convection oven and I''m tired of pretending it''s not',
    890, 567, 'ChefCroak', 'FoodiesUnite',
    '2024-09-15 19:45:00');

-- LabRat's posts (joined January 2025)
INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Latest Discoveries in Quantum Physics',
    'Breaking down the newest research in quantum mechanics and its implications.',
    423, 187, 'LabRat', 'ScienceLab',
    '2025-01-22 12:30:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Just published my first paper!!!',
    'After 2 years of research, my work on neural network optimization is finally out! Link in comments',
    789, 345, 'LabRat', 'ScienceLab',
    '2025-02-15 09:15:00');

-- Ribbit Admin's posts (joined February 2023)
INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Welcome to Ribbit!',
    'We''re excited to have you join our community...',
    1234, 789, 'Ribbit Admin', 'Official',
    '2023-03-01 00:00:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('New Features Coming Soon',
    'We''re working on some exciting updates...',
    987, 456, 'Ribbit Admin', 'Official',
    '2023-06-15 10:00:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Community Update: New Features',
    'Check out our latest platform improvements and upcoming changes...',
    890, 567, 'Ribbit Admin', 'Official',
    '2023-09-30 14:30:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Community Guidelines Reminder',
    'A friendly reminder about our community standards and how to keep Ribbit welcoming for everyone...',
    567, 234, 'Ribbit Admin', 'Official',
    '2024-01-01 09:00:00');

-- Other users' posts (starting from June 2023)
INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Building a Home Server Setup',
    'My journey setting up a home media and development server. Here''s what I learned.',
    278, 92, 'ServerMaster', 'TechTalk',
    '2023-06-20 15:45:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Desert Plants Survival Guide',
    'Tips for maintaining a thriving garden in arid climates.',
    189, 67, 'DesertBloom', 'GreenThumb',
    '2023-07-05 11:30:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Traditional vs Digital Art',
    'Breaking down the pros and cons of both mediums from years of experience.',
    445, 156, 'ArtisanSage', 'ArtistsCorner',
    '2023-08-12 14:20:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('International Street Food Series',
    'Exploring street food cultures around the world. First stop: Thailand!',
    367, 134, 'StreetFoodie', 'FoodiesUnite',
    '2023-10-25 17:15:00');

-- Additional Art Posts
INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Drew this at 3am instead of sleeping',
    'Just had to get this character design out of my head. Still not 100% happy with the lighting but...',
    876, 234, 'NightOwlArtist', 'ArtistsCorner',
    '2023-07-15 03:15:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Trying oils after 10 years of digital',
    'Ok wow this is HARD. Mad respect to traditional artists. Here''s my first attempt (be gentle)...',
    445, 122, 'DigitalDabbler', 'ArtistsCorner',
    '2023-08-22 16:45:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Mastering Digital Illustration',
    'A comprehensive guide to creating professional digital illustrations using industry-standard tools.',
    567, 234, 'DigitalDaVinci', 'ArtistsCorner',
    '2023-09-30 11:20:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('From Sketch to Masterpiece',
    'My journey through a 30-day art challenge. See the progression and lessons learned.',
    445, 189, 'SketchMaster', 'ArtistsCorner',
    '2023-10-15 14:30:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Color Theory Deep Dive',
    'Understanding color relationships and how to create harmonious palettes in your artwork.',
    389, 167, 'ColorWhisperer', 'ArtistsCorner',
    '2023-11-05 09:45:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('3D Modeling for Beginners',
    'Getting started with Blender - tips and tricks for creating your first 3D masterpiece.',
    298, 145, '3DArtist', 'ArtistsCorner',
    '2023-12-12 13:20:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Animation Fundamentals',
    'The 12 principles of animation and how to apply them in your digital work.',
    412, 178, 'AnimationPro', 'ArtistsCorner',
    '2024-01-08 15:45:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('AI Art: Friend or Foe?',
    'A balanced discussion on how AI is impacting the art community.',
    1876, 789, 'ArtPhilosopher', 'ArtistsCorner',
    '2024-02-20 10:30:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Urban Sketching Guide',
    'Tips and techniques for capturing city life in your sketchbook.',
    987, 345, 'UrbanArtist', 'ArtistsCorner',
    '2024-03-15 16:20:00');

-- Additional Book Club Posts
INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Classic Literature Analysis',
    'Deep diving into the themes of Jane Austen''s works and their modern relevance.',
    342, 156, 'LitScholar', 'BookClub',
    '2023-07-25 11:45:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Unpopular opinion: physical books > e-readers',
    'Yes they take up space. Yes they''re harder to travel with. But there''s just something about the smell...',
    789, 456, 'PageTurner', 'BookClub',
    '2023-08-30 19:15:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Building my reading nook! (progress pics)',
    'Finally converted that weird corner of my apartment into the cozy reading space of my dreams',
    445, 123, 'CozyReader', 'BookClub',
    '2023-09-18 14:30:00');

-- Film Posts
INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('The Evolution of CGI in Cinema',
    'From Jurassic Park to today: How CGI has transformed moviemaking.',
    567, 278, 'FilmBuff', 'CinemaSpot',
    '2023-10-10 13:30:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Hidden Gems: Underrated Films of 2023',
    'Great movies you might have missed this year.',
    445, 189, 'CinematicArt', 'CinemaSpot',
    '2023-11-15 14:45:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('The lost art of practical effects',
    'Modern CGI is amazing, but there''s something special about practical effects...',
    789, 345, 'FilmNoir', 'CinemaSpot',
    '2023-12-20 16:15:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Best Director''s Cuts?',
    'Sometimes longer is better. Here are some films that were improved by their director''s cut...',
    567, 234, 'CinematicArt', 'CinemaSpot',
    '2024-01-25 11:30:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Sound design appreciation thread',
    'Let''s talk about films with outstanding sound design. Starting with Dune...',
    456, 189, 'SoundGeek', 'CinemaSpot',
    '2024-02-28 15:45:00');

-- Pet Posts
INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Understanding Cat Body Language',
    'A comprehensive guide to what your cat is trying to tell you.',
    678, 234, 'CatWhisperer', 'PetPals',
    '2023-07-12 09:30:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('First-Time Dog Owner Guide',
    'Everything you need to know before getting your first dog.',
    789, 345, 'DogTrainer', 'PetPals',
    '2023-08-18 14:15:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('My rescue story',
    'After months of patience, my anxious rescue dog finally trusts me. Here''s what worked...',
    1432, 567, 'RescueHero', 'PetPals',
    '2023-09-25 16:45:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Raw diet myths debunked',
    'Veterinarian here! Let''s clear up some misconceptions about raw feeding...',
    876, 345, 'VetDoc', 'PetPals',
    '2023-11-02 10:30:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Apartment-friendly pets',
    'Not just cats and dogs! Here are some great pets for small living spaces...',
    567, 234, 'UrbanPets', 'PetPals',
    '2024-01-15 13:20:00');

-- Fitness Posts
INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Myth-Busting: Common Fitness Misconceptions',
    'Separating fact from fiction in the fitness world.',
    567, 289, 'FitCoach', 'FitnessZone',
    '2023-08-05 08:30:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Building a Home Gym on a Budget',
    'Smart ways to create your workout space without breaking the bank.',
    445, 178, 'HomeGymPro', 'FitnessZone',
    '2023-09-12 11:45:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('From couch to 5K - Success!',
    'Just finished my first 5K after being sedentary for years. Here''s my journey...',
    987, 345, 'RunnerNewbie', 'FitnessZone',
    '2023-10-20 15:30:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('The protein myth',
    'Sports nutritionist here! Let''s talk about how much protein you really need...',
    1023, 456, 'NutritionPro', 'FitnessZone',
    '2023-12-05 09:15:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Mobility work changed my life',
    'After years of lifting, adding mobility work made the biggest difference...',
    678, 234, 'FlexMaster', 'FitnessZone',
    '2024-02-10 14:45:00');

-- Science Posts
INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Space Exploration Updates',
    'Latest developments in space technology and upcoming missions.',
    567, 234, 'SpaceExplorer', 'ScienceLab',
    '2023-07-15 09:30:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('ELI5: The new quantum computing breakthrough',
    'Trying to understand the IBM announcement. Can someone break down what this means for the field?',
    567, 234, 'QuantumCurious', 'ScienceLab',
    '2023-08-22 14:15:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Cool physics demos for high school students?',
    'Need to make physics exciting for my 10th graders. What demonstrations blew your mind as a student?',
    345, 123, 'PhysicsTeacher', 'ScienceLab',
    '2023-09-30 11:45:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Fascinating lab accident today',
    'When two samples got mixed up, we discovered something unexpected...',
    456, 178, 'AccidentalScientist', 'ScienceLab',
    '2023-10-15 16:20:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Science communication tips',
    'How to explain complex concepts to non-scientists without dumbing it down',
    678, 234, 'SciComm', 'ScienceLab',
    '2023-11-28 13:45:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('DIY lab equipment hacks',
    'Budget-friendly solutions that actually work. #4 will surprise you',
    567, 145, 'FrugalScientist', 'ScienceLab',
    '2023-12-15 10:30:00');

-- Gaming Posts
INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('The Rise of Indie Games',
    'How independent developers are reshaping the gaming industry...',
    789, 345, 'IndieGamer', 'GamerHaven',
    '2023-08-10 15:30:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Next-Gen Console Comparison',
    'Detailed analysis of the latest gaming consoles: specs, games, and value.',
    1023, 567, 'TechGamer', 'GamerHaven',
    '2023-09-25 11:15:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('The Psychology of Game Design',
    'How game developers use psychology to create engaging experiences.',
    678, 234, 'GamePsych', 'GamerHaven',
    '2023-10-30 14:45:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('Finally beat Malenia!',
    'After 47 attempts, I finally did it! Here''s the build that worked for me...',
    1289, 445, 'EldenLord', 'GamerHaven',
    '2023-11-15 20:30:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('The state of game preservation',
    'With digital-only releases and server shutdowns, we''re losing gaming history...',
    876, 234, 'RetroGamer', 'GamerHaven',
    '2023-12-20 16:45:00');

INSERT INTO ripples (title, content, like_count, comment_count, author_username, pond_name, created_at)
VALUES 
    ('My first speedrun experience',
    'Decided to try speedrunning Hollow Knight. The community is amazing...',
    567, 156, 'SpeedRunner', 'GamerHaven',
    '2024-01-10 19:15:00');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM ripples;
-- +goose StatementEnd 