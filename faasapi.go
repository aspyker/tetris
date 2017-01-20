package main

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

const ip = "100.66.22.40"

type ResultMsg struct {
	Parameters ResultParameters `json:"parameters"`
}

type ResultParameters struct {
	BlockIndex int `json:"blockIndex"`
}

func getNextMinoIndex(gameid string) int {
	var jsonStr = []byte(`{}`)

	url := fmt.Sprintf("http://%s:7001/faas/randBlockIndex", ip)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))

	resultMsg := ResultMsg{}
	err = json.Unmarshal(body, &resultMsg)

	if err != nil {
		fmt.Printf("err = %v\n", err)
		panic("unable to demarshal")
	}

	//fmt.Printf("about to return block %v\n", resultMsg.Parameters.BlockIndex)
	return resultMsg.Parameters.BlockIndex
}

