package utils

import (
	// "fmt"

	"net/http"

	"golang.org/x/net/html"
)

/*
	Structs
*/

type LinkPreview struct {
	Title string
	Description string
	ImageUrl string // TODO
}

/*
	Internal
*/



// gets attr if other attribute name=val
// var getAttr func(*html.Node, name string) string
// returns "" if nothing found
func getAttr(n *html.Node, name, val, attr string) string {
	isDesired := false
	for _, a := range n.Attr {
		if a.Key == name && a.Val == val {
			isDesired = true
			break
		}
	}
	if isDesired == true {
		for _, a := range n.Attr {
			if a.Key == attr {
				return a.Val
			}
		}
	}

	return ""
}


func parseBody(n *html.Node, result *LinkPreview) *LinkPreview {
	// initiate if none
	// don't want client to have to
	if result == nil { result = new(LinkPreview) }
	// result := new(LinkPreview)

	// termination
	if result.Title != "" && result.Description != "" {
		return result
	}

	// description
	if n.Type == html.ElementNode && n.Data == "meta" {
		// get description from attribute
		description := getAttr(n, "name", "description", "content")
		if description != "" { 
			// fmt.Printf("description: %v\n", description) 
			result.Description = description
		}
	}

	// title
	if n.Type == html.ElementNode && n.Data == "title" {
		title := n.FirstChild.Data
		// fmt.Printf("title: %v\n", title)
		result.Title = title
	}


	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseBody(c, result)
	}


	return result
}

/*
	Exported
*/


func ParseLink(url string) (*LinkPreview, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	

	// get meta information
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	preview := parseBody(doc, nil)
	// fmt.Printf("%+v\n", preview)

	// println("done")
	



	return preview, nil

}