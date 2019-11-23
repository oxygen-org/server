/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-15 21:47:55
 * @LastEditTime: 2019-08-15 22:30:09
 * @LastEditors: Please set LastEditors
 */
package apiv1

import (
	"net/http"

	"github.com/labstack/echo"
)

// Images GET POST
func Images(e echo.Context) error{
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "images"})
}

// ImagesDeploy POST
func ImagesDeploy(e echo.Context) error{
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "images Deploy"})
}

// ImageConfirmation GET 查询某台主机是否某个镜像
func ImageConfirmation(e echo.Context) error{
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "images confirmation"})
}
