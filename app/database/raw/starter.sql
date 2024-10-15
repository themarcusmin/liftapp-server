-- FORMATS
INSERT INTO
    formats (id, display_name)
VALUES
    (1, "reps and weight"),
    (2, "reps and bodyweight"),
    (3, "reps, bodyweight and weight"),
    (4, "duration"),
    (5, "reps"),
    (6, "time under tension"),
    (7, "distance and time"),
    (8, "weight and time");

-- MUSCLES
INSERT INTO
    muscles (id, display_name)
VALUES
    (1, 'chest'),
    (2, 'leg'),
    (3, 'back');

-- EXERCISES
INSERT INTO
    exercises (id, display_name, format_id, parent_id)
VALUES
    -- CHEST: parent
    (1, 'Barbell Bench Press', 1, NULL),
    (2, 'Dumbbell Bench Press', 1, NULL),
    (3, 'Machine Bench Press', 1, NULL),
    (4, 'Cable Chest Fly', 1, NULL),
    (5, 'Push-Up', 2, NULL),
    (6, 'Pec Deck', 1, NULL),
    (7, 'Cable Crossovers', 1, NULL),
    (8, 'Bodyweight Chest Dips', 2, NULL),
    (9, 'Weighted Chest Dips', 3, NULL),
    -- CHEST: children
    (10, 'Incline Barbell Bench Press', 1, 1),
    (11, 'Decline Barbell Bench Press', 1, 1),
    (12, 'Incline Dumbbell Bench Press', 1, 2),
    (13, 'Decline Dumbbell Bench Press', 1, 2),
    (14, 'Incline Machine Bench Press', 1, 3),
    (15, 'Decline Machine Bench Press', 1, 3),
    (16, 'High Cable Chest Fly', 1, 4),
    (17, 'Low Cable Chest Fly', 1, 4),
    (18, 'Incline Push-Up', 2, 5),
    (19, 'Decline Push-Up', 2, 5),
    (20, 'Weighted Push-Up', 3, 5),
    (21, 'High Cable Crossover', 1, 7),
    (22, 'Low Cable Crossover', 1, 7),
    -- LEG: parent
    (23, 'Barbell Squat', 1, NULL),
    (24, 'Leg Extensions', 1, NULL),
    (25, 'Leg Press', 1, NULL),
    (26, 'Calf Raises', 1, NULL),
    (27, 'Lunges', 1, NULL),
    (28, 'Hip Thrusts', 1, NULL),
    (29, 'Leg Curls', 1, NULL),
    (30, 'Glute Kickbacks', 1, NULL),
    -- LEG: child
    (31, 'Front Squat', 1, 23),
    (32, 'Goblet Squat', 1, 23),
    (33, 'Barbell Lunges', 1, 27),
    (34, 'Dumbbell Lunges', 1, 27),
    -- BACK: parent
    (35, 'Bodyweight Pull-Ups', 2, NULL),
    (36, 'Weighted Pull-Ups', 1, NULL),
    (37, 'Lat Pull-down', 1, NULL),
    (38, 'Bent Rows', 1, NULL),
    (39, 'Seated Rows', 1, NULL),
    (40, 'One-Arm Dumbbell Row', 1, NULL),
    (41, 'Dumbbell Pull-Over', 1, NULL),
    -- BACK: child
    (42, 'T-Bar Bent Rows', 1, 38),
    (43, 'Dumbbell Bent Rows', 1, 38),
    (44, 'Barbell Bent Rows', 1, 38),
    (45, 'Cable Seated Rows', 1, 39), 
    (46, 'Machine Seated Rows', 1, 39),
    (47, 'One-Arm Cable Rows', 1, 39);

INSERT INTO
    exercise_muscles (exercise_id, muscle_id, is_primary)
VALUES
    -- CHEST
    (1, 1, 1), -- Barbell Bench Press
    (2, 1, 1), -- Dumbbell Bench Press
    (3, 1, 1), -- Machine Bench Press
    (4, 1, 1), -- Cable Chest Fly
    (5, 1, 1), -- Push-up
    (6, 1, 1), -- Pec Deck
    (7, 1, 1), -- Cable Crossovers
    (8, 1, 1), -- Bodyweight Chest Dips
    (9, 1, 1), -- Weighted Chest Dips
    (10, 1, 1), -- Incline Barbell Bench Press
    (11, 1, 1), -- Decline Barbell Bench Press
    (12, 1, 1), -- Incline Dumbbell Bench Press
    (13, 1, 1), -- Decline Dumbbell Bench Press
    (14, 1, 1), -- Incline Machine Bench Press
    (15, 1, 1), -- Decline Machine Bench Press
    (16, 1, 1), -- High Cable Chest Fly
    (17, 1, 1), -- Low Cable Chest Fly
    (18, 1, 1), -- Incline Push-Up
    (19, 1, 1), -- Decline Push-Up
    (20, 1, 1), -- Weighted Push-Up
    (21, 1, 1), -- High Cable Crossover
    (22, 1, 1), -- Low Cable Crossover
    -- LEG
    (23, 2, 1), -- Barbell Squat
    (24, 2, 1), -- Leg Extensions
    (25, 2, 1), -- Leg Press
    (26, 2, 1), -- Calf Raises
    (27, 2, 1), -- Lunges
    (28, 2, 1), -- Hip Thrusts
    (29, 2, 1), -- Leg Curls
    (30, 2, 1), -- Glute Kickbacks
    (31, 2, 1), -- Front Squat
    (32, 2, 1), -- Goblet Squat
    (33, 2, 1), -- Barbell Lunges
    (34, 2, 1), -- Dumbbell Lunges
    -- BACK
    (35, 3, 1), -- Bodyweight Pull-Ups
    (36, 3, 1), -- Weighted Pull-Ups
    (37, 3, 1), -- Lat Pull-down
    (38, 3, 1), -- Bent Rows
    (39, 3, 1), -- Seated Rows
    (40, 3, 1), -- One-Arm Dumbbell Row
    (41, 3, 1), -- Dumbbell Pull-Over
    (42, 3, 1), -- T-Bar Bent Rows
    (43, 3, 1), -- Dumbbell Bent Rows
    (44, 3, 1), -- Barbell Bent Rows
    (45, 3, 1), -- Cable Seated Rows
    (46, 3, 1), -- Machine Seated Rows
    (47, 3, 1); -- One-Arm Cable Rows