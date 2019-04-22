create table coindrop_public (
	id uuid DEFAULT uuid_generate_v4() UNIQUE,
	badge_id uuid REFERENCES coindrop_badges(id),
	created_at timestamp with time zone DEFAULT now(),
	updated_at timestamp with time zone
);
