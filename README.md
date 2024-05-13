# Aplicação Go com deploy na Google Cloud Run

## :notebook_with_decorative_cover: Sobre o Projeto

Aplicação desenvolvida no curso **Pós Graduação Go Expert - Full Cycle** na linguagem Go.
A aplicação recebe uma requisição de CEP, que após ser validado consulta uma API externa para coletar informações de endereço.
Com a informação da cidade referente ao CEP, é realizado uma consulta em outra API externa para coletar informações sobre o clima.
E por fim é retornado as temperaturas da localidade em Célsius, Fahrenheit e Kelvin.

## :wrench: Requisitos

### :sparkles: Funcionalidades
- Criação de ordens de compra com id, preço, taxa e preço final.
- Lista de ordens criadas
- Manipulação de eventos

### :computer: Tecnologias Aplicadas
* Go
* Net/Http
* Docker
* Google Cloud Run

## :arrow_forward: Executando a aplicação

Para realizar requisições na aplicação acesse: `https://climate-cep-go-3g7xlqkloq-uc.a.run.app/<cep>`
Exemplos:
- `https://climate-cep-go-3g7xlqkloq-uc.a.run.app/01001000`

Para executar a aplicação siga as instruções abaixo.

### Pré-requisitos

Primeiramente é necessário que possua instalado as seguintes ferramentas: Go, Git, Docker e Maker.
Além disto é bom ter um editor para trabalhar com o código como VSCode.
Cadastro na [API WeatherAPI](https://www.weatherapi.com/)

### Instalação

1. Faça uma cópia do repositório (HTTPS ou SSH)
   ```sh
   git clone https://github.com/flpnascto/clean-architecture-go
   ```
   ```sh
   git clone git@github.com:flpnascto/clean-architecture-go.git
   ```
2. Acessar a pasta do repositório local configurar variáveis de ambiente
   - No arquivo `./.env_example` adicionar a API KEY de WeatherAPI para **WEATHER_API_KEY**
   - Alterar o nome deste arquivo de `.env_example` para `.env`
3. Instanciar a aplicação em um container do docker
   ```sh
   make server
   ```
4. Executar exemplos de requisições para testar a aplicação
   ```sh
   make requests
   ```
5. Realizar requisições ao servidor com o comando `curl http://localhost:8080/<cep>`

