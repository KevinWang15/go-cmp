package cmpopts

import "github.com/google/go-cmp/cmp"

// WithFullVerbosity returns an Option that disables all truncations in the diff output.
func WithFullVerbosity() cmp.Option {
	return cmp.WithFullVerbosity()
}
