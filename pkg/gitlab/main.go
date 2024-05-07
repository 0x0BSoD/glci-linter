package gitlab

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/viper"
)

const (
	httpRequestTimeout uint = 15
)

type GitLabClient struct {
	Token      string
	ProjectID  int
	RepoPath   string
	ServerUrl  string
	RepoDir    string
	LinterPath string
	ShowMerged bool
}

func NewGitLabClient(repoDir string) *GitLabClient {
	res := GitLabClient{}

	server, path, _ := buildUrl(repoDir)
	res.RepoPath = path
	res.ServerUrl = server
	res.RepoDir = repoDir
	res.ShowMerged = viper.GetBool("show_merged")
	res.LinterPath = viper.GetString("linter_path")
	res.Token = viper.GetString("access_token")

	return &res
}

func (glc *GitLabClient) getProjectID() error {
	url := fmt.Sprintf("https://%s/api/v4/projects/%s", glc.ServerUrl, glc.RepoPath)
	httpClient, req, err := httpClientRequest(glc.Token, "GET", url, "")

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("[getProjectID] HTTP request error: %w", err)
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("[getProjectID] HTTP request error, code: %s", resp.Status)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("[getProjectID] Unable to parse response: %w", err)
	}

	var result GitLabProject
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return fmt.Errorf("[getProjectID] Unable to parse JSON response: %w", err)
	}

	glc.ProjectID = result.ID

	return nil
}

func (glc *GitLabClient) Lint() error {
	if err := glc.getProjectID(); err != nil {
		return fmt.Errorf("[Lint] Unable to get ProjectID: %w", err)
	}

	content, err := prepareComabinedCiYaml(glc.LinterPath, glc.RepoDir)
	if err != nil {
		return fmt.Errorf("[Lint] Can't build YAML: %w", err)
	}

	url := fmt.Sprintf("https://gitlab.com/api/v4/projects/%v/ci/lint?include_merged_yaml=false", glc.ProjectID)

	httpClient, req, err := httpClientRequest(glc.Token, "POST", url, content)

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("[Lint] HTTP request error: %w", err)
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("[Lint] HTTP request error, code: %s", resp.Status)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("[Lint] Unable to parse response: %w", err)
	}

	var result GitlabAPILintResponse
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return fmt.Errorf("[Lint] Unable to parse JSON response: %w", err)
	}

	// TODO: Return result
	if glc.ShowMerged && result.Valid {
		fmt.Println(result.MergedYaml)
	}

	fmt.Println("Valid: ", result.Valid)
	fmt.Println("Warnings: ", result.Warnings)
	fmt.Println("Errors: ", result.Errors)
	if !result.Valid {
		os.Exit(1)
	}

	return nil
}
