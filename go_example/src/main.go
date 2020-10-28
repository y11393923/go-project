package main

import (
	"controller"
	"fmt"
	"model"
	"service"
	"strconv"
	"strings"
	"time"
	"tools"
)

func main() {
	helloGo()
	controller.InitWeb()
}

func helloGo() {
	fmt.Println("Hello world!", time.Now())
	var name string = "abc"
	fmt.Println(name)

	var a, b = 1, 2
	fmt.Println(a + b)

	var c int
	var d string
	c, d = 5, "go"
	fmt.Println(c, d)

	e, f := 3, "go, go"
	fmt.Println(e, f)

	//_不要第一个返回值
	_, str := split("0-hello go")
	fmt.Println(str)

	//调用函数并遍历
	array := tools.Sort([]int{3, 4, 2, 1})
	for _, v := range array {
		fmt.Println(v)
	}
	var aa = "go"
	fmt.Println(model.Student{Id: 123, Name: &aa})
	//调用接口函数
	var stu = model.Student{Id: 0, Name: &aa}
	var human service.Humaner = stu
	human.SayHi()

	testMap()

	testSlice()

	testGoroutine()
}

/**
 *  同包调用首字母小写，不同包需要首字母大写
 */
func split(a string) (int, string) {
	array := strings.Split(a, "-")
	//字符串转换
	num, e := strconv.Atoi(array[0])
	if e != nil {
		//抛出异常(类似java中的throw)
		panic(e)
	}
	return num, array[1]
}

func testMap() {
	/* 声明变量，默认 map 是 nil */
	var countryCapitalMap map[string]string
	/* 使用 make 函数 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["American"] /*如果确定是真实的,则存在,否则不存在 */
	fmt.Println(capital)
	fmt.Println(ok)
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}

	/* 创建map */
	countryCapitalMap = map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	fmt.Println("原始地图")

	/* 打印地图 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*删除元素*/
	delete(countryCapitalMap, "France")

	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")

	/*打印地图*/
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

}

func testSlice() {
	//slice1 := make([]int, 3)
	//
	//s := slice1[:]

	result := make(chan int)

	go func() {
		sum := 0
		for i := 0; i < 10; i++ {
			sum = sum + i
		}
		result <- sum
	}()
	fmt.Print(<-result)
}

func testGoroutine() {
	printer := func(c chan int) {
		// 开始无限循环等待数据
		for {
			// 从channel中获取一个数据
			data := <-c
			// 将0视为数据结束
			if data == 0 {
				break
			}
			// 打印数据
			fmt.Println(data)
		}
		// 通知main已经结束循环(我搞定了!)
		c <- 0
	}
	// 创建一个channel
	c := make(chan int)
	// 并发执行printer, 传入channel
	go printer(c)

	for i := 1; i <= 10; i++ {
		// 将数据通过channel投送给printer
		c <- i
		//睡眠5秒
		//time.Sleep(5 * 1e9)
	}
	// 通知并发的printer结束循环(没数据啦!)
	c <- 0
	// 等待printer结束(搞定喊我!)
	<-c

}
