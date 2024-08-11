package webserver

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/eujandergois/mc-control/pkg/config"
)

type Server struct {
	Name         string
	Gamemode     string
	CommandBlock string
	MOTD         string
	Port         string
}

func GetServerProperties() Server {
	properties, err := config.LoadProperties("server.properties")
	if err != nil {
		fmt.Println("Erro ao carregar server.properties:", err)
		os.Exit(1)
	}

	serverName := properties["server-name"]
	if serverName == "" {
		serverName = "Magma Wave" // Valor padrão se não estiver definido em server.properties
	}

	serverGamemode := properties["gamemode"]
	serverCommandBlock := properties["enable-command-block"]
	serverMOTD := properties["motd"]
	serverPort := properties["query.port"]

	serverProperties := Server{
		Name:         serverName,
		Gamemode:     serverGamemode,
		CommandBlock: serverCommandBlock,
		MOTD:         serverMOTD,
		Port:         serverPort,
	}

	return serverProperties
}

func GetTemplates() (*template.Template, error) {
	tmpl, err := template.ParseFiles(
		filepath.Join("services", "cmd", "web", "templates", "index.html"),
		filepath.Join("services", "cmd", "web", "templates", "base.html"),
	)
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

func StartServer(ready chan<- bool, serverName string) {
	serverProperties := GetServerProperties()

	tmpl, err := GetTemplates()
	if err != nil {
		fmt.Println("Erro ao carregar templates:", err)
		close(ready)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "base.html", serverProperties)
		if err != nil {
			http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
			fmt.Println("Erro ao renderizar template:", err)
		}
	})

	http.HandleFunc("/api/properties", func(w http.ResponseWriter, r *http.Request) {
		serverProperties := GetServerProperties()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(serverProperties)
	})

	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			serverProperties := Server{
				Name:         r.FormValue("serverName"),
				Gamemode:     r.FormValue("gamemode"),
				CommandBlock: r.FormValue("commandBlock"),
				MOTD:         r.FormValue("motd"),
				Port:         r.FormValue("port"),
			}

			err := config.SaveProperties("server.properties", map[string]string{
				"server-name":          serverProperties.Name,
				"gamemode":             serverProperties.Gamemode,
				"enable-command-block": serverProperties.CommandBlock,
				"motd":                 serverProperties.MOTD,
				"query.port":           serverProperties.Port,
			})
			if err != nil {
				http.Error(w, "Erro ao salvar propriedades", http.StatusInternalServerError)
				fmt.Println("Erro ao salvar propriedades:", err)
				return
			}

			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join("services", "cmd", "web", "static")))))

	go func() {
		ready <- true
	}()

	err = http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor web:", err)
		close(ready)
	}
}
