# semver-generator

Github Action and golang binary for generating a semver version

## Requirements

- Golang
- Docker (optional)

## Usage

```shell
# build it
docker build --progress plain -t semver .

# run it with arguments
docker run --rm semver --input 0.1.2 --bump patch

# run it with environment variables
docker run --rm -e SEMVER_GENERATOR_INPUT=0.1.2 -e SEMVER_GENERATOR_BUMP=patch semver
```

## Releasing

Run the `bump-version` Github Actions workflow with the appropriate `bump type`. A docker image will be generated and pushed into GHCR, while the built binaries will be attached to the Github Release associated with the generated tag.
