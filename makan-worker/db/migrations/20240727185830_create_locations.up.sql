CREATE TABLE IF NOT EXISTS locations (
  id VARCHAR(48) PRIMARY KEY,
  google_place_id VARCHAR(128) NOT NULL,
  place_name VARCHAR(128) NOT NULL,
  google_maps_uri VARCHAR(128) NOT NULL,
  address TEXT NOT NULL,
  latitude FLOAT, 
  longitude FLOAT, 
  dish_type VARCHAR[],
  types VARCHAR[],
  rating FLOAT NOT NULL,
  user_rating_count INT NOT NULL,
  reviews VARCHAR[],
  tags VARCHAR[],
  summary_review TEXT NOT NULL, 
  phone_number VARCHAR(24),
  score FLOAT,
  created_at TIMESTAMP(0) DEFAULT NOW(),
  updated_at TIMESTAMP(0) DEFAULT NOW()
);