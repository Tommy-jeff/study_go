package studysteps

import (
	"fmt"

	"github.com/Tommy-jeff/study_go/requests"
)

func ErrorHandling() {
	// HTTP Request with error handling
	res1, err := requests.GoogleRequest()
	if err != nil {
		fmt.Printf("Error during googleRequest: %v\n", err)
		return
	}
	fmt.Printf("Response status code: %v\n %v\n", res1.StatusCode, res1.Header)

	// HTTP Request without error handling
	res2, _ := requests.GoogleRequest()
	
	fmt.Printf("Response status code: %v\n %v\n", res2.StatusCode, res2.Header)
}