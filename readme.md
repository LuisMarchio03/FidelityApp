# ACIM Fidelity App

![Logo](./.github/acim-backend.png)

The ACIM Fidelity App is an application that allows the registration of users, companies, loyalty cards, and points. Users can sign up, log in, update their profile, and recover their password. Companies can register, update their data, and list their information. Loyalty cards can be registered, updated, listed, and deleted, along with the associated points. The application was developed in GoLang, using SQLite3 as the database and Docker/Docker Compose for the development environment. The author of the project is Luis Gabriel Marchió Batista, and he can be contacted through various platforms, such as email, LinkedIn, Github, Instagram, Discord, and WhatsApp.

# Schemas

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
- **Finished** 
- **Point[]**

## Point

- **CardFidelityId**
- **Point: 1**
- **CreatedAt**
- **UpdatedAt**

# Features

- [x] Create Company
- [x] Create User
- [x] Create CardFidelity
- [x] List Companies
- [x] List Users
- [] List Users by Company
- [] List Users by Type
- [] List Users by Company and Type
- [x] List CardFidelity
- [x] Show Company
- [x] Show User
- [x] Show CardFidelity
- [x] Update Company
- [x] Update User
- [x] Update CardFidelity
- [x] Delete Company
- [x] Delete User
- [x] Delete CardFidelity
- [x] Login
- [x] Logout
- [x] Forgot Password
- [x] Reset Password
- [x] Update Profile
- [x] Update Password

# Requirements


## Schema user

- [] The user must be able to register in the application with name, email, password and type (user / employee / admin).

- [x] The user must be able to update name, email and password.

- [x] The user must be able to login with email and password.

<!-- - [] The user must be able to logout. -->

- [x] The user must be able to recover the password using the email.

- [x] The user must be able to reset the password.

- [x] The user must be able to update the profile (name, email and password).

- [x] The user must not be able to create a new user with an email already registered.

- [x] The user must not be able to update the email to an email already registered.

- [x] The user must not be able to create a new user with type employee and company not registered / not informed.

## Schema company

- [x] The user must be able to register a company with name, CNPJ, address, address number, address complement, address city, address state and address zip code.

- [x] The user must be able to update the company data.

- [x] The user must be able to list the companies.

- [x] The user must be able to show the company data.

- [x] The user must be able to delete the company.

- [x] The user must not be able to register a company with a CNPJ already registered.

- [x] The user must not be able to update the CNPJ to a CNPJ already registered.

<!-- - [] The user must not be able to delete a company with users registered. -->

# CardFidelity Schema

- [x] The user must be able to register a card fidelity with user, company, total points and points.

- [x] The user must be able to update the card fidelity data.

- [x] The user must be able to list the card fidelity.

- [x] The user must be able to show the card fidelity data.

- [x] The user must be able to delete the card fidelity.

<!-- - [] The user must not be able to register a card fidelity with a user already registered.

- [] The user must not be able to update the user to a user already registered.

- [] The user must not be able to delete a card fidelity with points registered. -->

# Point Schema

- [x] The user must be able to register a point with card fidelity, point and created at.

- [x] The user must be able to update the point data.

- [x] The user must be able to list the points.

- [x] The user must be able to show the point data.

- [x] The user must be able to delete the point.

<!-- - [] The user must not be able to register a point with a card fidelity already registered.

- [] The user must not be able to update the card fidelity to a card fidelity already registered.

- [] The user must not be able to delete a point with points registered. -->

# Technologies

- **GoLang**
- **SQLite3**
- **Docker**
- **Docker Compose**
- **Gin**
- **Gorm**

## Author

- **Luis Gabriel Marchió Batista** - [Luis_Marchio03](https://www.linkedin.com/in/luís-gabriel-marchió-batista-a0aa64206/)


## Contact

- **Email**: luisgabrielmarchio75@gmail.com
- **Linkedin**: https://www.linkedin.com/in/luís-gabriel-marchió-batista-a0aa64206/
- **Github**: https://github.com/LuisMarchio03/
- **Instagram**: https://www.instagram.com/luismarchio03.dev/
- **Discord**: Luís Gabriel Marchió#0305
- **Whatsapp**: +55 64 99991-8525
