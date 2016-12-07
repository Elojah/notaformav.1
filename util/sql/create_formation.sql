CREATE TABLE IF NOT EXISTS formation (
	"id" text primary key NOT NULL,
	"name" text,
	"sector" text,
	"globalrating" integer,
	"ncomments" integer,
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
