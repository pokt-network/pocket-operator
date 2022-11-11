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

package pocketvalidator

import (
	"strconv"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	nodesv1alpha1 "github.com/pokt-network/pocket-operator/apis/nodes/v1alpha1"
	"github.com/pokt-network/pocket-operator/apis/nodes/v1alpha1/pocketvalidator/mutate"
)

// +kubebuilder:rbac:groups=apps,resources=statefulsets,verbs=get;list;watch;create;update;patch;delete

// CreateStatefulSetCollectionNameParentName creates the StatefulSet resource with name parent.Name.
func CreateStatefulSetCollectionNameParentName(
	parent *nodesv1alpha1.PocketValidator,
	collection *nodesv1alpha1.PocketSet,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "StatefulSet",
			"metadata": map[string]interface{}{
				"name":      parent.Name,     //  controlled by field:
				"namespace": collection.Name, //  controlled by collection field:
			},
			"spec": map[string]interface{}{
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": parent.Name, //  controlled by field:
					},
				},
				"serviceName": parent.Name, //  controlled by field:
				"replicas":    1,           //  we can't really scale validators, because of 1:1 relationship with the private key
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app":        parent.Name, //  controlled by field:
							"v1-purpose": "validator",
						},
					},
					"spec": map[string]interface{}{
						// imagePullPolicy: Always
						"containers": []interface{}{
							map[string]interface{}{
								"name":  "pocket",
								"image": parent.Spec.PocketImage, //  controlled by field: pocketImage
								"args": []interface{}{
									"pocket",
									"-config=/configs/config.json",
									"-genesis=/genesis.json",
								},
								"ports": []interface{}{
									map[string]interface{}{
										"containerPort": parent.Spec.Ports.Consensus, //  controlled by field: ports.consensus
										"name":          "consensus",
									},
								},
								"env": []interface{}{
									map[string]interface{}{
										"name": "POCKET_BASE_PRIVATE_KEY",
										"valueFrom": map[string]interface{}{
											"secretKeyRef": map[string]interface{}{
												"name": parent.Spec.PrivateKey.SecretKeyRef.Name, //  controlled by field: privateKey.secretKeyRef.name
												"key":  parent.Spec.PrivateKey.SecretKeyRef.Key,  //  controlled by field: privateKey.secretKeyRef.key
											},
										},
									},
									map[string]interface{}{
										"name": "POCKET_CONSENSUS_PRIVATE_KEY",
										"valueFrom": map[string]interface{}{
											"secretKeyRef": map[string]interface{}{
												"name": parent.Spec.PrivateKey.SecretKeyRef.Name, //  controlled by field: privateKey.secretKeyRef.name
												"key":  parent.Spec.PrivateKey.SecretKeyRef.Key,  //  controlled by field: privateKey.secretKeyRef.key
											},
										},
									},
									map[string]interface{}{
										"name": "POCKET_P2P_PRIVATE_KEY",
										"valueFrom": map[string]interface{}{
											"secretKeyRef": map[string]interface{}{
												"name": parent.Spec.PrivateKey.SecretKeyRef.Name, //  controlled by field: privateKey.secretKeyRef.name
												"key":  parent.Spec.PrivateKey.SecretKeyRef.Key,  //  controlled by field: privateKey.secretKeyRef.key
											},
										},
									},
									map[string]interface{}{
										"name": "POSTGRES_USER",
										"valueFrom": map[string]interface{}{
											"secretKeyRef": map[string]interface{}{
												"name": parent.Spec.Postgres.User.SecretKeyRef.Name, //  controlled by field: postgres.user.secretKeyRef.name
												"key":  parent.Spec.Postgres.User.SecretKeyRef.Key,  //  controlled by field: postgres.user.secretKeyRef.key
											},
										},
									},
									map[string]interface{}{
										"name": "POSTGRES_PASSWORD",
										"valueFrom": map[string]interface{}{
											"secretKeyRef": map[string]interface{}{
												"name": parent.Spec.Postgres.Password.SecretKeyRef.Name, //  controlled by field: postgres.password.secretKeyRef.name
												"key":  parent.Spec.Postgres.Password.SecretKeyRef.Key,  //  controlled by field: postgres.password.secretKeyRef.key
											},
										},
									},
									map[string]interface{}{
										"name":  "POSTGRES_HOST",
										"value": parent.Spec.Postgres.Host, //  controlled by field: postgres.host
									},
									map[string]interface{}{
										"name":  "POSTGRES_PORT",
										"value": parent.Spec.Postgres.Port, //  controlled by field: postgres.port
									},
									map[string]interface{}{
										"name":  "POSTGRES_DB",
										"value": parent.Spec.Postgres.Database, //  controlled by field: postgres.database
									},
									map[string]interface{}{
										"name":  "POCKET_PERSISTENCE_POSTGRES_URL",
										"value": "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)",
									},
									map[string]interface{}{
										"name":  "POCKET_PERSISTENCE_NODE_SCHEMA",
										"value": parent.Spec.Postgres.Schema, //  controlled by field: postgres.schema
									},
								},
								"volumeMounts": []interface{}{
									map[string]interface{}{
										"name":      "config-volume",
										"mountPath": "/configs",
									},
									map[string]interface{}{
										"name":      "genesis-volume",
										"mountPath": "/genesis.json",
										"subPath":   "genesis.json",
									},
									map[string]interface{}{
										"name":      "" + parent.Name + "-blockstore", //  controlled by field:
										"mountPath": "/blockstore",
									},
								},
							},
						},
						"volumes": []interface{}{
							map[string]interface{}{
								"name": "config-volume",
								"configMap": map[string]interface{}{
									"name": "" + parent.Name + "-config", //  controlled by field:
								},
							},
							map[string]interface{}{
								"name": "genesis-volume",
								"configMap": map[string]interface{}{
									"name": "" + collection.Name + "-genesis", //  controlled by collection field:
								},
							},
						},
					},
				},
				"volumeClaimTemplates": []interface{}{
					map[string]interface{}{
						"metadata": map[string]interface{}{
							"name": "" + parent.Name + "-blockstore", //  controlled by field:
						},
						"spec": map[string]interface{}{
							"accessModes": []interface{}{
								"ReadWriteOnce",
							},
							"resources": map[string]interface{}{
								"requests": map[string]interface{}{
									"storage": "1Gi",
								},
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateStatefulSetCollectionNameParentName(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// CreateServiceCollectionNameParentName creates the Service resource with name parent.Name.
func CreateServiceCollectionNameParentName(
	parent *nodesv1alpha1.PocketValidator,
	collection *nodesv1alpha1.PocketSet,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      parent.Name,     //  controlled by field:
				"namespace": collection.Name, //  controlled by collection field:
				"labels": map[string]interface{}{
					"app": parent.Name, //  controlled by field:
				},
			},
			"spec": map[string]interface{}{
				"ports": []interface{}{
					map[string]interface{}{
						"port": parent.Spec.Ports.Consensus, //  controlled by field: ports.consensus
						"name": "consensus",
					},
				},
				"selector": map[string]interface{}{
					"app": parent.Name, //  controlled by field:
				},
			},
		},
	}

	return mutate.MutateServiceCollectionNameParentName(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete

// CreateConfigMapCollectionNameParentNameConfig creates the ConfigMap resource with name parent.name + -config.
func CreateConfigMapCollectionNameParentNameConfig(
	parent *nodesv1alpha1.PocketValidator,
	collection *nodesv1alpha1.PocketSet,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"name":      "" + parent.Name + "-config", //  controlled by field:
				"namespace": collection.Name,              //  controlled by collection field:
			},
			"data": map[string]interface{}{
				// controlled by field: ports.consensus
				"config.json": `{
  "base": {
    "root_directory": "/go/src/github.com/pocket-network",
    "private_key": ""
  },
  "consensus": {
    "max_mempool_bytes": 500000000,
    "pacemaker_config": {
      "timeout_msec": 5000,
      "manual": true,
      "debug_time_between_steps_msec": 1000
    },
    "private_key": ""
  },
  "utility": {
    "max_mempool_transaction_bytes": 1073741824,
    "max_mempool_transactions": 9000
  },
  "persistence": {
    "postgres_url": "",
    "node_schema": "validator",
    "block_store_path": "/blockstore"
  },
  "p2p": {
    "consensus_port": ` + strconv.Itoa(parent.Spec.Ports.Consensus) + `,
    "use_rain_tree": true,
    "is_empty_connection_type": false,
    "private_key": ""
  },
  "telemetry": {
    "enabled": true,
    "address": "0.0.0.0:9000",
    "endpoint": "/metrics"
  }
}
`,
			},
		},
	}

	return mutate.MutateConfigMapCollectionNameParentNameConfig(resourceObj, parent, collection, reconciler, req)
}
