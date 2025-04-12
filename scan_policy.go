// Copyright 2014-2022 Aerospike, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aerospike

import "time"

// ScanPolicy encapsulates parameters used in scan operations.
//
// Inherited Policy fields Policy.Txn are ignored in scan commands.
type ScanPolicy struct {
	MultiPolicy
}

// NewScanPolicy creates a new ScanPolicy instance with default values.
// Set MaxRetries for scans on server versions >= 4.9. All other
// scans are not retried.
//
// The latest servers support retries on individual data partitions.
// This feature is useful when a cluster is migrating and partition(s)
// are missed or incomplete on the first scan attempt.
//
// If the first scan attempt misses 2 of 4096 partitions, then only
// those 2 partitions are retried in the next scan attempt from the
// last key digest received for each respective partition.  A higher
// default MaxRetries is used because it's wasteful to invalidate
// all scan results because a single partition was missed.
func NewScanPolicy() *ScanPolicy {
	mp := *NewMultiPolicy()
	mp.TotalTimeout = 0

	return &ScanPolicy{
		MultiPolicy: mp,
	}
}

// copyQueryPolicy creates a new BasePolicy instance and copies the values from the source BasePolicy.
func copyScanPolicy(src *ScanPolicy) *ScanPolicy {
	if src == nil {
		return nil
	}

	response := NewScanPolicy()

	response.Txn = src.Txn
	response.FilterExpression = src.FilterExpression
	response.ReadModeAP = src.ReadModeAP
	response.ReadModeSC = src.ReadModeSC
	response.TotalTimeout = src.TotalTimeout
	response.SocketTimeout = src.SocketTimeout
	response.MaxRetries = src.MaxRetries
	response.ReadTouchTTLPercent = src.ReadTouchTTLPercent
	response.SleepBetweenRetries = src.SleepBetweenRetries
	response.SleepMultiplier = src.SleepMultiplier
	response.ExitFastOnExhaustedConnectionPool = src.ExitFastOnExhaustedConnectionPool
	response.SendKey = src.SendKey
	response.UseCompression = src.UseCompression
	response.ReplicaPolicy = src.ReplicaPolicy
	response.IncludeBinData = src.IncludeBinData

	return response
}

// applyConfigToQueryPolicy applies the dynamic configuration and generates a new policy. This function
// will NOT override any custom settings in the QueryPolicy.
func applyConfigToScanPolicy(policy *ScanPolicy, dynConfig *DynConfig) *ScanPolicy {
	config := dynConfig.config

	if config == nil && !dynConfig.configInitialized.Load() {
		// On initial load it is possible that the config is not yet loaded. This will kick things off to make sure
		// config is loaded.
		dynConfig.loadConfig()
		config = dynConfig.config
	}

	if config != nil && config.Dynamic != nil && config.Dynamic.Query != nil {
		var responsePolicy *ScanPolicy
		if policy != nil {
			// Copy the existing write policy to preserve any custom settings.
			responsePolicy = copyScanPolicy(policy)
		} else {
			responsePolicy = NewScanPolicy()
		}

		if config.Dynamic.Scan.ReadModeAp != nil {
			responsePolicy.ReadModeAP = mapReadModeAPToReadModeAP(*config.Dynamic.Scan.ReadModeAp)
		}
		if config.Dynamic.Scan.ReadModeSc != nil {
			responsePolicy.ReadModeSC = mapReadModeSCToReadModeSC(*config.Dynamic.Scan.ReadModeSc)
		}
		if config.Dynamic.Scan.TotalTimeout != nil {
			responsePolicy.TotalTimeout = time.Duration(*config.Dynamic.Scan.TotalTimeout)
		}
		if config.Dynamic.Scan.SocketTimeout != nil {
			responsePolicy.SocketTimeout = time.Duration(*config.Dynamic.Scan.SocketTimeout)
		}
		if config.Dynamic.Scan.MaxRetries != nil {
			responsePolicy.MaxRetries = *config.Dynamic.Scan.MaxRetries
		}
		if config.Dynamic.Scan.SleepBetweenRetries != nil {
			responsePolicy.SleepBetweenRetries = time.Duration(*config.Dynamic.Scan.SleepBetweenRetries)
		}
		if config.Dynamic.Scan.Replica != nil {
			responsePolicy.ReplicaPolicy = mapReplicaToReplicaPolicy(*config.Dynamic.Scan.Replica)
		}

		return responsePolicy
	} else {
		return policy
	}
}
