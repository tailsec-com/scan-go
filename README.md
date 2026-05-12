# Tailsec Go SDK

```bash
go get github.com/tailsec-com/scan-go
```

```go
client := scan.NewClient("k8s")
result, err := client.Scan("./manifests")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Found %d issues\n", result.Summary.Total)
```