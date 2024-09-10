# Logo Maker


## Development

in order to build for all systems uncomment in `.goreleaser.yaml` the Darwin position.

To check the `yaml` config run:

```bash
goreleaser check
```

To local-only release:
```
goreleaser release --snapshot --clean
```

To release to GitHub:

```bash
git tag -a v0.1.0 -m "First release"
git push origin v0.1.0
goreleaser release
```
