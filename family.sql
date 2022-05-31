CREATE SEQUENCE menu_id_seq START 1;

CREATE TABLE IF NOT EXISTS public.member
(
    id integer NOT NULL DEFAULT nextval('menu_id_seq'::regclass),
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    parent_id integer,
    is_deleted smallint NOT NULL DEFAULT 0,
    path character varying(255) COLLATE pg_catalog."default" NOT NULL,
    date_birth date DEFAULT '1949-08-01'::date,
    date_marry date DEFAULT '1949-08-01'::date,
    place_birth character varying(255) COLLATE pg_catalog."default",
    date_death date DEFAULT '1949-08-01'::date,
    place_death character varying(255) COLLATE pg_catalog."default",
    content character varying(255) COLLATE pg_catalog."default",
    honor character varying(255) COLLATE pg_catalog."default",
    family_simple character varying(255) COLLATE pg_catalog."default" DEFAULT 'new family'::character varying,
    marry_id integer DEFAULT 0,
    sex character varying COLLATE pg_catalog."default" DEFAULT '女'::character varying,
    CONSTRAINT test_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.member
    OWNER to postgres;

COMMENT ON COLUMN public.member.parent_id
    IS '父id';

COMMENT ON COLUMN public.member.path
    IS '从根开始到本身的路径';

COMMENT ON COLUMN public.member.date_birth
    IS '出生日期';

COMMENT ON COLUMN public.member.date_marry
    IS '结婚日期';

COMMENT ON COLUMN public.member.place_birth
    IS '出生地';

COMMENT ON COLUMN public.member.date_death
    IS '死亡时间';

COMMENT ON COLUMN public.member.place_death
    IS '死亡地点';

COMMENT ON COLUMN public.member.content
    IS '生平简介';

COMMENT ON COLUMN public.member.honor
    IS '荣誉';

COMMENT ON COLUMN public.member.family_simple
    IS '家族标示，代表所属';

COMMENT ON COLUMN public.member.marry_id
    IS '族内配偶id';