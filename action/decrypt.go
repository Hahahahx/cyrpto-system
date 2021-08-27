package action

import (
	"crypto-system/action/handle"
	"crypto-system/internal/context"
	"os"
	"path/filepath"
)

func Decrypt(opts *DecryptOptions) {

	var downloadFile string
	file, err := os.Open(opts.Filename)
	context.App.Logger.Error(err)

	fileInfo, err := file.Stat()
	context.App.Logger.Error(err)
	context.App.Logger.Log("fileSize : ", fileInfo.Size())

	key := handle.DecryptByLocalKey(string(opts.Key))

	// 查看是否指定了文件
	// 否则下载到files目录下
	// 并且命名为CID
	if opts.Newname != "" {
		downloadFile = opts.Newname
	} else {
		downloadFile = filepath.Join(context.App.Config.Path.Download(), filepath.Base(file.Name())+".decrypt")
	}

	handle.DecryptFileCache(file, downloadFile, key)

	context.App.Logger.Log(fileInfo.Name() + "解密完成 √")

}
