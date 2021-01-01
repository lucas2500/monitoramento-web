package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	// "io/ioutil"
	// reflect
)

const delay = 5

func main() {

	// O operador "_" indica que não há interesse em alguma variável retornada por uma função com múltiplos retornos
	// _, idade := funcaoComDoisRetornos()
	exibeMenu()
	comando := lerComando()

	// fmt.Println("O tipo da variável nome é: ", reflect.TypeOf(versao))

	switch comando {
	case 1:
		// O compilador não obriga o uso do "break", porém não vai reclamar se for colocado

		iniciarMonitoramento()

	case 2:
		fmt.Println("Exibindo logs...")
		imprimeLogs()
	case 0:
		fmt.Println("Saindo do programa...")

		// Encerra o programa com sucesso
		os.Exit(0)
	default:
		fmt.Println("Comando não reconhecido!!")

		// Encerra o programa com erro
		os.Exit(-1)
	}
}

func exibeMenu() {

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

}

func lerComando() int {

	comandoLido := 0

	// fmt.Scanf("%d", &comando)
	fmt.Scan(&comandoLido)
	fmt.Println("O comando inserido foi:", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {

	fmt.Println("Monitorando...")

	sites := lerSitesDoArquivo()

	for i, sites := range sites {

		testaSite(i, sites)

		time.Sleep(delay * time.Second)
	}
}

func testaSite(index int, site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro ao requisitar o site", site, "!!", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Requisição ao site", index, "-", site, "OK!!")
		registraLog(site, true)
	} else {
		fmt.Println("Falha na requisição ao site", index, "-", site)
		registraLog(site, false)
	}
}

func lerSitesDoArquivo() []string {

	var sites []string

	// arquivo, err := ioutil.ReadFile("sites.txt")
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Erro na abertura do arquivo!!", err)
		os.Exit(-1)
	}

	reader := bufio.NewReader(arquivo)

	for {

		linha, err := reader.ReadString('\n')

		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()

	return sites

}

func registraLog(site string, status bool) {

	// timeformat := time.Now().Format("20060102") + "_" + time.Now().Format("1504")
	// arquivo, err := os.OpenFile("log_"+timeformat+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Houve um erro ao abrir o arquivo!!", err)
		os.Exit(-1)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")
	comandoLido := 0

	if err != nil {
		fmt.Println("Houve um erro ao abrir o arquivo de logs!!", err)
	}

	fmt.Println(string(arquivo))

	fmt.Println("0 - Encerrar programa")
	fmt.Println("1 - Voltar ao menu anterior")

	fmt.Scan(&comandoLido)

	if comandoLido == 0 {
		os.Exit(0)
	} else {
		main()
	}

}
