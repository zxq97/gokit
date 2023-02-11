package host

import (
	"log"
	"testing"
)

func TestGetIP(t *testing.T) {
	ip, err := GetIP()
	if err != nil {
		t.Fatal(err)
	}
	log.Println(ip)
}
