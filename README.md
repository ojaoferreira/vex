# Vex CLI

Vex é uma ferramenta de linha de comando (CLI) desenvolvida em Go para facilitar operações e automações em projetos.

## Funcionalidades
- Gerenciamento de versões
- Integração com Git
- Operações de arquivos
- Logging customizado

## Instalação

### Usando Docker
```sh
docker build -t vex .
docker run --rm -it vex [comando]
```

### Build manual
```sh
go build -o vex ./main.go
./vex [comando]
```

## Uso
Execute `vex --help` para ver todos os comandos disponíveis.

## Estrutura do Projeto
```
├── cmd/            # Comandos principais do CLI
├── internal/       # Lógica interna e serviços
├── main.go         # Ponto de entrada
├── Dockerfile      # Build da imagem Docker
```

## Contribuição
Pull requests são bem-vindos! Para grandes mudanças, abra uma issue primeiro para discutir o que você gostaria de modificar.

## Licença
[Apache](LICENSE)
