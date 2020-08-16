package main

import (
	"flag"
	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
	"gitlab.com/pplayground/pet_tracking/user-service/router"

	//"gitlab.com/pplayground/pet_tracking/user-service/router"
	"log"
	"os"
)

var (
	corsAllowHeaders     = "Origin, Accept, Content-Type, Accept-Encoding, Authorization, Set-Cookie"
	corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE"
	corsAllowCredentials = "true"
)

func CORS(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)

		next(ctx)
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Cannot load environment file")
	}
	port := os.Getenv("PORT")
	log.Printf("User Service running on port %v ...\n", port)

	addr := flag.String("addr", ":"+port, "TCP address binding to ")
	r := router.New()
	router.Mount(r)

	if err := fasthttp.ListenAndServe(*addr, CORS(r.Handler)); err != nil {
		log.Println(err)
	}

}
