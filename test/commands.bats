#!/usr/bin/env bats

@test "Not supported command" {
    run ftp -vnp 127.0.0.1 2121 <<EOS
user anonymous anonymous
account aaa
EOS
    [ ${status} -eq 0 ]
    echo "${output}" | grep "^500 ACCT command not supported$"
}
