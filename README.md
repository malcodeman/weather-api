# Weather api

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/facebook/react/blob/master/LICENSE)
[![code style: prettier](https://img.shields.io/badge/code_style-prettier-ff69b4.svg)](https://github.com/prettier/prettier)

Keep in mind that `GetLocation` method won't work as expected when deployed because it will fetch the location of the server, not client.

This project serves as proxy server for `Dark Sky API` because they disable cors as a [security precaution](https://darksky.net/dev/docs/faq).

## Table of Contents

- [Usage](#usage)
- [License](#license)

## Usage

.env file should look like this:

```
DARKSKY_API_KEY=123
IPSTACK_API_KEY=123
PORT=5000
```

without docker:

```
go run main.go
```

or with docker:

```
docker build -t <image-name> .
docker run -p 5000:5000 <image-name>
```

## License

This project is [MIT licensed](./LICENSE).
