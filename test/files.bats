#  +-------------------------------------------------------------------------
#  | Copyright (C) 2017 Yunify, Inc.
#  +-------------------------------------------------------------------------
#  | Licensed under the Apache License, Version 2.0 (the "License");
#  | you may not use this work except in compliance with the License.
#  | You may obtain a copy of the License in the LICENSE file, or at:
#  |
#  | http://www.apache.org/licenses/LICENSE-2.0
#  |
#  | Unless required by applicable law or agreed to in writing, software
#  | distributed under the License is distributed on an "AS IS" BASIS,
#  | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  | See the License for the specific language governing permissions and
#  | limitations under the License.
#  +-------------------------------------------------------------------------

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
