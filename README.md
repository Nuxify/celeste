# The Celeste Project

Gomora service for Account Abstraction (AA) implementation using MPC SSS.

## Local Development

Setup the .env file first

```bash
cp .env.example .env
```

To bootstrap everything, run:

```bash
make
```

The command above will install, build, and run the binary

For manual install:

```bash
make install
```

For lint:

```bash
make lint
```

Just ensure you installed golangci-lint.

To test:

```bash
make test
```

For manual build:

```bash
make build

# The output for this is in bin/
```

## Docker Build

To build, run:

```bash
make run
```

To run the container:

```bash
make up
```

## Database Migration

Gomora uses go-migrate (https://github.com/golang-migrate/migrate) to handle migration. Download and change your migrate database command accordingly.

To create a schema, run:

```bash
make schema NAME=<init_schema>
```

To migrate up, run:

```bash
STEPS=<remove STEPS to apply all or specify step number> make migrate-up
```

To migrate down, run:

```bash
STEPS=<remove STEPS to apply all or specify step number> make migrate-down
```

To check migrate version, run:

```bash
make migrate-version
```

To force migrate, run:

```bash
STEPS=<specify step number> make migrate-force
```

## License

[MIT](https://choosealicense.com/licenses/mit/)

Made with ❤️ at [Nuxify](https://nuxify.tech)
