package action

import (
	"crypto-system/action/request"
	"fmt"
	"time"
)

func List() {

	start := time.Now() // 获取当前时间

	mdList := request.GetList()

	for _, md := range mdList {
		fmt.Println(md.ToString())
	}

	elapsed := time.Since(start)
	fmt.Println("该命令执行完成耗时：", elapsed)
}
