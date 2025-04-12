package aerospike

import (
	"testing"
	"time"

	dynconfig "github.com/aerospike/aerospike-client-go/v8/config"
	"github.com/stretchr/testify/require"
)

func TestApplyConfigQueryPolicy(t *testing.T) {
	// Create a dummy configuration in dynconfig.
	config := &DynConfig{
		config: &dynconfig.Config{
			Dynamic: &dynconfig.DynamicConfig{
				Query: &dynconfig.Query{
					ReadModeAp:          func() *dynconfig.ReadModeAp { d := dynconfig.All; return &d }(),
					ReadModeSc:          func() *dynconfig.ReadModeSc { d := dynconfig.LINEARIZE; return &d }(),
					TotalTimeout:        func() *dynconfig.Duration { d := dynconfig.Duration(time.Second * 5); return &d }(),
					SocketTimeout:       func() *dynconfig.Duration { d := dynconfig.Duration(time.Second * 3); return &d }(),
					MaxRetries:          func() *int { d := 3; return &d }(),
					SleepBetweenRetries: func() *dynconfig.Duration { d := dynconfig.Duration(time.Second * 2); return &d }(),
					Replica:             func() *dynconfig.Replica { d := dynconfig.PREFER_RACK; return &d }(),
					IncludeBinData:      func() *bool { d := false; return &d }(),
					RecordQueueSize:     func() *int { d := 50; return &d }(),
					ExpectedDuration:    func() *dynconfig.QueryDuration { d := dynconfig.SHORT; return &d }(),
				},
			},
		},
	}

	// Create an initial WritePolicy.
	policy := NewQueryPolicy()

	// Check defaults
	require.NotNil(t, policy)
	require.Equal(t, ReadModeAPOne, policy.ReadModeAP)
	require.Equal(t, ReadModeSCSession, policy.ReadModeSC)
	require.Equal(t, 0, int(policy.TotalTimeout.Milliseconds()))
	require.Equal(t, 30, int(policy.SocketTimeout.Seconds()))
	require.Equal(t, 5, policy.MaxRetries)
	require.Equal(t, 1, int(policy.SleepBetweenRetries.Milliseconds()))
	require.Equal(t, false, policy.SendKey)
	require.Equal(t, false, policy.SendKey)
	require.Equal(t, SEQUENCE, policy.ReplicaPolicy)
	require.Equal(t, false, policy.UseCompression)
	require.Equal(t, LONG, int(policy.ExpectedDuration)) // TODO: fix these

	// Apply the configuration.
	updatedPolicy := applyConfigToQueryPolicy(policy, config)

	// Validate the applied configuration.
	require.NotNil(t, updatedPolicy)
	require.Equal(t, ReadModeAPAll, updatedPolicy.ReadModeAP)
	require.Equal(t, ReadModeSCLinearize, updatedPolicy.ReadModeSC)
	require.Equal(t, 5000, int(updatedPolicy.TotalTimeout.Milliseconds()))
	require.Equal(t, 3, int(updatedPolicy.SocketTimeout.Seconds()))
	require.Equal(t, 3, updatedPolicy.MaxRetries)
	require.Equal(t, 2000, int(updatedPolicy.SleepBetweenRetries.Milliseconds()))
	require.Equal(t, false, updatedPolicy.SendKey)
	require.Equal(t, false, updatedPolicy.UseCompression)
	require.Equal(t, PREFER_RACK, updatedPolicy.ReplicaPolicy)
	require.Equal(t, false, updatedPolicy.IncludeBinData)
	require.Equal(t, SHORT, int(updatedPolicy.ExpectedDuration)) // TODO: fix these
}
func TestApplyConfigQueryPolicyWithSelectFields(t *testing.T) {
	// Create a dummy configuration in dynconfig.
	config := &DynConfig{
		config: &dynconfig.Config{
			Dynamic: &dynconfig.DynamicConfig{
				Query: &dynconfig.Query{
					TotalTimeout:        func() *dynconfig.Duration { d := dynconfig.Duration(time.Second * 5); return &d }(),
					SocketTimeout:       func() *dynconfig.Duration { d := dynconfig.Duration(time.Second * 3); return &d }(),
					SleepBetweenRetries: func() *dynconfig.Duration { d := dynconfig.Duration(time.Second * 2); return &d }(),
					Replica:             func() *dynconfig.Replica { r := dynconfig.PREFER_RACK; return &r }(),
				},
			},
		},
	}

	// Create an initial WritePolicy.
	policy := NewQueryPolicy()

	// Check defaults
	require.NotNil(t, policy)
	require.Equal(t, ReadModeAPOne, mapReadModeAPToReadModeAP(dynconfig.ONE))
	require.Equal(t, ReadModeSCLinearize, mapReadModeSCToReadModeSC(dynconfig.LINEARIZE))
	require.Equal(t, 0, int(policy.TotalTimeout.Milliseconds()))
	require.Equal(t, 30, int(policy.SocketTimeout.Seconds()))
	require.Equal(t, 5, policy.MaxRetries)
	require.Equal(t, 1, int(policy.SleepBetweenRetries.Milliseconds()))
	require.Equal(t, false, policy.SendKey)
	require.Equal(t, false, policy.SendKey)
	require.Equal(t, SEQUENCE, policy.ReplicaPolicy)
	require.Equal(t, false, policy.UseCompression)

	// Apply the configuration.
	updatedPolicy := applyConfigToQueryPolicy(policy, config)

	// Validate the applied configuration.
	require.NotNil(t, updatedPolicy)
	require.Equal(t, 5000, int(updatedPolicy.TotalTimeout.Milliseconds()))
	require.Equal(t, 3, int(updatedPolicy.SocketTimeout.Seconds()))
	require.Equal(t, 5, updatedPolicy.MaxRetries)
	require.Equal(t, 2000, int(updatedPolicy.SleepBetweenRetries.Milliseconds()))
	require.Equal(t, false, updatedPolicy.SendKey)
	require.Equal(t, false, updatedPolicy.UseCompression)
	require.Equal(t, PREFER_RACK, updatedPolicy.ReplicaPolicy)

}
