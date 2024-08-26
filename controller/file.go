package controller

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"mediahub/pkg/config"
	"mediahub/pkg/log"
	"mediahub/pkg/storage"
	"mediahub/pkg/zerror"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	sf     storage.StorageFactory
	log    log.ILogger
	config *config.Config
}

func NewController(sf storage.StorageFactory, log log.ILogger, cnf *config.Config) *Controller {
	return &Controller{
		sf:     sf,
		log:    log,
		config: config.GetConfig(),
	}
}

// func (c *Controller) Upload(r io.Reader, md5Digest []byte, dstPath string) (url string, err error) {
// 	return c.sf.CreateStorage().Upload(r, md5Digest, dstPath)
// }

func (c *Controller) Upload(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		c.log.Error(zerror.NewByErr(err))
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		c.log.Error(zerror.NewByErr(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	defer file.Close()
	conntent, err := io.ReadAll(file)
	if err != nil {
		c.log.Error(zerror.NewByErr(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	//校验格式
	if !isImage(io.NopCloser(bytes.NewReader(conntent))) {
		err = zerror.NewByMsg("仅支持jpg、png、gif格式")
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	md5Digest := calMD5Digest(conntent)
	filename := fmt.Sprintf("%x%s", md5Digest, path.Ext(fileHeader.Filename))
	filepath := "/public/" + filename
	// println(filepath)
	if userId != 0 {
		filepath = fmt.Sprintf("/%d/%s", userId, filepath)
	}

	s := c.sf.CreateStorage()

	url, err := s.Upload(io.NopCloser(bytes.NewReader(conntent)), md5Digest, filepath)
	if err != nil {
		c.log.Error(zerror.NewByErr(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"url": url,
	})

	//    conntent,err:=io.ReadAll()
	//    if err!=nil{
	// 	c.log.Error(zerror.NewByErr(err))

	//    }
	//生成短链接
	return
}

func isImage(r io.Reader) bool {
	_, _, err := image.Decode(r)
	if err != nil {
		return false
	} else {
		return true
	}
}

func calMD5Digest(msg []byte) []byte {
	m := md5.New()
	m.Write(msg)
	md5Digest := m.Sum(nil)
	return md5Digest
}
