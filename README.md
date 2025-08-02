# integration_api
Integration API Core Services

This application provides several resources for exploring data based on REST api calls

## Configuring workspace

### Installing dependêncies
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