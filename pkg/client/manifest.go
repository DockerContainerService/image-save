package client

import (
	"github.com/containers/image/v5/manifest"
	"strings"
)

func platformValidate(osFilterList, archFilterList []string, platform *manifest.Schema2PlatformSpec) bool {
	osMatched := true
	archMatched := true

	if len(osFilterList) != 0 && platform.OS != "" {
		osMatched = false
		for _, o := range osFilterList {
			if colonMatch(o, platform.OS, platform.OSVersion) {
				osMatched = true
			}
		}
	}

	if len(archFilterList) != 0 && platform.Architecture != "" {
		archMatched = false
		for _, a := range archFilterList {
			if colonMatch(a, platform.Architecture, platform.Variant) {
				archMatched = true
			}
		}
	}

	return osMatched && archMatched
}

func colonMatch(pat string, first string, second string) bool {
	if strings.Index(pat, first) != 0 {
		return false
	}

	return len(first) == len(pat) || (pat[len(first)] == ':' && pat[len(first)+1:] == second)
}
