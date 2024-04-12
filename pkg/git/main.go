package git

import (
	"github.com/go-git/go-git/v5"
)

type GitRepo struct {
	Path    string
	Remotes []*git.Remote
	repo    *git.Repository
}

func (r *GitRepo) OpenRepo() error {
	var err error
	r.repo, err = git.PlainOpen(r.Path)
	if err != nil {
		return err
	}

	return nil
}

func (r *GitRepo) GetRmotes() error {
	var err error
	r.Remotes, err = r.repo.Remotes()
	if err != nil {
		return err
	}

	return nil
}
