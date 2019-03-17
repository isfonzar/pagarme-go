<p align="center">
    <img width="120" src="http://www.gophergala.com/assets/img/fancy_gopher_renee.jpg">
</p>

# Pagarme-go ![Language Badge](https://img.shields.io/badge/Language-Go-blue.svg) ![License Badge](https://img.shields.io/badge/License-MIT-blue.svg) ![Status Badge](https://img.shields.io/badge/Status-Beta-brightgreen.svg)

Implementação não-oficial da API do Pagar.me para Go.

Você pode acessar a documentação oficial do Pagar.me acessando esse [link](https://docs.pagar.me).

## Índice

- [Instalação](#instalação)
- [Configuração](#configuração)
- [Transações](#transações)
  - [Criando uma transação](#criando-uma-transação)
  - [Capturando uma transação](#capturando-uma-transação)
  - [Estornando uma transação](#estornando-uma-transação)
    - [Estornando uma transação parcialmente](#estornando-uma-transação-parcialmente)
    - [Estornando uma transação com split](#estornando-uma-transação-com-split)
  - [Retornando transações](#retornando-transações)
  - [Retornando uma transação](#retornando-uma-transação)
  - [Retornando recebíveis de uma transação](#retornando-recebíveis-de-uma-transação)
- [Recebível](#recebível)
  - [Retornando recebíveis](#retornando-recebíveis)
  - [Retornando um recebível](#retornando-um-recebível)
- [Contas bancárias](#contas-bancárias)
  - [Criando uma conta bancária](#criando-uma-conta-bancária)
  - [Retornando uma conta bancária](#retornando-uma-conta-bancária)
  - [Retornando contas bancárias](#retornando-contas-bancárias)
- [Endpoints faltando](#endpoints-faltando)
- [Contribuições](#contribuições)

## Instalação

Utilize o `go get` para obter o SDK e adiciona-lo a seu GOPATH.

`go get github.com/isfonzar/pagarme-go`

Para atualizar o SDK, utilize `go get -u` para obter a última versão.

`go get -u github.com/isfonzar/pagarme-go`

## Configuração

Para instanciar um novo Client, basta fazer o seguinte:

```go
package main

import "github.com/isfonzar/pagarme-go/pkg/client"

func main () {
    client := *client.NewClient("SUA_CHAVE_DE_API", nil)
}
```

### Definindo headers customizados

1. Se necessário, é possível definir headers http customizados para os requests. Para isso basta informá-los durante a instanciação do objeto `Client`:

```go
package main

import "github.com/isfonzar/pagarme-go/pkg/client"

func main () {
	headers := &client.Headers{
        "X-PagarMe-User-Agent": "Pagar.me Checkout/2.0.0",
        "X-PagarMe-Version": "2017-08-28",
    }
	
    client := *client.NewClient("SUA_CHAVE_DE_API", headers)
}
```

E então, você pode poderá utilizar o cliente para fazer requisições ao Pagar.me, como nos exemplos abaixo.

## Transações

Nesta seção será explicado como utilizar transações no Pagar.me com essa biblioteca.

### Criando uma transação

```go
package main

import (
	"github.com/isfonzar/pagarme-go/pkg/client"
	"github.com/isfonzar/pagarme-go/pkg/transactions"
)

func main () {
    client := *client.NewClient("SUA_CHAVE_DE_API", nil)
    
    t := transactions.CreateTransaction{
        Amount:             100,
        CardHolderName:     "Morpheus Fishburne",
        CardExpirationDate: "1220",
        CardNumber:         "4242424242424242",
        CardCVV:            "123",
        PostbackURL:        "http://requestb.in/pkt7pgpk",
        Installments:       "1",
        Billing: transactions.Billing{
            Name:    "Morpheus Fishburne",
            Address: transactions.Address{
                Street:        "Avenida Brigadeiro Faria Lima",
                StreetNumber:  "1811",
                Neighborhood:  "Jardim Paulistano",
                Zipcode:       "01451001",
                Country:       "br",
                State:         "sp",
                City:          "Sao Paulo",
                Complementary: "Quinto andar",
            },
        },
        Shipping: transactions.Shipping{
            Name:         "Neo Reeves",
            DeliveryDate: "2000-12-21",
            Address:      transactions.Address{
                Street:        "Avenida Brigadeiro Faria Lima",
                StreetNumber:  "1811",
                Neighborhood:  "Jardim Paulistano",
                Zipcode:       "01451001",
                Country:       "br",
                State:         "sp",
                City:          "Sao Paulo",
                Complementary: "Quinto andar",
            },
        },
        Items: []transactions.Item{
            {
                Id:        "r123",
                Title:     "Red Pill",
                UnitPrice: 10000,
                Quantity:  1,
                Category:  "Pills",
                Venue:     "Matrix",
                Date:      "2000-12-21",
            },
            {
                Id:        "b123",
                Title:     "Blue Pill",
                UnitPrice: 10000,
                Quantity:  1,
                Category:  "Pills",
                Venue:     "Matrix",
                Date:      "2000-12-21",
            },
        },
        Customer: transactions.Customer{
            ExternalID: "#3311",
            Type:       "individual",
            Name:       "Morpheus Fishburne",
            Country:    "br",
            Email:      "mopheus@nabucodonozor.com",
            Documents: []transactions.Document{
                {
                    Number: "30621143049",
                    Type:   "cpf",
                },
            },
            PhoneNumbers: []string{
                "+5511999998888",
                "+5511888889999",
            },
        },
    }
	
    transaction, err := t.Create(client)
}
```

### Capturando uma transação

```go
t := transactions.Transaction{
    ID: 123456,
    Amount: 100,
}

transaction, err := t.Capture(client)
```

### Estornando uma transação

```go
t := transactions.Transaction{
    ID: 123456,
}

transaction, err := t.Refund(client)
```

Esta funcionalidade também funciona com estornos parciais, ou estornos com split. Por exemplo:

#### Estornando uma transação parcialmente

```go
t := transactions.Transaction{
    ID: 123456,
    Amount: 1000, // Valor parcial do estorno
}

transaction, err := t.Refund(client)
```

#### Estornando uma transação com split

```go
t := transactions.Transaction{
    ID: 123456,
    Amount: 1000, // Valor parcial do estorno
    RefundSplitRules: []transactions.RefundSplitRules{
        {
            ID: "sr_cj41w9m4d01ta316d02edaqav",
            Amount: 3000,
            RecipientID: "re_cj2wd5ul500d4946do7qtjrvk",
        },
        {
            ID: "sr_cj41w9m4e01tb316dl2f2veyz",
            Amount: 3153,
            RecipientID: "re_cj2wd5u2600fecw6eytgcbkd0",
            ChargeProcessingFee: true,
        },
    },
}

transaction, err := t.Refund(client)
```

### Retornando transações

```go
t := transactions.Transaction{}

transactions, err := t.GetList(client)
```

### Retornando uma transação

```go
t := transactions.Transaction{
	ID: 123456,
}

transaction, err := t.Get(client)
```


### Retornando recebíveis de uma transação

```go
t := transactions.Transaction{
	ID: 123456,
}

payables, err := t.ListPayables(client)
```






## Recebível

Objeto contendo os dados de um recebível. O recebível (payable) é gerado automaticamente após uma transação ser paga. Para cada parcela de uma transação é gerado um recebível, que também pode ser dividido por recebedor (no caso de um split ter sido feito).

### Retornando recebíveis

```php
<?php
$payables = $pagarme->payables()->getList();
```

Se necessário, você pode aplicar filtros na busca dos payables, por exemplo, você pode recuperar todos os payables de uma transação:

```php
<?php
$transactionPayables = $pagarme->payables()->getList([
    'transaction_id' => 'ID_DA_TRANSAÇÃO'
]);
```

### Retornando um recebível

```php
<?php
$payable = $pagarme->payables()->get([
    'id' => 'ID_DO_PAYABLE'
]);
```

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
    
    if err != nil {
        t.Errorf("Error when sending capture transaction request got: %s", err.Error())
    }
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
