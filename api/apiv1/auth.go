/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-15 21:47:55
 * @LastEditTime: 2019-08-15 22:27:39
 * @LastEditors: Please set LastEditors
 */
package apiv1

import (
	"net/http"

	"github.com/labstack/echo"
)

// Token GET 获取Token
func Token(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "auth ping"})

}

// AuthPing GET for testing
func AuthPing(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "auth ping"})
}

// AuthVerification get
func AuthVerification(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "auth v"})
}
