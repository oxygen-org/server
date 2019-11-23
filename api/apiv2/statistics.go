/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-15 21:47:55
 * @LastEditTime: 2019-08-15 22:29:00
 * @LastEditors: Please set LastEditors
 */
package apiv2

import (
	"net/http"

	"github.com/labstack/echo"
)

// StatResource GET 实时资源/每个节点资源的统计接口
func StatResource(e echo.Context) error{
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "stat resource"})
}

// StatResourceCluster 集群资源统计
func StatResourceCluster(e echo.Context) error{
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "stat cluster"})
}

// StatJob GET
func StatJob(e echo.Context) error{
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "stat job"})
}

// StateJobDistribute GET 根据参数num(x轴宽度)，来统计该时间段内有多少任务完成
func StateJobDistribute(e echo.Context) error{
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "stat job distribute"})

}

// StateDataset GET 统计文件的数量和大小
func StateDataset(e echo.Context) error{
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "stat dataset"})
}
