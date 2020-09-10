package cdn

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/qida/go/logs"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/cdn"
	"github.com/qiniu/api.v7/storage"
	"golang.org/x/net/context"
)

type QiNiu struct {
	Bucket string
	Url    string
	Mac    *qbox.Mac
	Config *storage.Config
}

func NewQiNiu(bucket string, url string, accessKey, secretKey string) *QiNiu {
	return &QiNiu{
		Bucket: bucket,
		Url:    url,
		Mac:    qbox.NewMac(accessKey, secretKey),
		Config: &storage.Config{
			Zone:          &storage.ZoneHuanan,
			UseHTTPS:      false,
			UseCdnDomains: true,
		},
	}
}

func (c *QiNiu) Upload(localFile io.Reader, size int64, file_name string) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: c.Bucket,
	}
	upToken := putPolicy.UploadToken(c.Mac)
	formUploader := storage.NewFormUploader(c.Config)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	err := formUploader.Put(context.Background(), &ret, upToken, file_name, localFile, size, &putExtra)
	if err != nil {
		return "", err
	}
	return ret.Key, nil
}

func (c *QiNiu) UploadFile(src_url string, file_data []byte) (url_file string, err error) {
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", c.Bucket, src_url),
	}
	upToken := putPolicy.UploadToken(c.Mac)
	formUploader := storage.NewFormUploader(c.Config)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": src_url,
		},
	}
	err = formUploader.Put(context.Background(), &ret, upToken, src_url, bytes.NewReader(file_data), int64(len(file_data)), &putExtra)
	if err == nil {
		url_file = ret.Key
		fmt.Printf("=====上传======\r\nKey:%s Hash:%s\r\n==============\r\n", ret.Key, ret.Hash)
		urlsToRefresh := []string{c.Url + url_file}
		cdnManager := cdn.NewCdnManager(c.Mac)
		fmt.Printf("=====刷新文件======\r\n%s\r\n==============\r\n", urlsToRefresh)
		_, err = cdnManager.RefreshUrls(urlsToRefresh)
	}
	return
}

func (c *QiNiu) MoveFile(src_url string, dst_url string) (err error, url_file string) {
	bucketManager := storage.NewBucketManager(c.Mac, c.Config)
	//如果目标文件存在，是否强制覆盖，如果不覆盖，默认返回614 file exists
	err = bucketManager.Move(c.Bucket, src_url, c.Bucket, dst_url, true)
	if err != nil {
		logs.Send2Dingf(logs.Rb错误, "移动文件：%s", err.Error())
		return
	}
	return
}

//返回私人有地址
func (c *QiNiu) GetPrivateMediaUrl(src_url string) (privateAccessURL string) {
	deadline := time.Now().Add(time.Minute * 60).Unix() //60分钟有效期
	privateAccessURL = storage.MakePrivateURL(c.Mac, c.Url, src_url, deadline)
	fmt.Println(privateAccessURL)
	return
}
