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
`language.Parse`, so the scan reports a real reachable-candidate finding. Findings are
deterministic candidates, never proofs of exploitability.

## Layout

- `bin/` — the prebuilt scanner (`nucleon`) and its analyzer (`nucleon-plugin`).
- `target/` — the sample Go module that gets scanned (vendored deps; offline).
- `.github/workflows/scan.yml` — runs the binary and publishes the tiers.
