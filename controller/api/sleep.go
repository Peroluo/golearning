package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func calcSquares(number int, squareop chan int) {
	fmt.Println("calcSquares")
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	fmt.Println("calcCubes")
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	time.Sleep(time.Second * 2)
	fmt.Println("sleep执行完")
	cubeop <- sum
}

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func run2() {
	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		fmt.Println(v, ok)
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}
}

func run() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	fmt.Println("走起")
	go calcCubes(number, cubech) // sleep
	fmt.Print(3)
	go calcSquares(number, sqrch)
	fmt.Print(4)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}

func run3() []string {
	c := make(chan int, 5)
	s := []string{}
	var num int = 0
	go func() {
		for i := 0; i < 4; i++ {
			// time.Sleep(1 * time.Second)
			c <- i
			num++
			v := "inner=>" + strconv.Itoa(num) + ";set=>" + strconv.Itoa(i)
			s = append(s, v)
		}
	}()
	fmt.Println(num)
	for i := 0; i < 4; i++ {
		fmt.Println(num)
		val := <-c
		num++
		v := "outer=>" + strconv.Itoa(num) + ";value=>" + strconv.Itoa(val)
		s = append(s, v)
	}

	return s
}

func run4() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}

func run5() {
	c := make(chan int, 2)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		c <- 4
		close(c)
	}()

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

// Sleep 控制器
func Sleep(c *gin.Context) {
	// run()
	// run2()
	// arr := run3()
	run5()
	c.JSON(200, gin.H{
		"data": "sleep",
	})
}
