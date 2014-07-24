package total //先用main  搞定后改成uniqID

var (
	num = make(chan uint64) //通道传递 总数
)

func init() { //这里是用于统计客户端登陆过的总数，只要登陆就加1，客户端退出不影响总数
	go func() {
		for total := uint64(1); ; total++ {
			num <- total
		}
	}()
}

func GetTotal() uint64 { //返回的是 客户端登陆总数目
	return <-num
}

/*func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(GetTotal())
	}
}*/
