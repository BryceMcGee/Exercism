package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

type Robot struct {
	name string
}

//making the namesList map global so that it is accessible in other functions
var namesList map[string]string

//init is initializing the namesList map
func init() {
	namesList = make(map[string]string)
}

//nameGen does most of the heavy lifting, actually assigning a name to the Robot struct instance
func (r *Robot) nameGen() (string, error) {
	if len(namesList) == 676000 {
		return "", errors.New("available names exhausted")
	}

	rand.Seed(time.Now().UnixNano())
	FA := string(unicode.ToUpper(rune('a' + rand.Intn(26))))
	SA := string(unicode.ToUpper(rune('a' + rand.Intn(26))))
	FN := fmt.Sprintf("%03d", rand.Intn(1000))
	name := FA + SA + FN

	_, exists := namesList[name]

	for exists == false {
		r.name = name
		namesList[name] = ""
		exists = true
	}

	//return name string, and nil, until there are no more available names, then return "", "no more names"
	return name, nil
}

//Name returns a string of the name value of the Robot struct instance. If no name is set, nameGen is called to give it one.
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		name, err := r.nameGen()
		if err != nil {
			return "", err
		}
		r.name = name
	}
	name := r.name

	return name, nil
}

//Reset removes a given name from the given robot.
func (r *Robot) Reset() {
	r.name = ""
}
