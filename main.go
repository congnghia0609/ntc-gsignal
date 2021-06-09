/**
 * @author nghiatc
 * @since Jun 07, 2021
 */
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"os"

	"github.com/congnghia0609/ntc-gsignal/gsignal"
)

type SignalServer struct {
}

func randSeq(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (s *SignalServer) AuthenticateRequest(params url.Values) (apiKey, room, sessionKey string, ok bool) {
	return "ABC", "ABC", randSeq(16), true
}

func (s *SignalServer) OnClientMessage(ApiKey, Room, SessionKey string, raw []byte) {
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
	}
	log.Printf("Signal server is running on port: %s", port)
	fmt.Println(gsignal.Start(&SignalServer{}, port))
}
