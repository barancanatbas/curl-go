package curl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	CmdCurl = "curl"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:            CmdCurl,
		HelpName:        CmdCurl,
		Action:          Action,
		Usage:           `create a http request`,
		Description:     `It's a curl package where you can send simple http requests. You can get your output as a file.`,
		SkipFlagParsing: true,
		HideHelp:        true,
		HideHelpCommand: true,
	}
}

func Action(c *cli.Context) error {
	var request Request
	request.FindParams(c.Args().Slice())

	resp, err := Curl(&request)
	if err != nil {
		return err
	}
	Print(&request, resp)

	return err
}

func Curl(request *Request) (*http.Response, error) {
	client := &http.Client{}
	body, err := json.Marshal(request.Body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(request.Method, request.Url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if request.Header != nil {
		for k, v := range request.Header {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)
	return resp, err
}

func Print(request *Request, resp *http.Response) {
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if request.OutputName != "" {
		CreateOutputFile(request.OutputName, string(bodyBytes))
	} else {
		bodyString := string(bodyBytes)
		fmt.Println("status code  >>>>", resp.StatusCode)
		fmt.Println("\nheader       >>>>", resp.Header)
		fmt.Println("\ncontent type >>>>", resp.Header["Content-Type"])
		fmt.Println("\nbody         >>>> \n", bodyString)
	}
}

func CreateOutputFile(fileName string, data string) error {
	if data == "" {
		return nil
	}

	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(data))
	return err
}
