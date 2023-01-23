INIT_OPTS=init \
	--workload-config .operator-builder/workload.yaml \
   	--repo github.com/pokt-network/pocket-operator \
	--controller-image ghcr.io/pokt-network/pocket-operator:latest \
    --skip-go-version-check

CREATE_OPTS=create api \
	--workload-config .operator-builder/workload.yaml \
	--controller \
	--resource# \
	#--force

operator-build:
	bash rebuild.sh

operator-dependencies:
	kubectl apply -f dependencies

operator-samples:
	cp samples/* ../config/samples/


