package wechat

import (
	"context"
	"github.com/xgpc/dsg/service/grpcService/proto"
)

// BindUser 绑定用户
func BindUser(mobile string, openID string, sysCode uint32) error {
	c := proto.NewWechatServiceClient(proto.GRPCConn)
	//调用函数
	_, err := c.BindUser(context.Background(), &proto.WechatMobile{
		Mobile:  mobile,
		OpenID:  openID,
		SysCode: sysCode,
	})
	return err
}
