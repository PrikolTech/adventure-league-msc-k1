INSERT INTO "user".role (title, description)
VALUES ('admin', 'администратор'),
    ('employee', 'сотрудник'),
    ('student', 'студент'),
    ('enrollee', 'абитуриент');

INSERT INTO "user".data_role (role_id, data_id)
VALUES ((SELECT id from "user".role WHERE title='admin'), (SELECT id from "user".data WHERE email='admin@mail.ru')),
((SELECT id from "user".role WHERE title='employee'), (SELECT id from "user".data WHERE email='employee@mail.ru')),
((SELECT id from "user".role WHERE title='student'), (SELECT id from "user".data WHERE email='student@mail.ru')),
((SELECT id from "user".role WHERE title='enrollee'), (SELECT id from "user".data WHERE email='enrollee@mail.ru'));