# UK Visa Sponsors CLI Checker

This Command Line Interface (CLI) tool is crafted to fetch the updated list of organisations authorised with a visa license from the UK government's official platform. For direct access, see the [Register of Licensed Sponsors](https://www.gov.uk/government/publications/register-of-licensed-sponsors-workers).

## Prerequisites
- Go >= 1.21 (to run the binary)
- Make (Optional)

## Installation
### Using HomeBrew
You can use brew to install and use this CLI command.

First tap the repository:
```shell
brew tap renatoaraujo/homebrew-renatoaraujo
```

Then update and install, using: 
```shell
brew update && brew install uk-visa-sponsors
```

Now have some fun just running: 
```shell
uk-visa-sponsors find -c [company-name]
```

### From the source
#### Using Makefile
Running the CLI tool is a breeze. Just punch in the following command in your terminal:
```shell
make find [company-name]
```
#### Binary (requires Go installed in the machine)
Build the binary, using:
```shell
go build -o uk-visa-sponsors
```

Run the binary:
```shell
./uk-visa-sponsors find -c [company-name]
```
_You might need to update the permissions `chmod +x ./uk-visa-sponsors`_

## Motivation
While navigating through job listings in the UK, I stumbled upon numerous opportunities, but many were ambiguous about visa sponsorship. This tool was birthed to alleviate such uncertainties, allowing users to promptly determine if a company is an accredited visa license holder.

> **Note:** Though this CLI offers a handy way to confirm a company's visa license status, it's always a good idea to directly refer to the [UK government website](https://www.gov.uk/government/publications/register-of-licensed-sponsors-workers). Some users might find this approach more straightforward and direct. Ensure you're always referencing the most current and precise source when it's pivotal.

#### Additional options
- `-d`, `--datasource`: string containing an extra datasource, for example an old version of the CSV or a sanitised one;
- `-h`, `--help`: To get help with the command

## Credits

* [Renato Araujo](https://www.linkedin.com/in/renatoraraujo/)

## License

The MIT License (MIT) - see [`LICENSE`](LICENSE) for more details
