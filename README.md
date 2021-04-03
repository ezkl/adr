# ADR CLI
Originally forked from https://github.com/marouni/adr

A minimalist command line tool written in Go to work with [ADRs][adr-og].


# Install
```bash
go get -u github.com/ezkl/adr/cmd/...
```

# Usage
## Initialize
Before creating any new ADR you need to choose a folder that will host your ADRs
and use the `init` sub-command to initialize the configuration :

```bash
adr init /home/user/my_adrs
```

## Creating a new ADR

```bash
adr new New feature
```

This will create a new numbered ADR in your ADR folder: `xxx-new-feature.md`.

[adr-og]: http://thinkrelevance.com/blog/2011/11/15/documenting-architecture-decisions
