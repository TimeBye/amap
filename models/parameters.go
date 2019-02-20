package models

type Parameters struct {
	Location string `form:"location"`
	Keywords string `form:"keywords"`
	Types    string `form:"types"`
	City     string `form:"city"`
	Radius   int    `form:"radius"`
	SortRule string `form:"sortrule"`
}
