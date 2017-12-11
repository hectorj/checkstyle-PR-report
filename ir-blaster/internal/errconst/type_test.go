package errconst_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"ir-blaster.com/ir-blaster/internal/errconst"
)

func TestError_Error(t *testing.T) {
	err := errconst.Error("test error")

	require.Equal(t, "test error", err.Error())
}
