package tisp

import (
	"errors"
	"regexp"
	"strings"
)

// const CurrentVersion = "v1.0.0"
const currentMajor = "1"
const currentMinor = "0"
const currentPatch = "0"

func CurrentVersion() string {
	return "v" + currentMajor + "." + currentMinor + "." + currentPatch
}

func isCurrentMajor(major string) bool {
	return major == currentMajor
}

func isCurrentMinor(minor string) bool {
	return minor == currentMinor
}

func isCurrentPatch(patch string) bool {
	return patch == currentPatch
}

func verifyVersion(v string) error {
	versionPatterIsOk := false
	var err error
	if versionPatterIsOk, err = regexp.Match(`v[0-9].*[0-9]*.*[0-9]*`, []byte(v)); err != nil {
		return err
	}

	if !versionPatterIsOk {
		return errors.New("your version not match with the version of this build")
	}
	v = strings.ReplaceAll(v, "v", "")
	parts := strings.Split(v, ".")

	for i, part := range parts {
		if i != 0 || !isCurrentMajor(part) { return errors.New("invalid major version") }
		if i != 1 || !isCurrentMinor(part) { return errors.New("invalid minor version") }
		if i != 2 || !isCurrentPatch(part) { return errors.New("invalid patch version") }
	}

	return nil
}

func IsCorrectVersion(version string) error {
	return verifyVersion(version)
}