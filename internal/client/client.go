package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gohook/internal/structs/discord"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const baseURL = "https://discord.com/api"

var client = &http.Client{}

func init() {
	// Add support for socks5 in http_proxy and all_proxy
	for _, key := range []string{"HTTP_PROXY", "ALL_PROXY"} {
		var val string
		if v, ok := os.LookupEnv(key); ok {
			val = v
		} else {
			val = os.Getenv(strings.ToLower(key))
		}

		if strings.HasPrefix(val, "socks5://") {
			_ = os.Setenv("HTTPS_PROXY", val)
			break
		}
	}
}

func ExecuteWebhook(eventResult *discord.Webhook, creds discord.Credentials) error {
	url := fmt.Sprintf("%s/webhooks/%s/%s", baseURL, creds.ID, creds.Token)

	body, err := json.Marshal(eventResult)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("could not create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response

	for i := range 10 {
		resp, err = client.Do(req)
		if err == nil {
			break
		}
		log.Println("Request failed, retrying... Attempt", i+1, "Error", err)
		time.Sleep(time.Second)
	}

	if err != nil {
		return fmt.Errorf("could not send request: %w", err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println("Failed to close response body", "err", err.Error())
		}
	}()

	if resp.StatusCode != http.StatusNoContent {
		buf := new(bytes.Buffer)
		_, _ = buf.ReadFrom(resp.Body)
		return fmt.Errorf("discord api error: %s", buf.String())
	}

	return nil
}
