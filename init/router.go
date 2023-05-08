package initialize

import (
	"github.com/gin-gonic/gin"
	"soleaf.xyz/yaowen/global"
	"soleaf.xyz/yaowen/handler"
	"soleaf.xyz/yaowen/model"
	"strconv"
)

func Routers()*gin.Engine{
	r := gin.Default()

	r.GET("/api/riskyTroves",riskyTroves)

	return r
}

//	api handler
func riskyTroves(c *gin.Context){
	pn := c.DefaultQuery("start","0")
	pSize := c.DefaultQuery("count","5")

	pnInt,_ := strconv.Atoi(pn)
	pSizeInt,_ := strconv.Atoi(pSize)

	lists := handler.GetList(pnInt, pSizeInt)

	result := make([]interface{},0,len(lists))

	for _,list := range lists{
		result = append(result,list)
	}

	c.JSON(200,gin.H{
		"data":result,
	})
}

var (
	addressChan = make(chan string,global.ADDRCHANNUM)
	itemChan = make(chan model.Data,global.WORKERNUM)
)

// 不断地根据addr去爬地址，把目标地址放到 addressChan
func getAddress(addr string)string{

	// 开个协程去做
	go func() {

	}()

	return "ok"
}

// 根据地址拿到健康值
func getHealth(addr string)int{
	return  0
}

// 生产者，消费者模型
func work(){

	for i := 0;i<global.ADDRCHANNUM;i++{
		go func() {
			for {
				select {
					case addr := <- addressChan:
						health := getHealth(addr)
						itemChan <- model.Data{
							Address: addr,
							Health:  int32(health),
						}
				}
			}
		}()
	}

	for i := 0;i<global.WORKERNUM;i++{
		go func() {
			for{
				select {
				case item:= <- itemChan:
					handler.Save(item)
				}
			}
		}()
	}

}

