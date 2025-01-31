#!/usr/bin/env bash
set -Eeuo pipefail

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

args=($*)

echo "Preparing e2e tests.."

export VPP_AGENT="${VPP_AGENT:-ligato/vpp-agent:latest}"
export TESTDATA_DIR="$SCRIPT_DIR/resources"
export GOTESTSUM_FORMAT="${GOTESTSUM_FORMAT:-testname}"
export DOCKER_BUILDKIT=1

testname="vpp-agent-e2e-test"
imgname="vpp-agent-e2e-tests"
etcdcontainername="e2e-test-etcd"
sharevolumename="share-for-vpp-agent-e2e-tests"

# Compile agentctl for testing
go build -o ./tests/e2e/agentctl.test \
	  -tags 'osusergo netgo' \
    -ldflags '-w -s -extldflags "-static"' \
    -trimpath \
    ./cmd/agentctl

# Compile testing suite
go test -c -o ./tests/e2e/e2e.test \
	  -tags 'osusergo netgo e2e' \
    -ldflags '-w -s -extldflags "-static"' \
    -trimpath \
    ./tests/e2e

# Build testing image
docker build \
    -f ./tests/e2e/Dockerfile.e2e \
    --tag "${imgname}" \
    ./tests/e2e

run_e2e() {
    gotestsum --raw-command -- \
        go tool test2json -t -p "e2e" \
        ./tests/e2e/e2e.test -test.v "$@"
}

cleanup() {
	echo "Cleaning up e2e tests.."
	set -x
	docker stop -t 1 "${testname}" 2>/dev/null
	docker rm -v "${testname}" 2>/dev/null
	set +x

  echo "Stopping etcd container if running"
  if [ "$(docker ps -a | grep "${etcdcontainername}")" ]; then
    set -x
    docker stop -t 1 "${etcdcontainername}" 2>/dev/null
    docker rm -v "${etcdcontainername}" 2>/dev/null
    set +x
  fi

  echo "Removing volume for sharing files between containers"
  if [ "$(docker volume ls | grep "${sharevolumename}")" ]; then
    set -x
    docker volume rm -f "${sharevolumename}"
    set +x
  fi

}

trap 'cleanup' EXIT

echo "Creating volume for sharing files between containers.."
if docker volume create "${sharevolumename}"
then
	echo >&2 -e "\e[32m...created\e[0m"
else
	res=$?
	echo >&2 -e "\e[31m...volume creation failed!\e[0m (exit code: $res)"
	exit $res
fi

vppver=$(docker run --rm -i "$VPP_AGENT" dpkg-query -f '${Version}' -W vpp)

echo "=========================================================================="
echo -e " E2E TEST - $(date) "
echo "=========================================================================="
echo "-    VPP_AGENT: $VPP_AGENT"
echo "-     image ID: $(docker inspect $VPP_AGENT -f '{{.Id}}')"
echo "-      created: $(docker inspect $VPP_AGENT -f '{{.Created}}')"
echo "-  VPP version: $vppver"
echo "--------------------------------------------------------------------------"

# Run e2e tests
#if run_e2e ${args[@]:-}
if docker run -it \
	--name "${testname}" \
	--pid=host \
	--privileged \
	--label io.ligato.vpp-agent.testsuite=e2e \
	--label io.ligato.vpp-agent.testname="${testname}" \
	--volume "${TESTDATA_DIR}":/testdata:ro \
	--volume /var/run/docker.sock:/var/run/docker.sock \
	--volume "${sharevolumename}":/test-share \
	--env TESTDATA_DIR \
	--env INITIAL_LOGLVL \
	--env VPP_AGENT \
	--env GOTESTSUM_FORMAT \
	${DOCKER_ARGS-} \
	"${imgname}" ${args[@]:-}
then
	echo >&2 "-------------------------------------------------------------"
	echo >&2 -e " \e[32mPASSED\e[0m (took: ${SECONDS}s)"
	echo >&2 "-------------------------------------------------------------"
	exit 0
else
	res=$?
	echo >&2 "-------------------------------------------------------------"
	echo >&2 -e " \e[31mFAILED!\e[0m (exit code: $res)"
	echo >&2 "-------------------------------------------------------------"
	exit $res
fi
