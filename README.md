# qsftp

[![License](http://img.shields.io/badge/license-apache%20v2-blue.svg)](https://github.com/yunify/qsftp/blob/master/LICENSE)

A ftp server that uses QingStor Object Storage as storage backend.

## Usage

### Create config

copy `qsftp.yaml.example` and edit as your desired.

### Run qsftp

```bash
# without -c, qsftp will run as qsftp -c qsftp.yaml
qsftp -c path/to/your/config
```

Everything is done!

## LICENSE

The Apache License (Version 2.0, January 2004).
