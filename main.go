package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
		return
	}
}

func main() {

	timeout := time.Duration(5 * time.Second)
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{
		Timeout: timeout,
		Jar:     jar,
	}

	tmon_url, tmon_password := getTmonSettings()
	tracker_data := url.Values{
		"action":        {"torrent_add"},
		"name":          {""},
		"url":           {"https://nnmclub.to/forum/viewtopic.php?p=11205627#11205627"},
		"path":          {"/home/media/storage"},
		"update_header": {"true"},
	}
	series_data := url.Values{
		"action":  {"serial_add"},
		"tracker": {"tracker"},
		"name":    {"series_name"},
		"hd":      {"hd"}, //  0 -SD, 1 -1080p, 2 - 720p
		"path":    {"/home/media/storage"},
	}
	setCookies(client, tmon_url, tmon_password)
	AddTitleToMonitor(client, tmon_url, tmon_password, tracker_data)
	AddTitleToMonitor(client, tmon_url, tmon_password, series_data)
}

func getTmonSettings() (tmon_url, tmon_password string) {
	tmon_password, exists := os.LookupEnv("TMON_PASSWORD")
	if !exists {
		log.Fatal("add \"TMON_PASSWORD\" variable to .env file")
		return
	}
	tmon_url, exists = os.LookupEnv("TMON_URL")
	if !exists {
		log.Fatal("add \"TMON_URL\" variable to .env file")
		return
	}
	return tmon_url, tmon_password
}

func setCookies(client *http.Client, tmon_url, tmon_password string) {
	login_data := url.Values{
		"action":   {"enter"},
		"password": {tmon_password},
		"remember": {"true"},
	}
	req, err := http.NewRequest(
		"POST",
		tmon_url+"/action.php",
		strings.NewReader(login_data.Encode()),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Content-Length", strconv.Itoa(len(login_data.Encode())))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	u, err := url.Parse(tmon_url + "/action.php")
	if err != nil {
		log.Fatal(err)
	}
	client.Jar.SetCookies(u, resp.Cookies())

	defer resp.Body.Close()
}

func AddTitleToMonitor(
	client *http.Client,
	tmon_url string,
	tmon_password string,
	payload url.Values,
) (msg string, err error) {

	req, err := http.NewRequest(
		"POST",
		tmon_url+"/action.php",
		strings.NewReader(payload.Encode()),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Content-Length", strconv.Itoa(len(payload.Encode())))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	msg = fmt.Sprintf("%v", res["msg"])
	if res["error"] == false {
		err = nil
	}
	if res["error"] == true {
		err = errors.New(msg)
	}
	return msg, err
}
