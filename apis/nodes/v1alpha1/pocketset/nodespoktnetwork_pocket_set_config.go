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
  "validators": [
    {
      "address": "0157a1d82da437eb6b2d0a612ebf934c3a54fb19",
      "public_key": "264a0707979e0d6691f74b055429b5f318d39c2883bb509310b67424252e9ef2",
      "paused": false,
      "status": 2,
      "service_url": "v1-validator1:8221",
      "staked_tokens": "1000000000000000",
      "missed_blocks": 0,
      "paused_height": -1,
      "unstaking_height": -1,
      "output": "0157a1d82da437eb6b2d0a612ebf934c3a54fb19"
    },
    {
      "address": "4cda991a51da75acf50e966c2716a7a2837d72eb",
      "public_key": "ee37d8c8e9cf42a34cfa75ff1141e2bc0ff2f37483f064dce47cb4d5e69db1d4",
      "paused": false,
      "status": 2,
      "service_url": "v1-validator2:8221",
      "staked_tokens": "1000000000000000",
      "missed_blocks": 0,
      "paused_height": -1,
      "unstaking_height": -1,
      "output": "4cda991a51da75acf50e966c2716a7a2837d72eb"
    },
    {
      "address": "67f6e8c48c62dc62a3706e7a8ba2164ca345d762",
      "public_key": "1ba66c6751506850ae0787244c69476b6d45fb857a914a5a0445a24253f7b810",
      "paused": false,
      "status": 2,
      "service_url": "v1-validator3:8221",
      "staked_tokens": "1000000000000000",
      "missed_blocks": 0,
      "paused_height": -1,
      "unstaking_height": -1,
      "output": "67f6e8c48c62dc62a3706e7a8ba2164ca345d762"
    },
    {
      "address": "b0cca84843f6f5a274150a98da66d78f0273f64e",
      "public_key": "f868bcc508133899cc47b612e4f7d9d5dacc90ce1f28214a97b651baa00bf6e4",
      "paused": false,
      "status": 2,
      "service_url": "v1-validator4:8221",
      "staked_tokens": "1000000000000000",
      "missed_blocks": 0,
      "paused_height": -1,
      "unstaking_height": -1,
      "output": "b0cca84843f6f5a274150a98da66d78f0273f64e"
    }
  ],
  "accounts": [],
  "pools": [
    {
      "name": "SERVICE_NODE_STAKE_POOL",
      "account": {
        "address": "97a8cc38033822da010422851062ae6b21b8e29d4c34193b7d8fa0f37b6593b6",
        "amount": "1000000000000000"
      }
    },
    {
      "name": "APP_STAKE_POOL",
      "account": {
        "address": "0834e9575c19d9d7d664ec460a98abb2435ece93440eb482c87d5b7259a8d271",
        "amount": "1000000000000000"
      }
    },
    {
      "name": "VALIDATOR_STAKE_POOL",
      "account": {
        "address": "664b5b682e40ee218e5feea05c2a1bb595ec15f3850c92b571cdf950b4d9ba23",
        "amount": "1000000000000000"
      }
    },
    {
      "name": "FISHERMAN_STAKE_POOL",
      "account": {
        "address": "733a711adb6fc8fbef6a3e2ac2db7842433053a23c751d19573ab85b52316f67",
        "amount": "1000000000000000"
      }
    },
    {
      "name": "DAO_POOL",
      "account": {
        "address": "aacaa24a69bcf2819ed97ab5ed8d1e490041e5c7ef9e1eddba8b5678f997ae58",
        "amount": "1000000000000000"
      }
    },
    {
      "name": "FEE_POOL",
      "account": {
        "address": "826f8e7911fa89b4bd6659c3175114d714c60bac63acc63817c0d3a4ed2fdab8",
        "amount": "1000000000000000"
      }
    }
  ],
  "fisherman": [],
  "service_nodes": [],
  "apps": [],
  "params": {}
}
`,
			},
		},
	}

	resourceObjs = append(resourceObjs, resourceObj)

	return resourceObjs, nil
}
