
-- Database to hold yugioh cards as well as user information

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false)
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warnings;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;




INSERT INTO yugioh_cards (name, level, attack, defense) VALUES ('Dark Magician', 7, 2500, 2100)
INSERT INTO yugioh_cards (name, level, attack, defense) VALUES ('Blue Eyes White Dragon', 8, 3000, 2500)
INSERT INTO yugioh_cards (name, level, attack, defense) VALUES ('Jinzo', 6, 2400, 1500)
INSERT INTO yugioh_cards (name, level, attack, defense) VALUES ('Man Eater Bug', 2, 450, 600)
INSERT INTO yugioh_cards (name, level, attack, defense) VALUES ('Time Wizard', 2, 500, 400)


CREATE TABLE public.users (
    id integer NOT NULL,
    first_name character varying(255),
    last_name character varying(255),
    email character varying(255),
    password character varying(255)
    created_at timestamp wihtout time zone,
    updated_at timestamp without time zone,
);

COPY public.yugioh_cards (id, name, level, attack, defense) FROM stdin;
1   Dark Magician   7   2500    2100   
2   Blue Eyes White Dragon  8   3000    2500
3   Jinzo   6   2400    1500
4   Man Eater Bug   2   450 600
5   Time Wizard     2   500     400
\.




SELECT pg_catalog.setval('public.users', 1, true);


CREATE TABLE
  public.yugioh_cards (
    id bigserial NOT NULL,
    name character varying(255) NOT NULL,
    level integer NULL,
    attack integer NULL,
    defense integer NULL
  );

INSERT INTO yugioh_cards (name, level, attack, defense) VALUES ('Dark Magician', 7, 2500, 2100)
INSERT INTO yugioh_cards (name, level, attack, defense) VALUES ('Blue Eyes White Dragon', 8, 3000, 2500)
INSERT INTO yugioh_cards (name, level, attack, defense) VALUES ('Jinzo', 6, 2400, 1500)
INSERT INTO yugioh_cards (name, level, attack, defense) VALUES ('Man Eater Bug', 2, 450, 600)
INSERT INTO yugioh_cards (name, level, attack, defense) VALUES ('Time Wizard', 2, 500, 400)



COPY public.yugioh_cards (id, name, level, attack, defense) FROM stdin;
1   Dark Magician   7   2500    2100
2   Blue Eyes White Dragon  8   3000    2500
3   Jinzo   6   2400    1500
4   Man Eater Bug   2   450     600
5   Time Wizard     2       500     400


SELECT pg_catalog.setval('public.yugioh_cards', 5, true);

ALTER TABLE
  public.yugioh_cards
ADD
  CONSTRAINT yugioh_cards_pkey PRIMARY KEY (id)


