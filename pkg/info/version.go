package info

import (
	"fmt"
	"strings"
)

// app build context
var verStr string
var brsStr string
var tagStr string
var uptStr string
var envStr string

// record system version
type version struct {
	ver string
	brs string
	tag string
	upt string
	env string
}

func (v *version) ToString() string {
	u := strings.Join(
		strings.Split(
			strings.Split(
				v.upt, " ",
			)[0], "/"),
		"",
	)
	p1 := strings.Join([]string{ServicesVersion, v.ver, u}, ".")
	ap := strings.Join([]string{p1, v.tag}, "@")
	return ap
}

func (v *version) Print() {
	fmt.Printf("Version: %s (%s) \n", v.ver, ServicesVersion)
	fmt.Printf("Git Branch: %s \n", v.brs)
	fmt.Printf("Git Commit: %s \n", v.tag)
	fmt.Printf("Build Data: %s \n", v.upt)
	fmt.Printf("Compiler Environment: %s \n", v.env)
}

func getVersion() *version {
	return &version{
		ver: verStr,
		brs: brsStr,
		tag: tagStr,
		upt: uptStr,
		env: envStr,
	}
}

var Version = getVersion()
