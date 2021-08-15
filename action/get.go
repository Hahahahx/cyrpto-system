package action

import (
	"crypto-system/action/utils"
	"crypto-system/internal/context"
	"crypto-system/internal/crypto"
	"fmt"
	"path/filepath"
	"time"
)

func Get(opts *GetOptions) {

	var (
		cacheFilePath, key string
	)

	start := time.Now() // 获取当前时间
	url := context.App.Config.Ipfs.GetFileURL(opts.CID)

	progress := func(length, downlen int64) {

	}

	// 查看是否指定了文件
	// 否则下载到files目录下
	// 并且命名为CID
	if opts.Newname != "" {
		cacheFilePath = opts.Newname
	} else {
		cacheFilePath = filepath.Join(context.App.Config.Path.Download(), opts.CID)
	}

	if opts.Decrypt {
		key = utils.DecryptByRemoteKey(opts.CID)
		utils.FileDownload(url, cacheFilePath, func(buf []byte) []byte {
			return crypto.AesCTR_Decrypt(buf, key)
		}, progress)
	} else {
		utils.FileDownload(url, cacheFilePath, func(buf []byte) []byte {
			return buf
		}, progress)
	}

	elapsed := time.Since(start)
	fmt.Println("该命令执行完成耗时：", elapsed)

}
