/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-15 21:47:55
 * @LastEditTime: 2019-08-15 22:28:31
 * @LastEditors: Please set LastEditors
 */
package apiv1

import (
	"net/http"

	"github.com/labstack/echo"
)

// Jobs GET POST PUT DELETE
func Jobs(e echo.Context) error{
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "jobs"})
}

// JobsDetail  GET 获取Jobs详情
func JobsDetail(e echo.Context) error{
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "jobs detail"})

}

// JobsRun POST 执行Job
func JobsRun(e echo.Context) error{
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "jobs run"})
}

// JobsKill  POST杀死Job
func JobsKill(e echo.Context) error{
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "jobs kill"})
}

// JobState GET PATCH PUT Job状态
func JobState(e echo.Context) error{
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "jobs state"})
}

// JobLog GET 获取任务日志
func JobLog(e echo.Context) error{
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "jobs log"})
}

// JobRPCService GET rpc
func JobRPCService(e echo.Context) error{
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "jobs RPCService"})
}

// JobLogClear Delete 日志
func JobLogClear(e echo.Context) error{
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "jobs log Clean"})
}

// JobEvents GET POST
func JobEvents(e echo.Context) error{
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": "jobs events"})
}
