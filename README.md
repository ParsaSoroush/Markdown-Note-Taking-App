# 📘 Markdown Note-Taking App

A simple RESTful API that allows users to upload/save markdown notes, check grammar, list notes, and render them as HTML.

🔗 **Project page:** https://roadmap.sh/projects/markdown-note-taking-app

---

## 🚀 Features
- Check grammar of a note  
- Save a note written in Markdown  
- List all saved notes  
- Render a Markdown note as HTML

---

## 📦 Installation & Setup

### 1️⃣ Clone the repository
```bash
git clone https://github.com/ParsaSoroush/Markdown-Note-Taking-App.git
cd markdown-notes
```

### 2️⃣ Install dependencies
```bash
go mod init markdown-notes
go get github.com/gin-gonic/gin
go get github.com/abadojack/whatlanggo
go get github.com/gomarkdown/markdown
```

### 3️⃣ Run the server
```bash
go run main.go
```

- Your API will available at:
```bash
http://localhost:8080
```

## 🧪 API Endpoints
| Method | Endpoint          | Description                     |
| ------ | ----------------- | ------------------------------- |
| POST   | `/check-grammar`  | Check the grammar of the note   |
| POST   | `/notes`          | Save a note written in Markdown |
| GET    | `/notes`          | List all saved notes            |
| GET    | `/notes/:id/html` | Render the note as HTML         |
---