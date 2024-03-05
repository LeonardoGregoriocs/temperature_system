# Temperature System

Sistema que receba um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin).

## Execução da Aplicação

Crie um arquivo na raiz do projeto com o nome *configs.json*, com o seguinte conteúdo:
```json
{
    "ApiKey": "your_api_key"
}
```

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


