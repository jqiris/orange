package db

import (
	"errors"

	"github.com/jqiris/kungfu/v2/rpc"
	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/model"
)

var (
	builder = rpc.DefaultReqBuilder().SetCodeType(rpc.CodeTypeJson).SetSuffix(rpc.JsonSuffix)
)

func UpdateMember(s rpc.ServerRpc, uid int, data map[string]any) error {
	dbServer := s.Find(constant.DbServer, uid)
	if dbServer == nil {
		return errors.New(constant.ErrNotFoundDbServer)
	}
	b := builder.Build()
	if err := s.Publish(b.SetMsgId(constant.DbMsgIdUpdateMember).SetServer(dbServer).SetReq(&model.DbUpdateMember{
		UserId:  uid,
		Updates: data,
	}).Build()); err != nil {
		return err
	}
	return nil
}
