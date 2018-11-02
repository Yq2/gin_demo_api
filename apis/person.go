package apis

import (
	"fmt"
	. "github.com/Yq2/gin_demo_api/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK,
		"Gin web 演示demo\n\n API列表 \n GET /persons 获取用户列表 \n POST /person  提交用户form {first_name:'YANG',last_name:'QIANG'} \n GET /person/:id 获取指定用户 \n PUT /person/:id 更新用户资料 \n DELETE /person/:id 删除指定用户 \n...\n 作者:github.com/Yq2")
}

func AddPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")
	p := Person{FirstName: firstName, LastName: lastName}
	ra, err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("成功添加用户 id. %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func GetPersonsApi(c *gin.Context) {
	var p Person
	persons, err := p.GetPersons()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"persons": persons,
	})
}

func GetPersonApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	p := Person{Id: id}
	person, err := p.GetPerson()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"person": person,
	})

}

func ModPersonApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	p := Person{Id: id}
	err = c.Bind(&p)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err := p.ModPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("更新用户#.%d资料成功 %d", p.Id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func DelPersonApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	p := Person{Id: id}
	ra, err := p.DelPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("删除用户#.%d #.%d", id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
