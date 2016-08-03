package lib

type SiteConfig struct {
	Title     string
	Introduce string
	Limit     int
	Theme     string
	URL       string
}

type AuthorConfig struct {
	Name string
}

type BuildConfig struct {
	Port    string
	Copy    []string
	Publish string
}
