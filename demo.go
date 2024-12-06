package main

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/rand"
	"crypto/rc4"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/http/cgi"
	"os"
)

func main1() {
	// ruleid: insecure-module-used
	cgi.Serve(http.FileServer(http.Dir("/usr/share/doc")))
}

func main1() {
	// ruleid: insecure-module-used
	cgi.Serve(http.FileServer(http.Dir("/usr/share/doc")))
}

func main2() {
	fmt.println("testing diff aware")

main2
