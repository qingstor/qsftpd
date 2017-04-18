#!/usr/bin/env bats

@test "Put file" {
    run ftp -vnp 127.0.0.1 2121 <<EOS
    run dd if=/dev/zero of=object_0 bs=1048576 count=1 <<EOS
user anonymous anonymous
put object_0
EOS
    [ ${status} -eq 0 ]
    echo "${output}" | grep "^local: object_0 remote: object_0$"
    echo "${output}" | grep "^227 Entering Passive Mode"
    echo "${output}" | grep "^226 Closing transfer connection$"
}

@test "List file" {
    run ftp -vnp 127.0.0.1 2121 <<EOS
user anonymous anonymous
ls
EOS
    [ ${status} -eq 0 ]
    echo "${output}" | grep "^227 Entering Passive Mode"
    echo "${output}" | grep "^150 Using transfer connection$"
    echo "${output}" | grep "^-rwxrwxrwx.*object_0$"
    echo "${output}" | grep "^226 Closing transfer connection$"
}

@test "Get file" {
    run ftp -vnp 127.0.0.1 2121 <<EOS
user anonymous anonymous
get object_0
EOS
    [ ${status} -eq 0 ]
    echo "${output}" | grep "^227 Entering Passive Mode"
    echo "${output}" | grep "^150 Using transfer connection$"
    echo "${output}" | grep "^226 Closing transfer connection$"
}

@test "Delete file" {
    run ftp -vnp 127.0.0.1 2121 <<EOS
user anonymous anonymous
del object_0
EOS
    [ ${status} -eq 0 ]
    echo "${output}" | grep "^250 Removed file /object_0$"
}
