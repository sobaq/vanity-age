# vanity-age

Mildly performant [age](https://github.com/FiloSottile/age) vanity public key brute-forcer.

## Installation and Usage
First clone this repository  
<!-- `git clone https://github.com/seaofmars/vanity-age`  -->
### Docker

`docker build -t vanity-age .`  
`docker run -it --rm -d -v $(pwd)/key.txt:/key.txt --name vanity-age vanity-age "query"`  

### Go  

```
go build
./vanity-age query
```

### Queries  
`query` is a valid `fnmatch` query. Try

```
./vanity-age
```

for information about fnmatch.

To simply generate a key which starts with specific characters, try:  
`./vanity-age "<0-51 alphanumeric characters>*`  

The output is identical to `age-keygen`.
