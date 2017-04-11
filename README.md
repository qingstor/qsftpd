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

## Not Supported Commands

```
+-------+------+-------------------+------+------+------------------+
| cmd   | FEAT | description       | type | conf | RFC#s/References |
|       | Code |                   |      |      | and Notes        |
+-------+------+-------------------+------+------+------------------+
| ABOR  | base | Abort             | s    | m    | 959              |
| ACCT  | base | Account           | a    | m    | 959              |
| ADAT  | secu | Authentication/   | a    | o    | 2228, 2773, 4217 |
|       |      | Security Data     |      |      |                  |
| CCC   | secu | Clear Command     | a    | o    | 2228             |
|       |      | Channel           |      |      |                  |
| CONF  | secu | Confidentiality   | a    | o    | 2228             |
|       |      | Protected Command |      |      |                  |
| ENC   | secu | Privacy Protected | a    | o    | 2228, 2773, 4217 |
|       |      | Command           |      |      |                  |
| EPRT  | nat6 | Extended Port     | p    | o    | 2428             |
| HELP  | base | Help              | s    | m    | 959              |
| LANG  | UTF8 | Language (for     | p    | o    | 2640             |
|       |      | Server Messages)  |      |      |                  |
| MIC   | secu | Integrity         | a    | o    | 2228, 2773, 4217 |
|       |      | Protected Command |      |      |                  |
| MLSD  | MLST | List Directory    | s    | o    | 3659             |
|       |      | (for machine)     |      |      |                  |
| MLST  | MLST | List Single       | s    | o    | 3659             |
|       |      | Object            |      |      |                  |
| MODE  | base | Transfer Mode     | p    | m    | 959              |
| REIN  | base | Reinitialize      | a    | m    | 959              |
| SMNT  | base | Structure Mount   | a    | o    | 959              |
| STOU  | base | Store Unique      | a    | o    | 959, 1123        |
| STRU  | base | File Structure    | p    | m    | 959              |
+-------+------+-------------------+------+------+------------------+
```

All commands can be found [here](https://tools.ietf.org/html/rfc5797).

If you need to use commands that we not supported, you can create issues to let us know.


## LICENSE

The Apache License (Version 2.0, January 2004).
