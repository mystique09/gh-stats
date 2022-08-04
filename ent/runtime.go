// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/mystique09/gh-profile/ent/schema"
	"github.com/mystique09/gh-profile/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[2].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescVisits is the schema descriptor for visits field.
	userDescVisits := userFields[3].Descriptor()
	// user.DefaultVisits holds the default value on creation for the visits field.
	user.DefaultVisits = userDescVisits.Default.(uint64)
}