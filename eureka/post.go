package eureka

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (c *Client) RegisterInstance(appId string, instanceInfo *InstanceInfo) error {
	values := []string{"apps", appId}
	path := strings.Join(values, "/")
	instance := &Instance{
		Instance: instanceInfo,
	}
	body, err := json.Marshal(instance)
	fmt.Println(string(body))
	if err != nil {
		return err
	}

	_, err = c.Post(path, body)
	return err
}
