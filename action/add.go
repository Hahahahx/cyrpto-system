package action

import (
	"crypto-system/action/request"
	"crypto-system/action/utils"
	"crypto-system/internal/context"
	"fmt"
	"os"
	"time"
)

func Add(opts *AddOptions) {

	start := time.Now() // 获取当前时间

	file, err := os.Open(opts.Filename)
	context.App.Logger.Error(err)
	defer file.Close()

	if opts.Encrypt {

		fileInfo, err := file.Stat()
		context.App.Logger.Error(err)

		md := &request.MateData{
			Name: fileInfo.Name(),
			Size: fileInfo.Size(),
		}
		res, frigerPrint := utils.VerifyMD5(file)

		if res["hasFile"].(bool) {
			context.App.Logger.Log("上传成功,CID: ", res["CID"])
			return
		}

		md.MD5 = frigerPrint

		cache, key := utils.EncryptFileCache(file)
		key = utils.EncryptByRemoteKey(key)

		md.Key = key

		cid, err := context.App.Ipfs.Add(cache)
		context.App.Logger.Error(err)
		// 上传完成后应该直接关闭文件
		// 下面的步骤需要删除缓存文件
		// 如果没有关闭则会报出文件已打开的错误
		cache.Close()

		md.CID = cid

		// todo：上传matedata数据
		request.UploadFile(md)
		context.App.Logger.Log("上传成功,CID: ", md.CID)

		os.Remove(cache.Name())

		elapsed := time.Since(start)
		fmt.Println("该命令执行完成耗时：", elapsed)
		return

	}

	cid, err := context.App.Ipfs.Add(file)
	context.App.Logger.Error(err)
	context.App.Logger.Log("上传成功,CID: ", cid)

	elapsed := time.Since(start)
	fmt.Println("该命令执行完成耗时：", elapsed)
}
