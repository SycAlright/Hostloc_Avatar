package main

/**
 * Syc <github.com/SycAlright>
 * Hostloc_Avatar
 */

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	INPUT = ""
	AGENT = ""
)

func main() {
	agent := flag.String("agent", "", "your agent")
	input := flag.String("input", "", "your input")
	avatar1 := flag.String("avatar1", "45.jpg", "45px avatar image path ")
	avatar2 := flag.String("avatar2", "120.jpg", "120px avatar image path")
	avatar3 := flag.String("avatar3", "200.jpg", "200px avatar image path")
	flag.Parse()
	AGENT = *agent
	INPUT = *input
	S := hexImage(*avatar1)
	M := hexImage(*avatar2)
	L := hexImage(*avatar3)
	uploadAvatar(S, M, L)
}

func hexImage(imagePath string) string {
	file, _ := os.Open(imagePath)
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		log.Println("Image Decode Fail!")
	}
	emptyBuff := bytes.NewBuffer(nil)
	png.Encode(emptyBuff, img)
	hex := fmt.Sprintf("%X", emptyBuff.Bytes())
	return hex
}

func uploadAvatar(s, m, l string) {
	timeUnix := time.Now().Unix()
	timeString := strconv.FormatInt(timeUnix, 10)
	data := "avatar1=" + l + "&avatar2=" + m + "&avatar3=" + s + "&urlReaderTS=" + timeString
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://hostloc.com/uc_server/index.php?m=user&inajax=1&a=rectavatar&appid=1&input="+INPUT+"&agent="+AGENT+"&avatartype=virtual", strings.NewReader(data))
	if err != nil {
		log.Println(err.Error())
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(body))
}
