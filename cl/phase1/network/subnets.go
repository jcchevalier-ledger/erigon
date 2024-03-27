package network

import (
	"github.com/ledgerwatch/erigon/cl/cltypes/solid"
	"github.com/ledgerwatch/erigon/cl/phase1/core/state"
)

// def compute_subnets_for_sync_committee(state: BeaconState, validator_index: ValidatorIndex) -> Set[uint64]:
//     next_slot_epoch = compute_epoch_at_slot(Slot(state.slot + 1))
//     if compute_sync_committee_period(get_current_epoch(state)) == compute_sync_committee_period(next_slot_epoch):
//         sync_committee = state.current_sync_committee
//     else:
//         sync_committee = state.next_sync_committee

// target_pubkey = state.validators[validator_index].pubkey
// sync_committee_indices = [index for index, pubkey in enumerate(sync_committee.pubkeys) if pubkey == target_pubkey]
// return set([
//
//	uint64(index // (SYNC_COMMITTEE_SIZE // SYNC_COMMITTEE_SUBNET_COUNT))
//	for index in sync_committee_indices
//
// ])

// ComputeSubnetsForSyncCommittee is used by the ValidatorClient to determine which subnets a validator should be subscribed to for sync committees.
// the function takes an extra syncCommitteeIndicies parameter to adapt to the Beacon API specs.
func ComputeSubnetsForSyncCommittee(s *state.CachingBeaconState, syncCommitteeIndicies []uint64, validatorIndex uint64) (subnets []uint64, err error) {
	cfg := s.BeaconConfig()
	var syncCommittee *solid.SyncCommittee
	if cfg.SyncCommitteePeriod(s.Slot()) == cfg.SyncCommitteePeriod(s.Slot()+1) {
		syncCommittee = s.CurrentSyncCommittee()
	} else {
		syncCommittee = s.NextSyncCommittee()
	}

	targetPublicKey, err := s.ValidatorPublicKey(int(validatorIndex))
	if err != nil {
		return nil, err
	}

	// make sure we return each subnet id, exactly once.
	alreadySeenSubnetIndex := make(map[uint64]struct{})

	committee := syncCommittee.GetCommittee()
	for _, index := range syncCommitteeIndicies {
		subnetIdx := uint64(index) / cfg.SyncCommitteeSize / cfg.SyncCommitteeSubnetCount
		if _, ok := alreadySeenSubnetIndex[subnetIdx]; ok {
			continue
		}
		if targetPublicKey == committee[index] {
			subnets = append(subnets, subnetIdx)
			alreadySeenSubnetIndex[subnetIdx] = struct{}{}
		}
	}
	return subnets, nil
}
