package handler

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	frameworkInter "gitlab.com/pplayground/pet_tracking/main-framework/Interfaces"
	frameworkRepo "gitlab.com/pplayground/pet_tracking/main-framework/handler/repository"
	frameworkUtils "gitlab.com/pplayground/pet_tracking/main-framework/utils"
	"gitlab.com/pplayground/pet_tracking/user-service/handler/model"
	"gitlab.com/pplayground/pet_tracking/user-service/handler/repository"
	"log"
	"strconv"
)

func UpdateUser(ctx *fasthttp.RequestCtx) {
	log.Println("---- UpdateUser function is triggered ----")
	bodyByte := ctx.Request.Body()
	var user model.User
	if err:= json.Unmarshal(bodyByte, &user); err != nil {
		log.Println(err)
		res := frameworkInter.CreateHTTPResponsePayload(nil, "Failed", "invalid input")
		frameworkUtils.HTTPResponse(ctx, res, 400)
		return
	}

	dbName := "db_petTracking"
	sql, err := frameworkRepo.ConnectSqlDB(dbName)
	if err != nil {
		log.Println(err)
		res := frameworkInter.CreateHTTPResponsePayload(nil, "Failed", "cannot connect to database")
		frameworkUtils.HTTPResponse(ctx, res, 400)
		return
	}
	defer frameworkRepo.DisConnectSqlDB(sql)

	param := ctx.UserValue("userId")
	userId, err := strconv.ParseInt(param.(string), 10, 64)
	if err != nil {
		log.Println(err)
		res := frameworkInter.CreateHTTPResponsePayload(nil, "Failed", "this request required userId as integer")
		frameworkUtils.HTTPResponse(ctx, res, 400)
		return
	}

	user.Id = int(userId)
	editUser := model.EditUser(user)
	err = repository.UpdateUser(sql, editUser)
	if err != nil {
		res := frameworkInter.CreateHTTPResponsePayload(nil, "Failed", err.Error())
		frameworkUtils.HTTPResponse(ctx, res, 400)
		return
	}
	res := frameworkInter.CreateHTTPResponsePayload(nil, "Success", nil)
	frameworkUtils.HTTPResponse(ctx, res, 200)
}