package https

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Res struct {
	Status  int
	context string
}

type Req struct {
	Target string
	Verb   string
	Dat    url.Values
}

type Server struct{}

// - Implement a simple HTTP server that serves static files and handles different HTTP methods (GET, POST, etc.).
func (s *Server) HandleReq(reqDat Req, res chan Res) {

	const timeout = 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	// showcase for deadline exceed, we wait req to timeout
	// time.Sleep(1000 + timeout)

	defer cancel()

	switch strings.ToUpper(reqDat.Verb) {
	case "GET":
		HttpHandler(ctx, reqDat, res)
	case "POST":
		HttpHandler(ctx, reqDat, res)
	default:
		panic("Method not allowed")
	}

}

func HttpHandler(ctx context.Context, reqDat Req, results chan<- Res) {

	var fc io.Reader

	if reqDat.Dat != nil {
		fc = strings.NewReader(reqDat.Dat.Encode())
	}

	req, err := http.NewRequestWithContext(ctx, strings.ToUpper(reqDat.Verb), reqDat.Target, fc)
	if err != nil {
		panic(err)
	}

	if reqDat.Dat != nil {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	client := http.DefaultClient // set client for create http response back
	resp, err := client.Do(req)
	if err != nil {
		results <- Res{context: fmt.Sprintf("Error making request to %s: %s", reqDat.Target, err.Error()), Status: 500}
		return
	}
	defer resp.Body.Close()

	results <- Res{context: fmt.Sprintf("Response from %s: %d", reqDat.Target, resp.StatusCode), Status: resp.StatusCode}
}
