# tg
a command line tool to perform tigo based push payment request and disbursement.


## supported features
- [init](#init)
- [push](#push)
- [disburse](#disburse)


## init
This offers a way to configure the tool with tigo integration details

```bash
 tg init --file=config.yaml

 tg init

 tg init push

 tg init disburse

```

## push

make push pay request to a single user

```bash

tg push --config=config-file.yml --file=customers.csv

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
