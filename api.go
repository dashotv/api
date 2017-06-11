package main

import "os"

func init() {
	api.Use(Auth(os.Getenv("TOKEN_SECRET")))
}
