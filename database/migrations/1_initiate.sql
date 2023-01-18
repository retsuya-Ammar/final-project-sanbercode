-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS public.users
(
    id bigint NOT NULL,
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    email character varying(255) COLLATE pg_catalog."default" NOT NULL,
    email_verified_at timestamp without time zone,
    password character varying(255) COLLATE pg_catalog."default" NOT NULL,
    remember_token character varying(100) COLLATE pg_catalog."default",
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    CONSTRAINT "PK_users" PRIMARY KEY (id),
    CONSTRAINT users_email_unique UNIQUE (email)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.users
    OWNER to postgres;

-- +migrate StatementEnd