-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS public.movies
(
    id SERIAL PRIMARY KEY, 
    name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    slug character varying(100) COLLATE pg_catalog."default" NOT NULL,
    category character varying(100) COLLATE pg_catalog."default" NOT NULL,
    video_url character varying(255) COLLATE pg_catalog."default" NOT NULL,
    thumbnail character varying(255) COLLATE pg_catalog."default" NOT NULL,
    rating double precision NOT NULL,
    is_featured boolean NOT NULL,
    created_at timestamp,
    updated_at timestamp
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.movies
    OWNER to postgres;

-- +migrate StatementEnd