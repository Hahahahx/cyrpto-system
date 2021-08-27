package action

import (
	"crypto-system/action/request"
	"crypto-system/utils"
	"fmt"
	"time"

	"github.com/fatih/color"
)

func List() {

	start := time.Now() // 获取当前时间

	mdList := request.GetList()

	fmt.Println("------------------------------------------------------------------------------------------------")
	for _, md := range mdList {
		if md.Encrypt {
			fmt.Println(utils.BackGroundString(color.BgRed, " 加密文件 ") + "\t" + md.ToString())
		} else {
			fmt.Println(utils.BackGroundString(color.BgGreen, " 普通文件 ") + "\t" + md.ToString())
		}

	}
	fmt.Println("------------------------------------------------------------------------------------------------")

	elapsed := time.Since(start)
	fmt.Println("该命令执行完成耗时：", elapsed)
}
