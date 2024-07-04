# Password Cracker

[![go](https://github.com/claudemuller/password-cracker.go/actions/workflows/go.yml/badge.svg)](https://github.com/claudemuller/password-cracker.go/actions/workflows/go.yml)

A password cracker written in Go.

## Crack a Password

### Incremental mode

- Similar to John the Ripper's incremental mode
- *This is the default mode*
- *The `maxlen` default is 4*

```bash
make run ARGS="-mode incremental -password <password_hash> -maxlen <brute_force_len>"

// Or...

make run ARGS="-password <password_hash>"
```

### Dictionary mode

- using a wordlist

```bash
make run ARGS="-mode dictionary -password <password_hash> -wordlist <the_wordlist.txt>"
```

## Run Tests

```bash
make test
```
