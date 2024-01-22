package system

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"

	"backend/api/system/v1"
)

// Pk 获取系统公钥
func (c *ControllerV1) Pk(ctx context.Context, req *v1.PkReq) (res *v1.PkRes, err error) {
	var publicKeySrc = "./data/public.pem"

	g.Log().Infof(ctx, "获取系统公钥")

	// 判断公钥文件是否存在
	if gfile.IsFile(publicKeySrc) {

		// 读取公钥
		publicKey := gfile.GetContents(publicKeySrc)

		// 过滤换行回车空格
		publicKey = gstr.Replace(publicKey, "\n", "", -1)
		publicKey = gstr.Replace(publicKey, "\r", "", -1)
		publicKey = gstr.Replace(publicKey, " ", "", -1)

		// 构建响应数据
		res = &v1.PkRes{
			Pk: publicKey,
		}

	} else {
		g.Log("system").Notice(ctx, "公钥文件不存在")

		// 构建响应数据
		res = &v1.PkRes{
			Pk: "公钥不存在",
		}
	}

	return res, nil
}
