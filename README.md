# GeekTrust Make Space

This repository contains solution to [Make Space](https://www.geektrust.com/coding/detailed/make-space)
problem on [geektrust.com](https://www.geektrust.com/).

> Note:
>
> - The solution intentionally contains one bug (It is fixed in the one that I submitted).
> - The directory structure is not as per the geektrust guidelines.

## Requirements

**go@1.20**

## How to build ?

Run command below to build the code from the project directory.

```sh
go build .
```

## How to run tests?

Run command below to run tests from the project directory.

```sh
go test ./...
```

## How to run application ?

Once you build the code you will find executable file named 'geektrust' in the current directory.

> Sample inputs
>
> Two sample inputs are provided in the [sample_input](./sample_input/) directory.

Run following command in the current directory to run the code with sample inputs.

```sh
./geektrust ./sample_input/input1.txt
```
