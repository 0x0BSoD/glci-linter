package git

import (
	"fmt"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
)

const HookPath = ".git/hooks/pre-commit"

type Repo struct {
	Path    string
	Remotes []*git.Remote
	repo    *git.Repository
}

func (r *Repo) OpenRepo() error {
	var err error
	r.repo, err = git.PlainOpen(r.Path)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetRmotes() error {
	var err error
	r.Remotes, err = r.repo.Remotes()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) CheckHook() (bool, string) {

	fullHookPath := path.Join(r.Path, HookPath)

	f, err := os.Lstat(fullHookPath)
	if os.IsNotExist(err) {
		return false, fullHookPath
	}

	if f.Mode()&os.ModeSymlink != os.ModeSymlink {
		fmt.Println("[hook] file already exist and it not SymLink")
		return true, fullHookPath
	}

	fmt.Println("[hook] file already exist")
	return true, fullHookPath
}

func (r *Repo) AddHook() {
	exist, fullHookPath := r.CheckHook()
	if exist {
		os.Exit(1)
	}

	glciBin, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = os.Symlink(glciBin, fullHookPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("[hook] created")
}

func (r *Repo) DeleteHook() {
	exist, fullHookPath := r.CheckHook()
	if !exist {
		os.Exit(1)
	}

	err := os.Remove(fullHookPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("[hook] deleted")
}
