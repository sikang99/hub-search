package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/segmentio/go-log"
	"github.com/tj/docopt"
)

type GitResponse struct {
	TotalCount        int  `json:"total_count"`
	IncompleteResults bool `json:"incomplete_results"`
	Items             []struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Owner    struct {
			Login             string `json:"login"`
			ID                int    `json:"id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
		} `json:"owner"`
		Private         bool    `json:"private"`
		HTMLURL         string  `json:"html_url"`
		Description     string  `json:"description"`
		Fork            bool    `json:"fork"`
		URL             string  `json:"url"`
		CreatedAt       string  `json:"created_at"`
		UpdatedAt       string  `json:"updated_at"`
		PushedAt        string  `json:"pushed_at"`
		Homepage        string  `json:"homepage"`
		Size            int     `json:"size"`
		StargazersCount int     `json:"stargazers_count"`
		WatchersCount   int     `json:"watchers_count"`
		Language        string  `json:"language"`
		ForksCount      int     `json:"forks_count"`
		OpenIssuesCount int     `json:"open_issues_count"`
		MasterBranch    string  `json:"master_branch"`
		DefaultBranch   string  `json:"default_branch"`
		Score           float64 `json:"score"`
	} `json:"items"`
}

var Version = "0.1.4"

const Usage = `
  Hub-Search for searching respositories in https://github.com

  Usage:
    hub-search <query>... [--lang=<type>][--sort=<method>][--order=<style>][--score=<sn>][--list=<ln>][--down][--text]
    hub-search -d | --down
    hub-search -t | --text
    hub-search -h | --help
    hub-search --version

  Options:
    --lang=<type>    implemenation language, default:ALL
    --sort=<method>  sort field, default: best match [stars|forks|updated]
    --order=<style>  style of sort order, default: desc [asc|desc]
    --score=<sn>     show items more than the score <sn>
    --list=<ln>      list top items below the number <ln>
    -d, --down       download packages searched
    -t, --text       display plain text without escape color characters
    -h, --help       output help information
    -v, --version    output version

`

func main() {
	args, err := docopt.Parse(Usage, nil, true, Version, false)
	//fmt.Println(args)
	log.Check(err)

	query := strings.Join(args["<query>"].([]string), " ")
	req := "https://api.github.com/search/repositories?q=" + url.QueryEscape(query)

	if args["--lang"] != nil {
		lang := args["--lang"].(string)
		req = req + "+language:" + lang
	}
	if args["--sort"] != nil {
		sort := args["--sort"].(string)
		req = req + "&sort=" + sort
	}
	if args["--order"] != nil {
		order := args["--order"].(string)
		req = req + "&order=" + order
	}

	var score float64
	if args["--score"] != nil {
		score, _ = strconv.ParseFloat(args["--score"].(string), 64)
	}
	var num int
	if args["--list"] != nil {
		num, _ = strconv.Atoi(args["--list"].(string))
	}

	var text bool
	if args["--text"] != nil {
		text = args["--text"].(bool)
	}
	var down bool
	if args["--down"] != nil {
		down = args["--down"].(bool)
	}

	//fmt.Println(req)

	res, err := http.Get(req)
	if err != nil {
		log.Fatalf("request failed: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("request error: %s", http.StatusText(res.StatusCode))
	}

	var body GitResponse
	log.Check(json.NewDecoder(res.Body).Decode(&body))
	//fmt.Println(body.Items)

	println()
	if len(body.Items) == 0 {
		fmt.Printf("  no package for '%s' in github.com.\n\n", query)
		return
	}

	for i, item := range body.Items {
		if score > 0 && score > item.Score {
			// filter item with lower score
			continue
		}
		if num > 0 && i >= num {
			break
		}

		// show in colored or plain text
		if text {
			fmt.Printf("  %s\n", item.FullName)
			fmt.Printf("  %s\n", description(item.Description))
			fmt.Printf("  %s  %s  %d  %.2f\n", item.UpdatedAt, language(item.Language), item.Size, item.Score)
			fmt.Printf("  %s\n\n", item.HTMLURL)
		} else {
			fmt.Printf("  \033[32;1m%s\033[m\n", item.FullName) // 32:green, 33:yellow, 36:cyan
			fmt.Printf("  %s\n", description(item.Description))
			fmt.Printf("  %s  \033[33;1m%s\033[m  %d  %.2f\n", item.UpdatedAt, language(item.Language), item.Size, item.Score)
			fmt.Printf("  %s\n\n", item.HTMLURL)
		}

		// download required
		if down {
			download(item.FullName)
		}
	}
	println()
}

// handle null string for description
func description(s string) string {
	if s == "" {
		return "no description"
	}
	return s
}

// handle null string for language
func language(s string) string {
	if s == "" {
		return "---"
	}
	return s
}

// strip down github domain string
func strip(s string) string {
	return strings.Replace(s, "github.com/", "", 1)
}

// download the package from github.com
func download(s string) error {
	pkg := "github.com/" + s
	fmt.Fprintf(os.Stderr, "> downloading (\033[32;1m%s\033[m) ...\n", s)
	r, err := exec.Command("go", "get", "-u", "-v", pkg).CombinedOutput()
	/*
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", r)
			//log.Check(err)
		}
	*/
	fmt.Fprintf(os.Stderr, "%s\n", r)
	return err
}
