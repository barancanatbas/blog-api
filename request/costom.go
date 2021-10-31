package request

import (
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

func JSON(c echo.Context, obj interface{}) (interface{}, error) {

	read, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return obj, err
	}

	err = json.Unmarshal(read, obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}
