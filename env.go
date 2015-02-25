package main

import (
	"fmt"
)

// TODO:
//   - check if entry already exists
//   - update value or ignore repeated entry

func addToEnv(env *[]string, key string, value string) error {
	ev := fmt.Sprintf("%s=%s", key, value)
	fmt.Println("adding %s to env", ev)
	*env = append(*env, ev)
	return nil
}
