package latestversion

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/njkevlani/aur-autoupdater/internal/version"
	"github.com/sirupsen/logrus"
)

// Currently, only fetches from GitHub
func GetLatestVersion(owner, repo string) (version.Version, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)

	resp, err := http.Get(url)
	if err != nil {
		logrus.WithError(err).WithField("url", url).Error("failed to get latest release info")
		return nil, err
	}

	defer resp.Body.Close()

	version := LatestGitHubVersion{}
	err = json.NewDecoder(resp.Body).Decode(&version)

	if err != nil {
		logrus.WithError(err).WithField("url", url).Error("failed to decode response")
		return nil, err
	}

	logrus.WithField("LatestGitHubVersion", fmt.Sprintf("%#v", version)).Info("LatestGitHubVersion")

	return &version, nil
}
