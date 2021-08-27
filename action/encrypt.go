package action

import (
	"crypto-system/action/handle"
	"crypto-system/internal/context"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func Encrypt(opts *EncryptOptions) {

	start := time.Now() // 获取当前时间

	file, err := os.Open(opts.Filename)
	context.App.Logger.Error(err)

	fileInfo, err := file.Stat()
	context.App.Logger.Error(err)

	cache, key := handle.EncryptFileCache(file)
	// 此处直接Close就行了
	// 不然下面移动文件的时候会发生
	// 文件尚未关闭的错误
	// 导致无法移动
	cache.Close()

	key = handle.EncryptByLocalKey(key)

	context.App.Logger.Log(fileInfo.Name() + "加密完成 √")
	context.App.Logger.Log("保存好您的文件密钥，它将是唯一能解开该文件的重要凭借:\n\t", key)

	// 保存密钥文件
	handle.FileSave([]byte(key), opts.Filename+".key")

	// 将缓存中的加密数据拷贝到files文件夹中
	downloadFile := filepath.Join(context.App.Config.Path.Download(), filepath.Base(file.Name())+".encrypt")
	err = os.Rename(handle.FileAbsPath(cache), downloadFile)

	context.App.Logger.Error(err, "remove file is err")

	elapsed := time.Since(start)
	fmt.Println("该命令执行完成耗时：", elapsed)

}
