package reddit

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

type Item struct {
	Title string
	URL   string
	Comments int `json:"num_comments"`
}

func (i Item) String() string {
	com := ""
	switch i.Comments {
	case 0:
		// nothing
	case 1:
		com = " (1 comment)"
	default:
		com = fmt.Sprintf(" (%d comments)", i.Comments)
	}
	return fmt.Sprintf("%s%s\n%s", i.Title, com, i.URL)
}

const UserAgent = "Golang Reddit Reader v0.13 (by /u/Ptk7l2)"

func Get(reddit string) ([]Item, error) {
	url := fmt.Sprintf("http://reddit.com/r/%s/top.json?limit=100&t=year", reddit)

	// Create a request and add the proper headers.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)

	// Handle the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	r := new(Response)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	items := make([]Item, len(r.Data.Children))
	for i, child := range r.Data.Children {
		items[i] = child.Data
	}

	return items, err
}
