package controller

import (
	"github.com/priyanka19-98/Wordpress-Operator/pkg/controller/wordpress"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, wordpress.Add)
}
