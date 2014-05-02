package packer

import (
	"bytes"
	"fmt"
)

// The git commit that is being compiled. This will be filled in by the
// compiler for source builds.
var GitCommit string

// The version of packer.
const Version = "0.6.0"

// Any pre-release marker for the version. If this is "" (empty string),
// then it means that it is a final release. Otherwise, this is the
// pre-release marker.
const VersionPrerelease = ""

type versionCommand byte

func (versionCommand) Help() string {
	return `usage: packer version

Outputs the version of Packer that is running. There are no additional
command-line flags for this command.`
}

func (versionCommand) Run(env Environment, args []string) int {
	env.Ui().Machine("version", Version)
	env.Ui().Machine("version-prelease", VersionPrerelease)
	env.Ui().Machine("version-commit", GitCommit)

	env.Ui().Say(VersionString())
	return 0
}

func (versionCommand) Synopsis() string {
	return "print Packer version"
}

// VersionString returns the Packer version in human-readable
// form complete with pre-release and git commit info if it is
// available.
func VersionString() string {
	var versionString bytes.Buffer
	fmt.Fprintf(&versionString, "Packer v%s", Version)
	if VersionPrerelease != "" {
		fmt.Fprintf(&versionString, ".%s", VersionPrerelease)

		if GitCommit != "" {
			fmt.Fprintf(&versionString, " (%s)", GitCommit)
		}
	}

	return versionString.String()
}
