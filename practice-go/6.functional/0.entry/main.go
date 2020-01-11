/*
函数式编程：
函数是一等公民：参数，变量，返回值，甚至可以为函数实现接口
高阶函数
函数->闭包

正统的函数编程
1.不可变性，不能有状态，只能有常量和函数
2.函数只能有一个参数
*/
package main

import (
	"fmt"
	"go-practice/6.functional/1.adder"
)

func main() {
	//这样就拿到了一个闭包：方法本身及其引用到的所有自由变量，以及自由变量所引用的一切，当都引用完了，就是一个完整体，
	//如下add不仅是个方法，还包含外部的自由变量
	//只要是同一个对象，那么里面的自由变量会一直保持他的状态
	add := adder.Add()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + 2 + ... + %d = %d \n", i, add(i))
	}

	//上面是比较活的函数编程，简单易懂，下面写一个标准的函数式编程
	//TODO 没怎么搞懂
	add2 := adder.Add2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, add2 = add2(i)
		fmt.Printf("0 + 1 + 2 + ... + %d = %d \n", i, s)
	}
}