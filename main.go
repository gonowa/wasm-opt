package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

const (
	containerOutDir = "/out"
	containerInDir  = "/in"
	dockerImage     = "gonowa/wasm-opt"
)

func main() {

	mArgs, srcDir, dstDir := buildArgs(os.Args[1:])

	var args = []string{"run", "-rm"}
	args = append(args, "-v", strings.Join([]string{srcDir, containerInDir}, ":"))

	//most likely
	if srcDir == dstDir {
		args = append(args, "-v", strings.Join([]string{dstDir, containerOutDir}, ":"))
	} else {
		if dstDir != "" {
			args = append(args, "-v", strings.Join([]string{dstDir, containerOutDir}, ":"))
		}
	}

	args = append(args, "-i", dockerImage)
	args = append(args, mArgs...)
	cmd := exec.Command("docker", args...)
	if Debug(args) {
		fmt.Fprintf(os.Stdout, "runing - docker %s\n", strings.Join(cmd.Args, " "))
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}

}

func buildArgs(args []string) (mArgs []string, srcDir, dstDir string) {
	for i := 0; i < len(args); i++ {
		if isFlag(args[i]) {
			if isOutput(args[i]) {
				mArgs = append(mArgs, args[i])
				//next arg is the output file
				//todo guard this
				if len(args) < i+1 {
					log.Fatalln("you must specify output file")
				}
				output := args[i+1]
				//get ouput path dir
				dstDir = filepath.Dir(AbsPath(output))
				mArgs = append(mArgs, path.Join(containerOutDir, filepath.Base(args[i+1])))
				i++
			} else {
				mArgs = append(mArgs, args[i])
			}
		} else { //input file
			srcDir = filepath.Dir(AbsPath(args[i]))
			mArgs = append(mArgs, path.Join(containerInDir, filepath.Base(args[i])))
		}
	}
	return
}

func isFlag(str string) bool {
	if strings.HasPrefix(str, "-") {
		return true
	}
	return false
}

func AbsPath(str string) string {
	abspath, err := filepath.Abs(str)
	if err != nil {
		panic(err)
	}
	return abspath
}

func isOutput(str string) bool {
	if str == "-o" || str == "--output" {
		return true
	}
	return false
}

func Debug(args []string) bool {
	for i := range args {
		if args[i] == "-d" || args[i] == "--debug" {
			return true
		}
	}
	return false
}
