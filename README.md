# git-events-operator

This is an operator (or more of a framework) that allows users to build their own specific flavor
of an operator and deploy to Kubernetes.

The framework makes it easy to trigger on arbitrary events and map them to arbitrary actions.

In the case of our first example we map a merge to master, to a rebrandly link

# Building

```bash
make container push
```

More information

```bash
make help
```

# Brokers

### GitHub

Required environmental variables

```
export GITHUB_USER="myuser"
export GITHUB_PASS="mypass"
```

# Actions

### Rebrandly

```
export REBRANDLY_API_KEY="mykey"
```

### Sparkpost

```
export SPARKPOST_API_KEY="mykey"
```