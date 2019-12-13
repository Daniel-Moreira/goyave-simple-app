package userrequest

import "github.com/System-Glitch/goyave/v2/validation"

var (
	User validation.RuleSet = validation.RuleSet{
		"name":  {"required", "string", "between:3,50"},
		"email": {"required", "email"},
		"phone": {"required", "regex:^([/+]55)?([0-9]{2}[0-9]{5}[0-9]{4})"},
		"uid":   {"required", "uuid"},
	}
)
