package main

import (
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	ontent := []byte("digraph graphname{a -> b;b -> c;a -> c;}")
	tmpfile, _ := ioutil.TempFile("", "example")
	tmpfile.Write(ontent)
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", tmpfile.Name()).Output()
	mode := int(0777)
	ioutil.WriteFile("outimg.png", cmd, os.FileMode(mode))

}
