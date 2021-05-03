package main

/**
 * Syc <github.com/SycAlright>
 * Hostloc_Avatar Code
 */

import (
	"bytes"
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

// Edit your data!
const (
	INPUT = "775f64jN3UVIp%SDxTzGCAwy66vZs1RU6%2Fe9cYo2%2FBdLuCG76%2FV5%2BZlkNDIWKtoQmLKbJUKj0qVklclJsoFmGvmX0sIQ5Jwlpcno9TjL%2FJoYSAjaHq2r1ObEbeJ9Xd"
	AGENT = "3651f000000000000000000000000a66"
)

func main() {
	S := hexImage("45.jpg")
	M := hexImage("120.jpg")
	L := hexImage("200.jpg")
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
