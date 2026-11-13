# Contributing to Karpenter (OpenShift Downstream)

This document covers contribution guidelines specific to the OpenShift downstream fork of [kubernetes-sigs/karpenter](https://github.com/kubernetes-sigs/karpenter). For upstream contribution guidelines, see [CONTRIBUTING.md](CONTRIBUTING.md). For AI-specific guidance, see [AGENTS.md](AGENTS.md).

## Related Resources

| Resource | Link |
|----------|------|
| Upstream repo | [kubernetes-sigs/karpenter](https://github.com/kubernetes-sigs/karpenter) |
| AWS provider (operand) | [openshift/aws-karpenter-provider-aws](https://github.com/openshift/aws-karpenter-provider-aws) |
| Azure provider (operand) | [openshift/Azure-karpenter-provider-azure](https://github.com/openshift/Azure-karpenter-provider-azure) |
| Operator repo | [openshift/karpenter-operator](https://github.com/openshift/karpenter-operator) |
| CI configuration | [openshift/release/.../kubernetes-sigs-karpenter/](https://github.com/openshift/release/tree/master/ci-operator/config/openshift/kubernetes-sigs-karpenter) |
| AI guidance | [AGENTS.md](AGENTS.md) |
| Upstream docs | [karpenter.sh](https://karpenter.sh) |
| OpenShift docs (ROSA AutoNode) | [Managing compute nodes using Red Hat build of Karpenter](https://docs.redhat.com/en/documentation/red_hat_openshift_service_on_aws/4/html/cluster_administration/managing-compute-nodes-using-red-hat-build-of-karpenter) |

This repo is a Go library (`sigs.k8s.io/karpenter`), not a product binary. OpenShift ships Karpenter through the AWS and Azure provider forks, which the [karpenter-operator](https://github.com/openshift/karpenter-operator) deploys. Those repos `replace` this module to `github.com/openshift/kubernetes-sigs-karpenter`. Changes here do not reach the product until those repos bump the replace.

## Review and Approval Policy

Every change in every pull request must be understood and approved by two humans. This can be the PR author and a reviewer, or, if the author used an AI tool and does not fully understand the contents of the PR, two human reviewers.

**Exception:** PRs authored by deterministic automation tools that are part of our CI and related systems (whose code has been reviewed by the OpenShift engineering org) can be merged with a single human review.

Every change should be closely scrutinized for bugs. This library is the scheduling and disruption core for every Karpenter provider. Review changes from multiple angles:

- **Product architecture**: Does this fit Karpenter's library contract with cloud providers, and OpenShift's AutoNode / operator deployment model?
- **Security**: Are there new attack surfaces, credential handling issues, or privilege escalations?
- **Thread safety**: Cluster state, the disruption queue, and the provisioner batcher are shared. Are they still correctly synchronized?
- **Regressions**: Could this break provisioning, disruption, NodeClaim lifecycle, or CRD validation?
- **Effects on other components**: How does this impact the AWS and Azure providers, the karpenter-operator (including the CRDs it embeds), or HyperShift AutoNode?

## Upstream Commit Convention

This is a downstream fork. All non-upstream commits must use one of the following prefixes so changes are not lost during the next upstream rebase:

- `UPSTREAM: <carry>: ` -- A change that should be kept (carried) indefinitely, or as long as it makes sense to do so
- `UPSTREAM: <drop>: ` -- A change that should be discarded during the next rebase cycle
- `UPSTREAM: 1234: ` -- A change carried until the rebase includes upstream PR #1234

Examples:

```
UPSTREAM: <carry>: Add OpenShift-specific e2e test targets
UPSTREAM: <drop>: Pin Go version to 1.25 until 1.26 builder is available
UPSTREAM: 5678: Backport fix for race condition in disruption queue
```

These prefixes are for commit titles, not PR titles.

## Upstream-First Policy

New feature work should be directed to [kubernetes-sigs/karpenter](https://github.com/kubernetes-sigs/karpenter). Downstream-only features are discouraged because they must be re-applied on every rebase. If a downstream-only change is necessary, use the `UPSTREAM: <carry>:` prefix and include a comment in the PR explaining why it cannot go upstream.

## PR Title Convention

PR titles should be prefixed with a Jira ticket reference:

```
AUTOSCALE-123: Fix the whatsit in the thingamajig
OCPBUGS-456: Correct nil pointer in scaler shutdown
NO-JIRA: Update Go module dependencies
```

The Jira prefix goes in the **PR title**. The upstream commit prefix goes in the **commit message**.

## PR Workflow

This repo uses [OpenShift CI (Prow)](https://docs.ci.openshift.org/) for continuous integration. GitHub Actions workflows in this repo are from upstream and are **not used** for our CI. PRs are automatically merged once all required tests pass and the correct labels are present.

### Required labels for merge

- `lgtm`: added by a reviewer via the `/lgtm` command. Any developer from the OpenShift org can add this after reviewing the PR.
- `approved`: added by an approver listed in the [OWNERS](OWNERS) file via the `/approve` command.

### Useful commands

Comment these on the PR:

| Command | Effect |
|---------|--------|
| `/lgtm` | Add the `lgtm` label after reviewing |
| `/lgtm cancel` | Remove the `lgtm` label |
| `/approve` | Add the `approved` label (OWNERS approvers only) |
| `/retest` | Re-run all failed required tests |
| `/retest-required` | Re-run only the failed required tests |
| `/test <test-name>` | Run a specific test, e.g. `/test unit`, `/test verify`, `/test e2e-kwok` |
| `/hold` | Prevent the PR from being merged |
| `/hold cancel` | Remove the hold and allow merging |
| `/verified` | Mark the PR as verified |
| `/cherry-pick release-4.22` | Create a cherry-pick PR to a release branch |

### Preventing premature merges

- Add the `WIP:` prefix to the PR title (e.g., `WIP: AUTOSCALE-123: Work in progress`). Prow adds the `do-not-merge/work-in-progress` label automatically.
- Use `/hold` to temporarily block merging while awaiting additional review or testing.

## Test Expectations

PRs should include tests to verify correctness and prevent future regressions:

- **Unit tests**: Required for new logic, bug fixes, and behavior changes. Run with `make test`. These cover `./pkg/...` with the race detector.
- **E2E tests**: Expected for new features or significant behavior changes. OpenShift CI runs the KWOK regression suite via `hack/openshift-kwok-e2e.sh` (`/test e2e-kwok`). Static capacity tests are currently skipped (`SKIP="StaticCapacity"`).

## Verified Label

Use `/verified` to indicate changes have been verified. Examples:

```none
/verified
```

Mark as verified by referencing specific test coverage:

```none
/verified by unit
/verified by e2e-aws-olm
/verified by E2Es
```

If verification will happen later (e.g., by QE or in a staging environment):

```none
/verified deferred to QE
```

## Generated Code

The following files are generated and should never be hand-edited:

| File(s) | Generator | Regenerate with |
|---------|-----------|-----------------|
| `**/zz_generated.deepcopy.go` | controller-gen via `go:generate` | `go generate ./...` (also run by `make verify`) |
| `pkg/scheduling/zz_generated.deepcopy.go` | controller-gen via `go:generate` | `go generate ./...` (also run by `make verify`) |
| `pkg/apis/crds/*.yaml` | controller-gen, then `hack/validation/*.sh` (yq) | `go generate ./...` then the validation scripts |
| `kwok/charts/crds/karpenter.sh_*.yaml` | Copied from `pkg/apis/crds`, then `hack/kwok/requirements.sh` | `make verify` |
| `kwok/charts/README.md` | helm-docs | `make verify` |
| `.github/dependabot.yaml` (lower section) | `hack/dependabot.sh` | `make verify` |

After modifying API types, run `make verify` (or `make openshift-verify` to match CI) and commit the generated and yq-patched files in the same PR. `go generate` alone is not enough: CRD OpenAPI is patched in place by the validation scripts.

## Development Quick Reference

| Task | Command |
|------|---------|
| Run unit tests | `make test` |
| Run all presubmit checks | `make presubmit` |
| Verify (generate, lint, vendor, git diff) | `make verify` |
| Verify as OpenShift CI does | `make openshift-verify` |
| Install CI toolchain (kubebuilder assets, ko, yq) | `make openshift-toolchain` |
| Format / license headers | handled by `make verify` (`nwa` + goimports) |
| Lint | `go tool -modfile=go.tools.mod golangci-lint-kube-api-linter run` (also part of `make verify`) |
| Vulnerability scan | `make vulncheck` |
| License check | `make licenses` |
| Node overlay memory tests | `make test-memory` |
| DRA KWOK driver unit tests | `make test-dra` |
| Install KWOK in a cluster | `make install-kwok` |
| Deploy KWOK Karpenter on OpenShift | `make apply-with-openshift` |
| Run e2e against a cluster | `TEST_SUITE=regression make e2etests` |
| OpenShift CI e2e (KWOK on a real cluster) | `./hack/openshift-kwok-e2e.sh` |
| Filter e2e tests | `FOCUS="DaemonSet" SKIP="Disruption" TEST_SUITE=regression make e2etests` |

There is no standalone `make generate`, `make fmt`, or `make lint` target. Generation, formatting, and linting are steps inside `make verify`.

## Pre-Submit Checklist

Before requesting review:

1. `make test`: run unit tests
2. `make openshift-verify`: match what OpenShift CI runs (`GOFLAGS=-mod=readonly`)
3. Review your diff for secrets, credentials, or debug code
4. Address any [CodeRabbit](https://coderabbit.ai/) review feedback. Responding with an explanation of why you are not acting on a suggestion is fine. The goal is to resolve straightforward issues so human reviewers can focus on the substantive aspects.

`make verify` / `make openshift-verify` delete `vendor/`, regenerate code, re-vendor, and fail in CI if the working tree is dirty. Run them before you push.

## Code Style

- Follow Go conventions for error strings: lowercase, no trailing punctuation, wrap with `fmt.Errorf("context: %w", err)`
- Use structured logging with logr/klog: constant messages, key-value pairs in lowerCamelCase
- Import ordering: stdlib, external packages, internal packages (separated by blank lines). Local prefix is `sigs.k8s.io/karpenter`
- License headers are managed by `nwa` (see `.nwa-config.yaml`). Do not hand-edit them.

## AI Code Review

Our repos use CodeRabbit for automated AI code review. CodeRabbit will post review comments on your PR automatically.
As a courtesy to the human reviewer who follows, please address CodeRabbit’s feedback before requesting human review. You do not need to accept every suggestion — responding with an explanation of why you are not taking action on a comment is perfectly acceptable. The goal is to resolve straightforward issues so that human reviewers can focus on the substantive aspects of the change.
