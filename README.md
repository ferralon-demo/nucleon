# nucleon

A demo of the Ferralon Assay — a free, deterministic dependency-advisory scan. This
repository runs the pinned [`ferralon-ai/ferralon-assay`](https://github.com/ferralon-ai/ferralon-assay)
Action against a small sample Go service on every push and pull request, and publishes the
results across GitHub's surfaces. The Action fetches a versioned scanner release at run
time and builds nothing — no scanner binary is committed here.

## What it produces

| Artifact | Where |
|---|---|
| `report.json`, `openvex.json`, `report.sarif.json`, `report.html` | the `scan-report` build artifact |
| Job summary + inline annotations | the Actions run page |
| SARIF findings | **Security ▸ Code scanning** |
| Pinned Issue dashboard | **Issues** |
| Sticky comment | on each pull request |

The sample target (`target/`) depends on `golang.org/x/text v0.3.6` and reaches
`language.Parse`, so the scan reports a real reachable-candidate finding. Each candidate
carries a graded reachability label (`attacker_tainted` ranks above `control_flow_only`),
the structured entry point and call path that reach the vulnerable symbol, and EPSS/KEV
exploitation-likelihood context (CISA KEV listing + EPSS score/percentile) used to rank the
findings. Grade and EPSS/KEV are evidence strength and in-the-wild context — not proofs of
exploitability. Findings are deterministic candidates, never proofs.

On the default branch the run persists a StateStore baseline to `refs/tegron/state`; pull
request runs stay stateless and never overwrite it.

## Layout

- `target/` — the sample Go module that gets scanned (vendored deps; offline).
- `.github/workflows/ferralon-assay.yml` — runs the pinned assay Action and publishes the surfaces.
