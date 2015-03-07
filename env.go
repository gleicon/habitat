package main

import (
	"fmt"
)

// TODO:
//   - check if entry already exists
//   - update value or ignore repeated entry

func addToEnv(env *[]string, key string, value string) error {
	ev := fmt.Sprintf("%s=%s", key, value)
	fmt.Printf("adding %s to env\n", ev)
	*env = append(*env, ev)
	return nil
}
