package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
	"time"
)

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func ReadLog() []byte {
	homeDir := UserHomeDir()
	filepath := homeDir + "/AppData/LocalLow/NavilaSoftware/Bunny eShop/output_log.txt"
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Can not read log file!")
		panic(err)
	}
	return content
}

func PostPastebin(log []byte) string {
	currentTimeUnix := fmt.Sprint(time.Now().Unix())
	urlValues := url.Values{
		"api_dev_key":           {"PASTEBIN_API_KEY"},                   //api_dev_key
		"api_paste_code":        {string(log)},                          //api_paste_code
		"api_option":            {"paste"},                              //api_option
		"api_paste_name":        {"Bunny eShop log " + currentTimeUnix}, //api_paste_name
		"api_paste_format":      {"text"},                               //api_paste_format
		"api_paste_private":     {"1"},                                  //api_paste_private
		"api_paste_expire_date": {"1W"},                                 //api_paste_expire_date
	}
	reqBody := urlValues.Encode()
	resp, err := http.Post("https://pastebin.com/api/api_post.php", "application/x-www-form-urlencoded", strings.NewReader(reqBody))
	if err != nil {
		fmt.Println("Can not post to pastebin, please check your internet connection!")
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func main() {
	fmt.Println("Bunny eShop log report, press any key to start...")
	fmt.Scanln()
	log := ReadLog()
	fmt.Println("Please send this link to maintainer: \n" + PostPastebin(log))
	fmt.Println("\nPress any key to exit...")
	fmt.Scanln()
}
