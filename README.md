# crt.sh Domain Finder

Get all the related domains and subdomains using [crt.sh](https://crt.sh/).

## Installation
If you have Go installed and configured (i.e. with $GOPATH/bin in your $PATH):

```bash
▶ go get -u github.com/shivamsaraswat/crt_sh_go
```

# Using the crt.sh Domain Finder
To run the crt.sh Domain Finder on a domain, use the '-d' flag and provide the domain as an argument:
```bash
▶ ./crt_sh_go -d example.com
```

For an overview of all commands use the following command:

```bash
Usage of ./crt_sh_go:
  -d string
        Domain to search for domains
  -o string
        Output file to write domains to
```