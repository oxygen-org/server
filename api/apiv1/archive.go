/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-15 21:47:55
 * @LastEditTime: 2019-08-15 22:28:08
 * @LastEditors: Please set LastEditors
 */
package apiv1

import (
	"net/http"

	"github.com/labstack/echo"
)

// Archive 文件
func Archive(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "archive"})
}

// ArchiveCheck 检查文件是否存在
func ArchiveCheck(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "archive check"})
}

// ArchiveDetail 文件详情
func ArchiveDetail(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "archive detail"})
}
