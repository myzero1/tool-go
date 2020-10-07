package z1err

import (
	"errors"
	"fmt"
	"log"
	"testing"
)

/*
	http://c.biancheng.net/view/124.html
	https://www.cnblogs.com/zhangboyu/p/7911190.html

	错误和异常是可以相互转换的（形式上和逻辑上）

	错误处理的正确姿势
	姿势一：失败的原因只有一个时，不使用error
	姿势二：没有失败时，不使用error
	姿势三：error应放在返回值类型列表的最后
	姿势四：错误值统一定义，而不是跟着感觉走
	姿势五：错误逐层传递时，层层都加日志
	姿势六：错误处理使用defer
	姿势七：当尝试几次可以避免失败时，不要立即返回错误
	姿势八：当上层函数不关心错误时，建议不返回error
	姿势九：当发生错误时，不忽略有用的返回值

	异常处理的正确姿势
	姿势一：在程序开发阶段，坚持速错
	姿势二：在程序部署后，应恢复异常避免程序终止
	姿势三：对于不应该出现的分支，使用异常处理
	姿势四：针对入参不应该有问题的函数，使用panic设计

*/
func TestHandleError(t *testing.T) {
	fmt.Println(`----------test split line,TestHandleError-----------`)
	err := HandleError()
	if nil != err {
		log.Printf(`=========`)
		log.Printf(`%+v`, err)
		log.Printf(`---------------------------`)
		log.Println(StackSkipPrint(err))
		log.Printf(`---------------------------`)
		log.Println(StackSkipPrint(err, 7))
		log.Printf(`=========`)
		log.Printf(`++++++++++++++`)
		err2 := errors.New("this is a error in TestHandleError \n this is new line. ")
		log.Printf(`%+v`, err2)
		log.Printf(`---------------------------`)
		log.Println(StackSkipPrint(err2))
		log.Printf(`---------------------------`)
		log.Println(StackSkipPrint(err2, 7))
		log.Printf(`++++++++++++++`)
	}
}

func TestHandleException(t *testing.T) {
	fmt.Println(`----------test split line,TestHandleException-----------`)
	HandleException()
}

func TestCheckErr(t *testing.T) {
	fmt.Println(`----------test split line,TestCheckErr-----------`)
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
			fmt.Println(`--------------------`)
			fmt.Printf(`%+v`, r)
		}
	}()

	noerr, _ := CheckErr(nil)
	if noerr {
		noerr, _ = CheckErr(nil)
		if noerr {
			err := errors.New(`this is a error in TestCheckErr`)
			CheckErr(err, true, "just a addon msg")
		}
	}
}

func HandleException() {
	defer Handle(nil, func(err error) {
		log.Println(`++++++++++++`)
		log.Println(StackSkipPrint(err))
		log.Println(`++++++++++++`)
	})

	Check(nil)
	Check(nil)
	Check(nil)
	HandleException2()
	return
}

func HandleException2() {
	err4 := errors.New(`This err in HandleException2`)
	Check(err4)
}

func HandleError() (err error) {
	defer Handle(&err)
	Check(nil)
	Check(nil)
	Check(nil)
	err4 := errors.New("This err in HandleError \n This is new line")
	Check(err4)
	return
}
