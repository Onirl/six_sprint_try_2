package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func ReturnHTML(w http.ResponseWriter, r *http.Request) {
	html := "./index.html"
	if _, err := os.Stat(html); os.IsNotExist(err) {
		http.Error(w, "Файл не найден", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, html)
}

func UploadHTML(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		file, headerFile, err := r.FormFile("myFile") //Получаем файл, метаданные
		if err != nil {
			http.Error(w, "Ошибка получения файла!", http.StatusInternalServerError)
			return
		}
		defer file.Close() //Закрываем файл по завершению
		fileBytes, err := io.ReadAll(file)
		fileContent := string(fileBytes)
		converString := service.ConvertFileData(fileContent)
		pathFile := "../" + time.Now().Format("02_01_2006_15_04_05") + filepath.Ext(headerFile.Filename)
		fileRes, err := os.OpenFile(pathFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
		if err != nil {
			http.Error(w, "Ошибка создания файла!", http.StatusInternalServerError)
			return
		}
		defer fileRes.Close()
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		fileRes.WriteString(converString)
	}
}
