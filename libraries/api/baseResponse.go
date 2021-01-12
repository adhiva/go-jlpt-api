package api

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseJSON -- set response to json format
func ResponseJSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
	return
}

func ResponseJSONWithStatusCode(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
	return
}

// ResponseCreated -- Set response for create process
func ResponseCreated(c *gin.Context, message string) {
	if message == "" {
		message = "Resource Created"
	}
	c.JSON(http.StatusCreated, gin.H{"message": message})
	return
}

// ResponseUpdated -- Set response for update process
func ResponseUpdated(c *gin.Context, message string) {
	if message == "" {
		message = "Resource Updated"
	}
	c.JSON(http.StatusAccepted, gin.H{"message": message})
	return
}

func ResponseSuccessNoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
	return
}

func ResponseText(c *gin.Context, statusCode int, message string) {
	if message == "" {
		message = "Success"
	}

	c.String(statusCode, message)
	return
}

// ResponseDeleted -- Set response for delete process
func ResponseDeleted(c *gin.Context, message string) {
	if message == "" {
		message = "Resource Deleted"
	}
	c.JSON(http.StatusAccepted, gin.H{"message": message})
	return
}

// ResponseError -- Set response for error
func ResponseError(c *gin.Context, message interface{}, statusCode int) {
	c.JSON(statusCode, gin.H{"errors": message})
	return
}

// ResponseFailValidation -- Set response for fail validation
func ResponseFailValidation(c *gin.Context, message interface{}) {
	ResponseError(c, message, 422)
	return
}

// ResponseUnauthorized -- Set response not authorized
func ResponseUnauthorized(c *gin.Context, message string) {
	if message == "" {
		message = "Unauthorized"
	}
	c.JSON(http.StatusUnauthorized, gin.H{"errors": message})
	return
}

// ResponseNotFound -- Set response not found
func ResponseNotFound(c *gin.Context, message string) {
	if message == "" {
		message = "Resource Not Found"
	}
	c.JSON(http.StatusNotFound, gin.H{"errors": message})
	return
}

// ResponseMethodNotAllowed -- Set response method not allowed
func ResponseMethodNotAllowed(c *gin.Context, message string) {
	if message == "" {
		message = "Method Not Allowed"
	}

	c.JSON(http.StatusMethodNotAllowed, gin.H{"errors": message})
	return
}

func ResponseRedirect(c *gin.Context, url string) {
	if url == "" {
		return
	}

	c.Redirect(http.StatusFound, url)
	return
}

func ResponseAttachment(c *gin.Context, contentType, filename string, b *bytes.Buffer) {
	c.Header("Content-Description", "Attachment")
	c.Header("Content-Disposition", "attachment;filename="+filename)
	c.Data(http.StatusOK, contentType, b.Bytes())
	return
}

func ResponseXML(c *gin.Context, statusCode int, data interface{}) {
	c.XML(statusCode, data)
	return
}
