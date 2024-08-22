package main

import (
    "bufio"
    "context"
    "fmt"
    "log"
	"time"
	"strings"
    "os"
    "github.com/chromedp/chromedp"
	"github.com/gen2brain/beeep"
)

func main() {
    // Solicitar credenciais ao usuário
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
        chromedp.Flag("headless", false), // Não rodar em modo headless
		chromedp.Flag("force-device-scale-factor", "1"), // Definir escala de dispositivo
		chromedp.Flag("force-device-scale-factor", "1"), // Força a escala do dispositivo
		chromedp.Flag("start-maximized", true), // Iniciar o navegador maximizado

	)
    reader := bufio.NewReader(os.Stdin)
    
    //fmt.Print("Digite o nome de usuário: ")
    //username, _ := reader.ReadString('\n')
    //username = username[:len(username)-1] // Remove o caractere de nova linha
    
    //fmt.Print("Digite a senha: ")
    //password, _ := reader.ReadString('\n')
    //password = password[:len(password)-1] // Remove o caractere de nova linha

	// Solicitar armazém ao usuário
    fmt.Print("Digite o armazém onde você está (WOLFF ou LYOR): ")
    armazem, _ := reader.ReadString('\n')
    armazem = armazem[:len(armazem)-1] // Remove o caractere de nova linha
    armazem = strings.TrimSpace(strings.ToUpper(armazem)) // Normalizar a entrada

	// Solicitar número de carga e notas fiscais
	fmt.Print("Digite o número da carga: ")
	carga, _ := reader.ReadString('\n')
	carga = carga[:len(carga)-1] // Remove o caractere de nova linha

	// Solicitar número das notas fiscais separadas por vírgula
	fmt.Print("Digite os números das notas fiscais (separados por vírgula): ")
	notas, _ := reader.ReadString('\n')
	notas = notas[:len(notas)-1] // Remove o caractere de nova linha
	notasArray := strings.Split(notas, ",") // Divide os números das notas fiscais em um array


	// Inicializar o contexto do chromedp com as opções configuradas
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()


    // Executar as tarefas de automação
    err := chromedp.Run(ctx,
        chromedp.Navigate("https://rojemachml.seniorcloud.com.br/siltwms"),
        chromedp.WaitVisible(`#LoginDialog_loginText`), // Ajuste os seletores conforme a página
        chromedp.SendKeys(`#LoginDialog_loginText`, `kennedyy`),
        chromedp.SendKeys(`#LoginDialog_passwordText`, `1234`),
		chromedp.Click(`#LoginDialog_armazemComboBox`),
		chromedp.WaitVisible(`#LoginDialog_armazemComboBox-LYOR`), // Ajuste os seletores conforme a página
		chromedp.Sleep(1 * time.Second), // Aguarda a abertura do combobox

        // Selecionar o armazém informado
        chromedp.ActionFunc(func(ctx context.Context) error {
            fmt.Printf("Selecionando armazém: %s\n", armazem) // Debug: exibe o armazém selecionado
            switch armazem {
            case "WOLFF":
                if err := chromedp.Click(`.x-combo-list-item:nth-of-type(2)`).Do(ctx); err != nil { // O segundo item
                    return err
                }
            case "LYOR":
                if err := chromedp.Click(`.x-combo-list-item:nth-of-type(1)`).Do(ctx); err != nil { // O primeiro item
                    return err
                }
            default:
                return fmt.Errorf("armazém não reconhecido: %s", armazem)
            }
            return nil
        }),
        chromedp.Click(`#LoginDialog_loginButton`),
        chromedp.Sleep(2 * time.Second),
		chromedp.Click(`.sample-box:nth-of-type(9)`),

		// Selecionar o campo de filtro e escrever o número de carga
		chromedp.WaitVisible(`#filter-COLETA`), // Ajuste o seletor CSS para o campo de filtro
		chromedp.SendKeys(`#filter-COLETA`, carga+string('\r')), // Substitua pelo seletor do campo de filtro

        chromedp.ActionFunc(func(ctx context.Context) error {
            for i := 0; i < 39; i++ {
                if err := chromedp.KeyEvent("\t").Do(ctx); err != nil {
                    return err
                }
                // Espera um curto intervalo entre os TABs para garantir a movimentação do foco
                time.Sleep(100 * time.Millisecond)
            }
            return nil
        }),

		chromedp.Click(`#tb-Embarque-AutorizarNotaFiscal`),

		chromedp.Sleep(8 * time.Second),

		chromedp.WaitVisible(`#SiltTransfere_buscarComboBoxComboArrow`),
		chromedp.Click(`#SiltTransfere_buscarComboBoxComboArrow`),
		chromedp.WaitVisible(`#SiltTransfere_buscarComboBox-NOTAFISCAL`),
		chromedp.Click(`#SiltTransfere_buscarComboBox-NOTAFISCAL`),

		// Iterar sobre cada nota fiscal, buscar, pressionar Enter e clicar no botão
		chromedp.ActionFunc(func(ctx context.Context) error {
			for _, nota := range notasArray {
				// Substitua pelo seletor do campo de busca de nota fiscal
				err := chromedp.Run(ctx,
					// Limpar o campo de busca
			        chromedp.Click(`#SiltTransfere_buscarText`),        
					// Executar JavaScript para selecionar e apagar o texto
					chromedp.Evaluate(`
					var field = document.querySelector('#SiltTransfere_buscarText');
					field.focus();
					field.select(); // Seleciona todo o texto
					field.value = ''; // Limpa o texto
				`, nil),
					chromedp.SendKeys(`#SiltTransfere_buscarText`, strings.TrimSpace(nota)),
					chromedp.KeyEvent("\r"), // Pressiona Enter após digitar a nota
					chromedp.Sleep(2 * time.Second), // Aguarda 2 segundo
					chromedp.Click(`.x-grid3-cell-inner:first-child`), // Substitua pelo seletor do botão desejado
					chromedp.Sleep(2 * time.Second), // Aguarda 2 segundos antes de continuar
				)
				if err != nil {
					log.Printf("Erro ao processar a nota %s: %v", nota, err)
				}
			}
			return nil
		}),

    )
    if err != nil {
        log.Fatalf("Erro ao executar tarefas de automação: %v", err)
    }

    fmt.Println("Login realizado com sucesso!")

	// Adicionar carga e notas fiscais (aqui você pode personalizar a adição com base na página)
	fmt.Printf("Número da carga: %s\n", carga)
	
	// Enviar notificação para o usuário
	err = beeep.Alert("Concluído", "A carga solicitada foi finalizada com sucesso!", "")
	if err != nil {
		log.Fatalf("Erro ao enviar notificação: %v", err)
	}

	// Manter o navegador aberto por 10 minutos para visualização
	fmt.Println("Aguardando 10 minutos para visualização do navegador...")
	time.Sleep(10 * time.Minute)
}