load('ext://restart_process', 'docker_build_with_restart')

IMAGE_NAME = "pocket-operator"
DOCKERFILE = '''FROM debian:bullseye
WORKDIR /
COPY ./bin/manager /
CMD ["/manager"]
'''

# Makes sure kustomize is installed and sets the image name (so tilt can intervene & control the container lifecycle)
local('make kustomize')
k8s_yaml(kustomize('config/default-localnet',  kustomize_bin='bin/kustomize'))

### Monitors for changes in .operator-builder directory and rebuilds the operator from template
operator_builder_deps = ['.operator-builder/nodes.pokt.network/']
local_resource('operator-builder-watch-and-template', "cd .operator-builder && make operator-build", deps=operator_builder_deps)

### Builds and updates the operator binary on cluster
kubebuilder_deps = ['controllers', 'main.go', 'api', 'internal']
local_resource('kubebuilder-watch-and-compile', "make generate; GOOS=linux go build -o bin/manager main.go", deps=kubebuilder_deps, ignore=['*/*/zz_generated.deepcopy.go'])

docker_build_with_restart(IMAGE_NAME, '.',
    dockerfile_contents=DOCKERFILE,
    entrypoint='/manager',
    only=['./bin/manager'],
    live_update=[
        sync('./bin/manager', '/manager'),
    ]
)
