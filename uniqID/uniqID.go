package uniqID //uniqID    先用main 测试
//??  分开serverID 和clientID  貌似有点画蛇添足？直接返回totalID更好？？
// 暂时先分开  后面再看
import (
	"fmt"
)

const maxclient = 20000 //第一期任务每个服务器下20000台机子，用编号为0服务器测试
const firstserver = 100000

var totalID int // 这个ID 前面是serverID（通过%10万取得）   后面是clientID

func DelAServerID() {
	totalID = totalID - 1*firstserver
}

func DelAClientID() {
	totalID = totalID - 1
}

func GetServerID() int {
	totalID = totalID + 1*firstserver
	return int(totalID / firstserver)

}

func GetClientID() int {
	if (totalID % firstserver) < maxclient {
		totalID = totalID + 1
		return int(totalID % firstserver)
	} else {
		//this.totalID = this.totalID + 1*firstserver
		return 0
	}
}

/*
func main() {
	fmt.Println(GetServerID())
	for i := 0; i < 2; i++ {
		fmt.Println(GetClientID())
	}
	DelAClientID()
	DelAClientID()
	fmt.Println(GetClientID())
	DelAClientID()
	DelAClientID()
	//fmt.Println(GetClientID())
	fmt.Println(GetClientID())
}
*/
