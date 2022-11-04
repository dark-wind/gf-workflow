// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf-workflow/internal/dao/internal"
)

// internalProcessDefinesDao is internal type for wrapping internal DAO implements.
type internalProcessDefinesDao = *internal.ProcessDefinesDao

// processDefinesDao is the data access object for table process_defines.
// You can define custom methods on it to extend its functionality as you wish.
type processDefinesDao struct {
	internalProcessDefinesDao
}

var (
	// ProcessDefines is globally public accessible object for table process_defines operations.
	ProcessDefines = processDefinesDao{
		internal.NewProcessDefinesDao(),
	}
)

// Fill with you ideas below.
