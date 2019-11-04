package cloudcix

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"netsoc/cloud/services"
	"netsoc/cloud/services/cloudcix/models"

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

func (*CloudCIXService) sendAPIRequest(uri string, method string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Auth-Token", viper.GetString("cloud.token"))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CloudCIXService) GetVMs() ([]models.VM, error) {
	resp, err := c.sendAPIRequest("https://api.cloudcix.com/DNS/v1/VM/", "GET")
	if err != nil {
		return nil, err
	}

	body := struct {
		Content []models.VM `json:"content"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body.Content, nil
}

func (c *CloudCIXService) GetProjects() ([]models.Project, error) {
	resp, err := c.sendAPIRequest("https://api.cloudcix.com/DNS/v1/Project/", "GET")
	if err != nil {
		return nil, err
	}

	body := struct {
		Content []models.Project `json:"content"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body.Content, nil
}
