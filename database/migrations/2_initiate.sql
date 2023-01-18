-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS public.subscription_plans
(
    id bigint NOT NULL,
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    price integer NOT NULL,
    active_period_in_months smallint NOT NULL,
    features text COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT "PK_subscription_plans" PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.subscription_plans
    OWNER to postgres;


-- +migrate StatementEnd
