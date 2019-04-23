// Hello.go
package main

import (
	"os"
	"fmt"
	"os/exec"
	"path/filepath"
)

//中文
func main() {
	fmt.Println("Hello World! Welcome to PHP Start!");
	
	for ;; {
		
		//Execute_php();
		//Execute_php2();
		Execute_php3();
	}
	

}//

func Execute_php() {
	defer PrintError("Execute_php()"); //其实可以不用
	
	//arg := []string{}
	////arg := []string{ " -b 127.0.0.1:9000 " };
	//arg := []string{ "-b", "127.0.0.1:9000" };
	//arg := []string{ "" };
	//cmd := exec.Command("test", arg...);
	////cmd := exec.Command("D:\\Program Files (x86)\\NuSphere\\PhpED\\php54\\php-cgi.exe", arg...);
	//cmd := exec.Command("D:\\Program Files (x86)\\NuSphere\\PhpED\\php54\\php-cgi.exe", " -b 127.0.0.1:9000 " );
	//cmd := exec.Command("D:\\Program Files (x86)\\NuSphere\\PhpED\\php54\\php-cgi.exe" -b 127.0.0.1:9000" );
	//会向 cmd.Stdout和cmd.Stderr写入信息,其实cmd.Stdout==cmd.Stderr,具体可见源码
	
	//ok//cmd = exec.Command("cmd.exe", "-l");
	//ok//cmd = exec.Command("D:\\Program Files (x86)\\NuSphere\\PhpED\\php54\\php-cgi.exe", "-b9000");
	//ok//cmd := exec.Command("D:\\Program Files (x86)\\NuSphere\\PhpED\\php54\\php-cgi.exe", "-b", "127.0.0.1:9000");
	//这样并不会象在 windows cmd.exe 里执行那样产生文件重定向
	//cmd := exec.Command("D:\\Program Files (x86)\\NuSphere\\PhpED\\php54\\php-cgi.exe", "-baaa", "127.0.0.1:9000>d:\\1.txt");
	//这样并不会象在 windows cmd.exe 里执行那样产生文件重定向
	cmd := exec.Command("D:\\Program Files (x86)\\NuSphere\\PhpED\\php54\\php-cgi.exe", "-baaa", "127.0.0.1:9000", ">", "d:\\1.txt");
	//exec.Command("D:\\Program Files (x86)\\NuSphere\\PhpED\\php54\\php-cgi.exe", "-b", "127.0.0.1:9000"); //这样的话会不等待，更可怕
	
	output, err := cmd.CombinedOutput();
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("命令行结果:\n%v\n\n%v\n\n%v", string(output), cmd.Stdout, cmd.Stderr); //如果命令行的输出很多的话按道理应该会有问题
}//

func Execute_php2() {
	
	defer PrintError("Execute_php2()"); //其实可以不用
	
//应该打开一个文件test.txt作为命令"echo "hello world""的标准输出和错误输出：	
//package main
 
//import (
//    "fmt"
//    "os"
//    "os/exec"
//)
 
//func main() {
//    f, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY, 0777)
//    if err != nil {
//        fmt.Println(err)
//        return
//    }
//    defer f.Close()
//    cmd := exec.Command(
//        "echo", "hello world")
//    cmd.Stdout = f
//    cmd.Stderr = f //可选
//    cmd.Start()
//}	


    //cmd := exec.Command("echo", "hello world");
	////cmd := exec.Command("D:\\Program Files (x86)\\NuSphere\\PhpED\\php54\\php-cgi.exe", "-baaa", "127.0.0.1:9000>d:\\1.txt");
	cmd := exec.Command("D:\\Program Files (x86)\\NuSphere\\PhpED\\php54\\php-cgi.exe", "-b", "127.0.0.1:9000");
    //这样就可以把 php 命令行的结果重定向了
	cmd.Stdout = os.Stdout;
    cmd.Stderr = os.Stderr; //可选
    cmd.Start(); //这个不会等待
	
	cmd.Wait(); //等待进程结束//必须要有这个,否则会立即跳过去了
	
	
}//

func Execute_php3() {
	
	defer PrintError("Execute_php3()"); //其实可以不用
	
	//其实执行一个 bat 文件是最好的,因为路径可能会变

	////cmd := exec.Command("D:\\Program Files (x86)\\NuSphere\\PhpED\\php54\\php-cgi.exe", "-b", "127.0.0.1:9000");
	cmd := exec.Command(GetCurrentPath() + "/" + "php_start.bat");
    //这样就可以把 php 命令行的结果重定向了
	cmd.Stdout = os.Stdout;
    cmd.Stderr = os.Stderr; //可选
    cmd.Start(); //这个不会等待
	
	cmd.Wait(); //等待进程结束//必须要有这个,否则会立即跳过去了
	
	
}//

//golang 获取当前程序执行路径
func GetCurrentPath() (string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]));
	if err != nil {
		//log.Fatal(err)
		return "";
	}
	fmt.Println(dir);
	
	return dir;
}//
