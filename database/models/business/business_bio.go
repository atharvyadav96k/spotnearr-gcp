package business_models

type BusinessBio struct {
	ID              string  `json:"_id" firestore: "_id"`
	BusinessID      string  `json:"business_id" firestore: "business_id"`
	ProfileImageUrl string  `json:"profileImage,omitempty"`
	CoverImageUrl   string  `json:"coverImage,omitempty"`
	Bio             string  `json:"bio,omitempty"`
	Rating          float64 `json:"rating,omitempty"`
	ReviewCount     int64   `json:"reviewCount,omitempty"`
	WeeklyVisits    int64   `json:"weeklyVisits,omitempty"`
	Address         string  `json:"address,omitempty"`
	FollowerCount   *int64  `json:"followerCount,omitempty"`
}
