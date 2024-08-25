# AMA Application - Backend

Este é o backend da aplicação AMA (Ask Me Anything) desenvolvido com Golang + sockets.

## Tecnologias Utilizadas

- **Golang**
- **Socket**

## Configuração do Ambiente

1. Crie um arquivo `.env` na raiz do projeto. Esse arquivo conterá variáveis de ambiente necessárias para a aplicação.

2. Use o seguinte modelo como exemplo para o conteúdo do arquivo `.env`:

   ```env
   # Exemplo de variáveis de ambiente
   WSRS_DATABASE_PORT=5432
   WSRS_DATABASE_NAME="wsrs"
   WSRS_DATABASE_USER="postgres"
   WSRS_DATABASE_PASSWORD="123456789"
   WSRS_DATABASE_HOST="localhost"
   ```

## Como Rodar

1. Clone o repositório:
   ```bash
   git clone https://github.com/leobritto95/semana-tech-go-react-server
   ```
2. Instale as dependências:
   ```bash
   go mod tidy
   ```
3. Execute a aplicação:
   ```bash
   go run main.go
   ```

## Uso com Docker Compose (Opcional)

Se preferir rodar o banco de dados com Docker Compose, siga estas etapas:

1. Certifique-se de ter o **Docker** e o **Docker Compose** instalados.

2. Crie o arquivo `.env` conforme descrito na seção **Configuração do Ambiente**.
3. Execute o seguinte comando para subir os containers:
   ```bash
   docker-compose up --build
   ```
4. Isso irá iniciar todos os serviços definidos no arquivo `compose.yml`.
5. Para parar os containers:
   ```bash
   docker-compose down
   ```

## Deploy para Produção

1. Compile o código:
   ```bash
   go build -o ama-backend
   ```
2. Execute o binário gerado:
   ```bash
    ./ama-backend
   ```

## Frontend

O frontend da aplicação está disponível no repositório AMA [semana-tech-go-react-web](https://github.com/leobritto95/semana-tech-go-react-web).
