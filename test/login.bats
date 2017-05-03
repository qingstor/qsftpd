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
