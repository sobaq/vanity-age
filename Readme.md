# vanity-age

Mildly performant [age](https://github.com/FiloSottile/age) vanity public key brute-forcer.

# Usage

```
go build
./vanity-age query
```

Where `query` is a valid `fnmatch` query. Try

```
./vanity-age
```

for information about fnmatch.

The output is identical to `age-keyen`.