# integration_api
Integration API Core Services

This application provides several resources for exploring data based on REST api calls

# API Info:
### Author: Lucas R. Quadros (lucas-10101)
### Version: 0.1

### Installing dependencies
- [AWS SAM cli](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-sam-cli.html)
    - Tested with version >= 1.142.1
    
- Install golang runtime and compiler
    - [Install instructions](https://go.dev/doc/install)
    - Version must be >= 1.19

- Install swagger tools to generate OpenAPI 2.0 specification documents 
    - run: `go install github.com/go-swagger/go-swagger/cmd/swagger@latest`
    - Add the binary to path
        - (linux) `echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.bashrc && source ~/.bashrc`
        - (windows) You need to add the go bin folder to you environment variables manually


### Generating OpenAPI 2.0 Specification (swagger)


### I18N: Message translation features

Message translations are provided by I18nService and I18nLanguagePack, that are simple loaders to embbeded translation files under static/translations

Each translation is a json object with `message:translation` format with the file named as lang_name.json (`en_us.json`)

Since the golang does not provide an default request context, the translations are available after the `translation_aware` context manager, injecting
inside the request context the necessary translations for the request, validating the available languages and translation support under query parameters and headers

### Request context chain based on middlewares

This application is designed to handle requests using an middleware chain, each step handling part of context and required features before and after the request is supplied, managing authorization, logging, supplying external resource access and translations to all requests, in a thread-safe manner.

### Unit tests and integration tests

***

### API Logging

*** AWS DynamoDB is used for storage of logs, using TTL and providing simple storage for api metrics and usage

### Static folder and embbeded file system

### Running this application

This application is intended to run in two differente environments, lambda functions and direct deployment or tests. You can run as a standalone golang application, building and running in your local machine or deploying as a lambda function to AWS. This process is automatically managed by the application, using the correct environment based on env variable hints