package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/abadojack/whatlanggo"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
)

var notesDir = "notes"

func main() {
	os.MkdirAll(notesDir, os.ModePerm)

	r := gin.Default()

	r.POST("/check-grammar", checkGrammar)
	r.POST("/notes", saveNote)
	r.GET("/notes", listNotes)
	r.GET("/notes/:id/html", renderNoteHTML)

	r.Run(":8080")
}

type NoteRequest struct {
	Text string `json:"text"`
}

func checkGrammar(c *gin.Context) {
	var req NoteRequest
	if err := c.BindJSON(&req); err != nil {
		return
	}
	info := whatlanggo.Detect(req.Text)

	c.JSON(200, gin.H{
		"language": info.Lang.String(),
		"is_reliable": info.IsReliable(),
		"note": "⚠ Grammar check here is not full grammar correction — basic language detection only.",
	})
}

func saveNote(c *gin.Context) {
	var req NoteRequest
	if err := c.BindJSON(&req); err != nil {
		return
	}

	id := len(listFiles()) + 1
	filename := filepath.Join(notesDir, fmt.Sprintf("%d.md", id))
	err := os.WriteFile(filename, []byte(req.Text), 0644)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to save note"})
		return
	}

	c.JSON(201, gin.H{"message": "note saved", "id": id})
}

func listNotes(c *gin.Context) {
	files := listFiles()
	c.JSON(200, gin.H{"notes": files})
}

func renderNoteHTML(c *gin.Context) {
	id := c.Param("id")
	filename := filepath.Join(notesDir, id+".md")

	data, err := os.ReadFile(filename)
	if err != nil {
		c.JSON(404, gin.H{"error": "note not found"})
		return
	}

	html := markdown.ToHTML(data, nil, nil)
	c.Data(200, "text/html; charset=utf-8", html)
}

func listFiles() []string {
	files, _ := os.ReadDir(notesDir)
	var list []string

	for _, f := range files {
		if !f.IsDir() {
			id := f.Name()
			list = append(list, id[:len(id)-3])
		}
	}
	return list
}
