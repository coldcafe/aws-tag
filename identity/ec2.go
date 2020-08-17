package identity

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type InstanceIdentity struct {
	RegionID     string `json:"region"`
	InstanceID   string `json:"instanceId"`
	PrivateIPV4  string `json:"privateIp"`
	ImageId      string `json:"imageId"`
	ZoneID       string `json:"availabilityZone"`
	InstanceType string `json:"instance-type"`
}

func GetInstanceIdentity() (identity InstanceIdentity, err error) {
	metaUrl := "http://169.254.169.254/latest/dynamic/instance-identity/document"
	resp, err := http.DefaultClient.Get(metaUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	metaBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(metaBytes, &identity)
	if err != nil {
		return
	}

	return
}
