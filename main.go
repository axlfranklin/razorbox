package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gookit/color"
	"github.com/gosuri/uiprogress"
)

var (
	serviceMap = map[int]string{
		21:    "FTP",
		22:    "SSH",
		23:    "Telnet",
		25:    "SMTP",
		53:    "DNS",
		80:    "HTTP",
		110:   "POP3",
		135:   "RPC",
		139:   "NetBios SMB",
		443:   "HTTPS",
		445:   "SMB",
		3306:  "MySQL",
		3389:  "Remote Desktop",
		5432:  "PostgreSQL",
		6379:  "Redis",
		8080:  "HTTP Alt",
		27017: "MongoDB",

		// Portas no banco de dados do código
	}
)

func scanPort(target string, port int, timeout time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	addr := fmt.Sprintf("%s:%d", target, port)
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return // Aqui retorna a porta fechada para informar
	}
	defer conn.Close()

	service := serviceMap[port]
	if service == "" {
		service = "Serviço Desconhecido"
	}

	banner, _ := getBanner(conn)
	if banner != "" {
		color.Green.Printf("Porta %d (%s) está aberta - %s\n", port, service, banner)
	} else {
		color.Green.Printf("Porta %d (%s) está aberta\n", port, service)
	}
}

func getBanner(conn net.Conn) (string, error) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	reader := bufio.NewReader(conn)
	banner, err := reader.ReadString('\n')
	return strings.TrimSpace(banner), err
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite o host alvo (exemplo: host.com): ")
	scanner.Scan()
	target := scanner.Text()

	fmt.Print("Digite o range de portas (exemplo: 1-1000): ")
	scanner.Scan()
	portRange := scanner.Text()

	// Parse port range
	ports := strings.Split(portRange, "-")
	startPort, err := strconv.Atoi(ports[0])
	if err != nil {
		fmt.Println("Porta inicial inválida")
		return
	}
	endPort, err := strconv.Atoi(ports[1])
	if err != nil {
		fmt.Println("Porta final inválida.")
		return
	}

	fmt.Print("Escolha o intervalo (seconds): ")
	scanner.Scan()
	timeoutSeconds, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Intervalo inválido")
		return
	}
	timeout := time.Duration(timeoutSeconds) * time.Second

	var wg sync.WaitGroup
	fmt.Printf("Iniciando scan de portas no alvo: %s...\n", target)

	uiprogress.Start()
	bar := uiprogress.AddBar(endPort - startPort + 1).AppendCompleted().PrependElapsed()

	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			scanPort(target, port, timeout, &wg)
			bar.Incr()
		}(port)
	}

	wg.Wait()
	uiprogress.Stop()
	fmt.Println("Finalizado!")
}
