# Logo Maker


## Example usage

```bash
logo-maker -name "Example App" -path "./result"
```

Installation

```bash
go install github.com/tomek7667/Logo-Maker/v3@latest
```



## Development

in order to build for all systems uncomment in `.goreleaser.yaml` the Darwin position.

To check the `yaml` config run:

```bash
goreleaser check
```

To local-only release:
```bash
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
