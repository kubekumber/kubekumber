<p align="center">
  <img alt="Kubekumber Logo" src="https://avatars.githubusercontent.com/u/120406404?s=400&u=3c809a57af65ab8a94f13ed0485aa83ca6f3e346&v=4" height="140" />
</p>

# Kubekumber
Introducing Kubekumber, a good tool for a bad practice. The bad practice being manually having to run a `kubectl` command against hundreds of kube clusters that you or your company manage because some Executive said ["We need to cloud native our containerization strategy"](https://www.reddit.com/r/kubernetes/comments/dtsg4z/dilbert_on_kubernetes/). Behold, Kubekumber!
## Badges
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/kubekumber/kubekumber)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/kubekumber/kubekumber/goreleaser)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/kubekumber/kubekumber)


## What does it do?

Very simply, `kubekumber` allows to to run an arbitrary kubectl command against any clusters that match a `grep`-style regex. For example, say you wanted to find out which of your staging clusters has cert-manager running in the `infra` namespace:

    kubekumber -r "stage" -c "get pods -n infra" --verbose

Or, maybe you need to remove a CustomResourceDefinition from your legacy clusters, with names like `denver:staging` or `frankfurt:production`? Kubekumbers got you covered:

    kubekumber -r "[a-z]*:[a-z]*" -c "delete crd <unholy_crd_from_that_thing_you_tried>"

Think of it as that terrible bash script you use more than you like, but now with a memorable name!

## Dependencies

Kubekumber uses your `kubectl` and `kubectx`, so make sure those are installed.

## Install
### MacOS/Linux with Homebrew
Add the tap:

    brew tap kubekumber/tools

and then install with:

    brew install kubekumber

Or, do it in one line:

    brew install kubekumber/tools/kubekumber

## TODO
- [x] Homebrew
- [ ] Testing
- [ ] Versioning & Author in CLI
- [ ] Moar CI/CD
- [ ] Documentation
- [ ] Website?
- [ ] Medium Article
- [ ] Sponsoring
- [ ] Vendor `kubectx` directly into the app

## Shoulders of Giants, y'all
Big thanks to [John Arundel](https://github.com/bitfield) for their [script](https://github.com/bitfield/script) library, wonderful stuff!

Also thanks to the folks over at `urfave` for the [cli](https://github.com/urfave/cli)! 

Other projects used: `goreleaser`, `task`, `brew`, and all the good stuff at GitHub!