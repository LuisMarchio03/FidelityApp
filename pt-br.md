# ACIM Backend

![Logo](./.github/acim-backend.png)

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
- [] Criar usuário
- [x] Listar empresas
- [] Listar usuários
- [] Listar usuários por empresa
- [] Listar usuários por tipo
- [] Listar usuários por empresa e tipo
- [] Mostrar empresa
- [] Mostrar usuário
- [] Atualizar empresa
- [] Atualizar usuário
- [x] Excluir empresa
- [] Excluir usuário
- [] Login
- [] Logout
- [] Esqueci minha senha
- [] Redefinir senha
- [] Atualizar perfil
- [] Atualizar senha

# Requisitos

## Schema user

- [x] O usuário deve ser capaz de se registrar no aplicativo com nome, email, senha e tipo (usuário / funcionário).

- [x] O usuário deve ser capaz de atualizar nome, email e senha.

- [] O usuário deve ser capaz de fazer login com email e senha.

- [] O usuário deve ser capaz de fazer logout.

- [] O usuário deve ser capaz de recuperar a senha usando o email.

- [] O usuário deve ser capaz de redefinir a senha.

- [] O usuário deve ser capaz de atualizar o perfil (nome, email e senha).

- [] O usuário não deve ser capaz de criar um novo usuário com um email já registrado.

- [] O usuário não deve ser capaz de atualizar o email para um email já registrado.

- [] O usuário não deve ser capaz de criar um novo usuário com tipo de funcionário e empresa não registrada / não informada.

## Schema company

- [x] O usuário deve ser capaz de registrar uma empresa com nome, CNPJ, endereço, número do endereço, complemento do endereço, cidade do endereço, estado do endereço e código postal do endereço.

- [x] O usuário deve ser capaz de atualizar os dados da empresa.

- [x] O usuário deve ser capaz de listar as empresas.

- [x] O usuário deve ser capaz de mostrar os dados da empresa.

- [x] O usuário deve ser capaz de excluir a empresa.

- [] O usuário não deve ser capaz de registrar uma empresa com um CNPJ já registrado.

- [] O usuário não deve ser capaz de atualizar o CNPJ para um CNPJ já registrado.

- [] O usuário não deve ser capaz de excluir uma empresa com usuários registrados.

# CardFidelity Schema

- [x] O usuário deve ser capaz de registrar um cartão fidelidade com usuário, empresa, pontos totais e pontos.

- [x] O usuário deve ser capaz de atualizar os dados do cartão fidelidade.

- [x] O usuário deve ser capaz de listar os cartões fidelidade.

- [x] O usuário deve ser capaz de mostrar os dados do cartão fidelidade.

- [x] O usuário deve ser capaz de excluir o cartão fidelidade.

- [] O usuário não deve ser capaz de registrar um cartão fidelidade com um usuário já registrado.

- [] O usuário não deve ser capaz de atualizar o usuário para um usuário já registrado.

- [] O usuário não deve ser capaz de excluir um cartão fidelidade com pontos registrados.

# Point Schema

- [x] O usuário deve ser capaz de registrar um ponto com cartão fidelidade, ponto e criado em.

- [x] O usuário deve ser capaz de atualizar os dados do ponto.

- [x] O usuário deve ser capaz de listar os pontos.

- [x] O usuário deve ser capaz de mostrar os dados do ponto.

- [x] O usuário deve ser capaz de excluir o ponto.

- [] O usuário não deve ser capaz de registrar um ponto com um cartão fidelidade já registrado.

- [] O usuário não deve ser capaz de atualizar o cartão fidelidade para um cartão fidelidade já registrado.

- [] O usuário não deve ser capaz de excluir um ponto com pontos registrados.

## Autor

- **Luis Gabriel Marchió Batista** - [Luis_Marchio03](https://www.linkedin.com/in/luís-gabriel-marchió-batista-a0aa64206/)


## Contato

- **Email**: luisgabrielmarchio75@gmail.com
- **Linkedin**: https://www.linkedin.com/in/luís-gabriel-marchió-batista-a0aa64206/
- **Github**: https://github.com/LuisMarchio03/
- **Instagram**: https://www.instagram.com/luismarchio03.dev/
- **Discord**: Luís Gabriel Marchió#0305
- **Whatsapp**: +55 64 99991-8525
