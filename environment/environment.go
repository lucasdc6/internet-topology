package environment

import (
	"fmt"
	"os"
	"strings"
)

// GetDebug - Get general debug value from environment variables
func GetDebug() bool {
	if len(os.Getenv("INTERNET_TOPOLOGY_DEBUG")) != 0 {
		return true
	}
	return false
}

// GetDebugFor - Get module debug value from environment variables
func GetDebugFor(module string) bool {
	env := fmt.Sprintf("INTERNET_TOPOLOGY_%s_DEBUG", strings.ToUpper(module))
	if len(os.Getenv(env)) != 0 {
		return true
	}
	return false
}
