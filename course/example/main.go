package main

import (
	"errors"
	"fmt"
	"os"

	courseSdk "github.com/juanjoaquin/back-g-sdk/course"
)

func main() {
	courseTrans := courseSdk.NewHttpClient("http://localhost:8082", "")

	user, err := courseTrans.Get("3427e338-dc98-4daa-8144-aed411fa4520")

	if err != nil {
		if errors.As(err, &courseSdk.ErrNotFound{}) {
			fmt.Println("Not found:", err.Error())
			os.Exit(1)
		}
		fmt.Println("internal server error:", err.Error())
		os.Exit(1)
	}

	fmt.Println(user)
}
