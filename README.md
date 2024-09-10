# Logo Maker


## Example usage



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
# on windows:
$env:GITHUB_TOKEN = 'ghp_xxxxx'; goreleaser r --clean
```



### Debugging

```bash
go run main -name "Essa Bessa" -debug -path "./test"
```
