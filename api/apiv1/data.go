/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-15 21:47:55
 * @LastEditTime: 2019-08-15 22:29:35
 * @LastEditors: Please set LastEditors
 */
package apiv1

import (
	"net/http"

	"github.com/labstack/echo"
)

// Datasets 数据集 GET/POST/DELETE
func Datasets(e echo.Context) error{
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "datasets"})
}

// DatasetsCheck 检查数据集是否重复 get only
func DatasetsCheck(e echo.Context) error{
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "datasets check"})

}

// DatasetsDetail 数据集详情 GET PATCH
func DatasetsDetail(e echo.Context) error{
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "datasets detail"})
}
