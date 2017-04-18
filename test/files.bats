#!/usr/bin/env bats

@test "Put file" {
    temp=$(mktemp)
    dd if=/dev/zero of=${temp} bs=1048576 count=1
    cd $(dirname ${temp})
    run ftp -vnp 127.0.0.1 2121 <<EOS
user anonymous anonymous
put $(basename ${temp}) qsftp.bin
EOS
    [ ${status} -eq 0 ]
    echo "${output}" | grep "^local: $(basename ${temp}) remote: qsftp.bin$"
    rm -f ${temp}
}

@test "List file" {
    run ftp -vnp 127.0.0.1 2121 <<EOS
user anonymous anonymous
ls
EOS
    [ ${status} -eq 0 ]
    echo "${output}" | grep "^-rwxrwxrwx.*qsftp.bin$"
}

@test "Get file" {
    temp=$(mktemp)
    run ftp -vnp 127.0.0.1 2121 <<EOS
user anonymous anonymous
get qsftp.bin ${temp}
EOS
    [ ${status} -eq 0 ]
    ls -l ${temp} | grep "1048576"
    rm -f ${temp}
}

@test "Delete file" {
    run ftp -vnp 127.0.0.1 2121 <<EOS
user anonymous anonymous
del qsftp.bin
EOS
    [ ${status} -eq 0 ]
    echo "${output}" | grep "^250 Removed file /qsftp.bin$"
}
