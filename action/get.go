package action

import (
	"crypto-system/action/utils"
	"crypto-system/internal/context"
	"fmt"
	"os"
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
		fmt.Println(length, "======>", downlen)
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
		// fmt.Println(key)
		cache := filepath.Join(context.App.Config.Path.Cache(), opts.CID)

		// 由于AES-CTR分块必须加密时候多大，解密的时候就得多大
		// 所以只能下载文件下来然后再去读取文件解密
		downloadFile := utils.FileDownload(url, cache, func(buf []byte, len int64) []byte {
			return buf
			// return crypto.AesCTR_crypter(buf, key)
		}, progress)

		downloadFile.Close()

		// 必须得把文件关闭再打开才可以顺利读取
		// 不然下载下来的文件hash会有问题，导致解密失败
		downloadFile, err := os.Open(downloadFile.Name())
		context.App.Logger.Error(err)
		utils.DecryptFileCache(downloadFile, cacheFilePath, key)
		downloadFile.Close()
		os.Remove(downloadFile.Name())

	} else {
		downloadFile := utils.FileDownload(url, cacheFilePath, func(buf []byte, len int64) []byte {
			return buf
		}, progress)
		defer downloadFile.Close()
	}

	elapsed := time.Since(start)
	fmt.Println("该命令执行完成耗时：", elapsed)

}
