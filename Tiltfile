load('ext://restart_process', 'docker_build_with_restart')

IMAGE_NAME = "pocket-operator"
DOCKERFILE = '''FROM debian:bullseye
USER 65532:65532
WORKDIR /
COPY ./bin/manager /
CMD ["/manager"]
'''

def yaml():
    return local('make kustomize && cd config/manager; ../../bin/kustomize edit set image controller=' + IMAGE_NAME + '; cd ../..; ./bin/kustomize build config/default')

def binary():
    return 'GOOS=linux go build -o bin/manager main.go'

local("make manifests; make generate")

local_resource('Operator: CRDs', "make manifests; make install", deps=["api"])

k8s_yaml(yaml())

deps = ['controllers', 'main.go']
deps.append('api')
deps.append('internal')

local_resource('Operator: Watch&Compile', "make generate; " + binary() , deps=deps, ignore=['*/*/zz_generated.deepcopy.go'])

docker_build_with_restart(IMAGE_NAME, '.',
    dockerfile_contents=DOCKERFILE,
    entrypoint='/manager',
    only=['./bin/manager'],
    live_update=[
        sync('./bin/manager', '/manager'),
    ]
)
