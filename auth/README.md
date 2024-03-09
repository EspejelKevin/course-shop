
## API Reference

#### Create user

```http
  POST /auth-management/api/v1/signup
```
| Body | Description |
| :--- | :---------- |
|`true`| **Required** |

**Estructura del body**

```json
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

```http
  POST /auth-management/api/v1/login
```

| Body | Description |
| :--- | :---------- |
|`true`| **Required** |

```json
{
    "email": "test@gmail.com",
    "password: "alphanum y len min 8"
}
```

