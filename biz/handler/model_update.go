package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/go-playground/validator/v10"
	"github.com/siddontang/go/log"
	"net/http"
	"time"
)

type Request struct {
	Name string    `json:"name" validate:"required"`
	Data ModelData `json:"data" validate:"required"`
}

type ModelData struct {
	ModelPath     string    `json:"model_path" validate:"required,startswith=hdfs://"`
	IsActive      bool      `json:"is_active" validate:"boolean"`
	LastTrainTime time.Time `json:"last_train_time"`
}

type RequestMap map[string]interface{}

func (m *RequestMap) ToRequest() (req Request) {
	reqBytes, _ := json.Marshal(*m)
	_ = json.Unmarshal(reqBytes, &req)
	return req
}

func (r *Request) ToMap() map[string]interface{} {
	reqMap := make(map[string]interface{})
	reqBytes, err := json.Marshal(r)
	if err != nil {
		log.Errorf("json marshal error: %v", err.Error())
		return nil
	}
	err = json.Unmarshal(reqBytes, &reqMap)
	if err != nil {
		log.Errorf("json unmarshal error: %v", err.Error())
		return nil
	}
	return reqMap
}

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

var validFields = map[string]bool{
	"name":            true,
	"data":            true,
	"model_path":      true,
	"is_active":       true,
	"last_train_time": true,
}

// UpdateModel 更新模型
func UpdateModel(_ context.Context, c *app.RequestContext) {
	log.Info("start update model...")

	var raw RequestMap
	if err := c.Bind(&raw); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": fmt.Sprintf("parm err: %v", err),
		})
		return
	}

	if err := validateFields(raw); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	log.Info("pass validate fields!")

	if err := validate.Struct(raw.ToRequest()); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": fmt.Sprintf("validation error %v", err.Error()),
		})
		return
	}
	log.Info("pass bind and validate!")

	log.Infof("received request: %v", raw)

	c.JSON(consts.StatusOK, map[string]interface{}{
		"code":    consts.StatusOK,
		"message": "success",
	})
}

func validateFields(data map[string]interface{}) error {
	for key, value := range data {
		if _, ok := validFields[key]; !ok {
			return fmt.Errorf("invalid field: %s", key)
		}
		if nestedMap, ok := value.(map[string]interface{}); ok {
			if err := validateFields(nestedMap); err != nil {
				return err
			}
		}
	}

	return nil
}
