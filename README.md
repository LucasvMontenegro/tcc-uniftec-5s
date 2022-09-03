# tcc-uniftec-5s
This project refers to a backend that provides an implementation of a 5S framework.

5S is a system for organizing spaces so work can be performed efficiently, effectively, and safely. This system focuses on putting everything where it belongs and keeping the workplace clean, which makes it easier for people to do their jobs without wasting time or risking injury.

The term 5S comes from five Japanese words:

Seiri
Seiton
Seiso
Seiketsu
Shitsuke

In English, these words are often translated to:

Sort
Set in Order
Shine
Standardize
Sustain

Each S represents one part of a five-step process that can improve the overall function of a business.

---
## Environment Variables

<br>
To run this project is necessary to add some environment variables.

<br>

> This is an Example Environments from Development environment

| KEY                                      | VALUE                                            | OPTIONAL | Description                                |
| ---------------------------------------- | ------------------------------------------------ | -------- | ------------------------------------------ |
| ENVIRONMENT                              | local                                            |          | Environment name |
| LOG_LEVEL                                | DEBUG                                            |          | Log Level                                  |
| SERVER_PORT                              | 3000                                             |          | Server Port. (used to run the http server) |
| DB_NAME                                  | sample                                           |          | Database name                              |
| DB_HOST                                  | localhost                                        |          | Database Hostname                          |
| DB_PORT                                  | 5432                                             |          | Database port service                      |
| DB_USER                                  | admin                                            |          | Database User                              |
| DB_PASSWORD                              | admin                                            |          | Database Password                          |

---

<br>

## Running the app Locally

You can run the app by setting the environment variables above in a `.env` file in the project root path. You can copy the `.env.example` and change its name to `.env`. Then, you can run the following command do set up the other services needed (Postgres):

```

docker compose up -d

```

Done that, run:

```
make run
```

---