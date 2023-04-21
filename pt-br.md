# ACIM Backend

![Logo](./.github/acim-backend.png)

ACIM Fidelity App é um aplicativo que permite o registro de usuários, empresas, cartões de fidelidade e pontos. Os usuários podem se cadastrar, fazer login, atualizar seu perfil e recuperar sua senha. As empresas podem se registrar, atualizar seus dados e listar suas informações. Os cartões de fidelidade podem ser registrados, atualizados, listados e excluídos, juntamente com os pontos associados a eles. O aplicativo foi desenvolvido em GoLang, utilizando o SQLite3 como banco de dados e Docker/Docker Compose para o ambiente de desenvolvimento. O autor do projeto é Luis Gabriel Marchió Batista e ele pode ser contatado por meio de várias plataformas, como e-mail, LinkedIn, Github, Instagram, Discord e WhatsApp.


# Esquemas

## Company

- **Name**
- **CNPJ**
- **Address**
- **Address Number**
- **Address Complement**
- **Address City**
- **Address State**
- **Address Zip Code**
- **User[]**

## User

- **Name**
- **Address**
- **Phone**
- **Email**
- **Password**
- **CompanyId (optional)**
- **Type (user / employee)**

## CardFidelity

- **UserId**
- **CompanyId**
- **TotalPoints**
- **Point[]**
- **Finished** 

## Point

- **CardFidelityId**
- **Point: 1**
- **CreatedAt**
- **UpdatedAt**

# Funcionalidades

- [x] Criar empresa
- [x] Criar usuário
- [x] Listar empresas
- [x] Listar usuários
<!-- - [x] Listar usuários por empresa
- [x] Listar usuários por tipo
- [x] Listar usuários por empresa e tipo -->
- [x] Mostrar empresa
- [x] Mostrar usuário
- [x] Atualizar empresa
- [x] Atualizar usuário
- [x] Excluir empresa
- [x] Excluir usuário
- [x] Login
- [x] Logout
- [x] Esqueci minha senha
- [x] Redefinir senha
- [x] Atualizar perfil
- [x] Atualizar senha
- [x] Controle de acesso (usuário / funcionário / admin) (middleware)

# Requisitos

## Schema user

- [x] O usuário deve ser capaz de se registrar no aplicativo com nome, email, senha e tipo (usuário / funcionário / admin).

- [x] O usuário deve ser capaz de atualizar nome, email e senha.

- [x] O usuário deve ser capaz de fazer login com email e senha.

<!-- - [] O usuário deve ser capaz de fazer logout. -->

- [x] O usuário deve ser capaz de recuperar a senha usando o email.

- [x] O usuário deve ser capaz de redefinir a senha.

- [x] O usuário deve ser capaz de atualizar o perfil (nome, email e senha).

- [] O usuário deve ser capaz de atualizar a foto do perfil.

- [x] O usuário não deve ser capaz de criar um novo usuário com um email já registrado.

- [x] O usuário não deve ser capaz de atualizar o email para um email já registrado.

- [x] O usuário não deve ser capaz de criar um novo usuário com tipo de funcionário e empresa não registrada / não informada.


## Schema company

- [x] O usuário deve ser capaz de registrar uma empresa com nome, CNPJ, endereço, número do endereço, complemento do endereço, cidade do endereço, estado do endereço e código postal do endereço.

- [x] O usuário deve ser capaz de atualizar os dados da empresa.

- [x] O usuário deve ser capaz de listar as empresas.

- [x] O usuário deve ser capaz de mostrar os dados da empresa.

- [x] O usuário deve ser capaz de excluir a empresa.

- [x] O usuário não deve ser capaz de registrar uma empresa com um CNPJ já registrado.

- [x] O usuário não deve ser capaz de atualizar o CNPJ para um CNPJ já registrado.

<!-- - [] O usuário não deve ser capaz de excluir uma empresa com usuários registrados. -->

# CardFidelity Schema

- [x] O usuário deve ser capaz de registrar um cartão fidelidade com usuário, empresa, pontos totais e pontos.

- [x] O usuário deve ser capaz de atualizar os dados do cartão fidelidade.

- [x] O usuário deve ser capaz de listar os cartões fidelidade.

- [x] O usuário deve ser capaz de mostrar os dados do cartão fidelidade.

- [x] O usuário deve ser capaz de excluir o cartão fidelidade.

<!-- - [] O usuário não deve ser capaz de registrar um cartão fidelidade com um usuário já registrado.

- [] O usuário não deve ser capaz de atualizar o usuário para um usuário já registrado.

- [] O usuário não deve ser capaz de excluir um cartão fidelidade com pontos registrados. -->

# Point Schema

- [x] O usuário deve ser capaz de registrar um ponto com cartão fidelidade, ponto e criado em.

- [x] O usuário deve ser capaz de atualizar os dados do ponto.

- [x] O usuário deve ser capaz de listar os pontos.

- [x] O usuário deve ser capaz de mostrar os dados do ponto.

- [x] O usuário deve ser capaz de excluir o ponto.
<!-- 
- [] O usuário não deve ser capaz de registrar um ponto com um cartão fidelidade já registrado.

- [] O usuário não deve ser capaz de atualizar o cartão fidelidade para um cartão fidelidade já registrado.

- [] O usuário não deve ser capaz de excluir um ponto com pontos registrados. -->

## Autor

- **Luis Gabriel Marchió Batista** - [Luis_Marchio03](https://www.linkedin.com/in/luís-gabriel-marchió-batista-a0aa64206/)


## Contato

- **Email**: luisgabrielmarchio75@gmail.com
- **Linkedin**: https://www.linkedin.com/in/luís-gabriel-marchió-batista-a0aa64206/
- **Github**: https://github.com/LuisMarchio03/
- **Instagram**: https://www.instagram.com/luismarchio03.dev/
- **Discord**: Luís Gabriel Marchió#0305
- **Whatsapp**: +55 64 99991-8525
