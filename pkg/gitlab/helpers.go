package gitlab

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/0x0BSoD/glci-linter/pkg/git"
	"github.com/0x0BSoD/glci-linter/pkg/helpers"
)

func prepareComabinedCiYaml(linterPath, path string) (string, error) {
	os.Chdir(path)
	cmd := exec.Command(fmt.Sprintf("%s", linterPath), "--preview")
	out, err := cmd.CombinedOutput()
	if err != nil {
		if out != nil {
			fmt.Fprint(os.Stderr, "Valid: false\n")
			fmt.Fprint(os.Stderr, "Warnings: []\n")
			fmt.Fprintf(os.Stderr, "Errors: %s\n", string(out))
			os.Exit(1)
		}
		return "", err
	}

	popCounter := 0
	for k, v := range strings.Split(string(out), "\n") {
		if v == "---" {
			popCounter = k
		}
	}

	fname := fmt.Sprintf("/tmp/gitlab-ci.%v.tmp", time.Now().Unix())
	err = os.WriteFile(fname, out, 0644)
	if err != nil {
		return "", err
	}

	f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()

	if popCounter > 0 {
		for lines := 0; lines < popCounter; lines++ {
			_, err = helpers.PopLine(f)
			if err != nil {
				return "", err
			}
		}
	}

	f1, err := os.ReadFile(fname)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var reqParams = APILintRequest{Content: string(f1)}
	reqBody, _ := json.Marshal(reqParams)

	return string(reqBody), nil
}

func prepareCiYaml(path string) (string, error) {
	ciFile, err := os.ReadFile(fmt.Sprintf("%s/.gitlab-ci.yml", path))
	if err != nil {
		return "", err
	}

	var reqParams = APILintRequest{Content: string(ciFile)}
	reqBody, _ := json.Marshal(reqParams)

	return string(reqBody), nil
}

func buildURL(repoDir string) (string, string, error) {
	repo := git.Repo{
		Path: repoDir,
	}

	if err := repo.OpenRepo(); err != nil {
		return "", "", err
	}

	if err := repo.GetRmotes(); err != nil {
		return "", "", err
	}

	path := ""
	for _, r := range repo.Remotes {
		remote := strings.Split(r.String(), "\n")[0]
		path = strings.Fields(remote)[1]
	}

	server := strings.TrimPrefix(strings.Split(path, ":")[0], "git@")
	path = strings.Replace(strings.TrimSuffix(strings.Split(path, ":")[1], ".git"), "/", "%2F", -1)

	return server, path, nil
}
