package handler

import (
	"github.com/valyala/fasthttp"
	frameworkInter "gitlab.com/pplayground/pet_tracking/main-framework/Interfaces"
	frameworkRepo "gitlab.com/pplayground/pet_tracking/main-framework/handler/repository"
	frameworkUtils "gitlab.com/pplayground/pet_tracking/main-framework/utils"
	"gitlab.com/pplayground/pet_tracking/user-service/handler/repository"
	"log"
	"strconv"
)

func GetUsers(ctx *fasthttp.RequestCtx) {
	log.Println("---- GetUsers function is triggered ----")

	dbName := "db_petTracking"
	sql, err := frameworkRepo.ConnectSqlDB(dbName)
	if err != nil {
		log.Println(err)
		res := frameworkInter.CreateHTTPResponsePayload(nil, "Failed", "Cannot connect to database.")
		frameworkUtils.HTTPResponse(ctx, res, 400)
		return
	}
	defer frameworkRepo.DisConnectSqlDB(sql)

	userList, err := repository.GetUsers(sql)
	if err != nil {
		res := frameworkInter.CreateHTTPResponsePayload(nil, "Failed", err.Error())
		frameworkUtils.HTTPResponse(ctx, res, 400)
		return
	}
	res := frameworkInter.CreateHTTPResponsePayload(userList, "Success", nil)
	frameworkUtils.HTTPResponse(ctx, res, 200)
	return
}

func GetUser(ctx *fasthttp.RequestCtx) {
	log.Println("---- GetUser function is triggered ----")

	param := ctx.UserValue("userId")
	userId, err := strconv.ParseInt(param.(string), 10, 64)
	if err != nil {
		log.Println(err)
		res := frameworkInter.CreateHTTPResponsePayload(nil, "Failed", "This request required userId as integer.")
		frameworkUtils.HTTPResponse(ctx, res, 400)
		return
	}

	dbName := "db_petTracking"
	sql, err := frameworkRepo.ConnectSqlDB(dbName)
	if err != nil {
		log.Println(err)
		res := frameworkInter.CreateHTTPResponsePayload(nil, "Failed", "Cannot connect to database.")
		frameworkUtils.HTTPResponse(ctx, res, 400)
		return
	}
	defer frameworkRepo.DisConnectSqlDB(sql)

	user, err := repository.GetUser(sql, userId)
	if err != nil {
		res := frameworkInter.CreateHTTPResponsePayload(nil, "Failed", err.Error())
		frameworkUtils.HTTPResponse(ctx, res, 400)
		return
	}
	res := frameworkInter.CreateHTTPResponsePayload(user, "Success", nil)
	frameworkUtils.HTTPResponse(ctx, res, 200)
	return




}
