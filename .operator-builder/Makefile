INIT_OPTS=init \
	--workload-config .operator-builder/workload.yaml \
   	--repo github.com/pokt-network/pocket-operator \
	--controller-image pokt-network/pocket-operator:latest \
    --skip-go-version-check

CREATE_OPTS=create api \
	--workload-config .operator-builder/workload.yaml \
	--controller \
	--resource# \
	#--force

# TMP_DIR=../.tmp-operator/pocket-operator

# MOVE_EXCEPTIONS_REGEX=README.md|.gitignore|^\.$$|^\.\.$$

operator-clean:
	rm -rf ../*

operator-init:
	cd .. && operator-builder $(INIT_OPTS)

operator-create:
	cd .. && operator-builder $(CREATE_OPTS)

# operator-move-files:
# 	# ls -a $(TMP_DIR) | grep -Ev '($(MOVE_EXCEPTIONS_REGEX))'
# 	FILES=$(cd $(TMP_DIR) && ls -a | grep -Ev '($(MOVE_EXCEPTIONS_REGEX))')
# 	echo $(FILES)
# 	# cd $(TMP_DIR) && mv $(ls -a | grep -Ev '($(MOVE_EXCEPTIONS_REGEX))') ../

operator-dependencies:
	kubectl apply -f dependencies

operator-samples:
	cp samples/* ../config/samples/

checkout:
	git checkout \
		../CONTRIBUTING.md \
		../LICENSE.md \
		../README.md \
		../.gitignore \
		../Tiltfile

