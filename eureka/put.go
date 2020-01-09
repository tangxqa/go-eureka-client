package eureka

import (
	asynclog "github.com/alecthomas/log4go"
	"github.com/tangxqa/gogate/discovery"
	"net/http"
	"strings"
)

func (c *Client) SendHeartbeat(appId, instanceId string) error {
	values := []string{"apps", appId, instanceId}
	path := strings.Join(values, "/")
	resp, err := c.Put(path, nil)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case http.StatusNotFound:
		return newError(ErrCodeInstanceNotFound,
			"Instance resource not found when sending heartbeat", 0)
	case http.StatusBadGateway:
		asynclog.Info("receive 502, send regist instance info...!")
		discovery.SendRegistInstanceInfo()

	}
	return nil
}
