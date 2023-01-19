#!/usr/bin/env bash

MOVE_EXCEPTIONS_REGEX="README.md|.gitignore|.dockerignore|^\.$|^\.\.$|^.operator-builder$"

# due to the way operator-builder works, the directory must be named pocket-operator
TMP_DIR_PARENT=$(mktemp -d)
TMP_DIR=$(mkdir -p ${TMP_DIR_PARENT}/pocket-operator && echo ${TMP_DIR_PARENT}/pocket-operator)

# Create a symlink to the current directory so that operator-builder reflects the correct `workloadConfigPath`
CURRENT_DIR=$(pwd)
ln -s "${CURRENT_DIR}" "${TMP_DIR}/.operator-builder"

INIT_OPTS="init --workload-config=.operator-builder/workload.yaml --repo github.com/pokt-network/pocket-operator --controller-image ghcr.io/pokt-network/pocket-operator:latest --skip-go-version-check"
CREATE_OPTS="create api --workload-config .operator-builder/workload.yaml --controller --resource"

cd "${TMP_DIR}"
ls -la
operator-builder ${INIT_OPTS}
operator-builder ${CREATE_OPTS}
# FILES=$()

ls -a | grep -Ev "(${MOVE_EXCEPTIONS_REGEX})" | xargs -I {} cp -R {} "${CURRENT_DIR}/.."
# cd "${CURRENT_DIR}"

# rm -rf "${TMP_DIR_PARENT}"
