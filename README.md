# Temperature System

Sistema que receba um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin).

## Execução da Aplicação

Você pode executar a aplicação facilmente utilizando Docker. Basta seguir o exemplo abaixo:

```bash
docker build -t temperature_system .
``` 
```bash
docker run temperature_system 
```

## Requisição 

- Realizar uma request do tipo POST para localhost:8080/weather com o body abaixo:

```json
{
    "cep": "xxxxxxxx"
}
```


