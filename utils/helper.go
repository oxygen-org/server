/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-15 21:47:55
 * @LastEditTime: 2019-08-15 21:47:55
 * @LastEditors: your name
 */
package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SuccessRespond x
func SuccessRespond(c *gin.Context, msg gin.H) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": msg})
}

// FailedRespond x
func FailedRespond(c *gin.Context, msg gin.H) {
	c.JSON(http.StatusOK, gin.H{"code": 1, "error": msg})
}
