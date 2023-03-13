# crt.sh Domain Finder

Get all the related domains and subdomains using [crt.sh](https://crt.sh/).

## Installation
To install, use the following command:

```bash
▶ go install github.com/shivamsaraswat/crt_sh_go@latest
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