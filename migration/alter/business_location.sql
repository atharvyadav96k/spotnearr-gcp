ALTER TABLE business_locations
ADD COLUMN geohash TEXT;

CREATE INDEX idx_business_locations_geohash
ON business_locations(geohash);