# Datamkr
A CLI for genrating mock data into postgres and csv.
Datamkr defines data definitions in yaml files that can be checked into source control.

## Table of Contents
1. [Stdout Usage](#stdout-usage)
2. [Postgres Usage](#postgres-usage)
3. [Installation](#installation)

# Stdout Usage
To start using datamkr, first initiate the repo with:
```
datamkr init
```
This creates a few this:

1. `datamkr.yml` a config file
2. `datasets/` a datasets directory to store dataset definitions (can be changed to any path)
3. `datasets/demo.yml` an example dataset definition

Now we can create mock data using the demo defintion:
```
datamkr make demo
```
This prints out 5 rows of data with the columns defined in `datasets/demo.yml`

The verbose version of this command would use the flags `--to stdout` & `--num 5`
```
datamkr make demo --to stdout --num 5
```

All the flags and options can be seen with the `--help` flag or by typing only `datamkr`.

# Postgres Usage
Datamkr supports populating mock data into postgres as well as creating dataset definitions from a postgres table.

To do this, we will first add our postgres instance as a `storage alias` in our config which will let use easily talk to postgres without typing a postgres connection string over and over again.

Edit the config file (`./datamkr.yml`) should look something like below for connecting to a local postgres db named `test` running on port `5432` accessible by db role `postgres`
```
datamkr:
  datasetDir: test/datasets
  storage:
    postgres_local:
      connection: postgresql://postgres@localhost:5432/test?sslmode=disable
      type: postgres
  version: 1
```

### Dataset Creation
Now that we have our postgres storage alias setup we can generate a dataset definition for the table users.
```
datamkr add users --from postgres_local
```

### Mock Data Generation
Now that the dataset is created we can populate the table with 100 records by running the following:
```
datamkr make users -n 100
```

The table should now be populated with mock data.

# Installation
Instructions for Mac OS
## 1. Download Binary

### Intel
```
curl -LO "https://awle-testing.s3.amazonaws.com/datamkr/amd64/datamkr"
```

### Apple Silicon
```
curl -LO "https://awle-testing.s3.amazonaws.com/datamkr/arm64/datamkr"
```

## 2. Make Executable

```
chmod +x ./datamkr
```

## 3. Move to PATH
Move executable to a location on your `PATH`

```
sudo mv ./datamkr /usr/local/bin/datamkr
```

To view and add `/usr/local/bin` to `PATH`
```
echo $PATH
export PATH=$PATH:/usr/local/bin/
```

## 4. Test
It should be successfully installed, check by running
```
datamkr
```