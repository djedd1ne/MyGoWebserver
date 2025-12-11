// database connection
package utils

import (
	"log"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Substitute DB details
const (
	host		= "192.168.1.243"
	port		= 6543
	user		= "postgres"
	password	= "secret"
)
