package errconst_test

import (
	"testing"
	"ir-blaster.com/ir-blaster/internal/errconst"
	"github.com/stretchr/testify/require"
)

func TestError_Error(t *testing.T) {
	err := errconst.Error("test error")

	require.Equal(t, "test error", err.Error())
}
