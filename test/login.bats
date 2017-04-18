#!/usr/bin/env bats

@test "Login as anonymous user" {
    run ftp -vnp 127.0.0.1 2121 <<EOS
user anonymous anonymous
bye
EOS
    [ ${status} -eq 0 ]
    echo "${output}" | grep "^230 Password ok, continue$"
}

@test "Login as unknown user" {
    run ftp -vnp 127.0.0.1 2121 <<EOS
user unknown unknown
bye
EOS
    [ ${status} -eq 0 ]
    echo "${output}" | grep "^430 Invalid username or password$"
}
