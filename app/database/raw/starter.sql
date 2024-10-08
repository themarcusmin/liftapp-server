INSERT INTO formats (id, display_name) VALUES
(1, "reps+weight");

INSERT INTO exercises (id, display_name, format_id) VALUES 
(1, 'back squat', 1), (2, 'bench Press', 1);

INSERT INTO muscles (id, display_name) VALUES 
(1, 'quadriceps'), (2, 'glutes'), (3, "abductors"), (4, 'hamstrings'), (5, 'abs'), (6, 'erector spinae'), (7, 'calves'), (8, 'hip flexors'),
(9, 'chest'), (10, 'anterior deltoids'), (11, 'triceps');

INSERT INTO exercise_muscles (exercise_id, muscle_id, is_primary) VALUES 
(1, 1, 1), (1, 2, 1), (1, 3, 0), (1, 4, 0), (1, 5, 0), (1, 6, 0), (1, 7, 0), (1, 8, 0),
(2, 9, 1), (2, 10, 0), (2, 11, 0);



