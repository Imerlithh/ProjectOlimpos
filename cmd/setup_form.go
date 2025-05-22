package cmd

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"ProjectOlimpos/config"
	"gopkg.in/yaml.v3"
)

func setupFirstRunForm() {
	http.HandleFunc("/first-run", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tmpl := template.Must(template.ParseFiles("web/templates/first_run.html"))
			tmpl.Execute(w, nil)
			return
		}

		port, err := strconv.Atoi(r.FormValue("db_port"))
		if err != nil {
			http.Error(w, "Port sayısal olmalıdır", http.StatusBadRequest)
			return
		}

		var cfg config.Config
		cfg.Database.Type = r.FormValue("db_type")
		cfg.Database.Name = r.FormValue("db_name")
		cfg.Database.User = r.FormValue("db_user")
		cfg.Database.Password = r.FormValue("db_pass")
		cfg.Database.Host = r.FormValue("db_host")
		cfg.Database.Port = port
		cfg.Database.Extra = r.FormValue("db_extra")

		os.MkdirAll("config", 0700)
		data, err := yaml.Marshal(&cfg)
		if err != nil {
			http.Error(w, "YAML oluşturulamadı: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if err := os.WriteFile(filepath.Join("config", "config.yaml"), data, 0600); err != nil {
			http.Error(w, "config.yaml yazılamadı: "+err.Error(), http.StatusInternalServerError)
			return
		}

		go func() {
			startApplication()
		}()
		w.Write([]byte("Kurulum tamamlandı. Uygulama başlatılıyor..."))
	})

	fmt.Println("Kurulum için http://localhost:8080/first-run adresine gidin.")
	http.ListenAndServe(":8080", nil)
}
