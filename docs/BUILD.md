## Building

```bash
git clone https://github.com/patrickhuber/terraform-provider-pivotal-om
cd terrafor-provider-pivotal-om
go build -o terraform-provider-pivotal-om
```

```powershell
git clone https://github.com/patrickhuber/terraform-provider-pivotal-om
cd terrafor-provider-pivotal-om
go build -o terraform-provider-pivotal-om.exe
```

### Requirements

* Requires go 1.11 (for go module support)
* Set the GO111MODULE environment variable
  - bash: `export GO111MODULE=on`
  - powershell: `$env:GO111MODULE="on"`
  - optional: set bash profile or global windows environment variable