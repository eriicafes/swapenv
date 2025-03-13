package args

import "github.com/spf13/cobra"

type validator[T any] struct {
	validator func(cmd *cobra.Command, args []string) (T, error)
	parsed    T
}

// New returns a new args validator.
func New[T any](fn func(cmd *cobra.Command, rawArgs []string) (T, error)) validator[T] {
	return validator[T]{validator: fn}
}

// Validate validates the arg from the raw cobra positional args.
func (a *validator[T]) Validate(cmd *cobra.Command, rawArgs []string) (err error) {
	a.parsed, err = a.validator(cmd, rawArgs)
	return err
}

// Run returns a cobra command func with the parsed args.
func (a *validator[T]) Run(fn func(cmd *cobra.Command, args T)) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		fn(cmd, a.parsed)
	}
}
