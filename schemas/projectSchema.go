package schemas

type ProjectCreateSchema struct {
	Title       string   `json:"title" validate:"required,max=20"` //! want alphanum with space
	Tagline     string   `json:"tagline" validate:"required"`
	Description string   `json:"description" validate:"required,max=500"`
	Tags        []string `json:"tags" validate:"dive,alpha"`
	Category    string   `json:"category" validate:"alpha,required"`
	IsPrivate   bool     `json:"isPrivate" validate:"boolean"`
	Links       []string `json:"links" validate:"dive,url"`
}

type ProjectUpdateSchema struct {
	Title        string   `json:"title" validate:"alphanum,max=20"` //! Remove later
	Tagline      string   `json:"tagline" validate:"alphanum,max=40"`
	CoverPic     string   `json:"coverPic" validate:"image"`
	Description  string   `json:"description" validate:"alphanum,max=500"`
	Page         string   `json:"page"`
	Tags         []string `json:"tags" validate:"alpha,dive"`
	IsPrivate    bool     `json:"isPrivate" validate:"boolean"`
	Links        []string `json:"links" validate:"url,dive"`
	PrivateLinks []string `json:"privateLinks" validate:"url,dive"`
}
