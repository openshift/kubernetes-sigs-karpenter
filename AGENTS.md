# AGENTS.md: openshift/kubernetes-sigs-karpenter

This file provides AI-specific guidance for working in the OpenShift downstream fork of [kubernetes-sigs/karpenter](https://github.com/kubernetes-sigs/karpenter). For contribution guidelines, see [CONTRIBUTING_OPENSHIFT.md](CONTRIBUTING_OPENSHIFT.md).

## Project Overview

This repo is the Karpenter core library (`module sigs.k8s.io/karpenter`). It is not a product binary. Cloud providers import it and implement `cloudprovider.CloudProvider`. The in-tree KWOK provider is for e2e only.

OpenShift consumers `replace` this module to `github.com/openshift/kubernetes-sigs-karpenter`:

| Repo | Role |
|------|------|
| [openshift/aws-karpenter-provider-aws](https://github.com/openshift/aws-karpenter-provider-aws) | AWS provider binary (operand) |
| [openshift/Azure-karpenter-provider-azure](https://github.com/openshift/Azure-karpenter-provider-azure) | Azure provider binary (operand) |
| [openshift/karpenter-operator](https://github.com/openshift/karpenter-operator) | Deploys the operand and applies CRDs. Embeds NodePool / NodeClaim YAML in `pkg/assets/crds/`. |

HyperShift AutoNode runs Karpenter as a Hosted Control Plane component against Hosted Cluster NodePool/NodeClaim objects.

A change here does not reach a cluster until those consumers bump their `replace`. CRD changes also need the operator's embedded YAML updated.

## Upstream / Downstream Relationship

This repo tracks [kubernetes-sigs/karpenter](https://github.com/kubernetes-sigs/karpenter). Rebasebot rebases `openshift/kubernetes-sigs-karpenter:main` onto the latest upstream tag. Config lives in [openshift/release](https://github.com/openshift/release/tree/master/ci-operator/config/openshift/kubernetes-sigs-karpenter).

### Downstream-only or OpenShift-owned

- `OWNERS`, `OWNERS_ALIASES`
- `.ci-operator.yaml`
- `hack/openshift-toolchain.sh`
- `hack/openshift-kwok-e2e.sh`
- Makefile targets `openshift-verify`, `openshift-toolchain`, `build-with-openshift`, `apply-with-openshift`
- `CONTRIBUTING_OPENSHIFT.md`, `AGENTS.md`, `.coderabbit.yaml`
- `.ko.yaml`

Go version in `go.mod` are carry patches on files that also exist upstream.

Everything else is upstream, including `.github/workflows/` (not used by OpenShift CI). Prefer sending feature and bug-fix work upstream. A downstream patch in upstream-owned files has to be re-applied on every rebase.

### Rebase Cycle

Periodically, the `main` branch is rebased onto a newer upstream release. During a rebase:

- `UPSTREAM: <drop>:` commits are discarded
- `UPSTREAM: <carry>:` commits are re-applied
- `UPSTREAM: 1234:` commits are dropped if upstream PR 1234 is now included, otherwise re-applied

## Common Pitfalls

1. **Do NOT review any code that is not a carry-patch.** That code is upstream, and should be taken up there first, not a blocker for downstream.

2. **Add a carry prefix to every downstream commit.** Use `UPSTREAM: <carry>:`, `UPSTREAM: <drop>:`, or `UPSTREAM: 1234:`.

3. **Do not use GitHub Actions.** OpenShift jobs are `unit`, `verify`, and `e2e-kwok` in [openshift/release](https://github.com/openshift/release/tree/master/ci-operator/config/openshift/kubernetes-sigs-karpenter).

4. **Match CI with `make openshift-verify`, not `go generate` alone.** That target sets `GOFLAGS=-mod=readonly` and runs the same generate / lint / vendor / dirty-tree checks as Prow. Do not hand-edit generated files.

5. **Do not modify `vendor/` directly.** `make verify` deletes it and regenerates it.

6. **Do not add product features that belong in a provider or the operator.** Credentials, NodeClass types, and ClusterOperator status live in the provider and operator repos.

7. **Do not assume a change here is in the product.** AWS, Azure, and operator `replace` pins must move, and operator-embedded CRDs must be copied, before clusters see the change.

8. **OpenShift E2Es can skip certain tests that exist in the upstream.** `hack/openshift-kwok-e2e.sh` sets `SKIP="StaticCapacity"`. This can be subject to change when we want to enable the feature across OpenShift.

9. **Keep the Makefile `JUNIT_REPORT` line.** OpenShift e2e writes junit under `$ARTIFACT_DIR`.

## Human-in-the-Loop Triggers

Stop and consult a human before:

- Modifying CRD API types: the operator copies those manifests
- Go version or `.ci-operator.yaml` bumps
- Rebase-related decisions (whether a carry is still needed)
- Anything that should instead go upstream

## Paired Changes

| If you change... | Also update... |
|-----------------|----------------|
| CRD types or generated CRD YAML | Follow-up in `openshift/karpenter-operator` `pkg/assets/crds/`, and bump `replace` in the AWS and Azure provider forks |
| OpenShift e2e skip/focus | `hack/openshift-kwok-e2e.sh` |
| `go.mod` / `go.tools.mod` | `go mod vendor`; tools live in `go.tools.mod` |

## Further Reading

- [CONTRIBUTING_OPENSHIFT.md](CONTRIBUTING_OPENSHIFT.md)
- [CONTRIBUTING.md](CONTRIBUTING.md)
- Upstream docs: [karpenter.sh](https://karpenter.sh)
