package main

import (
	"fmt"
	"slices"
)

type BuilderFedora struct{}

var (
	BaseDNFInstall = []string{
		"dnf", "install",
	}
	BaseDNFClean = []string{
		"dnf", "clean", "all",
	}
	FlagsDNF = []string{
		"--assumeyes",
		"--disablerepo=*",
		"--enablerepo=fedora",
		"--enablerepo=updates",
		"--nodocs",
		"--setopt", "install_weak_deps=False",
		"--installroot", InstallRoot,
	}
	InstallRoot = "/mnt"
)

func (m *BuilderFedora) Install(buildMount *Directory, releasever string, packages []string) *Directory {
	extraFlags := []string{
		"--releasever", releasever,
	}

	return dag.Container().
		From(fmt.Sprintf("registry.fedoraproject.org/fedora:%s", releasever)).
		WithMountedDirectory(InstallRoot, buildMount).
		WithExec(slices.Concat(BaseDNFInstall, FlagsDNF, extraFlags, packages)).
		Directory(InstallRoot)
}

func (m *BuilderFedora) CleanAll(buildMount *Directory, releasever string) *Directory {
	extraFlags := []string{
		"--releasever", releasever,
	}

	return dag.Container().
		From(fmt.Sprintf("registry.fedoraproject.org/fedora:%s", releasever)).
		WithMountedDirectory(InstallRoot, buildMount).
		WithExec(slices.Concat(BaseDNFClean, FlagsDNF, extraFlags)).
		Directory(InstallRoot)
}
