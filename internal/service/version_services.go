package service

import (
	"fmt"
	"strconv"
	"strings"
	"vex/internal/core"
)

func ParseVersion(raw string) (*core.Version, error) {
	version := strings.Split(raw, ".")
	if len(version) != 3 {
		return nil, fmt.Errorf("formato de versão inválido")
	}

	// Incrementa a versão major
	major, _ := strconv.Atoi(version[0])
	minor, _ := strconv.Atoi(version[1])
	patch, _ := strconv.Atoi(version[2])

	return &core.Version{Major: major, Minor: minor, Patch: patch}, nil
}

func BumpVersion(v *core.Version, level string) *core.Version {
	switch level {
	case "major":
		v.Major++
		v.Minor = 0
		v.Patch = 0
	case "minor":
		v.Minor++
		v.Patch = 0
	case "patch":
		v.Patch++
	}

	return v
}
