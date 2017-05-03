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

@test "Create directory" {
    run ftp -vnp 127.0.0.1 <<EOS
user anonymous anonymous
mkdir qsftp-test
mkdir qsftp-test/nested
EOS
    [ ${status} -eq 0 ]
    echo "${output}" | grep "^257 Created dir /qsftp-test$"
    echo "${output}" | grep "^257 Created dir /qsftp-test/nested$"
}

@test "List directory" {
    run ftp -vnp 127.0.0.1 <<EOS
user anonymous anonymous
ls qsftp-test
EOS
    [ ${status} -eq 0 ]
    echo "${output}" | grep "^d---------.*nested$"
}

@test "Change directory" {
    run ftp -vnp 127.0.0.1 <<EOS
user anonymous anonymous
cd qsftp-test
ls
EOS
    [ ${status} -eq 0 ]
    echo "${output}" | grep "^250 CD worked on /qsftp-test$"
    echo "${output}" | grep "^d---------.*nested$"
}

@test "Print directory" {
    run ftp -vnp 127.0.0.1 <<EOS
user anonymous anonymous
cd qsftp-test
pwd
EOS
    [ ${status} -eq 0 ]
    echo "${output}" | grep "^250 CD worked on /qsftp-test$"
}

@test "Remove directory" {
    run ftp -vnp 127.0.0.1 <<EOS
user anonymous anonymous
rmdir qsftp-test/nested
rmdir qsftp-test
EOS
    [ ${status} -eq 0 ]
    echo "${output}" | grep "^250 Deleted dir /qsftp-test/nested$"
    echo "${output}" | grep "^250 Deleted dir /qsftp-test$"
}
