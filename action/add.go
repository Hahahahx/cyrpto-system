package action

import (
	"crypto-system/action/handle"
	"crypto-system/action/request"
	"crypto-system/internal/context"
	"fmt"
	"io"
	"os"
	"time"
)

func Add(opts *AddOptions) {

	start := time.Now() // 获取当前时间

	file, err := os.Open(opts.Filename)
	context.App.Logger.Error(err)
	// defer file.Close()

	fileInfo, err := file.Stat()
	context.App.Logger.Error(err)

	md := &request.MateData{
		Name: fileInfo.Name(),
		Size: fileInfo.Size(),
	}

	md.MD5 = ""
	md.Encrypt = opts.Encrypt

	if opts.Encrypt {
		res, frigerPrint := handle.VerifyMD5(file)

		// MD5校验以后会改变文件内容？？
		// 总之不把它关闭在重新打开所生成的加密文件
		// 和未经过MD5直接加密的文件就不一致了！
		// 目前不知道是什么原因导致的，但是重新打开文件应该就没问题了
		file.Close()

		md.MD5 = frigerPrint
		if res["hasFile"].(bool) {
			context.App.Logger.Log("上传成功,CID: ", res["CID"])
			return
		}

		file, err = os.Open(opts.Filename)
		context.App.Logger.Error(err)

		cache, key := handle.EncryptFileCache(file)

		// 上传完成后应该直接关闭文件
		// 下面的步骤需要删除缓存文件
		// 如果没有关闭则会报出文件已打开的错误
		cache.Close()
		key = handle.EncryptByRemoteKey(key)

		md.Key = key

		// 此处重新打开了cache文件
		// 如果不这样做ipfs add就没法读取到文件内容
		// 上传的始终是空文件
		// todo：通过优化可以将上面作为并发读写，这样就不需要重新打开了
		read, err := os.Open(cache.Name())
		context.App.Logger.Error(err)

		upload(read, md)

		os.Remove(cache.Name())

		elapsed := time.Since(start)
		fmt.Println("该命令执行完成耗时：", elapsed)
		return

	}

	file, err = os.Open(opts.Filename)
	context.App.Logger.Error(err)

	upload(file, md)

	elapsed := time.Since(start)
	fmt.Println("该命令执行完成耗时：", elapsed)
}

func upload(file io.Reader, md *request.MateData) {

	cid, err := context.App.Ipfs.Add(file)

	context.App.Logger.Error(err)
	md.CID = cid

	// 上传matedata数据
	request.UploadFile(md)
	context.App.Logger.Log("上传成功,CID: ", cid)
}
