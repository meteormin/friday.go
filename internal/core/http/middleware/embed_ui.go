package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meteormin/friday.go/internal/core"
	"github.com/meteormin/friday.go/ui"
	"io/fs"
	"mime"
	"path/filepath"
)

func EmbedUI(router fiber.Router) {
	router.Use(func(ctx *fiber.Ctx) error {
		path := ctx.Path()
		if len(path) > 0 && path[0] == '/' {
			path = path[1:]
		}

		// 기본 경로 ("/" 또는 빈 문자열)는 index.html로 매핑
		if path == "" {
			path = "index.html"
		}

		// embed FS에서 요청 파일 읽기 시도
		data, err := fs.ReadFile(ui.FS, path)
		if err != nil {
			// 요청 파일이 없으면 SPA 라우팅을 위해 index.html을 읽음
			data, err = fs.ReadFile(ui.FS, "index.html")
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}
			path = "index.html"
		}

		ext := filepath.Ext(path)
		contentType := mime.TypeByExtension(ext)
		if contentType == "" {
			contentType = "text/plain"
		}

		ctx.Set("Content-Type", contentType)
		core.Logger().Debugf("Serving file: %s, %s", path, contentType)
		return ctx.Send(data)
	})
}
