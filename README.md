# Password Cracker

[![go](https://github.com/claudemuller/password-cracker.go/actions/workflows/go.yml/badge.svg)](https://github.com/claudemuller/password-cracker.go/actions/workflows/go.yml)

A password cracker written in Go.

## Crack a Password

### Incremental Mode

- Similar to John the Ripper's incremental mode
- *This is the default mode*
- *The `maxlen` default is 4*

```bash
make run ARGS="-mode incremental -password <password_hash> -maxlen <brute_force_len>"

// Or...

make run ARGS="-password <password_hash>"
```

### Dictionary Mode

- using a wordlist

```bash
make run ARGS="-mode dictionary -password <password_hash> -wordlist <the_wordlist.txt>"
```

### Rainbow Table Mode

- using a Rainbow Table

```bash
make run ARGS="-mode rainbow -password <password_hash> -rainbow <rainbow_table.dat>"
```

#### Generate Rainbow Table

- using a wordlist

```bash
make run ARGS="-mode genRainbow -wordlist <the_wordlist.txt> -out <rainbow_table.dat>"
```

## Run Tests

```bash
make test
```
