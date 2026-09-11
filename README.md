# nucleon

A deterministic dependency-advisory scanner. This repository runs the prebuilt `nucleon`
binary against a small sample Go service on every push and pull request, and publishes the
results across GitHub's surfaces.

## What it produces

| Artifact | Where |
|---|---|
| `report.json`, `openvex.json`, `report.sarif.json`, `report.html` | the `scan-report` build artifact |
| Job summary + inline annotations | the Actions run page |
| SARIF findings | **Security ▸ Code scanning** |
| Pinned Issue dashboard | **Issues** |
| Sticky comment | on each pull request |
| Static HTML report | **GitHub Pages** (push to `main`) |

The sample target (`target/`) depends on `golang.org/x/text v0.3.6` and reaches
`language.Parse`, so the scan reports a real reachable-candidate finding. Each candidate
carries a graded reachability label (`attacker_tainted` ranks above `control_flow_only`),
the structured entry point and call path that reach the vulnerable symbol, and EPSS/KEV
exploitation-likelihood context (CISA KEV listing + EPSS score/percentile) used to rank the
findings. Grade and EPSS/KEV are evidence strength and in-the-wild context — not proofs of
exploitability. Findings are deterministic candidates, never proofs.

## Layout

- `bin/` — the prebuilt scanner (`nucleon`) and its analyzer (`nucleon-plugin`).
- `target/` — the sample Go module that gets scanned (vendored deps; offline).
- `.github/workflows/scan.yml` — runs the binary and publishes the tiers.

<!-- demo PR to exercise the sticky comment surface -->
