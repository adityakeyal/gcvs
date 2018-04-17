package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type CvsInfo struct {
	cvsRoot string

	nonRecursiveCommand []string
	recursiveCommand    []string
	nonRecursiveModule  []string
	recursiveModule     []string
}

func (cvsInfo *CvsInfo) GetCommand() [][]string {

	lenNonRecursiveModule := len(cvsInfo.nonRecursiveModule)
	totalLength := lenNonRecursiveModule + len(cvsInfo.recursiveModule)
	cmd := make([][]string, totalLength)

	for i, x := range cvsInfo.nonRecursiveModule {
		cmd[i] = make([]string, (len(cvsInfo.nonRecursiveCommand) + 1))
		copy(cmd[i], cvsInfo.nonRecursiveCommand)
		cmd[i][len(cvsInfo.nonRecursiveCommand)] = x

		//cmd[i] = strings.Join(cvsInfo.nonRecursiveCommand, " ") + " " + x
	}

	for i, x := range cvsInfo.recursiveModule {
		cmd[i] = make([]string, (len(cvsInfo.recursiveCommand) + 1))
		copy(cmd[i], cvsInfo.recursiveCommand)
		cmd[i][len(cvsInfo.recursiveCommand)] = x
		//cmd[lenNonRecursiveModule+i] = strings.Join(cvsInfo.recursiveCommand, " ") + " " + x
	}

	return cmd
}

type configuration struct {
	CvsRoot string
	Modules []string
}

func NewCVSInfo(tagName string) CvsInfo {

	cvsInfo := CvsInfo{}

	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := configuration{}
	err := decoder.Decode(&conf)

	if err != nil {
		panic("Ahhhhhhh.........")

	} else {

		cvsInfo.cvsRoot = conf.CvsRoot
		cvsInfo.recursiveModule = conf.Modules

	}

	//read the details from a stupid configuration file somewhere below

	cvsInfo.nonRecursiveCommand = []string{"-d", cvsInfo.cvsRoot, "checkout", "-P", "-l", "-r", tagName, "--"}
	cvsInfo.recursiveCommand = []string{"-d", cvsInfo.cvsRoot, "checkout", "-P", "-r", tagName, "--"}

	cvsInfo.nonRecursiveModule = []string{"istar-gv/devel/", "istar-gv/devel/gv-core/", "istar-gv/devel/gv-deployments/"}

	return cvsInfo

}
