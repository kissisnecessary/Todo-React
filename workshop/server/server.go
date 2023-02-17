
package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Hyperledger-TWGC/fabric-gm-plugins/workshop"
	restful "github.com/emicklei/go-restful/v3"
)

var pubFile = "pub.pem"
var path = "./"
var Key workshop.SM2
var err error

type Anything map[string]interface{}

func main() {
	path = os.Args[1]
	ws := new(restful.WebService)
	ws.Route(ws.POST("/verify").To(verify))
	ws.Route(ws.GET("/encrypt").To(encrypt))
	ws.Route(ws.GET("/sm4").To(sm4))

	restful.Add(ws)
	Key, err = workshop.LoadFromPubPem(path + pubFile)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("start server")
	}
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))