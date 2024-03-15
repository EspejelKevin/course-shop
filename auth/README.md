
# Auth Microservice

**This microservice is responsible for user authentication and authorization. In addition, it has operations to validate and verify email and cell phone**

`Created with Go and its Gin framework`




## API Reference

#### Create user

```
  POST /auth-management/api/v1/signup
```
| Body | Description |
| :--- | :---------- |
|`true`| **Required.** |

```
{
    "email": "test@gmail.com",
    "password: "alphanum y len min 8",
    "phone": "+5217731230000",
    "rol": "student or teacher or admin",
    "name": "test",
    "lastname": "test"
}
```

#### Log in user

```
  POST /auth-management/api/v1/login
```

| Body | Description |
| :--- | :---------- |
|`true`| **Required.** |

```
{
    "email": "test@gmail.com",
    "password: "alphanum y len min 8"
}
```

#### Validate token

```
  GET /auth-management/api/v1/validations/token
```

| Header | Description |
| :--- | :---------- |
|`Authorization`| **Required.** Bearer {token} |


#### Validate email

```
  POST /auth-management/api/v1/validations/email
```

| Body | Description |
| :--- | :---------- |
|`true`| **Required.** |

```
{
    "code": "UvPSosXgs0tqTBymqbXY"
}
```

#### Validate phone

```
  POST /auth-management/api/v1/validations/phone
```

| Body | Description |
| :--- | :---------- |
|`true`| **Required.** |

```
{
    "code": "UvPSosXgs0tqTBymqbXY"
}
```

#### Confirm email

```
  POST /auth-management/api/v1/confirmations/email
```

| Body | Description |
| :--- | :---------- |
|`true`| **Required.** |

```
{
    "code": "UvPSosXgs0tqTBymqbXY",
    "from": "testfrom@gmail.com",
    "name": "test",
    "subject": "Verification Code",
    "to": "testto@gmail.com"
}
```

#### Confirm phone

```
  POST /auth-management/api/v1/confirmations/phone
```

| Header | Description |
| :--- | :---------- |
|`Authorization`| **Required.** Bearer {token} |

## Environment Variables

To run this project, you will need to add the following environment variables to your .sh file

Comando:

```bash
source env.sh
```

Script env.sh:
```bash
#!/bin/bash

export PORT=":3000"
export DRIVER_NAME=mysql
export DATA_SOURCE_NAME={user}:{pass}@tcp(127.0.0.1:3306)/{dbname}
export NAMESPACE=auth-management
export VERSION_API=v1
export SECRET_KEY=supersecret
export TIME_EXPIRATION=30
export SMTP_HOST=smtp.gmail.com
export SMTP_USER=test@gmail.com
export SMTP_PASS={your pass}
export SMTP_PORT=465
export EMAIL_FROM=test@gmail.com
export TWILIO_ACCOUNT_SID={twilioSID}
export TWILIO_AUTH_TOKEN={twilioToken}
export PHONE_FROM={twilioPhone}
```

