package model

type Bucket struct {
	ID      int      `json:"id" form:"id"`
	UserID  int      `json:"userid" form:"userid"`
	Adverts []Advert `json:"adverts" form:"adverts"`
	Count   int      `json:"count" form:"count"`
}
