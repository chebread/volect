package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "modernc.org/sqlite"
)

type ResultData struct {
	Name string
}

func main() {
	adminId := os.Getenv("VOLUNTEER_ADMIN_ID")
	adminPw := os.Getenv("VOLUNTEER_ADMIN_PW")

	if adminId == "" || adminPw == "" {
		log.Fatal("오류: VOLUNTEER_ADMIN_ID 또는 VOLUNTEER_ADMIN_PW 환경변수가 설정되지 않았습니다")
	}

	db, err := sql.Open("sqlite", "./reviews.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `CREATE TABLE IF NOT EXISTS reviews (id INTEGER PRIMARY KEY, name TEXT, content TEXT);`
	db.Exec(sqlStmt)

	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			name := r.FormValue("name")
			content := r.FormValue("content")

			_, err := db.Exec("INSERT INTO reviews(name, content) VALUES(?, ?)", name, content)
			if err != nil {
				http.Error(w, "저장 실패", 500)
				return
			}

			tmpl, err := template.ParseFiles("result.html")
			if err != nil {
				http.Error(w, "템플릿 에러", 500)
				return
			}
			tmpl.Execute(w, ResultData{Name: name})
		}
	})

	http.HandleFunc("/admin/download", func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || user != adminId || pass != adminPw {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "접근 권한이 없습니다", 401)
			return
		}

		rows, err := db.Query("SELECT id, name, content FROM reviews")
		if err != nil {
			http.Error(w, "DB 조회 실패", 500)
			return
		}
		defer rows.Close()

		w.Header().Set("Content-Type", "text/csv; charset=utf-8")
		w.Header().Set("Content-Disposition", "attachment; filename=reviews.csv")

		w.Write([]byte{0xEF, 0xBB, 0xBF})
		writer := csv.NewWriter(w)
		writer.Write([]string{"번호", "이름", "소감내용"})

		for rows.Next() {
			var id int
			var name, content string
			rows.Scan(&id, &name, &content)
			writer.Write([]string{fmt.Sprintf("%d", id), name, content})
		}
		writer.Flush()
	})

	log.Println("서버 시작: 80번 포트")
	log.Fatal(http.ListenAndServe(":80", nil))
}
