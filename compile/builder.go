package compile

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type Compiler struct {
	codeDir string
}

// NewCompiler
//
//	@param codeDir
//	@return *Compiler
func NewCompiler(codeDir string) *Compiler {
	return &Compiler{
		codeDir: codeDir,
	}
}

func (c *Compiler) runCommand(cmd *exec.Cmd) error {
	stdout, err := cmd.StdoutPipe()
	fmt.Println(c.codeDir, cmd.String())
	if err != nil {
		fmt.Println("cmd.StdoutPipe: ", err)
		return err
	}
	cmd.Stderr = os.Stderr
	cmd.Dir = c.codeDir
	err = cmd.Start()
	if err != nil {
		return err
	}
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}
	err = cmd.Wait()
	return err

}
func (c *Compiler) Prepare() error {
	cmd := exec.Command("go", "mod", "tidy")

	return c.runCommand(cmd)
}

func (c *Compiler) Build(target string) error {
	cmd := exec.Command("go", "build")
	cmd.Args = append(cmd.Args, "-o", target)
	return c.runCommand(cmd)

}
