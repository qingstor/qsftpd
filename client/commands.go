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
}
