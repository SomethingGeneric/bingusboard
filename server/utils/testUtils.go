package utils

import "github.com/stretchr/testify/mock"

// Anything is a mock matcher that matches any argument.
var Anything = mock.MatchedBy(func(interface{}) bool { return true })
