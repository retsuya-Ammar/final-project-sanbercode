-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS public.movies
(
    id bigint NOT NULL,
    name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    slug character varying(100) COLLATE pg_catalog."default" NOT NULL,
    category character varying(100) COLLATE pg_catalog."default" NOT NULL,
    video_url character varying(255) COLLATE pg_catalog."default" NOT NULL,
    thumbnail character varying(255) COLLATE pg_catalog."default" NOT NULL,
    rating double precision NOT NULL,
    is_featured smallint NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT "PK_movies" PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.movies
    OWNER to postgres;

-- +migrate StatementEnd