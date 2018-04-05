# Golang ERP

#### Installing


```bash
go get -v github.com/tools/godep
go get -v github.com/gin-gonic/gin
go get -v github.com/jinzhu/gorm
go get -v github.com/json-iterator/go
go get -v github.com/lib/pq
go get -v github.com/inconshreveable/log15
go get -v github.com/kardianos/govendor
go get -v golang.org/x/crypto/bcrypt
```

#### Run

```bash
go run main.go
```

#### Test

```bash
curl -X POST \
  http://localhost:8080/v2/company/create \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "odenktools",
  "email": "odenktools86@gmail.com",
  "password": "odenktools86",
  "telephone": "0229218391",
  "code": "odk86"
}'
```

```bash
curl -X POST \
  http://localhost:8080/v2/company/login \
  -H 'Content-Type: application/json' \
  -d '{
  "email": "odenktools86@gmail.com",
  "password": "odenktools86"
}'
```


```bash
curl -X GET \
  http://localhost:8080/v2/company
```