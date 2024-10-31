# yaws
Yet Another Web Store

## Running

To run the application, you need to have `docker` and `docker-compose` installed.

```bash
make dbup
```

This will start the database and the application. The application will be available at `http://localhost:8080`.

```bash
make run
```

## API Documentation

The API documentation is available at `http://localhost:8080/docs/`.


## Assumptions / Limitations

- Webhook authentication is done using a hardcoded secret key
- Customers are stored in the local database
- GORM associations are not used
- The application is not production-ready

## TODOs

- Order -> Product association
- Tests
