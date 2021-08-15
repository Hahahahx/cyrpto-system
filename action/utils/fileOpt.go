package utils

import (
	"crypto-system/internal/context"
	"errors"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// 通过文件名获取文件绝对路径
func FileAbsPath(file *os.File) string {

	absPath, err := filepath.Abs(filepath.Dir(file.Name()))
	context.App.Logger.Error(err)

	// 有时候文件名file.Name()本身是一个相对于当前执行文件的路径
	// 所以如果文件不是在当前目录下那么就会出现其他的路径
	// 除非直接使用file.Stat()中的Name()
	// 否则避免出现奇怪的问题需要加一层Base
	return filepath.Join(absPath, filepath.Base(file.Name()))
}

// 文件循环遍历读取buf,通过全局配置的chunk
func FileForEach(file *os.File, fn func(buf []byte)) {
	info, _ := file.Stat()
	filesize := info.Size()
	filechunk := context.App.Config.File.Chunk

	blocks := uint64(math.Ceil(float64(filesize) / float64(filechunk)))

	for i := uint64(0); i < blocks; i++ {

		// 末尾数据小于chunk的话
		blocksize := int(math.Min(float64(filechunk), float64(filesize-int64(i*uint64(filechunk)))))

		if filesize < filechunk {
			blocksize = int(filesize)
		}

		buf := make([]byte, blocksize)
		file.Read(buf)
		fn(buf)
	}
}

// 文件下载工具
func FileDownload(url, filename string, decrypt func(buf []byte) []byte, progressBar func(length, downlen int64)) {
	var (
		written, fsize int64
		buf            = make([]byte, context.App.Config.File.Chunk)
	)

	client := new(http.Client)
	// client.Timeout=time.Second * 60 // 设置超时时间
	//  get方法获取资源
	resp, err := client.Get(url)

	context.App.Logger.Error(err)

	fsize, err = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)

	context.App.Logger.Error(err)

	if FileIsExist(filename, fsize) {
		return
	}

	file, err := os.Create(filename)

	context.App.Logger.Error(err)

	defer file.Close()

	if resp.Body == nil {
		context.App.Logger.Error(errors.New("request body is null"))
	}

	defer resp.Body.Close()

	// 下列流程其实可以直接使用io.copyBuffer()
	// 但是我们下载的时候需要对文件进行解密
	// 使用流加密，直接流解密即可，读多少解密多少写多少
	for {
		// 每次下载
		nr, er := resp.Body.Read(buf)
		if nr > 0 {
			// 下载多少解密多少写入多少
			// 解密使用回调函数留给外部决定是否解密
			// 即是否是加密文件
			// 如果不是加密文件直接返回buf即可
			// 如果是加密文件则在回调中完成解密
			nw, ew := file.Write(decrypt(buf[0:nr]))
			// 数据长度大于0
			if nw > 0 {
				written += int64(nw)
			}
			// 写入出错
			if ew != nil {
				err = ew
				break
			}
		}
		// 读取下载的数据时，数据长度不等于写入的数据长度
		if er != nil {
			// 空数据
			if er != io.EOF {
				err = er
			}
			break
		}
		// 没有异常，回调进度条
		progressBar(fsize, written)
	}
	file.Close()
	context.App.Logger.Error(err)
}

// 判断文件是否存在
// 存在只有一种可能性，那就是文件已经下载好了
// 如果是半成品就会直接删除，当作没有文件
func FileIsExist(filename string, filesize int64) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		context.App.Logger.Error(err)
	}

	if filesize == info.Size() {
		context.App.Logger.Info("file is download already")
		return true
	}

	del := os.Remove(filename)
	if del != nil {
		context.App.Logger.Error(errors.New("unknown error"))
	}

	return false
}

// 将内存中的数据保存到files文件夹中
func FileSave(data []byte, filename string) {

	download := filepath.Join(context.App.Config.Path.Download(), filename)

	file, err := os.OpenFile(download, os.O_CREATE|os.O_WRONLY, 0777)
	context.App.Logger.Error(err)
	defer file.Close()

	_, err = file.Write(data)
	context.App.Logger.Error(err)
}
