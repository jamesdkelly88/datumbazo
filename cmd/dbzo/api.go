package main

import (
	"fmt"
	"io"
	"net/http"
)

func invokeAPICall(path string) (string, int, error) {
	url := fmt.Sprintf("http://%s:%d%s", cfg.Client.Hostname, cfg.Client.Port, path)
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", -1, err
	}
	req.SetBasicAuth(cfg.Client.Username, cfg.Client.Password)
	resp, err := client.Do(req)
	if err != nil {
		return "", -1, err
	}
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", -1, err
	}
	s := string(bodyText)
	return s, resp.StatusCode, nil
}
