# Password Cracker

A password cracker written in Go.

## Crack a Password

Incremental mode (similar to John the Ripper's incremental mode)
- *This is the default mode*
- *The `maxlen` default is 4*

```bash
make run ARGS="-mode incremental -password <password_hash> -maxlen <max_length_to_brute_force>"

// Or...

make run ARGS="-password <password_hash>"
```

Dictionary mode - using a wordlist

```bash
make run ARGS="-mode dictionary -password <password_hash> -wordlist <the_wordlist.txt>"
```

## Run Tests

```bash
make test
```
