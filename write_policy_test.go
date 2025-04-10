package aerospike

import (
	"testing"
	"time"

	dynconfig "github.com/aerospike/aerospike-client-go/v8/config"
	"github.com/stretchr/testify/require"
)

func TestApplyConfig(t *testing.T) {
	// Create the full config.
	config := &DynConfig{
		config: &dynconfig.Config{
			Dynamic: &dynconfig.DynamicConfig{
				Write: &dynconfig.Write{
					TotalTimeout:        func() *dynconfig.Duration { r := dynconfig.Duration(5000 * time.Millisecond); return &r }(),
					SocketTimeout:       func() *dynconfig.Duration { d := dynconfig.Duration(time.Second * 3); return &d }(),
					MaxRetries:          func() *int { r := 3; return &r }(),
					DurableDelete:       func() *bool { r := true; return &r }(),
					SleepBetweenRetries: func() *dynconfig.Duration { d := dynconfig.Duration(time.Second * 2); return &d }(),
					SendKey:             func() *bool { r := true; return &r }(),
					Replica:             func() *dynconfig.Replica { r := dynconfig.PREFER_RACK; return &r }(),
				},
			},
		},
	}

	// Create an initial WritePolicy.
	policy := NewWritePolicy(0, 0)

	// Check defaults
	require.NotNil(t, policy)
	require.Equal(t, 1000, int(policy.TotalTimeout.Milliseconds()))
	require.Equal(t, 30, int(policy.SocketTimeout.Seconds()))
	require.Equal(t, 0, policy.MaxRetries)
	require.Equal(t, false, policy.DurableDelete)
	require.Equal(t, 1, int(policy.SleepBetweenRetries.Milliseconds()))
	require.Equal(t, false, policy.SendKey)

	// Apply the configuration.
	updatedPolicy := applyConfigToWritePolicy(policy, config)

	// Validate the applied configuration.
	require.NotNil(t, updatedPolicy)
	require.Equal(t, 5000, int(updatedPolicy.TotalTimeout.Milliseconds()))
	require.Equal(t, 3, int(updatedPolicy.SocketTimeout.Seconds()))
	require.Equal(t, 3, updatedPolicy.MaxRetries)
	require.Equal(t, true, updatedPolicy.DurableDelete)
	require.Equal(t, 2000, int(updatedPolicy.SleepBetweenRetries.Milliseconds()))
	require.Equal(t, true, updatedPolicy.SendKey)
	require.Equal(t, PREFER_RACK, updatedPolicy.ReplicaPolicy)
}
func TestApplyConfigWithSelectFields(t *testing.T) {
	// Create the full config.

	config := &DynConfig{
		config: &dynconfig.Config{
			Dynamic: &dynconfig.DynamicConfig{
				Write: &dynconfig.Write{
					SocketTimeout:       func() *dynconfig.Duration { d := dynconfig.Duration(time.Second * 3); return &d }(),
					MaxRetries:          func() *int { r := 3; return &r }(),
					DurableDelete:       func() *bool { r := true; return &r }(),
					SleepBetweenRetries: func() *dynconfig.Duration { d := dynconfig.Duration(time.Second * 2); return &d }(),
					SendKey:             func() *bool { r := false; return &r }(),
					Replica:             func() *dynconfig.Replica { r := dynconfig.PREFER_RACK; return &r }(),
				},
			},
		},
	}

	// Create an initial WritePolicy.
	policy := NewWritePolicy(0, 0)

	// Check defaults
	require.NotNil(t, policy)
	require.Equal(t, 1000, int(policy.TotalTimeout.Milliseconds()))
	require.Equal(t, 30, int(policy.SocketTimeout.Seconds()))
	require.Equal(t, 0, policy.MaxRetries)
	require.Equal(t, false, policy.DurableDelete)
	require.Equal(t, 1, int(policy.SleepBetweenRetries.Milliseconds()))
	require.Equal(t, false, policy.SendKey)

	// Apply the configuration.
	updatedPolicy := applyConfigToWritePolicy(policy, config)

	// Validate the applied configuration.
	require.NotNil(t, updatedPolicy)
	require.Equal(t, 3, int(updatedPolicy.SocketTimeout.Seconds()))
	require.Equal(t, 3, updatedPolicy.MaxRetries)
	require.Equal(t, true, updatedPolicy.DurableDelete)
	require.Equal(t, 2000, int(updatedPolicy.SleepBetweenRetries.Milliseconds()))
	require.Equal(t, false, updatedPolicy.SendKey)
	require.Equal(t, PREFER_RACK, updatedPolicy.ReplicaPolicy)
}
