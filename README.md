# Go Modules Remote Import Path Proxy
Proxy your Go Module\`s Import Path from your own domain to a public host (e.g. github.com).   
For example Uber (built by their own): `go.uber.org/atomic` resolves to the Git Repository `https://github.com/uber-go/atomic.git`.   

Please note this project is still heavily in **work-in-progress**, but you can already deploy it and give it a try.

## Example
```yaml
modules:
  - package: go.example.com/abc
    type: git
    target: https://github.com/StiviiK/abc.git
    sources:
      - https://github.com/StiviiK/abc
      - https://github.com/StiviiK/abc/tree/main{/dir}
      - https://github.com/StiviiK/abc/tree/main{/dir}/{file}#L{line}
```
Now `go.example.com/abc` resolves to `https://github.com/StiviiK/abc.git` and can be used as before using the new namespace.

If you want further Information how the Go Tooling handles this functionality have a look at [the Go Docs](https://pkg.go.dev/cmd/go#hdr-Remote_import_paths).
