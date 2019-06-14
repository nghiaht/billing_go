package bhandler

import (
	"billing/models"
	"billing/tools"
	"database/sql"
	"fmt"
)

type RegisterHandler struct {
	Db *sql.DB
}

func (*RegisterHandler) GetType() byte {
	return 0xF1
}
func (h *RegisterHandler) GetResponse(request *BillingData) *BillingData {
	var response BillingData
	response.PrepareResponse(request)
	//读取请求信息
	var opData []byte
	//用户名
	offset := 0
	usernameLength := request.OpData[offset]
	tmpLength := int(usernameLength)
	offset++
	username := request.OpData[offset : offset+tmpLength]
	//超级密码
	offset += tmpLength
	tmpLength = int(request.OpData[offset])
	offset++
	superPassword := string(request.OpData[offset : offset+tmpLength])
	//密码
	offset += tmpLength
	tmpLength = int(request.OpData[offset])
	offset++
	password := string(request.OpData[offset : offset+tmpLength])
	//注册IP
	offset += tmpLength
	tmpLength = int(request.OpData[offset])
	offset++
	registerIP := string(request.OpData[offset : offset+tmpLength])
	//email
	offset += tmpLength
	tmpLength = int(request.OpData[offset])
	offset++
	email := string(request.OpData[offset : offset+tmpLength])
	//
	regResult := models.GetRegisterResult(h.Db, string(username), password, superPassword, email)
	tools.LogMessage(fmt.Sprintf("user [%v](%v) try to register from %v : %v", string(username), email, registerIP, regResult == 1))
	opData = append(opData, usernameLength)
	opData = append(opData, username...)
	opData = append(opData, regResult)
	response.OpData = opData
	return &response
}
