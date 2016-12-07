CREATE TABLE IF NOT EXISTS comment_formation_vote_user (
	"userid" text NOT NULL,
	"commentid" text NOT NULL,
	"formationid" text NOT NULL,
	"vote" boolean NOT NULL
);
