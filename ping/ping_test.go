package ping

import "fmt"

func Example1() {
	// Output:
	pinger := Default(NoopLogger{})
	pp := PingIPTask{}
	pp2 := PingIPTask{}
	pp.New("baidu.com", 5, NoopLogger{}, pinger)
	pp2.New("baidu.com", 5, NoopLogger{}, pinger)
	pp.Start()
	pp2.Start()
	fmt.Println(pp.Rst())
	fmt.Println(pp2.Rst())
}
