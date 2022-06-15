package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"time"

	. "github.com/ehsanhajian/k8s-pingpong/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config AppConfig
}

func main() {
	var c AppConfig

	config := c.LoadConf()

	ticker := time.NewTicker(time.Duration(config.PingInterval) * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:

				for _, server := range c.Servers {

					pingOut := ping(server.IP)
					fmt.Println(pingOut)
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	serverPort := fmt.Sprintf(":%s", c.ServerPort)

	r.Run(serverPort)
}

func ping(ip string) string {
	out, _ := exec.Command("ping", ip, "-c 5", "-i 3", "-w 10").Output()
	return string(out)
	// if strings.Contains(string(out), "Destination Host Unreachable") {
	// 	fmt.Println("TANGO DOWN")
	// } else {
	// 	fmt.Println("IT'S ALIVEEE")
	// }

}
