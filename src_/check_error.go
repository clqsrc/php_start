// 实现 try except 的异常捕捉, golang 有个好处,同一包的的代码可以放在多个文件中
package main

import (
//	"strconv"
//	"reflect"
//	"database/sql"
	"fmt"
//	"container/list"

	//_ "github.com/bmizerany/pq"
	//_"github.com/lib/pq" //驱动的写法一定要这样写,否则会当做无效的导入
)

func CheckErr(err error) {
    if err != nil {
        panic(err)  //相当于抛出一个异常,没有这个代码的话就不知道是什么错误了
    }
}//


func PrintError(funcName string){ // 必须要先声明defer，否则不能捕获到panic异常

    ////fmt.Println("c")

    if err:=recover();err!=nil{

		fmt.Print("err:[" + funcName + "]");
        fmt.Println(err) // 这里的err其实就是panic传入的内容，55

    }

    ////fmt.Println("d")

}//

//示例
////取一个数据结果集
//func (self * HMysql) Query() {
//	
//	//--------------------------------------------------
//	//异常处理要先写
//	defer PrintError("HMysql.Query"); //因为 golang 是自动释放的,所以实际上可以用一个公用的错误输出,大多数时候不需要另外处理
//	
//	defer fmt.Println("有资源要释放,写在这里");
//
//	//--------------------------------------------------
//    db, err := sql.Open("postgres", "host=127.0.0.1 user=postgres password=postgres dbname=postgres sslmode=disable")
//
//    self.checkErr(err)
//
//
//
//}//


