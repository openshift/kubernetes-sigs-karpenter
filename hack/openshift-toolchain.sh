#!/usr/bin/env bash
set -exuo pipefail

# (maxcao13): K8S_VERSION will follow the version of the k8s.io/api module.
export GOFLAGS="-mod=readonly"
K8S_VERSION=$(go list -m -f "{{ .Version }}" k8s.io/api | awk -F'[v.]' '{printf "1.%d", $3}')
KUBEBUILDER_ASSETS="/usr/local/kubebuilder/bin"

main() {
    rm -rf vendor
    trap "go mod vendor" EXIT
    tools
    kubebuilder
}

tools() {
    YQ_VERSION="v4.52.2"
    curl -sSL "https://github.com/mikefarah/yq/releases/download/${YQ_VERSION}/yq_linux_$(go env GOARCH)" -o "${GOPATH:-$HOME/go}/bin/yq"
    chmod +x "${GOPATH:-$HOME/go}/bin/yq"
    go mod download
    go mod download -modfile=go.tools.mod

    # install ko because it is not in the tools.mod file
    go install github.com/google/ko@latest
}

kubebuilder() {
    if ! mkdir -p ${KUBEBUILDER_ASSETS}; then
      sudo mkdir -p ${KUBEBUILDER_ASSETS}
      sudo chown $(whoami) ${KUBEBUILDER_ASSETS}
    fi
    arch=$(go env GOARCH)
    ln -sf $(go tool -modfile=go.tools.mod setup-envtest use -p path "${K8S_VERSION}" --arch="${arch}" --bin-dir="${KUBEBUILDER_ASSETS}")/* ${KUBEBUILDER_ASSETS}
    find $KUBEBUILDER_ASSETS
}

main "$@"
