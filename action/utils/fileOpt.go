package utils

import (
	"crypto-system/internal/context"
	"math"
	"os"
	"path/filepath"
)

// 通过文件名获取文件绝对路径
func FileAbsPath(c *context.Request, file *os.File) string {

	absPath, err := filepath.Abs(filepath.Dir(file.Name()))
	c.App.Logger.Error(err)

	// 有时候文件名file.Name()本身就是一个绝对路径..
	// 避免出现奇怪的问题加一层Base
	return filepath.Join(absPath, filepath.Base(file.Name()))
}

// 文件循环遍历读取buf,通过全局配置的chunk
func FileForEach(c *context.Request, file *os.File, fn func(buf []byte)) {
	info, _ := file.Stat()
	filesize := info.Size()
	filechunk := c.App.Config.File.Chunk

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
