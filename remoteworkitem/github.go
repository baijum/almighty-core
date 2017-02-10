package remoteworkitem

import (
	"encoding/json"
	"log"
	"time"

	"github.com/almighty/almighty-core/configuration"
	"github.com/baijum/templog"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// githubFetcher provides issue listing
type githubFetcher interface {
	listIssues(query string, opts *github.SearchOptions) (*github.IssuesSearchResult, *github.Response, error)
	rateLimit()
}

// GithubTracker represents the Github tracker provider
type GithubTracker struct {
	URL   string
	Query string
}

// GithubIssueFetcher fetch issues from github
type githubIssueFetcher struct {
	client *github.Client
}

// ListIssues list all issues
func (f *githubIssueFetcher) listIssues(query string, opts *github.SearchOptions) (*github.IssuesSearchResult, *github.Response, error) {
	return f.client.Search.Issues(query, opts)
}

func (f *githubIssueFetcher) rateLimit() {
	if f.client.Rate().Remaining < 10 {
		sleep := f.client.Rate().Reset.Unix() - time.Now().Unix()
		time.Sleep(time.Duration(sleep))
	}
}

// Fetch tracker items from Github
func (g *GithubTracker) Fetch() chan TrackerItemContent {
	f := githubIssueFetcher{}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: configuration.GetGithubAuthToken()},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	f.client = github.NewClient(tc)
	return g.fetch(&f)
}

func (g *GithubTracker) fetch(f githubFetcher) chan TrackerItemContent {
	item := make(chan TrackerItemContent)
	go func() {
		opts := &github.SearchOptions{
			Sort:  "updated",
			Order: "asc",
			ListOptions: github.ListOptions{
				PerPage: 100,
			},
		}
		for {
			f.rateLimit()
			result, response, err := f.listIssues(g.Query, opts)
			if _, ok := err.(*github.RateLimitError); ok {
				log.Println("reached rate limit", err)
				templog.Println("reached rate limit", err.Error())
				break
			}
			if err != nil {
				templog.Println("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
			}
			issues := result.Issues
			for _, l := range issues {
				id, err := json.Marshal(l.URL)
				if err != nil {
					templog.Println("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee 1")
				}
				lu, err := json.Marshal(l.UpdatedAt)
				if err != nil {
					templog.Println("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee 2")
				}
				lut, err := time.Parse("\"2006-01-02T15:04:05Z\"", string(lu))
				if err != nil {
					templog.Println("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee 3")
				}
				content, err := json.Marshal(l)
				if err != nil {
					templog.Println("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee 4")
				}
				templog.Println2(string(id), lut.Format("2006-01-02T15:04"))
				item <- TrackerItemContent{ID: string(id), Content: content, LastUpdated: &lut}
			}
			if response.NextPage == 0 {
				templog.Println("fffffffffffffffffffffffffffffffffffffffffffffffffffffffff", g.Query)
				f.rateLimit()
				break
			}
			opts.ListOptions.Page = response.NextPage
		}
		close(item)
	}()
	return item
}
