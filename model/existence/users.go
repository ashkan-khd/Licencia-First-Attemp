package existence

type user struct {
	Username    string `json:"username" pg:"username,pk,notnull"`
	Password    string `json:"password" pg:"password,notnull"`
	Description string `json:"description" pg:"description"`
}

type person struct {
	FirstName   string `json:"firstname" pg:"firstname"`
	LastName    string `json:"lastname" pg:"lastname"`
	Email       string `json:"email" pg:"email"`
	PhoneNumber string `json:"phonenumber" pg:"phonenumber"`
	Address     string `json:"addr" pg:"addr"`
}

type image struct {
	Title string `json:"title" pg:"title"`
	Url   string `json:"url" pg:"url"`
}

type Employer struct {
	tableName struct{} `pg:"employers"`
	user
	person
	image
}
