# Automação de Navegador com Go e Chromedp

Este projeto implementa uma automação de navegador simples utilizando Go e o pacote `chromedp`. A automação realiza login em um sistema, seleciona um armazém, busca por uma carga específica e associa à ela as notas fiscais informadas pelo usuário. Ao final, uma notificação é enviada ao usuário, e o navegador permanece aberto para visualização.

## Pré-requisitos

Antes de executar este projeto, certifique-se de ter os seguintes requisitos instalados:

- **Go** (versão 1.16 ou superior)
- **Chromedp** (para controle do navegador)
- **Beeep** (para notificações)

## Instalação

Siga os passos abaixo para configurar o projeto:

1. Clone este repositório:
   ```bash
   git clone https://github.com/KennedyAFK/go-browser-automation
   ```

2. Navegue até o diretório do projeto:
   ```bash
   cd go-browser-automation
   ```

3. Instale as dependências do projeto:
   ```bash
   go mod tidy
   ```

## Uso

Para executar o programa, siga as instruções abaixo:

1. No terminal, execute o programa:
   ```bash
   go run main.go
   ```

2. Siga as instruções no terminal:
   - Forneça o nome do armazém.
   - Insira o número da carga.
   - Digite os números das notas fiscais separadas por vírgula.

3. O programa realizará as seguintes ações automaticamente:
   - Login em um sistema.
   - Seleção do armazém.
   - Processamento das notas fiscais.
   
4. Após a execução, uma notificação será exibida, informando que a carga foi finalizada com sucesso. O navegador permanecerá aberto por 10 minutos para visualização.

## Personalização

- **Credenciais de Login**: As credenciais (`teste` e `teste`) estão codificadas no script. Para uso em produção, é recomendável solicitar as credenciais ao usuário ou utilizar variáveis de ambiente.
  
- **URL de Navegação**: A URL padrão para navegação é `where/to?.com`. Altere-a conforme necessário.

- **Seletores CSS**: Certifique-se de que os seletores CSS no código correspondam aos elementos da página que você está automatizando.

## Dependências

O projeto utiliza as seguintes bibliotecas:

- **Chromedp**: Para controle e manipulação do navegador.
- **Beeep**: Para envio de notificações do sistema.

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir [issues](https://github.com/KennedyAFK/go-browser-automation/issues) ou enviar [pull requests](https://github.com/KennedyAFK/go-browser-automation/pulls).

## Licença

Este projeto está licenciado sob a [MIT License](https://opensource.org/licenses/MIT).