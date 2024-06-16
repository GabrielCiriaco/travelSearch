# Travel Search

## Descrição

O projeto Travel Search é um sistema que permite buscar voos e hotéis utilizando a API da Amadeus e a técnica de web scraping no Google Flights. Além disso, as informações obtidas são armazenadas em um banco de dados PostgreSQL. A comunicação entre o servidor e o cliente é feita utilizando Protobuf.

## Estrutura do Projeto

- **api.py**: Conecta-se à API da Amadeus para obter informações sobre voos.
- **crawler.py**: Usa Selenium para buscar passagens aéreas no Google Flights.
- **dataBase.py**: Gerencia a conexão e operações com o banco de dados PostgreSQL.
- **server.py**: Inicia um servidor socket para aceitar conexões de clientes e processar solicitações.
- **travelsearch.proto**: Define as mensagens Protobuf para comunicação entre os componentes.

## Instalação

### Pré-requisitos

- Python 3.x
- PostgreSQL
- Google Chrome
- ChromeDriver (compatível com a versão do Google Chrome instalada)
- Bibliotecas Python: `requests`, `selenium`, `psycopg2`, `protobuf`

### Passos para Instalação

1. **Clone o repositório:**

    ```
    git clone https://github.com/seu-usuario/travel-search.git
    cd travel-search
    ```

2. **Instale as dependências:**

    ```
    pip install requests selenium psycopg2 protobuf
    ```

3. **Compile o arquivo Protobuf:**

    ```
    protoc --python_out=. travelSearch.proto
    ```

4. **Configure o banco de dados PostgreSQL:**

    Crie um banco de dados e configure as credenciais em `dataBase.py`.

## Uso

### Iniciando o Servidor

Para iniciar o servidor, execute:

```
python server.py
```

