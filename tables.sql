-- Drop table

-- DROP TABLE public.accountability

CREATE TABLE IF NOT EXISTS public.accountability (
	id uuid NOT NULL,
	start_tl int8 NOT NULL,
	end_tl int8 NULL,
	description varchar NULL,
	CONSTRAINT accountability_pkey PRIMARY KEY (id, start_tl)
);
CREATE UNIQUE INDEX accountability_tl ON public.accountability USING btree (id, start_tl, end_tl DESC);
CREATE UNIQUE INDEX roleadditionalcontent_tl ON public.accountability USING btree (id, start_tl, end_tl DESC);

-- Drop table

-- DROP TABLE public.circledirectmember

CREATE TABLE IF NOT EXISTS public.circledirectmember (
	start_tl int8 NULL,
	end_tl int8 NULL,
	x uuid NULL,
	y uuid NULL
);
CREATE INDEX circledirectmember_x_start_tl ON public.circledirectmember USING btree (x, start_tl, end_tl DESC);
CREATE INDEX circledirectmember_y_start_tl ON public.circledirectmember USING btree (y, start_tl, end_tl DESC);

-- Drop table

-- DROP TABLE public.commandevent

CREATE TABLE IF NOT EXISTS public.commandevent (
	timeline int8 NULL,
	id uuid NULL,
	issuer uuid NULL,
	commandtype varchar NULL,
	"data" bytea NULL
);

-- Drop table

-- DROP TABLE public."domain"

CREATE TABLE IF NOT EXISTS public."domain" (
	id uuid NOT NULL,
	start_tl int8 NOT NULL,
	end_tl int8 NULL,
	description varchar NULL,
	CONSTRAINT domain_pkey PRIMARY KEY (id, start_tl)
);
CREATE UNIQUE INDEX domain_tl ON public.domain USING btree (id, start_tl, end_tl DESC);

-- Drop table

-- DROP TABLE public."event"

CREATE TABLE IF NOT EXISTS public."event" (
	id uuid NOT NULL,
	sequencenumber bigserial NOT NULL,
	eventtype varchar NOT NULL,
	category varchar NOT NULL,
	streamid varchar NOT NULL,
	"timestamp" timestamptz NOT NULL,
	"version" int8 NOT NULL,
	"data" bytea NULL,
	metadata bytea NULL,
	CONSTRAINT event_category_streamid_version_key UNIQUE (category, streamid, version),
	CONSTRAINT event_pkey PRIMARY KEY (sequencenumber)
);
CREATE INDEX event_category ON public.event USING btree (category);
CREATE INDEX event_streamid ON public.event USING btree (streamid, version);

-- Drop table

-- DROP TABLE public."member"

CREATE TABLE IF NOT EXISTS public."member" (
	id uuid NOT NULL,
	start_tl int8 NOT NULL,
	end_tl int8 NULL,
	isadmin bool NULL,
	username varchar NULL,
	fullname varchar NULL,
	email varchar NULL,
	CONSTRAINT member_pkey PRIMARY KEY (id, start_tl)
);
CREATE UNIQUE INDEX member_tl ON public.member USING btree (id, start_tl, end_tl DESC);

-- Drop table

-- DROP TABLE public.memberavatar

CREATE TABLE IF NOT EXISTS public.memberavatar (
	id uuid NOT NULL,
	start_tl int8 NOT NULL,
	end_tl int8 NULL,
	image bytea NULL,
	CONSTRAINT memberavatar_pkey PRIMARY KEY (id, start_tl)
);
CREATE UNIQUE INDEX memberavatar_tl ON public.memberavatar USING btree (id, start_tl, end_tl DESC);

-- Drop table

-- DROP TABLE public.memberevent

CREATE TABLE IF NOT EXISTS public.memberevent (
	timeline int8 NULL,
	id uuid NULL,
	command uuid NULL,
	cause uuid NULL,
	eventtype varchar NULL,
	memberid uuid NULL,
	"data" bytea NULL
);

-- Drop table

-- DROP TABLE public.membermatch

CREATE TABLE IF NOT EXISTS public.membermatch (
	memberid uuid NULL,
	matchuid varchar NULL
);

-- Drop table

-- DROP TABLE public.membertension

CREATE TABLE IF NOT EXISTS public.membertension (
	start_tl int8 NULL,
	end_tl int8 NULL,
	x uuid NULL,
	y uuid NULL
);
CREATE INDEX membertension_x_start_tl ON public.membertension USING btree (x, start_tl, end_tl DESC);
CREATE INDEX membertension_y_start_tl ON public.membertension USING btree (y, start_tl, end_tl DESC);

-- Drop table

-- DROP TABLE public.migration_eventstore

CREATE TABLE IF NOT EXISTS public.migration_eventstore (
	"version" int4 NOT NULL,
	"time" timestamptz NOT NULL
);

-- Drop table

-- DROP TABLE public.migration_readdb

CREATE TABLE IF NOT EXISTS public.migration_readdb (
	"version" int4 NOT NULL,
	"time" timestamptz NOT NULL
);

-- Drop table

-- DROP TABLE public."password"

CREATE TABLE IF NOT EXISTS public."password" (
	memberid uuid NULL,
	"password" varchar NULL
);

-- Drop table

-- DROP TABLE public."role"

CREATE TABLE IF NOT EXISTS public."role" (
	id uuid NOT NULL,
	start_tl int8 NOT NULL,
	end_tl int8 NULL,
	roletype varchar NOT NULL,
	"depth" int4 NOT NULL,
	"name" varchar NULL,
	purpose varchar NULL,
	CONSTRAINT role_pkey PRIMARY KEY (id, start_tl)
);
CREATE UNIQUE INDEX role_tl ON public.role USING btree (id, start_tl, end_tl DESC);

-- Drop table

-- DROP TABLE public.roleaccountability

CREATE TABLE IF NOT EXISTS public.roleaccountability (
	start_tl int8 NULL,
	end_tl int8 NULL,
	x uuid NULL,
	y uuid NULL
);
CREATE INDEX roleaccountability_x_start_tl ON public.roleaccountability USING btree (x, start_tl, end_tl DESC);
CREATE INDEX roleaccountability_y_start_tl ON public.roleaccountability USING btree (y, start_tl, end_tl DESC);

-- Drop table

-- DROP TABLE public.roleadditionalcontent

CREATE TABLE IF NOT EXISTS public.roleadditionalcontent (
	id uuid NOT NULL,
	start_tl int8 NOT NULL,
	end_tl int8 NULL,
	"content" varchar NULL,
	CONSTRAINT roleadditionalcontent_pkey PRIMARY KEY (id, start_tl)
);

-- Drop table

-- DROP TABLE public.roledomain

CREATE TABLE IF NOT EXISTS public.roledomain (
	start_tl int8 NULL,
	end_tl int8 NULL,
	x uuid NULL,
	y uuid NULL
);
CREATE INDEX roledomain_x_start_tl ON public.roledomain USING btree (x, start_tl, end_tl DESC);
CREATE INDEX roledomain_y_start_tl ON public.roledomain USING btree (y, start_tl, end_tl DESC);

-- Drop table

-- DROP TABLE public.roleevent

CREATE TABLE IF NOT EXISTS public.roleevent (
	timeline int8 NULL,
	id uuid NULL,
	command uuid NULL,
	cause uuid NULL,
	eventtype varchar NULL,
	roleid uuid NULL,
	"data" bytea NULL
);

-- Drop table

-- DROP TABLE public.rolemember

CREATE TABLE IF NOT EXISTS public.rolemember (
	start_tl int8 NULL,
	end_tl int8 NULL,
	x uuid NULL,
	y uuid NULL,
	focus varchar NULL,
	nocoremember bool NULL,
	electionexpiration timestamptz NULL
);
CREATE INDEX rolemember_x_start_tl ON public.rolemember USING btree (x, start_tl, end_tl DESC);
CREATE INDEX rolemember_y_start_tl ON public.rolemember USING btree (y, start_tl, end_tl DESC);

-- Drop table

-- DROP TABLE public.rolerole

CREATE TABLE IF NOT EXISTS public.rolerole (
	start_tl int8 NULL,
	end_tl int8 NULL,
	x uuid NULL,
	y uuid NULL
);
CREATE INDEX rolerole_x_start_tl ON public.rolerole USING btree (x, start_tl, end_tl DESC);
CREATE INDEX rolerole_y_start_tl ON public.rolerole USING btree (y, start_tl, end_tl DESC);

-- Drop table

-- DROP TABLE public.roletension

CREATE TABLE IF NOT EXISTS public.roletension (
	start_tl int8 NULL,
	end_tl int8 NULL,
	x uuid NULL,
	y uuid NULL
);
CREATE INDEX roletension_x_start_tl ON public.roletension USING btree (x, start_tl, end_tl DESC);
CREATE INDEX roletension_y_start_tl ON public.roletension USING btree (y, start_tl, end_tl DESC);

-- Drop table

-- DROP TABLE public.sequencenumber

CREATE TABLE IF NOT EXISTS public.sequencenumber (
	sequencenumber int8 NOT NULL,
	CONSTRAINT sequencenumber_pkey PRIMARY KEY (sequencenumber)
);

-- Drop table

-- DROP TABLE public.streamversion

CREATE TABLE IF NOT EXISTS public.streamversion (
	streamid varchar NOT NULL,
	category varchar NOT NULL,
	"version" int8 NOT NULL,
	CONSTRAINT streamversion_pkey PRIMARY KEY (streamid)
);

-- Drop table

-- DROP TABLE public.tension

CREATE TABLE IF NOT EXISTS public.tension (
	id uuid NOT NULL,
	start_tl int8 NOT NULL,
	end_tl int8 NULL,
	title varchar NULL,
	description varchar NULL,
	closed bool NULL,
	closereason varchar NULL,
	CONSTRAINT tension_pkey PRIMARY KEY (id, start_tl)
);
CREATE UNIQUE INDEX tension_tl ON public.tension USING btree (id, start_tl, end_tl DESC);

-- Drop table

-- DROP TABLE public.timeline

CREATE TABLE IF NOT EXISTS public.timeline (
	"timestamp" timestamptz NOT NULL,
	groupid uuid NOT NULL,
	aggregatetype varchar NOT NULL,
	aggregateid varchar NOT NULL,
	CONSTRAINT timeline_pkey PRIMARY KEY (groupid)
);
CREATE INDEX timeline_aggregateid ON public.timeline USING btree (aggregateid);
CREATE INDEX timeline_aggregatetype ON public.timeline USING btree (aggregatetype);
CREATE INDEX timeline_ts ON public.timeline USING btree ("timestamp");

