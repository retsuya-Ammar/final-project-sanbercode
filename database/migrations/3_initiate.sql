-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS public.user_subscriptions
(
    id SERIAL PRIMARY KEY,
    user_id bigint NOT NULL,
    subscription_plan_id bigint NOT NULL,
    price integer NOT NULL,
    expired_date timestamp without time zone,
    payment_status character varying(10) COLLATE pg_catalog."default" NOT NULL,
    "snapToken" character varying(255) COLLATE pg_catalog."default",
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    CONSTRAINT user_subscriptions_subscription_plan_id_foreign FOREIGN KEY (subscription_plan_id)
        REFERENCES public.subscription_plans (id) MATCH SIMPLE
        ON UPDATE RESTRICT
        ON DELETE RESTRICT,
    CONSTRAINT user_subscriptions_user_id_foreign FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE RESTRICT
        ON DELETE RESTRICT
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.user_subscriptions
    OWNER to postgres;


-- +migrate StatementEnd