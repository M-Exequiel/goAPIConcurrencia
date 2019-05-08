package go_api

type User struct {
	Id int64 `json:"id"`
	Nickname string `json:"nickname"`
	CountryId string `json:"country_id"`
	SiteId string `json:"site_id"`
}

func (u *User) SaberIdSitio() string{
	return u.SiteId
}
