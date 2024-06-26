# semver-generator

Github Action and golang binary for generating a semver version

## Requirements

- Golang
- Docker (optional)

## Usage

```shell
# build it
docker build --progress plain -t semver-generator .

# run it with arguments
docker run --rm semver-generator --input 0.1.2 --bump patch
```

If executed in a Github Actions environment, a Github Output of `version` will be written.

## Releasing

Run the `bump-version` Github Actions workflow with the appropriate `bump type`. A docker image will be generated and pushed into GHCR, while the built binaries will be attached to the Github Release associated with the generated tag.
