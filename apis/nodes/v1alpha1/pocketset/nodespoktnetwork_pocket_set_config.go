/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pocketset

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	nodesv1alpha1 "github.com/lander2k2/pocket-v1-operator/apis/nodes/v1alpha1"
)

// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete

// CreateConfigMapParentNameParentNameGenesis creates the !!start parent.Name !!end-genesis ConfigMap resource.
func CreateConfigMapParentNameParentNameGenesis(
	parent *nodesv1alpha1.PocketSet,
) ([]client.Object, error) {

	resourceObjs := []client.Object{}
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"name":      "" + parent.Name + "-genesis", //  controlled by field:
				"namespace": parent.Name,                   //  controlled by field:
			},
			"data": map[string]interface{}{
				"genesis.json": `{
  "consensus": {
    "genesis_time": {
      "seconds": 1660514050,
      "nanos": 296019000
    },
    "chain_id": "testnet",
    "max_block_bytes": 4000000
  },
  "utility": {
    "pools": [
      {
        "address": "ValidatorStakePool",
        "amount": "100000000000000"
      },
      {
        "address": "ServiceNodeStakePool",
        "amount": "100000000000000"
      },
      {
        "address": "FishermanStakePool",
        "amount": "100000000000000"
      },
      {
        "address": "DAO",
        "amount": "100000000000000"
      },
      {
        "address": "FeeCollector",
        "amount": "0"
      },
      {
        "address": "AppStakePool",
        "amount": "100000000000000"
      }
    ],
    "accounts": [
      {
        "address": "22b44cdddc6c3ca5e561cda2cae385a06e609671",
        "amount": "100000000000000"
      },
      {
        "address": "fe8cd7ca148ff08bcffb84b42d92f6c868980afa",
        "amount": "100000000000000"
      },
      {
        "address": "386a3b207085a9b44c4e5aaedf101990bea96eed",
        "amount": "100000000000000"
      },
      {
        "address": "d6990c242da15dddf43fbeb9169de110d5785b18",
        "amount": "100000000000000"
      },
      {
        "address": "4a1226601a0f3cbba852a8086f14f9a00e5a6e82",
        "amount": "100000000000000"
      },
      {
        "address": "efaf951a927b323251da5ab92c8927781395ce03",
        "amount": "100000000000000"
      },
      {
        "address": "3b67d1a6d63c295c312a2f013150f01f15f6d785",
        "amount": "100000000000000"
      }
    ],
    "applications": [
      {
        "address": "3b67d1a6d63c295c312a2f013150f01f15f6d785",
        "public_key": "d1e0ed14f7e5e98560f90cfaecc883e918558c5f458fc6aa37c799096a5045e6",
        "chains": ["0001"],
        "generic_param": "1000000",
        "staked_amount": "1000000000000",
        "paused_height": -1,
        "unstaking_height": -1,
        "output": "3b67d1a6d63c295c312a2f013150f01f15f6d785"
      }
    ],
    "validators": [
      {
        "address": "22b44cdddc6c3ca5e561cda2cae385a06e609671",
        "public_key": "79254a4bc46bf1182826145b0b01b48bab4240cd30e23ba90e4e5e6b56961c6d",
        "generic_param": "v1-validator1:8080",
        "staked_amount": "1000000000000",
        "paused_height": -1,
        "unstaking_height": -1,
        "output": "22b44cdddc6c3ca5e561cda2cae385a06e609671",
        "actor_type": 3
      },
      {
        "address": "fe8cd7ca148ff08bcffb84b42d92f6c868980afa",
        "public_key": "91dd4fd53e8e27020d62796fe68b469fad5fa5a7abc61d3eb2bd98ba16af1e29",
        "generic_param": "v1-validator2:8080",
        "staked_amount": "1000000000000",
        "paused_height": -1,
        "unstaking_height": -1,
        "output": "fe8cd7ca148ff08bcffb84b42d92f6c868980afa",
        "actor_type": 3
      },
      {
        "address": "386a3b207085a9b44c4e5aaedf101990bea96eed",
        "public_key": "3e5e4bbed5f98721163bb84445072a9202d213f1e348c5e9e0e2ea83bbb7e3aa",
        "generic_param": "v1-validator3:8080",
        "staked_amount": "1000000000000",
        "paused_height": -1,
        "unstaking_height": -1,
        "output": "386a3b207085a9b44c4e5aaedf101990bea96eed",
        "actor_type": 3
      },
      {
        "address": "d6990c242da15dddf43fbeb9169de110d5785b18",
        "public_key": "6c207cea1b1bf45dad8f8973d57291d3da31855254d7f1ed83ec3e06cabfe6b7",
        "generic_param": "v1-validator4:8080",
        "staked_amount": "1000000000000",
        "paused_height": -1,
        "unstaking_height": -1,
        "output": "d6990c242da15dddf43fbeb9169de110d5785b18",
        "actor_type": 3
      }
    ],
    "service_nodes": [
      {
        "address": "4a1226601a0f3cbba852a8086f14f9a00e5a6e82",
        "public_key": "cf05669b65132a6bd1069bcddcf47c349dd84986a03782a55081b438e15aa1bb",
        "chains": ["0001"],
        "generic_param": "node1.consensus:8080",
        "staked_amount": "1000000000000",
        "paused_height": -1,
        "unstaking_height": -1,
        "output": "4a1226601a0f3cbba852a8086f14f9a00e5a6e82",
        "actor_type": 1
      }
    ],
    "fishermen": [
      {
        "address": "efaf951a927b323251da5ab92c8927781395ce03",
        "public_key": "2beb3626a95b347057d8df86e5a4fc8b917775bd2f6594f99f0df7f25b6bae0d",
        "chains": ["0001"],
        "generic_param": "node1.consensus:8080",
        "staked_amount": "1000000000000",
        "paused_height": -1,
        "unstaking_height": -1,
        "output": "efaf951a927b323251da5ab92c8927781395ce03",
        "actor_type": 2
      }
    ],
    "params": {
      "blocks_per_session": 4,
      "app_minimum_stake": "15000000000",
      "app_max_chains": 15,
      "app_baseline_stake_rate": 100,
      "app_unstaking_blocks": 2016,
      "app_minimum_pause_blocks": 4,
      "app_max_pause_blocks": 672,
      "service_node_minimum_stake": "15000000000",
      "service_node_max_chains": 15,
      "service_node_unstaking_blocks": 2016,
      "service_node_minimum_pause_blocks": 4,
      "service_node_max_pause_blocks": 672,
      "service_nodes_per_session": 24,
      "fisherman_minimum_stake": "15000000000",
      "fisherman_max_chains": 15,
      "fisherman_unstaking_blocks": 2016,
      "fisherman_minimum_pause_blocks": 4,
      "fisherman_max_pause_blocks": 672,
      "validator_minimum_stake": "15000000000",
      "validator_unstaking_blocks": 2016,
      "validator_minimum_pause_blocks": 4,
      "validator_max_pause_blocks": 672,
      "validator_maximum_missed_blocks": 5,
      "validator_max_evidence_age_in_blocks": 8,
      "proposer_percentage_of_fees": 10,
      "missed_blocks_burn_percentage": 1,
      "double_sign_burn_percentage": 5,
      "message_double_sign_fee": "10000",
      "message_send_fee": "10000",
      "message_stake_fisherman_fee": "10000",
      "message_edit_stake_fisherman_fee": "10000",
      "message_unstake_fisherman_fee": "10000",
      "message_pause_fisherman_fee": "10000",
      "message_unpause_fisherman_fee": "10000",
      "message_fisherman_pause_service_node_fee": "10000",
      "message_test_score_fee": "10000",
      "message_prove_test_score_fee": "10000",
      "message_stake_app_fee": "10000",
      "message_edit_stake_app_fee": "10000",
      "message_unstake_app_fee": "10000",
      "message_pause_app_fee": "10000",
      "message_unpause_app_fee": "10000",
      "message_stake_validator_fee": "10000",
      "message_edit_stake_validator_fee": "10000",
      "message_unstake_validator_fee": "10000",
      "message_pause_validator_fee": "10000",
      "message_unpause_validator_fee": "10000",
      "message_stake_service_node_fee": "10000",
      "message_edit_stake_service_node_fee": "10000",
      "message_unstake_service_node_fee": "10000",
      "message_pause_service_node_fee": "10000",
      "message_unpause_service_node_fee": "10000",
      "message_change_parameter_fee": "10000",
      "acl_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "blocks_per_session_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "app_minimum_stake_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "app_max_chains_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "app_baseline_stake_rate_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "app_staking_adjustment_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "app_unstaking_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "app_minimum_pause_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "app_max_paused_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "service_node_minimum_stake_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "service_node_max_chains_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "service_node_unstaking_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "service_node_minimum_pause_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "service_node_max_paused_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "service_nodes_per_session_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "fisherman_minimum_stake_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "fisherman_max_chains_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "fisherman_unstaking_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "fisherman_minimum_pause_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "fisherman_max_paused_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "validator_minimum_stake_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "validator_unstaking_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "validator_minimum_pause_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "validator_max_paused_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "validator_maximum_missed_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "validator_max_evidence_age_in_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "proposer_percentage_of_fees_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "missed_blocks_burn_percentage_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "double_sign_burn_percentage_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_double_sign_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_send_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_stake_fisherman_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_edit_stake_fisherman_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_unstake_fisherman_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_pause_fisherman_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_unpause_fisherman_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_fisherman_pause_service_node_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_test_score_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_prove_test_score_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_stake_app_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_edit_stake_app_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_unstake_app_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_pause_app_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_unpause_app_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_stake_validator_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_edit_stake_validator_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_unstake_validator_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_pause_validator_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_unpause_validator_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_stake_service_node_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_edit_stake_service_node_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_unstake_service_node_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_pause_service_node_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_unpause_service_node_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
      "message_change_parameter_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45"
    }
  }
}
`,
			},
		},
	}

	resourceObjs = append(resourceObjs, resourceObj)

	return resourceObjs, nil
}
