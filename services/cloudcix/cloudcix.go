package cloudcix

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"netsoc/cloud/models"
	"netsoc/cloud/services"
	"netsoc/cloud/services/cloudcix/cloudcix_models"

	"github.com/Strum355/log"
	"github.com/spf13/viper"
)

type CloudCIXService struct {
	services.Service
}

func (*CloudCIXService) CreateService() error {
	log.WithFields(log.Fields{
		"user": viper.GetString("cloud.email"),
	}).Info("Logging in as user")
	jsonData := map[string]string{"email": viper.GetString("cloud.email"), "password": viper.GetString("cloud.password"), "api_key": viper.GetString("cloud.api_key")}
	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post("https://membership.api.cloudcix.com/auth/login/", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.WithError(err).Error("Error getting response.")
		return err
	}

	if response.StatusCode != 201 {
		log.WithFields(log.Fields{
			"response_code": response.StatusCode,
		}).Error("CloudCIX create token response not 201.")
		return errors.New("Response not 201")
	}

	body := make(map[string]string)
	json.NewDecoder(response.Body).Decode(&body)

	viper.Set("cloud.token", body["token"])
	log.Info("CloudCIX token set.")
	return nil
}

func (*CloudCIXService) sendAPIRequest(uri string, method string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("X-Auth-Token", viper.GetString("cloud.token"))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CloudCIXService) GetVMs() ([]models.VM, error) {
	resp, err := c.sendAPIRequest("https://api.cloudcix.com/DNS/v1/VM/", "GET", nil)
	if err != nil {
		return nil, err
	}

	body := struct {
		Content []cloudcix_models.VM `json:"content"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	vms := make([]models.VM, 0)

	for _, i := range body.Content {
		vms = append(vms, models.VM{
			ID:        i.ID,
			Username:  "undefined",
			Project:   "undefined",
			Name:      i.Name,
			RAM:       i.RAM,
			CPU:       i.CPU,
			State:     i.State,
			Created:   i.Created,
			Updated:   i.Updated,
			ImageID:   i.ImageID,
			DNS:       i.DNS,
			ProjectID: i.ProjectID,
		})
	}

	return vms, nil
}

func (c *CloudCIXService) CreateVM(vm cloudcix_models.VM) error {
	jsonValue, _ := json.Marshal(vm)
	response, err := c.sendAPIRequest("https://api.cloudcix.com/DNS/v1/VM/", "POST", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.WithError(err).Error("Error getting response.")
		return err
	}

	fmt.Println(fmt.Sprint(response.StatusCode))

	r, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(r))
	return nil
}

func (c *CloudCIXService) GetClouds() {
	response, err := c.sendAPIRequest("https://api.cloudcix.com/DNS/v1/Cloud/", "GET", nil)
	if err != nil {
		log.WithError(err).Error("Error getting response.")
		return
	}

	fmt.Println(fmt.Sprint(response.StatusCode))

	r, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(r))
	return
}

func (c *CloudCIXService) GetProjects() ([]cloudcix_models.Project, error) {
	resp, err := c.sendAPIRequest("https://api.cloudcix.com/DNS/v1/Project/", "GET", nil)
	if err != nil {
		return nil, err
	}

	body := struct {
		Content []cloudcix_models.Project `json:"content"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body.Content, nil
}
