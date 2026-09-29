# Huma API and OpenAPI compatibility check

This is a tiny Huma API with two endpoints:

- `GET /greeting?name=Magnus`
- `GET /health`

Run the API with `go run .` (listens on `:8888`) or print its generated OpenAPI document with `go run . openapi`.

## Check for breaking changes

`openapi/baseline.yaml` is the contract to preserve. Run:

```sh
./scripts/check-breaking.sh
```

The script generates the current spec from Huma and runs `oasdiff breaking` against the baseline. If `oasdiff` is already installed, it uses that; otherwise it runs the pinned `v1.20.0` release through `go run` (which needs network access the first time). A breaking change exits non-zero. To compare against another baseline, pass its path as the first argument.

When an API change is intentional, review the generated spec before updating the baseline:

```sh
go run . openapi > openapi/baseline.yaml
```

Then commit the baseline alongside the API change so future checks use the new contract.
