package v1

import (
	"github.com/gin-gonic/gin"
)

/*
tag.go will implement CRUD about tags
*/

func GetTags(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "send by router"})
}
func AddTag(c *gin.Context) {

}
func EditTag(c *gin.Context) {

}
func DeleteTag(c *gin.Context) {

}
