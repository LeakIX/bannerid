package bannerid

import (
	"errors"
	"strings"
)

type SoftwareModule struct {
	Name string `json:"name"`
	Version string `json:"version"`
}
type Software struct {
	Name string `json:"name"`
	Version string `json:"version"`
	Modules []SoftwareModule `json:"modules"`
	OS string `json:"os"`
}

func ParseWebServerBanner(banner string) (Software, error) {
	software := Software{}
	bannerParts := strings.Split(banner, " ")
	for bannerPartPos, bannerPart := range bannerParts {
		if strings.HasPrefix(bannerPart,"(") && strings.HasSuffix(bannerPart, ")") {
			if bannerPartPos == 1 {
				software.OS = strings.Trim(bannerPart,"()")
			}
			continue
		}
		softwarePair := strings.Split(bannerPart, "/")
		if len(softwarePair) == 2 && bannerPartPos == 0 {
			software.Name = softwarePair[0]
			software.Version = softwarePair[1]
		} else if len(softwarePair) == 2 && bannerPartPos > 0 {
			software.Modules = append(software.Modules, SoftwareModule{
				Name:    softwarePair[0],
				Version: softwarePair[1],
			})
		} else if len(softwarePair) == 1 && bannerPartPos == 0 && len(bannerParts) ==1 {
			software.Name = banner
		}
	}
	if len(software.Name) > 0 {
		return software, nil
	}
	return Software{}, errors.New("failed to parse banner")

}