Rode o comando a seguir para rodar o MySQL através do Docker 🐳

```bash
    docker run --name curso-golang -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=cursogolang -p 3306:3306 -d mysql:latest
```

User:       **root**
Password:   **123456**