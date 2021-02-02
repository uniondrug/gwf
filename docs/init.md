# Initialize application

```text
cd ~/Desktop && \
go mod init sketch
```

### MOD

```text
module sketch

go 1.15

require (
	github.com/uniondrug/gwf v1.0.0
)
```

### Build Path

> main.go

```text
import (
	"github.com/uniondrug/gwf/util"
)

func main() {
	util.NewPath().Build()
}
```