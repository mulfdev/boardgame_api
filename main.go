package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://www.yucata.de/Services/YucataService.svc/GetRankingList", nil)

	if err != nil {
		log.Fatal("could not make request", err)

	}

	req.Header.Add("ASP.NET_SessionId", "p4cpevtdw3tvrhjfdb1enlw3")
	req.Header.Add("Yucata", "D271A1E82EB14A74D0788E706E682A9FC3B62AF60915656E8EF51D7C36914877B85F2C796ACD81BFBF53B877D12BEF584D310CEF6FF4F35F1E6035FCD77A505B06685FE3DE6E31517E87122DF4282E5494B01A780B800624FB7C4FAB15DDE6834F3054B2")

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	parsedBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s", parsedBody[1:500])
}
