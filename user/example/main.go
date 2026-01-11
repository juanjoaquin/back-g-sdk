package main

import (
	"errors"
	"fmt"
	"os"

	userSdk "github.com/juanjoaquin/back-g-sdk/user"
)

func main() {
	userTrans := userSdk.NewHttpClient("http://localhost:8082", "")

	user, err := userTrans.Get("69d420cb-1186-4954-ac44-9d307d090f0d")

	if err != nil {
		if errors.As(err, &userSdk.ErrNotFound{}) {
			fmt.Println("Not found:", err.Error())
			os.Exit(1)
		}
		fmt.Println("internal server error:", err.Error())
		os.Exit(1)
	}

	fmt.Println(user)
}
