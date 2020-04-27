package www

//SiteRepository is an interface that contains functions for working with sites
type SiteRepository interface {
	GetQuestion() (interface{}, error)
}
