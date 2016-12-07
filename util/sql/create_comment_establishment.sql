CREATE TABLE IF NOT EXISTS comment_establishment (
	"id" text primary key NOT NULL,
	"author" text,
	"role" text,
	"content" text,
	"parentid" text,
	"globalrating" integer,
	"upvote" integer,
	"downvote" integer,
	"hygienerating" integer,
	"sizerating" integer,
	"adminrating" integer,
	"accessibilityrating" integer,
	"environmentalrating" integer,
	"stuffrating" integer
);
