//package gingetemp
//
//import(
//	"github.com/gin-gonic/gin"
//	"k2io.com/keshav/config"
//	"net/http"
//)
//
//
//func Getemps(c *gin.Context) {
//	emps := config.GetEmps()
//
//	if emps == nil || len(emps) == 0 {
//		c.AbortWithStatus(http.StatusNotFound)
//	} else {
//		c.IndentedJSON(http.StatusOK, emps)
//	}
//}