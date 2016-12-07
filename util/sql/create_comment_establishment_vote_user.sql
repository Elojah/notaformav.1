CREATE TABLE IF NOT EXISTS comment_establishment_vote_user (
	"userid" text NOT NULL,
	"commentid" text NOT NULL,
	"establishmentid" text NOT NULL,
	"vote" boolean NOT NULL
);
