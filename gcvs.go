//  CVS Command Wrapper - Helops you perform a checkout using CVS NT installed on the system
// Takes a configuration as input
//

package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/ivpusic/grpool"
)

var tagName string

func init() {
	flag.StringVar(&tagName, "tag", "", "Tag to checkout")
	flag.StringVar(&tagName, "t", "", "Tag to checkout")
}

func main() {

	//test fail

	flag.Parse()

	if tagName == "" {
		flag.Usage()
		log.Fatal()
	}

	//check if we need to create a new directory or not

	dir := fetchDirectory()

	cvsinfo := NewCVSInfo(tagName)

	commands := cvsinfo.GetCommand()
	// number of workers, and size of job queue
	pool := grpool.NewPool(4, len(commands))
	defer pool.Release()

	// how many jobs we should wait
	pool.WaitCount(len(commands))

	executeJob := func(count int) func() {
		return func() {

			module := commands[count][len(commands[count])-1]
			fmt.Printf("Started : %v \r\n", module)

			defer pool.JobDone()

			time.Sleep(time.Duration(1*count%5) * time.Second)

			cmd := exec.Command("cvs", commands[count]...)
			var out bytes.Buffer
			var errOut bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &errOut
			cmd.Dir = dir

			err := cmd.Run()
			if err != nil {
				fmt.Printf(errOut.String())
				fmt.Printf(err.Error())
			}

			//fmt.Printf(out.String())
			fmt.Printf("Completed : %v \r\n", module)

		}
	}

	// checkout actual files
	for i := 0; i < len(commands); i++ {

		count := i

		pool.JobQueue <- executeJob(count)
	}

	pool.WaitAll()

}

func fetchDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if !strings.Contains(dir, tagName) {
		fmt.Printf("Creating Directory %v \r\n", tagName)
		os.Mkdir(tagName, 0755)
		return path.Join(dir, tagName)
	}

	return dir
}
