#!/usr/bin/env bats

@test "Create directory" {
    run ftp -vnp 127.0.0.1 2121 <<EOS
user anonymous anonymous
mkdir qsftp-test
mkdir qsftp-test/nested
EOS
    [ ${status} -eq 0 ]
    echo "${output}" | grep "^257 Created dir /qsftp-test$"
    echo "${output}" | grep "^257 Created dir /qsftp-test/nested$"
}

@test "List directory" {
    run ftp -vnp 127.0.0.1 2121 <<EOS
user anonymous anonymous
ls qsftp-test
EOS
    [ ${status} -eq 0 ]
    echo "${output}" | grep "^d---------.*nested$"
}

@test "Change directory" {
    run ftp -vnp 127.0.0.1 2121 <<EOS
user anonymous anonymous
cd qsftp-test
ls
EOS
    [ ${status} -eq 0 ]
    echo "${output}" | grep "^250 CD worked on /qsftp-test$"
    echo "${output}" | grep "^d---------.*nested$"
}

@test "Print directory" {
    run ftp -vnp 127.0.0.1 2121 <<EOS
user anonymous anonymous
cd qsftp-test
pwd
EOS
    [ ${status} -eq 0 ]
    echo "${output}" | grep "^250 CD worked on /qsftp-test$"
}

@test "Remove directory" {
    run ftp -vnp 127.0.0.1 2121 <<EOS
user anonymous anonymous
rmdir qsftp-test/nested
rmdir qsftp-test
EOS
    [ ${status} -eq 0 ]
    echo "${output}" | grep "^250 Deleted dir /qsftp-test/nested$"
    echo "${output}" | grep "^250 Deleted dir /qsftp-test$"
}
