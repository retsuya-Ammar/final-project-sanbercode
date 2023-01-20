-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS public.subscription_plans
(
    id SERIAL PRIMARY KEY,
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    price integer NOT NULL,
    active_period_in_months smallint NOT NULL,
    features text COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp,
    updated_at timestamp
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.subscription_plans
    OWNER to postgres;


-- +migrate StatementEnd
