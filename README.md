# Multilinebeat

Welcome to Multilinebeat.

Multilinebeat is use for merge multiline from client.Some time ,We need a way to merger log
in server not in client because the client is too heavy to run Regx to merge multiline log.

Multilinebeat accept a json format like this

    {
        "message":"xxxx",
        "timestamp":"xxxx"
    }

we can config multilinebeat group message By `groupkey`,and it will test if the line the json
 field where we config in `messagefiledkey`is match the `multilineregx`,if match,We will merge
  the multiline log and send it

Ensure that this folder is at the following location:
`${GOPATH}/github.com/kira8565`

## Getting Started with Multilinebeat

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Init Project
To get running with Multilinebeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Multilinebeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/kira8565/multilinebeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Multilinebeat run the command below. This will generate a binary
in the same directory with the name multilinebeat.

```
make
```


### Run

To run Multilinebeat with debugging output enabled, run:

```
./multilinebeat -c multilinebeat.yml -e -d "*"
```


### Test

To test Multilinebeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `etc/fields.yml`.
To generate etc/multilinebeat.template.json and etc/multilinebeat.asciidoc

```
make update
```


### Cleanup

To clean  Multilinebeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Multilinebeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/github.com/kira8565
cd ${GOPATH}/github.com/kira8565
git clone https://github.com/kira8565/multilinebeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make package
```

This will fetch and create all images required for the build process. The hole process to finish can take several minutes.
