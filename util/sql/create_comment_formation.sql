CREATE TABLE IF NOT EXISTS comment_formation (
	"id" text primary key NOT NULL,
	"author" text,
	"role" text,
	"content" text,
	"parentid" text,
	"globalrating" integer,
	"upvote" integer,
	"downvote" integer,
	"qualityrating" integer,
	"teachersrating" integer,
	"affordabilityrating" integer,
	"headcountrating" integer,
	"monitoringrating" integer,
	"equipmentrating" integer,
	"externalrating" integer,
	"professionalisationrating" integer,
	"salaryrating" integer,
	"recognitionrating" integer,
	"ambiancerating" integer,
	"extraactivityrating" integer
);
