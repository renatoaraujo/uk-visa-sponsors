# UK Visa Sponsors CLI Checker

This Command Line Interface (CLI) tool is crafted to fetch the updated list of organisations authorised with a visa license from the UK government's official platform. For direct access, see the [Register of Licensed Sponsors](https://www.gov.uk/government/publications/register-of-licensed-sponsors-workers).

## Motivation
While navigating through job listings in the UK, I stumbled upon numerous opportunities, but many were ambiguous about visa sponsorship. This tool was birthed to alleviate such uncertainties, allowing users to promptly determine if a company is an accredited visa license holder.

> **Note:** Though this CLI offers a handy way to confirm a company's visa license status, it's always a good idea to directly refer to the [UK government website](https://www.gov.uk/government/publications/register-of-licensed-sponsors-workers). Some users might find this approach more straightforward and direct. Ensure you're always referencing the most current and precise source when it's pivotal.

## Prerequisites

- Go 1.21.x
- Make (Optional)

## How to Use

Running the CLI tool is a breeze. Just punch in the following command in your terminal:

```bash
make find [company-name]
```

If you prefer not to employ `make`, no worries! Construct and execute directly from the binary:

```bash
go build -o sponsors . && chmod +x ./sponsors && ./sponsors find -c [company-name]
```
