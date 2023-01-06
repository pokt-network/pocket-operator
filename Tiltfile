load('ext://restart_process', 'docker_build_with_restart')

IMAGE_NAME = "pocket-operator"
DOCKERFILE = '''FROM debian:bullseye
USER 65532:65532
WORKDIR /
COPY ./bin/manager /
CMD ["/manager"]
'''

def yaml():
    return local('make kustomize && cd config/manager; ../../bin/kustomize edit set image controller=' + IMAGE_NAME + '; cd ../..; ./bin/kustomize build config/default', quiet = True)

def binary():
    return 'GOOS=linux go build -o bin/manager main.go'

local("make manifests; make generate")

local_resource('Operator: CRDs', "make manifests; make install", deps=["api"])

k8s_yaml(yaml())

### Monitors for changes in .operator-builder directory and rebuilds the operator from template

operator_builder_deps = ['.operator-builder/nodes.pokt.network']
local_resource('Operator: operator-builder watch & template', "cd .operator-builder && make operator-build", deps=operator_builder_deps)

### Builds and updates the operator binary on cluster

kubebuilder_deps = ['controllers', 'main.go', 'api', 'internal']

local_resource('Operator: kubebuilder Watch & Compile', "make generate; " + binary() , deps=kubebuilder_deps, ignore=['*/*/zz_generated.deepcopy.go'])

docker_build_with_restart(IMAGE_NAME, '.',
    dockerfile_contents=DOCKERFILE,
    entrypoint='/manager',
    only=['./bin/manager'],
    live_update=[
        sync('./bin/manager', '/manager'),
    ]
)
