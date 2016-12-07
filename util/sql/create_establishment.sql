CREATE TABLE IF NOT EXISTS establishment (
	"id" text primary key NOT NULL,
	"name" text,
	"sigle" text,
	"type" text,
	"sector" text,
	"url" text,
	"globalrating" integer,
	"ncomments" integer,
	"hygienerating" integer,
	"sizerating" integer,
	"adminrating" integer,
	"accessibilityrating" integer,
	"environmentalrating" integer,
	"stuffrating" integer
);
