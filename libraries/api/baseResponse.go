package api

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
)

var statusCode = 200

// SetStatusCode -- Set status code
func SetStatusCode(statCode int) int {
	statusCode := statCode
	return statusCode
}

// ResponseJSON -- set response to json format
func ResponseJSON(c *gin.Context, data interface{}) {
	c.JSON(statusCode, data)
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
	statusCode := SetStatusCode(201)
	c.JSON(statusCode, gin.H{"message": message})
	return
}

// ResponseUpdated -- Set response for update process
func ResponseUpdated(c *gin.Context, message string) {
	if message == "" {
		message = "Resource Updated"
	}
	statusCode := SetStatusCode(200)
	c.JSON(statusCode, gin.H{"message": message})
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
	statusCode := SetStatusCode(200)
	c.JSON(statusCode, gin.H{"message": message})
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
	statusCode := SetStatusCode(401)
	c.JSON(statusCode, gin.H{"errors": message})
	return
}

// ResponseNotFound -- Set response not found
func ResponseNotFound(c *gin.Context, message string) {
	if message == "" {
		message = "Resource Not Found"
	}
	statusCode := SetStatusCode(404)
	c.JSON(statusCode, gin.H{"errors": message})
	return
}

// ResponseMethodNotAllowed -- Set response method not allowed
func ResponseMethodNotAllowed(c *gin.Context, message string) {
	if message == "" {
		message = "Method Not Allowed"
	}
	statusCode := SetStatusCode(405)
	c.JSON(statusCode, gin.H{"errors": message})
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
	c.Data(statusCode, contentType, b.Bytes())
	return
}

func ResponseXML(c *gin.Context, statusCode int, data interface{}) {
	statusCode = SetStatusCode(statusCode)
	c.XML(statusCode, data)
	return
}
