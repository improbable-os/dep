package vsolver

import "fmt"

type ProjectIdentifier struct {
	LocalName   ProjectName
	NetworkName string
}

func (i ProjectIdentifier) less(j ProjectIdentifier) bool {
	if i.LocalName < j.LocalName {
		return true
	}
	if j.LocalName < i.LocalName {
		return false
	}

	return i.NetworkName < j.NetworkName
}

func (i ProjectIdentifier) eq(j ProjectIdentifier) bool {
	if i.LocalName != j.LocalName {
		return false
	}
	if i.NetworkName == j.NetworkName {
		return true
	}

	if (i.NetworkName == "" && j.NetworkName == string(j.LocalName)) ||
		(j.NetworkName == "" && i.NetworkName == string(i.LocalName)) {
		return true
	}

	return false
}

func (i ProjectIdentifier) netName() string {
	if i.NetworkName == "" {
		return string(i.LocalName)
	}
	return i.NetworkName
}

func (i ProjectIdentifier) errString() string {
	if i.NetworkName == "" || i.NetworkName == string(i.LocalName) {
		return string(i.LocalName)
	}
	return fmt.Sprintf("%s (from %s)", i.LocalName, i.NetworkName)
}

func (i ProjectIdentifier) normalize() ProjectIdentifier {
	if i.NetworkName == "" {
		i.NetworkName = string(i.LocalName)
	}

	return i
}

type ProjectName string

type ProjectAtom struct {
	Ident   ProjectIdentifier
	Version Version
}

var emptyProjectAtom ProjectAtom

type atomWithPackages struct {
	atom ProjectAtom
	pl   []string
}

type ProjectDep struct {
	Ident      ProjectIdentifier
	Constraint Constraint
	Packages   []string
}

// completeDep (name hopefully to change) provides the whole picture of a
// dependency - the root (repo and project, since currently we assume the two
// are the same) name, a constraint, and the actual packages needed that are
// under that root.
type completeDep struct {
	// The base ProjectDep
	pd ProjectDep
	// The specific packages required from the ProjectDep
	pl []string
}

type Dependency struct {
	Depender ProjectAtom
	Dep      ProjectDep
}

// ProjectInfo holds manifest and lock for a ProjectName at a Version
type ProjectInfo struct {
	N ProjectName
	V Version
	Manifest
	Lock
}
