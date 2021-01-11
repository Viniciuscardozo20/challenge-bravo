# Desafio Challenge Bravo

Candidato: Vinicius da Silva Cardozo

LinkedIn: https://www.linkedin.com/in/vinicius-cardozo-669a15136/

Recrutadoras: Paloma Andrade e Beatriz Cristine

## Requesitos

* [Docker-compose](https://docs.docker.com/compose/install/)

## Setup

Após clonar o repositório, é necessário ir no caminho raiz e rodar o seguinte comando 

    docker-compose up

Isso subirá a API na porta `8082`

## Rotas

### Adicionar uma nova moeda

METÓDO `POST`v1/

    /v1/addCurrency
    
#### Exemplo

No terminal execute

    curl --location --request POST 'http://localhost:8082/v1/addCurrency' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "initials": "AUD", 
      }'

Ou em alguma outra interface tipo Postman ou Insomnia
**Body** (application/json):

    {
	    "initials": "AUD"
    }

Saída esperada

_Status http_: `200 OK`
```json
{
  "status": "Created",
  "data": {
    "initials": "AUD"
  }
}
```

### Converter moedas

METÓDO `GET`

    /v1/currencyConvert

**Parametros**: 

* From: Sigla da moeda base
* To: Sigla da moeda para ser convertida
* Amount: Quantidade a ser convertida   

#### Exemplo

No terminal execute

    curl --location --request GET 'http://localhost:8082/v1/currencyConvert?from=USD&to=BRL&amount=70'

Ou em alguma outra interface preferência, tipo Postman ou Insomnia

Saída esperada

_Status http_: `200 OK`
```json
{
  "status": "success",
  "data": {
    "From": "USD",
    "To": "BRL",
    "Amount": 70,
    "Value": 383.074899284
  }
}
```

### Deletar uma moeda

METÓDO `DELETE`v1/

    /v1/deleteCurrency/:currency

**Parametros**: 

* currency: Sigla da moeda a ser deletada
    
#### Exemplo

No terminal execute

    curl --location --request DELETE 'http://localhost:8082/v1/deleteCurrency/AUD' 

Ou em alguma outra interface de sua preferência, tipo Postman ou Insomnia

Saída esperada

_Status http_: `200 OK`
```json
{
  "status": "success"
}
```

### Receber todas moedas

METÓDO `GET`v1/

    /v1/getCurrencies

**Parametros**: 

* currency: Sigla da moeda a ser deletada
    
#### Exemplo

No terminal execute

    curl --location --request GET 'http://localhost:8082/v1/getCurrencies' 

Ou em alguma outra interface de sua preferência, tipo Postman ou Insomnia

Saída esperada

_Status http_: `200 OK`
```json
{
  "AUD": 1.2976239415,
  "BRL": 5.4724985612,
  "BTC": 0.00002886,
  "ETH": 0.0009379,
  "EUR": 0.8221655841,
  "USD": 1
}
```

## Rodar os testes
### Requesitos

* GoLang 1:15

