
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	name character varying(25) NOT NULL,
	profession character varying(25) NOT NULL,
	dream character varying(25) NOT NULL,
	hobby character varying(25) NOT NULL,
	website_url TEXT NOT NULL,
	phone character varying(15) DEFAULT NULL,
	avatar_url TEXT NOT NULL,
	email character varying(40) NOT NULL,
	password TEXT NOT NULL,
	token TEXT NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NOT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;
