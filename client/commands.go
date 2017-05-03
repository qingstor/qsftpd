// +-------------------------------------------------------------------------
// | Copyright (C) 2017 Yunify, Inc.
// +-------------------------------------------------------------------------
// | Licensed under the Apache License, Version 2.0 (the "License");
// | you may not use this work except in compliance with the License.
// | You may obtain a copy of the License in the LICENSE file, or at:
// |
// | http://www.apache.org/licenses/LICENSE-2.0
// |
// | Unless required by applicable law or agreed to in writing, software
// | distributed under the License is distributed on an "AS IS" BASIS,
// | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// | See the License for the specific language governing permissions and
// | limitations under the License.
// +-------------------------------------------------------------------------

package client

// CommandDescription defines which function should be used and if it should be
// open to anyone or only logged in users.
type CommandDescription struct {
	Open bool           // Open to clients without auth.
	Fn   func(*Handler) // Function to handle it.
}

var commandsMap map[string]*CommandDescription

func init() {
	// This is shared between FTPServer instances as there's no point in making
	// the FTP commands behave differently between them.
	// Whole commands can be found here: https://tools.ietf.org/html/rfc5797

	commandsMap = make(map[string]*CommandDescription)

	// Authentication.
	commandsMap["USER"] = &CommandDescription{Fn: (*Handler).handleUSER, Open: true}
	commandsMap["PASS"] = &CommandDescription{Fn: (*Handler).handlePASS, Open: true}

	// TLS handling.
	commandsMap["AUTH"] = &CommandDescription{Fn: (*Handler).handleAUTH, Open: true}
	commandsMap["PROT"] = &CommandDescription{Fn: (*Handler).handlePROT, Open: true}
	commandsMap["PBSZ"] = &CommandDescription{Fn: (*Handler).handlePBSZ, Open: true}

	// Misc.
	commandsMap["FEAT"] = &CommandDescription{Fn: (*Handler).handleFEAT, Open: true}
	commandsMap["SYST"] = &CommandDescription{Fn: (*Handler).handleSYST, Open: true}
	commandsMap["NOOP"] = &CommandDescription{Fn: (*Handler).handleNOOP, Open: true}
	commandsMap["OPTS"] = &CommandDescription{Fn: (*Handler).handleOPTS, Open: true}

	// File access.
	commandsMap["SIZE"] = &CommandDescription{Fn: (*Handler).handleSIZE}
	commandsMap["STAT"] = &CommandDescription{Fn: (*Handler).handleSTAT}
	commandsMap["MDTM"] = &CommandDescription{Fn: (*Handler).handleMDTM}
	commandsMap["RETR"] = &CommandDescription{Fn: (*Handler).handleRETR}
	commandsMap["STOR"] = &CommandDescription{Fn: (*Handler).handleSTOR}
	commandsMap["APPE"] = &CommandDescription{Fn: (*Handler).handleAPPE}
	commandsMap["DELE"] = &CommandDescription{Fn: (*Handler).handleDELE}
	commandsMap["RNFR"] = &CommandDescription{Fn: (*Handler).handleRNFR}
	commandsMap["RNTO"] = &CommandDescription{Fn: (*Handler).handleRNTO}
	commandsMap["ALLO"] = &CommandDescription{Fn: (*Handler).handleALLO}
	commandsMap["REST"] = &CommandDescription{Fn: (*Handler).handleREST}
	commandsMap["SITE"] = &CommandDescription{Fn: (*Handler).handleSITE}

	// Directory handling.
	commandsMap["CWD"] = &CommandDescription{Fn: (*Handler).handleCWD}
	commandsMap["PWD"] = &CommandDescription{Fn: (*Handler).handlePWD}
	commandsMap["CDUP"] = &CommandDescription{Fn: (*Handler).handleCDUP}
	commandsMap["NLST"] = &CommandDescription{Fn: (*Handler).handleLIST}
	commandsMap["LIST"] = &CommandDescription{Fn: (*Handler).handleLIST}
	commandsMap["MKD"] = &CommandDescription{Fn: (*Handler).handleMKD}
	commandsMap["RMD"] = &CommandDescription{Fn: (*Handler).handleRMD}

	// Connection handling.
	commandsMap["TYPE"] = &CommandDescription{Fn: (*Handler).handleTYPE}
	commandsMap["PASV"] = &CommandDescription{Fn: (*Handler).handlePASV}
	commandsMap["EPSV"] = &CommandDescription{Fn: (*Handler).handlePASV}
	commandsMap["PORT"] = &CommandDescription{Fn: (*Handler).handlePORT}
	commandsMap["QUIT"] = &CommandDescription{Fn: (*Handler).handleQUIT, Open: true}

	// Not Supported command.
	commandsMap["ABOR"] = nil
	commandsMap["ACCT"] = nil
	commandsMap["ADAT"] = nil
	commandsMap["CCC"] = nil
	commandsMap["CONF"] = nil
	commandsMap["ENC"] = nil
	commandsMap["EPRT"] = nil
	commandsMap["HELP"] = nil
	commandsMap["LANG"] = nil
	commandsMap["MIC"] = nil
	commandsMap["MLSD"] = nil
	commandsMap["MLST"] = nil
	commandsMap["MODE"] = nil
	commandsMap["REIN"] = nil
	commandsMap["SMNT"] = nil
	commandsMap["STOU"] = nil
	commandsMap["STRU"] = nil

}
