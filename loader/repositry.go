package loader

type Repositry struct {
	Name        string
	Language    string
	Owner       string
	Description string
	URL         string
	StarNum     string
}

func (r *Repositry) Setter(name, language, owner, description, url, starnum string) {
	r.Name = name
	r.Language = language
	r.Owner = owner
	r.Description = description
	r.URL = url
	r.StarNum = starnum
}
