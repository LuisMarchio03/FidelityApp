# ACIM Fidelity App

![Logo](./.github/acim-backend.png) 

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

# Features

- [x] Create Company
- [] Create User
- [x] List Companies
- [] List Users
- [] List Users by Company
- [] List Users by Type
- [] List Users by Company and Type
- [] Show Company
- [] Show User
- [] Update Company
- [] Update User
- [] Delete Company
- [] Delete User
- [] Login
- [] Logout
- [] Forgot Password
- [] Reset Password
- [] Update Profile
- [] Update Password

# Requirements


## Schema user

- [] The user must be able to register in the application with name, email, password and type (user / employee).

- [] The user must be able to update name, email and password.

- [] The user must be able to login with email and password.

- [] The user must be able to logout.

- [] The user must be able to recover the password using the email.

- [] The user must be able to reset the password.

- [] The user must be able to update the profile (name, email and password).

- [] The user must not be able to create a new user with an email already registered.

- [] The user must not be able to update the email to an email already registered.

- [] The user must not be able to create a new user with type employee and company not registered / not informed.

## Schema company

- [] The user must be able to register a company with name, CNPJ, address, address number, address complement, address city, address state and address zip code.

- [] The user must be able to update the company data.

- [] The user must be able to list the companies.

- [] The user must be able to show the company data.

- [] The user must be able to delete the company.

- [] The user must not be able to register a company with a CNPJ already registered.

- [] The user must not be able to update the CNPJ to a CNPJ already registered.

- [] The user must not be able to delete a company with users registered.



## Author

- **Luis Gabriel Marchió Batista** - [Luis_Marchio03](https://www.linkedin.com/in/luís-gabriel-marchió-batista-a0aa64206/)


## Contact

- **Email**: luisgabrielmarchio75@gmail.com
- **Linkedin**: https://www.linkedin.com/in/luís-gabriel-marchió-batista-a0aa64206/
- **Github**: https://github.com/LuisMarchio03/
- **Instagram**: https://www.instagram.com/luismarchio03.dev/
- **Discord**: Luís Gabriel Marchió#0305
- **Whatsapp**: +55 64 99991-8525
