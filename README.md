# API Bancária Completa em GoLang

Esta é uma API completa desenvolvida em GoLang para serviços bancários. Ela oferece uma ampla gama de funcionalidades, incluindo transações, verificação de saldo, histórico de transações e muito mais. Tudo isso é feito com segurança, utilizando criptografia para proteger os dados do usuário.

## Recursos Principais

- Transações bancárias seguras com criptografia de ponta a ponta.
- Consulta de saldo em tempo real.
- Histórico de transações detalhado.
- Autenticação robusta para proteger as contas dos usuários.
- Suporte para diferentes tipos de conta (corrente, poupança, etc.).
- Notificações em tempo real para transações importantes.

## Instalação e Execução

Siga estas etapas para configurar e executar a API em sua máquina local:

1. Clone este repositório:

```
git clone https://github.com/Lbnkz/Projeto-Transacoes.git
```

2. Acesse o diretório do projeto:

```
cd projeto-transacoes
```

3. Instale as dependências:

```
go mod tidy
```

4. Configure as variáveis de ambiente:

Antes de iniciar a API, é importante configurar as variáveis de ambiente, como chaves de criptografia, configurações de banco de dados e outras configurações específicas. Revise o arquivo `.env` e forneça os valores apropriados.

5. Inicie o servidor:

```
go run main.go
```

## Endpoints Disponíveis

A API oferece os seguintes endpoints:

- `GET /usuarios/{id}/saldo`: Retorna o saldo atual de um usuário.
- `GET /usuarios/{id}/transacoes`: Retorna o histórico de transações de um usuário.
- `GET /transacoes/{id}`: Retorna os detalhes de uma transação específica.
- `POST /usuarios`: Cria um novo usuário.
- `POST /login`: Autentica um usuário e gera um token de acesso.
- `POST /transacoes`: Realiza uma nova transação entre contas.
- `POST /notificacoes`: Envia uma notificação para um usuário.

## Uso da API

Você pode usar qualquer cliente HTTP para interagir com a API, como cURL, Postman ou até mesmo sua própria aplicação Go. Certifique-se de seguir as práticas recomendadas de segurança ao lidar com informações confidenciais, como tokens de acesso.

## Contribuição

Contribuições são bem-vindas! Se você encontrar problemas, tiver sugestões de melhorias ou quiser adicionar novos recursos, sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).

---