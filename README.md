# tg
a command line tool to perform tigo based push payment request and disbursement.


## supported features
- [config](#config)
- [push](#push)
- [disburse](#disburse)


## config
This offers a way to configure the tool with tigo integration details

```bash
 tg config --file=config.yaml

 tg config

 tg config push

 tg config disburse

```

## push

make push pay request to a single user

```bash

tg push --phone=0712XXXXXX --amount=10000 --remarks="donation from cmd"

```

make push pay request to multiple users at once

```bash

tg push --file=customers.csv

```


## disburse
pay someone or a number of people at once from command line

```bash

 tg disburse --file=loans.csv

 tg disburse --phone="07123CCCCC" --amount=10000

```
