ALTER TABLE locations
ADD COLUMN summary_review_food TEXT;

ALTER TABLE locations
RENAME COLUMN summary_review TO summary_review_place;