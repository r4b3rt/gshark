package search

import (
	"github.com/madneal/gshark/search/codesearch"
	"github.com/madneal/gshark/search/githubsearch"
	"github.com/madneal/gshark/search/gitlabsearch"
	"github.com/urfave/cli"
	"strings"
	"time"
)

func Scan(ctx *cli.Context) {
	var scanMode string
	// seconds
	var Interval time.Duration = 900

	if ctx.IsSet("mode") {
		scanMode = strings.ToLower(ctx.String("mode"))
	}

	if ctx.IsSet("time") {
		Interval = time.Duration(ctx.Int("time"))
	}

	switch scanMode {
	case "github":
		// use go keyword or not
		for {
			githubsearch.RunTask(Interval)
		}
	case "searchcode":
		for {
			codesearch.RunTask(Interval)
		}
	case "gitlab":
		for {
			gitlabsearch.RunTask(Interval)
		}
	case "all":
		for {
			gitlabsearch.RunTask(Interval)
			codesearch.RunTask(Interval)
			githubsearch.RunTask(Interval)
		}
	default:
		for {
			githubsearch.RunTask(Interval)
			//go gobuster.RunTask(Interval)
		}
	}
}
