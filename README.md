# laravel-make

---

Scaffolds a new Laravel project using `laravel new`, creates a `docker-compose.yml` file with services for 
the database (MariaDB) and Redis, then updates the `.env` file to use `redis` as a queue driver instead of the default.

## Usage

```bash
laravel-make <project-name>
```

## Setup

```bash
go mod tidy
```

## Building the project

```bash
task build
```

---

## Changelog

Please see [CHANGELOG](CHANGELOG.md) for more information on what has changed recently.

## Contributing

Please see [CONTRIBUTING](.github/CONTRIBUTING.md) for details.

## Security Vulnerabilities

Please review [our security policy](../../security/policy) on how to report security vulnerabilities.

## Credits

- [Patrick Organ](https://github.com/patinthehat)
- [All Contributors](../../contributors)

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
