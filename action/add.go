package action

import (
	"bytes"
	"crypto-system/action/request"
	"crypto-system/internal/context"
	"fmt"
	"os"
	"time"
)

func Add(c *context.Request) {

	start := time.Now() // 获取当前时间
	filename := c.Cli.Args().First()

	isEncrypt := c.Cli.Bool("e")

	file, err := os.Open(filename)
	c.App.Logger.Error(err)

	if isEncrypt {

		fileInfo, err := file.Stat()
		c.App.Logger.Error(err)

		md := &request.MateData{
			Name: fileInfo.Name(),
			Size: fileInfo.Size(),
		}

		buf := make([]byte, fileInfo.Size())

		file.Read(buf)

		file.Close()

		res, frigerPrint := verifyMD5(c, buf)

		if res["hasFile"].(bool) {
			c.App.Logger.Log("上传成功,CID: ", res["CID"])
			return
		}

		md.MD5 = frigerPrint

		buf, key := encryptFile(c, buf)

		key = remoteEncryptKey(c, key)

		md.Key = key

		read := bytes.NewReader(buf)
		cid, err := c.App.Ipfs.Add(read)

		c.App.Logger.Error(err)

		md.CID = cid

		// todo：上传matedata数据
		request.UploadFile(c, md)
		c.App.Logger.Log("上传成功,CID: ", md.CID)
		elapsed := time.Since(start)
		fmt.Println("该命令执行完成耗时：", elapsed)
		return

	}
	cid, err := c.App.Ipfs.Add(file)
	c.App.Logger.Error(err)

	c.App.Logger.Log("上传成功,CID: ", cid)

	elapsed := time.Since(start)
	fmt.Println("该命令执行完成耗时：", elapsed)
}
