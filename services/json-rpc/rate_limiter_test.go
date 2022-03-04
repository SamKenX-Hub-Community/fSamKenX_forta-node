package json_rpc_test

import (
	"testing"
	"time"

	json_rpc "github.com/forta-protocol/forta-node/services/json-rpc"
	"github.com/stretchr/testify/require"
)

const testClientID = "1"

func TestRateLimiting(t *testing.T) {
	r := require.New(t)
	rateLimiter := json_rpc.NewRateLimiter(2, time.Hour)
	reachedLimit := rateLimiter.CheckLimit(testClientID)
	r.False(reachedLimit)
	reachedLimit = rateLimiter.CheckLimit(testClientID)
	r.False(reachedLimit)
	reachedLimit = rateLimiter.CheckLimit(testClientID)
	r.True(reachedLimit)
}
