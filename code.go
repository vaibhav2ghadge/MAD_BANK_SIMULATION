package main
import(
	"fmt"
	"os"
	"strings"
	"strconv"
	"time"
	"sync"
	)
type Bank struct {
	customer int
	cashier chan int

}
type Utility struct {
	 wg sync.WaitGroup
}
type Branch struct {
	Banker Bank
	utility Utility
	repository
}
type Write interface {
	service(int,int)
}
type repository interface {
	Write
}
func (b *Branch)service(cno int,t int) {
	cashier_no := <-b.Banker.cashier
	fmt.Print(time.Now().Format("\n2006-01-02 15:04:05")," --> ","Cashier ",cashier_no,": Customer ",cno," Started")
	time.Sleep(time.Second*time.Duration(t))
	fmt.Print(time.Now().Format("\n2006-01-02 15:04:05")," --> ","Cashier ",cashier_no,": Customer ",cno," Completed")
	b.Banker.cashier<-cashier_no
	b.utility.wg.Done()
}
func main() {
	fmt.Print(time.Now().Format("2006-01-02 15:04:05")," --> Bank Simulation Started")
	args := os.Args
	customer,_:= strconv.Atoi(strings.Split(args[2],"=")[1])
	cashiers,_ := strconv.Atoi(strings.Split(args[1],"=")[1])
	bank := Bank{customer:customer,cashier:make(chan int,cashiers)}
	branch := &Branch{Banker:bank}
	t,_:= strconv.Atoi(strings.Split(args[3],"=")[1])
	for i:=1;i<=cashiers;i++ {
		branch.Banker.cashier<-i
	}
	for i:=1;i<=customer;i++ {
		branch.utility.wg.Add(1)
		go branch.service(i,t)		
	}
	branch.utility.wg.Wait()
	fmt.Print(time.Now().Format("\n2006-01-02 15:04:05")," --> Bank Simulated Ended\n")

}
