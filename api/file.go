package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"time"
)

type FileInfo struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255;not null"`
	Path      string `gorm:"size:255;not null"`
	Size      int64  `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// @BasePath /api/v1
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} success
// @Router /upload [post]
func UploadFile(c *gin.Context) {
	//file, err := c.FormFile("file")
	//if err != nil {
	//	c.JSON(500, gin.H{
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//
	//filePath := "/Users/tongkaiqiang/uploadFile/uploadFile" + file.Filename
	//if err := c.SaveUploadedFile(file, filePath); err != nil {
	//	c.JSON(500, gin.H{
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//
	//fileinfo, err := os.Stat(filePath)
	//if err != nil {
	//	c.JSON(500, gin.H{
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//
	//f := &FileInfo{
	//	Name:      file.Filename,
	//	Path:      filePath,
	//	Size:      fileinfo.Size(),
	//	CreatedAt: time.Now(),
	//	UpdatedAt: time.Now(),
	//}
	//
	//if err := db.Create(f).Error; err != nil {
	//	c.JSON(500, gin.H{
	//		"message": err.Error(),
	//	})
	//	return
	//}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

// @BasePath /api/v1
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} success
// @Router /download/:id [get]
func DownloadFile(c *gin.Context) {
	//id := c.Param("id")
	//
	//var f FileInfo
	//
	//if err := db.First(&f, id).Error; err != nil {
	//	c.JSON(500, gin.H{
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//
	//c.Header("Content-Disposition", "attachment; filename="+f.Name)
	//c.Header("Content-Type", "application/octet-stream")
	//c.File(f.Path)

	c.JSON(200, gin.H{
		"message": "success",
	})
}
