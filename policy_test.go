package aerospike

import (
	"testing"
	"time"

	dynconfig "github.com/aerospike/aerospike-client-go/v8/config"
	"github.com/stretchr/testify/require"
)

func TestApplyConfigBasePolicy(t *testing.T) {
	// Create a dummy configuration in dynconfig.
	config := &dynconfig.Config{
		Dynamic: &dynconfig.DynamicConfig{
			Read: &dynconfig.Read{
				ReadModeAp:          func() *dynconfig.ReadModeAp { r := dynconfig.All; return &r }(),
				ReadModeSc:          func() *dynconfig.ReadModeSc { r := dynconfig.LINEARIZE; return &r }(),
				TotalTimeout:        func() *dynconfig.Duration { d := dynconfig.Duration(time.Second * 5); return &d }(),
				SocketTimeout:       func() *dynconfig.Duration { d := dynconfig.Duration(time.Second * 3); return &d }(),
				MaxRetries:          func() *int { r := 3; return &r }(),
				SleepBetweenRetries: func() *dynconfig.Duration { d := dynconfig.Duration(time.Second * 2); return &d }(),
				Replica:             func() *dynconfig.Replica { r := dynconfig.PREFER_RACK; return &r }(),
			},
		},
	}

	// Create an initial WritePolicy.
	policy := NewPolicy()

	// Check defaults
	require.NotNil(t, policy)
	require.Equal(t, ReadModeAPOne, policy.ReadModeAP)
	require.Equal(t, ReadModeSCSession, policy.ReadModeSC)
	require.Equal(t, 1000, int(policy.TotalTimeout.Milliseconds()))
	require.Equal(t, 30, int(policy.SocketTimeout.Seconds()))
	require.Equal(t, 2, policy.MaxRetries)
	require.Equal(t, 1, int(policy.SleepBetweenRetries.Milliseconds()))
	require.Equal(t, false, policy.SendKey)
	require.Equal(t, false, policy.SendKey)
	expectedReplicaPolicy := mapReplicaToReplicaPolicy(dynconfig.SEQUENCE)
	require.Equal(t, expectedReplicaPolicy, policy.ReplicaPolicy)
	require.Equal(t, false, policy.UseCompression)

	// Apply the configuration.
	updatedPolicy := applyConfigToBasePolicy(policy, config)

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
}
func TestApplyConfigBasePolicyWithSelectFields(t *testing.T) {
	// Create a dummy configuration in dynconfig.
	config := &dynconfig.Config{
		Dynamic: &dynconfig.DynamicConfig{
			Read: &dynconfig.Read{
				TotalTimeout:        func() *dynconfig.Duration { d := dynconfig.Duration(time.Second * 5); return &d }(),
				SocketTimeout:       func() *dynconfig.Duration { d := dynconfig.Duration(time.Second * 3); return &d }(),
				SleepBetweenRetries: func() *dynconfig.Duration { d := dynconfig.Duration(time.Second * 2); return &d }(),
				Replica:             func() *dynconfig.Replica { r := dynconfig.PREFER_RACK; return &r }(),
			},
		},
	}

	// Create an initial WritePolicy.
	policy := NewPolicy()

	// Check defaults
	require.NotNil(t, policy)
	require.Equal(t, ReadModeAPOne, mapReadModeAPToReadModeAP(dynconfig.ONE))
	require.Equal(t, ReadModeSCLinearize, mapReadModeSCToReadModeSC(dynconfig.LINEARIZE))
	require.Equal(t, 1000, int(policy.TotalTimeout.Milliseconds()))
	require.Equal(t, 30, int(policy.SocketTimeout.Seconds()))
	require.Equal(t, 2, policy.MaxRetries)
	require.Equal(t, 1, int(policy.SleepBetweenRetries.Milliseconds()))
	require.Equal(t, false, policy.SendKey)
	require.Equal(t, false, policy.SendKey)
	require.Equal(t, SEQUENCE, policy.ReplicaPolicy)
	require.Equal(t, false, policy.UseCompression)

	// Apply the configuration.
	updatedPolicy := applyConfigToBasePolicy(policy, config)

	// Validate the applied configuration.
	require.NotNil(t, updatedPolicy)
	require.Equal(t, 5000, int(updatedPolicy.TotalTimeout.Milliseconds()))
	require.Equal(t, 3, int(updatedPolicy.SocketTimeout.Seconds()))
	require.Equal(t, 2, updatedPolicy.MaxRetries)
	require.Equal(t, 2000, int(updatedPolicy.SleepBetweenRetries.Milliseconds()))
	require.Equal(t, false, updatedPolicy.SendKey)
	require.Equal(t, false, updatedPolicy.UseCompression)
	require.Equal(t, PREFER_RACK, updatedPolicy.ReplicaPolicy)

}
