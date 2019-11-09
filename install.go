package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func cmd(cmd string) string {
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil && err.Error() != "exit status 1" && err.Error() != "exit status 2" {
		log.Println("[error]", err, "("+cmd+")")
	}

	return string(out)
}

func getTrueIP() string {
	res, err := http.Get("https://v4.ident.me/")
	if err != nil {
		log.Println("[error]", err)
		return ""
	}

	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Println("[error]", err)
		}
	}()

	ip, err := ioutil.ReadAll(res.Body)

	if net.ParseIP(string(ip)) != nil && err == nil {
		return string(ip)
	}

	return ""
}

func base64EncodeStripped(s string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(s))
	return strings.TrimRight(encoded, "=")
}

func main() {
	log.Println("         Starting | Начало работы")

	ip := getTrueIP()

	if ip == "" {
		log.Println("Couldn't identify IP | Не удалось определить IP")
		return
	}

	cmd("apt update")
	cmd("apt -y install shadowsocks-libev")

	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	var b strings.Builder
	for i := 0; i < 10; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	port := strconv.Itoa(rand.Intn(65535-49152) + 49152)
	password := b.String()

	config := `
{
   "server":"0.0.0.0",
   "server_port":` + port + `,
   "password":"` + password + `",
   "timeout":300,
   "mode":"udp",
   "method":"chacha20-ietf-poly1305",
   "nameserver":"1.1.1.1",
   "no_delay":true,
   "reuse_port":true,
   "fast_open":true,
   "workers":1
}`

	f, err := os.Create("/etc/shadowsocks-libev/config.json")
	if err != nil {
		log.Println(err)
	}
	_, _ = f.WriteString(config)

	cmd("service shadowsocks-libev restart")

	log.Println("Program completed | Программа завершена")

	server := "chacha20-ietf-poly1305:" + password + "@" + ip + ":" + port

	fmt.Println("\n\n\nPlain:\nss://" + server)

	b64 := base64EncodeStripped(server)

	k, _ := os.Create("./ss-key.txt")
	_, _ = k.WriteString("ss://" + b64 + "\n")

	fmt.Println("\nEncoded:\nss://" + b64 + "\n\n\n")
}
