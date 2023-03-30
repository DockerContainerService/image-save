package client

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

type repoUrl struct {
	url string

	registry  string
	namespace string
	project   string
	tag       string
}

func parseRepoUrl(url string) (*repoUrl, error) {
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
		logrus.Infof("Using default tag: latest")
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
			registry:  "registry.hub.docker.com",
			namespace: slice[0],
			project:   repo,
			tag:       tag,
		}, nil
	} else {
		return &repoUrl{
			url:       url,
			registry:  "registry.hub.docker.com",
			namespace: "library",
			project:   repo,
			tag:       tag,
		}, nil
	}
}
