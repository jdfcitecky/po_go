package utils

import "math"

type Page struct {
	//which page is using now
	Page int `json:"page"`
	//number of record in a page
	Size int `json:"size"`
	//number of total record
	Total int
}

func (page *Page) GetPage() int {
	//max number of pages
	max := int(math.Ceil(float64(page.Total) / float64(page.Size)))
	//get max if current page number exceed max
	if page.Page > max {
		page.Page = max
	}
	return page.Page
}

//read the record start
func (page *Page) GetStart() int {
	return (page.Page - 1) * page.Size
}
