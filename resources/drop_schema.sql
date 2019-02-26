-- psql [db name] < resources/drop_schema.sql

DROP TABLE IF EXISTS coindrop_auth;
DROP TABLE IF EXISTS coindrop_reddit;
DROP TABLE IF EXISTS coindrop_stackoverflow;
DROP TABLE IF EXISTS coindrop_tasks;
DROP TABLE IF EXISTS coindrop_quiz_results;
DROP TABLE IF EXISTS coindrop_quizzes;
DROP TABLE IF EXISTS coindrop_user_tasks;
