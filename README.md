<p align="center">
    <img width="120" src="http://www.gophergala.com/assets/img/fancy_gopher_renee.jpg">
</p>

# Pagarme-go ![Language Badge](https://img.shields.io/badge/Language-Go-blue.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/isfonzar/pagarme-go)](https://goreportcard.com/report/github.com/isfonzar/pagarme-go) ![License Badge](https://img.shields.io/badge/License-MIT-blue.svg) ![Status Badge](https://img.shields.io/badge/Status-Beta-brightgreen.svg)

pagarme-go é um SDK para a API do [Pagar.me](https://pagar.me) escrito em Go.

Você pode acessar a documentação oficial do Pagar.me acessando esse [link](https://docs.pagar.me).

## Índice

- [Instalação](#instalação)
- [Configuração](#configuração)
  - [Definindo headers customizados](#definindo-headers-customizados)
- [Endpoints faltando](#endpoints-faltando)
- [Contribuições](#contribuições)
  - [Bug Reports & Feature Requests](#bug-reports--feature-requests)
  - [Desenvolvimento](#desenvolvimento)

## Instalação

Utilize o `go get` para obter o SDK e adiciona-lo a seu GOPATH.

`go get github.com/isfonzar/pagarme-go`

Para atualizar o SDK, utilize `go get -u` para obter a última versão.

`go get -u github.com/isfonzar/pagarme-go`

## Configuração

Para instanciar um novo Client, basta fazer o seguinte:

```go
package main

import (
    "net/http"
    
    "github.com/isfonzar/pagarme-go/pkg/client"
)

func main () {
    client := *client.NewClient(http.Client{}, "SUA_CHAVE_DE_API", nil)
}
```

### Definindo headers customizados

1. Se necessário, é possível definir headers http customizados para os requests. Para isso basta informá-los durante a instanciação do objeto `Client`:

```go
package main

import (
    "net/http"
    
    "github.com/isfonzar/pagarme-go/pkg/client"
)

func main () {
	headers := &client.CustomHeaders{
        "X-PagarMe-User-Agent": "Pagar.me Checkout/2.0.0",
        "X-PagarMe-Version": "2017-08-28",
    }
	
    client := *client.NewClient(http.Client{}, "SUA_CHAVE_DE_API", headers)
}
```

E então, você pode poderá utilizar o cliente para fazer requisições ao Pagar.me.

## Endpoints faltando

Alguns endpoints ainda nao foram implementados neste SDK (contribuições são bem-vindas)

É possível usar o metodo **Send** do pacote **client** para fazer requests para os endpoints ainda não implementados.

```go
package main

import (
    "net/http"
    
    "github.com/isfonzar/pagarme-go/pkg/client"
)

func main() {
    client := client.NewClient(http.Client{}, "PAGARME_API_KEY", nil)
    
    params := struct{} // Sua struct de request
    responseObj := struct{} // Struct de resposta
    
    req, err := client.NewRequest("METODO_HTTP", "URL_DO_ENDPOINT", params)
    response := &responseObj{}
    if err != nil {
        return response, nil, err
    }
    errorResp, err := c.Send(req, response)
}
```

## Contribuições

### Bug Reports & Feature Requests

Por favor, use o [issue tracker](https://github.com/isfonzar/pagarme-go/issues) para reportar bugs e/ou pedir implementações de features.

### Desenvolvimento

Pull-requests são bem vindos. Para ajudar no desenvolvimento, basta fazer:

```bash
$ git clone git@github.com:isfonzar/pagarme-go.git
$ cd pagarme-go/
$ export PAGARME_API_KEY=ak_test_su4k3yd0p4g4rm3
$ make tests
```
