# DynamoDB Storage for OAuth 2.0

> Based on the https://github.com/go-oauth2/mongo

[![License][License-Image]][License-Url]

## Install

``` bash
$ go get -u github.com/morage/go-oauth2-dynamodb
```

## Usage (specifying credentials)

``` go
package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/morage/go-oauth2-dynamodb"
)

func main() {
	manager := manage.NewDefaultManager()
	conf, err := oauth2dynamodb.NewConfig(
		aws.NewConfig(), // however you want to config
		"oauth2_basic", // Oauth2 basic table name
		"oauth2_access", // Oauth2 access table name
		"oauth2_refresh", // Oauth2 refresh table name
	)
	if err != nil {
		...
		return
	}
	manager.MustTokenStorage(
		oauth2dynamodb.NewTokenStore(conf),
	)
	// ...
}
```

## Usage (with IAM Role configured for ec2 or Lambda)

``` go
package main

import (
	"github.com/contamobi/go-oauth2-dynamodb"
	"github.com/contamobi/go-oauth2/manage"
)

func main() {
	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(
		dynamo.NewTokenStore(dynamo.NewConfig(
			"us-east-1", // AWS Region
			"", // Emtpy
			"", // Emtpy
			"", // Emtpy
			"oauth2_basic", // Oauth2 basic table name
                        "oauth2_access", // Oauth2 access table name
                        "oauth2_refresh", // Oauth2 refresh table name

		)),
	)
	// ...
}
```

## Run tests

### Start dynamodb local
```
java -Djava.library.path=./DynamoDBLocal_lib -jar DynamoDBLocal.jar -sharedDb
```

### Export env variables
```
export AWS_REGION=us-east-1
export DYNAMODB_ENDPOINT='http://localhost:8000'
export AWS_ACCESS_KEY_ID=AKIA******
export AWS_SECRET_ACCESS_KEY=**************
```

### Run tests
```
go test
```

## MIT License

```
Copyright (c) 2018 Conta.MOBI
```

[License-Url]: http://opensource.org/licenses/MIT
[License-Image]: https://img.shields.io/npm/l/express.svg
