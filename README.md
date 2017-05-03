# qsftpd

[![Build Status](https://travis-ci.org/yunify/qsftpd.svg?branch=master)](https://travis-ci.org/yunify/qsftpd)
[![Go Report Card](https://goreportcard.com/badge/github.com/yunify/qsftpd)](https://goreportcard.com/report/github.com/yunify/qsftpd)
[![License](http://img.shields.io/badge/license-apache%20v2-blue.svg)](https://github.com/yunify/qsftpd/blob/master/LICENSE)

A FTP server that persists all data to QingStor Object Storage.

## Usage

``` bash
$ qsftpd --help
A FTP server that persists all data to QingStor Object Storage.

Usage:
  qsftpd [flags]

Flags:
  -c, --config string   Specify config file (default "qsftpd.yaml")
```

### Create configuration

Here's an example config file named `qsftpd.yaml.example` in the project root directory, copy it to `qsftpd.yaml` and change the settings.

### Run qsftpd

Run the FTP server.

``` bash
$ qsftpd -c path/to/your/config.yaml
[2017-04-12T03:24:40.541Z #2527]  INFO -- : Listening... 127.0.0.1:2121
[2017-04-12T03:24:40.541Z #2527]  INFO -- : Starting...
[2017-04-12T03:24:49.330Z #2527]  INFO -- : FTP Client connected: ftp.connected, id: 76e209d6a89448279e947a7babe0097d, RemoteAddr: 127.0.0.1:51788, Total: 1
......
```

___Note:__ When you upload large files, please set the timeout time of FTP client long enough to avoid connection disruption._

## Not Supported Commands

Currently, the commands listed below are not supported. You can submit issue to request new features.

| Command |           Description             |
|---------|-----------------------------------|
|  ABOR   | Abort                             |
|  ACCT   | Account                           |
|  ADAT   | Authentication / Security Data    |
|  CCC    | Clear Command Channel             |
|  CONF   | Confidentiality Protected Command |
|  ENC    | Privacy Protected Command         |
|  EPRT   | Extended Port                     |
|  HELP   | Help                              |
|  LANG   | Language (for Server Messages)    |
|  MIC    | Integrity Protected Command       |
|  MLSD   | List Directory (for machine)      |
|  MLST   | List Single Object                |
|  MODE   | Transfer Mode                     |
|  REIN   | Reinitialize                      |
|  SMNT   | Structure Mount                   |
|  STOU   | Store Unique                      |
|  STRU   | File Structure                    |

___Note:__ All FTP commands can be found here ([https://tools.ietf.org/html/rfc5797](https://tools.ietf.org/html/rfc5797))._

## References

- [QingStor Documentation](https://docs.qingcloud.com/qingstor/index.html)
- [QingStor Guide](https://docs.qingcloud.com/qingstor/guide/index.html)
- [QingStor APIs](https://docs.qingcloud.com/qingstor/api/index.html)
- [FTP Command and Extension Registry](https://tools.ietf.org/html/rfc5797)

## Statement

This project is highly inspired by [`fclairamb/ftpserver`](https://github.com/fclairamb/ftpserver) which is a fork of [`andrewarrow/paradise_ftp`](https://github.com/andrewarrow/paradise_ftp).

Thanks to [Andrew Arrow](andrew@0x7a69.com) and [Florent Clairambault](florent@clairambault.fr), and the original license can be found [here](./license.txt).

## LICENSE

The Apache License (Version 2.0, January 2004).
