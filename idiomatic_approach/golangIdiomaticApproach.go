package golangcommunityidiomaticapproach

import "sync"

// Write golang code that is more idiomatic to Go community

// Constant declaration
// unlike other programming language, constant in go can be smaller letter

const scalar = 0.1 // and its correct
const (
	Scalar  = 0.1 // for export we can also use first letter capital
	Version = 1.0
)

// Variable grouping
func varIdiomatic() int {
	var (
		scalar  = 0.1 // grouping all variable in one () is a good practice
		version = 1.0
		x       = 1
	)
	return x
}

func parseIntFromString(s string) (int, error) {  // here var name "s", in golang keep var name small
	// login
	retrun 10, nil
}

// Functions that panic
func MustParseIntFromString(s string) int { // use Must prefix if your func panic
	// logic 
	panic("ops")
	return 10 
}

// Struct initialization
type Vector struct {
	x int 
	y int 
}

func main(){
	vector := Vector(  // when initizlize vector initialize with var name
		y: 1,
        x: 2,
	)
}

// Mutex grouping
import (
	"net"
	"sync"
)

type Server struct {
	listenAddr string 
	isRunning  bool 
	
	peerMu sync.RWMutex        // while mutex declaration, group mutex with var name
	peers map[string]net.Conn

	otherMu sync.RWMutex	   // while mutex declaration, group mutex with var name
	others map[string]net.Conn 
}


// Interface declarations/naming

type Getter interface {
	Get()
}

type Putter interface {
	Put()
}

type Storer interface {  // interface need to end with "er", like storer, writer, fetcher etc
	Getter
	Putter
}

// Function grouping

func VeryImportantFuncExported() {} // put exporter and import func/var on top 
func simpleUtil() {}
func veryImportantFunc(){}

func main () {}

// Http handler naming
func handleGetUserById() {} // if the func is called in some api, name it with prefix "handle"
func hangleResizeImage() {}


// Enums
type Suit byte 

const (
    SuitHarts Suit = iota
    SuitDiamond
    SuitHeart
	SuitClub
}

func main() {}


// Constructor
type Order struct {
	Size float64
}

func NewOrder(size float64) *Order {  // right after declaring type, declare constructor with "New.."
	return &Order{
        Size: size,
    }
}
// or
func New(size float64) *Order {  // or use only "New" prefix
 	return &Order{
        Size: size,
    }
} 

func main() {}