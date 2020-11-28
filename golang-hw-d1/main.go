package main

import "C"
import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
)

func main() {
	context.Background()
	router := gin.Default()
	http.HandleFunc("ss", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	router.GET("/role", Get)
	router.GET("/role/:id", GetOne)
	router.POST("/role", Post)
	router.PUT("/role/:id", Put)
	router.DELETE("/role/:id", Delete)
	router.Run(":8080")
}

// 取得全部資料
func Get(c *gin.Context) {
	c.JSON(http.StatusOK, Data)
}

// 取得單筆資料
func GetOne(c *gin.Context) {
	// Validate input Id
	id, err := Search(c.Param("id"))
	if err != nil {
		if id == -1 {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else if id == -2 {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else if id == -3 {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Return role
	} else {
		c.JSON(http.StatusOK, Data[id])
		return
	}
}

// 新增單筆資料
func Post(c *gin.Context) {
	// Convert POST body to Role object
	var newRole Role
	err := c.BindJSON(&newRole)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Find a new role id for the new objects
	var max uint
	for _, role := range Data {
		if role.ID > max {
			max = role.ID
		}
	}
	// Add the new role
	newRole.ID = max + 1
	Data = append(Data, newRole)
	c.JSON(http.StatusOK, newRole)
}

type RoleVM struct {
	ID      uint   `json:"id"`      // Key
	Name    string `json:"name"`    // 角色名稱
	Summary string `json:"summary"` // 介紹
}

// 更新單筆資料
func Put(c *gin.Context) {
	// Validate input Id
	id, err := Search(c.Param("id"))
	if err != nil {
		if id == -1 {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else if id == -2 {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else if id == -3 {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		// Parse the POST body
		var role RoleVM
		err := c.BindJSON(&role)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			Data[id].ID = role.ID
			Data[id].Name = role.Name
			Data[id].Summary = role.Summary
			c.JSON(http.StatusOK, role)
			return
		}
	}
}

// 刪除單筆資料
func Delete(c *gin.Context) {
	id, err := Search(c.Param("id"))
	// Validate input Id
	if err != nil {
		if id == -1 {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else if id == -2 {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else if id == -3 {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Delete role
	} else {
		Data[id] = Data[len(Data)-1]
		// TODO: Verify if GC is required
		// Data[len(Data)-1] = nil
		Data = Data[:len(Data)-1]
		c.JSON(http.StatusOK, gin.H{"status": "success"})
		return
	}
}

// Search a role through its index, return its slice Id
// Positive return value: Slice Id
// Return -1: role not found
// Return -2: strconv.Atoi convert failed
// Return -3: must be an integer greater than 0
// TODO: Convert error id and string to ERROR enum variables
func Search(inputId string) (sliceId int, err error) {
	// Convert input "string" to "integer" and validate the integer
	idInt, err := strconv.Atoi(inputId)
	if err != nil {
		return -2, err
	} else if idInt < 1 {
		return -3, errors.New("must be an integer greater than 0")
	}

	// Convert id from "int" tp "uint" to align with role.ID in Data
	id := uint(idInt)

	// Loop through Data and find the corresponding slideId for the character id
	for i, role := range Data {
		if role.ID == id {
			return i, nil
		}
	}

	// Return error if id not found in Data
	return -1, errors.New("role ID " + strconv.FormatUint(uint64(id), 10) + " not found")
}
