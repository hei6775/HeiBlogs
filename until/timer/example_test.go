package timer

import (
	"testing"

	"log"
	"fmt"
	"time"
)

func TestTimer(t *testing.T){
	//10缓存
	d := NewDispatcher(10)

	//timer 1
	d.AfterFunc(1, func() {
		log.Print("My name is Leaf")
	})

	//timer 2
	time2 := d.AfterFunc(1, func() {
		log.Print("will not print")
	})
	time2.Stop()

	//dispatch(调度器)
	(<-d.ChanTimer).Cb()

	//Output:
	//My name is Leaf
}

//cronExpr测试
func TestCronExpr(t *testing.T){
	cronExpr,err := NewCronExpr("30 * * * *")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("cronExpr",*cronExpr)
	fmt.Println(cronExpr.Next(time.Date(
		2000,1,1,
		20,30,5,
		0,time.UTC,
	)))

	tt := time.Date(
		2000,1,1,
		20,30,5,
		0,time.UTC,
	)
	fmt.Println("tt:",tt)
	tt = tt.Truncate(time.Second).Add(time.Second)
	fmt.Println("tt operation:",tt)
	fmt.Println("uint tt second:",uint(tt.Second()))
	fmt.Printf("binary[%b], value[%v] \n",1<<uint(tt.Second()),1<<uint(tt.Second()))
	fmt.Printf("binary[%b], value[%v] \n",1,1)
	fmt.Println(1<<uint(tt.Second())&1)
	fmt.Println(1<<uint(tt.Second())&1 == 0 )
	//Output:
	//cronExpr {1 1073741824 16777215 4294967294 8190 127}
	//2000-01-01 21:30:00 +0000 UTC
	//tt: 2000-01-01 20:30:05 +0000 UTC
	//tt operation: 2000-01-01 20:30:06 +0000 UTC
	//uint tt second: 6
	//binary[1000000], value[64]
	//binary[1], value[1]
	//0
	//true
}

func TestCron(t *testing.T){
	d := NewDispatcher(10)

	// cron expr
	cronExpr,err := NewCronExpr("* * * * * *")
	if err != nil {
		return
	}
	//cron
	var c *Cron
	c = d.CronFunc(cronExpr, func() {
		fmt.Println("My name is Leaf")
		c.Stop()
	})

	//dispatch
	(<-d.ChanTimer).Cb()

}