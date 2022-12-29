package args

import "github.com/spf13/cobra"

type args[T any] struct {
	validator func(cmd *cobra.Command, args []string) (T, error)
	Fields    T
}

// Create new args object with fields and validator
func New[T any](validator func(cmd *cobra.Command, rawArgs []string) (T, error)) args[T] {
	return args[T]{
		validator: validator,
	}
}

// Cobra Positional Args function
func (a *args[T]) Validate(cmd *cobra.Command, rawArgs []string) error {
	fields, err := a.validator(cmd, rawArgs)

	if err != nil {
		return err
	}

	// mutate fields after successful validation
	a.Fields = fields
	return nil
}
