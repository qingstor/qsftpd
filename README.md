# qsftp

[![Build Status](https://travis-ci.org/yunify/qsftp.svg?branch=master)](https://travis-ci.org/yunify/qsftp)
[![Go Report Card](https://goreportcard.com/badge/github.com/yunify/qsftp)](https://goreportcard.com/report/github.com/yunify/qsftp)
[![License](http://img.shields.io/badge/license-apache%20v2-blue.svg)](https://github.com/yunify/qsftp/blob/master/LICENSE)

A FTP server that persists all data to QingStor Object Storage.

## Usage

``` bash
$ qsftp --help
A FTP server that persists all data to QingStor Object Storage.

Usage:
  qsftp [flags]

Flags:
  -c, --config string   Specify config file (default "qsftp.yaml")
```

### Create configuration

Here's an example config file named `qsftp.yaml.example` in the project root directory, copy it to `qsftp.yaml` and change the settings.

### Run qsftp

Run the FTP server

``` bash
$ qsftp -c path/to/your/config.yaml
[2017-04-12T03:24:40.541Z #2527]  INFO -- : Listening... 127.0.0.1:2121
[2017-04-12T03:24:40.541Z #2527]  INFO -- : Starting...
[2017-04-12T03:24:49.330Z #2527]  INFO -- : FTP Client connected: ftp.connected, id: 76e209d6a89448279e947a7babe0097d, RemoteAddr: 127.0.0.1:51788, Total: 1
......
```

## Not Supported Commands

Currently, the commands listed below are not supported. You can submit issue to request new features.

```
+---------+-----------------------------------+
| Command |           Description             |
+---------+-----------------------------------+
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
+---------+-----------------------------------+
```

___Note:__ All FTP commands can be found here ([https://tools.ietf.org/html/rfc5797](https://tools.ietf.org/html/rfc5797))._

## LICENSE

The Apache License (Version 2.0, January 2004).
