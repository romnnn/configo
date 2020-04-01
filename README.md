## configo

[![Build Status](https://travis-ci.com/romnnn/configo.svg?branch=master)](https://travis-ci.com/romnnn/configo)
[![GitHub](https://img.shields.io/github/license/romnnn/configo)](https://github.com/romnnn/configo)
[![GoDoc](https://godoc.org/github.com/romnnn/configo?status.svg)](https://godoc.org/github.com/romnnn/configo)
[![Test Coverage](https://codecov.io/gh/romnnn/configo/branch/master/graph/badge.svg)](https://codecov.io/gh/romnnn/configo)
[![Release](https://img.shields.io/github/release/romnnn/configo)](https://github.com/romnnn/configo/releases/latest)

Dead simple handling of config structs containing optional values for your golang app!


#### Usage

```golang
import (
	"fmt"
	opt "github.com/romnnn/configo"
)

type myAppConfig struct {
	EnableExperimental *opt.Flag
}

var defaultConfig myAppConfig = myAppConfig{
	EnableExperimental: opt.Set(false),
}

func main() {
	userConfig := myAppConfig{
		EnableExperimental: opt.Set(true),
	}
	// Merge user config with the default config
	// Will only set values that have NOT been set
	// Otherwise, use OverrideConfig
	opt.MergeConfig(&userConfig, defaultConfig)
	didEnableExperimental := opt.Enabled(userConfig.EnableExperimental)
	fmt.Printf("EnableExperimental=%t\n", didEnableExperimental) // true
}
```

For more examples, see `examples/`.


#### Development

######  Prerequisites

Before you get started, make sure you have installed the following tools::

    $ python3 -m pip install -U cookiecutter>=1.4.0
    $ python3 -m pip install pre-commit bump2version invoke ruamel.yaml halo
    $ go get -u golang.org/x/tools/cmd/goimports
    $ go get -u golang.org/x/lint/golint
    $ go get -u github.com/fzipp/gocyclo
    $ go get -u github.com/mitchellh/gox  # if you want to test building on different architectures

**Remember**: To be able to excecute the tools downloaded with `go get`, 
make sure to include `$GOPATH/bin` in your `$PATH`.
If `echo $GOPATH` does not give you a path make sure to run
(`export GOPATH="$HOME/go"` to set it). In order for your changes to persist, 
do not forget to add these to your shells `.bashrc`.

With the tools in place, it is strongly advised to install the git commit hooks to make sure checks are passing in CI:
```bash
invoke install-hooks
```

You can check if all checks pass at any time:
```bash
invoke pre-commit
```

Note for Maintainers: After merging changes, tag your commits with a new version and push to GitHub to create a release:
```bash
bump2version (major | minor | patch)
git push --follow-tags
```

#### Note

This project is still in the alpha stage and should not be considered production ready.
