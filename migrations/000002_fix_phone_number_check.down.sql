ALTER TABLE todo.users DROP CONSTRAINT users_phone_number_check;

ALTER TABLE todo.users
    ADD CONSTRAINT users_phone_number_check CHECK (
        phone_number ~ '^+[0-9]+$'
        AND
        char_length(phone_number) BETWEEN 10 AND 15
    );
