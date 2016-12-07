SELECT pid
FROM (
	SELECT establishment.id as pid,
		to_tsvector('french', establishment.name) ||
		to_tsvector('french', establishment.sector) as document
	FROM establishment) p_search
WHERE p_search.document @@ to_tsquery('Université & science');
