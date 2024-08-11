// É o pacote responsável por reunir as funcionalidades necessárias para execução
// do servidor e das demais serviços que forem necessários.
package start

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/eujandergois/mc-control/pkg/bukkit"
	"github.com/eujandergois/mc-control/pkg/config"
	webserver "github.com/eujandergois/mc-control/pkg/web"
)

// openBrowser é responsável por abrir o navegador padrão na url fornecida
// como parametro
//
// Deprecated: Em breve esta função será substituída por interface interna.
func openBrowser(url string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		fmt.Printf("Unsupported platform: %s\n", runtime.GOOS)
		return
	}

	err := cmd.Start()
	if err != nil {
		fmt.Printf("Failed to open browser: %v\n", err)
	}
}

// waitForServer tenta uma conexão na url fornecida como parametro
// repetidamente até obter sucesso e quebrar o loop.
func waitForServer(url string) {
	for {
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == http.StatusOK {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
}

// StartApplication é responsavel por implementar os métodos necessários para
// iniciar o servidor spigot e os demais serviços necessários para o funcionamento.
func StartApplication() {

	bukkitSettings, err := bukkit.GetBukkitSettings()
	if err != nil {
		fmt.Println("Erro ao carregar 'bukkit.yml':", err)
		os.Exit(1)
	}

	fmt.Println("\n\n\n --->", bukkitSettings.Settings.ShutdownMessage)

	// Carregar propriedades do arquivo server.properties
	properties, err := config.LoadProperties("server.properties")
	if err != nil {
		fmt.Println("Erro ao carregar server.properties:", err)
		os.Exit(1)
	}

	// Exemplo de como acessar uma propriedade
	serverName := properties["server-name"]
	if serverName == "" {
		serverName = "Magma Wave" // Valor padrão se não estiver definido em server.properties
	}

	// Canal para sinalizar quando o servidor web estiver pronto
	ready := make(chan bool)

	// Inicia o servidor web
	go webserver.StartServer(ready, serverName)

	// Espera até que o servidor web esteja pronto
	<-ready

	// Verifica se a API está disponível antes de abrir o navegador
	waitForServer("http://localhost:3333/api/properties")

	// Abre o navegador
	openBrowser("http://localhost:3333/")

	// Executa o comando para iniciar o servidor Spigot
	cmd := exec.Command("java", "-Xms3G", "-Xmx3G", "-XX:+UseG1GC", "-jar", "spigot.jar", "--nogui")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor Spigot:", err)
		os.Exit(1)
	}

	// Aguardar a conclusão do servidor Spigot
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Erro durante a execução do servidor Spigot:", err)
		os.Exit(1)
	}

	fmt.Println("Servidor Spigot encerrado.")
}
