/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-15 21:47:55
 * @LastEditTime: 2019-08-15 22:28:51
 * @LastEditors: Please set LastEditors
 */
package apiv1

import (
	"net/http"

	"github.com/labstack/echo"
)

// Resources GET
func Resources(e echo.Context) error{
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "resource"})
}

// ResourceMonitor GET POST 资源管理
func ResourceMonitor(e echo.Context) error{
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "resource monitor"})

}

// ResourceStations PATCH GET
func ResourceStations(e echo.Context) error{
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "resource stations"})
}

// ResourceDispatcher POST GET
func ResourceDispatcher(e echo.Context) error{
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "resource dispatcher"})
}
