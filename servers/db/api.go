package db

import (
	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/orange/database"
	"github.com/jqiris/orange/model"
)

func (d *DbServer) UpdateMember(req *model.DbUpdateMember) {
	logger.Infof("UpdateMember req:%+v", req)
	if err := database.UpdateUserMember(req.UserId, req.Updates); err != nil {
		logger.Error(err)
	}
}
