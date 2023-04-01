package client

import (
	"fmt"
	"strings"
)

type repoUrl struct {
	url string

	registry  string
	namespace string
	project   string
	tag       string

	username string
	password string
	insecure bool
}

func parseRepoUrl(url, mirror string) (*repoUrl, error) {
	// split to registry/namespace/repoAndTag
	slice := strings.SplitN(url, "/", 3)

	// parse project and tag
	var repo, tag string
	repoAndTag := slice[len(slice)-1]
	s := strings.Split(repoAndTag, ":")
	if len(s) > 2 {
		return nil, fmt.Errorf("invalid repository url: %v", url)
	} else if len(s) == 2 {
		repo = s[0]
		tag = s[1]
	} else {
		fmt.Printf("Using default tag: latest\n")
		repo = s[0]
		tag = "latest"
	}

	// parse registry
	if len(slice) == 3 {
		return &repoUrl{
			url:       url,
			registry:  slice[0],
			namespace: slice[1],
			project:   repo,
			tag:       tag,
		}, nil
	} else if len(slice) == 2 {
		// first string is a domain
		if strings.Contains(slice[0], ".") {
			return &repoUrl{
				url:       url,
				registry:  slice[0],
				namespace: "",
				project:   repo,
				tag:       tag,
			}, nil
		}

		return &repoUrl{
			url:       url,
			registry:  mirror,
			namespace: slice[0],
			project:   repo,
			tag:       tag,
		}, nil
	} else {
		return &repoUrl{
			url:       url,
			registry:  mirror,
			namespace: "library",
			project:   repo,
			tag:       tag,
		}, nil
	}
}
