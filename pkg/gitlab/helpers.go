package gitlab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/0x0BSoD/glci-linter/pkg/git"
)

func popLine(f *os.File) ([]byte, error) {
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(make([]byte, 0, fi.Size()))

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(buf, f)
	if err != nil {
		return nil, err
	}

	line, err := buf.ReadBytes('\n')
	if err != nil && err != io.EOF {
		return nil, err
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}
	nw, err := io.Copy(f, buf)
	if err != nil {
		return nil, err
	}
	err = f.Truncate(nw)
	if err != nil {
		return nil, err
	}
	err = f.Sync()
	if err != nil {
		return nil, err
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}
	return line, nil
}

func checkJSciLinter() error {
	_, err := exec.LookPath("gitlab-ci-local")
	if err != nil {
		return err
	}

	return nil
}

func prepareComabinedCiYaml(path string) (string, error) {
	os.Chdir(path)
	cmd := exec.Command("gitlab-ci-local", "--preview")
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

	_, err = popLine(f)
	if err != nil {
		return "", err
	}

	f1, err := os.ReadFile(fname)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var reqParams = GitlabAPILintRequest{Content: string(f1)}
	reqBody, _ := json.Marshal(reqParams)

	return string(reqBody), nil
}

func prepareCiYaml(path string) (string, error) {
	ciFile, err := os.ReadFile(fmt.Sprintf("%s/.gitlab-ci.yml", path))
	if err != nil {
		return "", err
	}

	var reqParams = GitlabAPILintRequest{Content: string(ciFile)}
	reqBody, _ := json.Marshal(reqParams)

	return string(reqBody), nil
}

func buildUrl(repoDir string) (string, string, error) {
	repo := git.GitRepo{
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
