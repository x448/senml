# Disco SenML - Debloated cisco/senml that passes tests

## SenML
[RFC 8428 Sensor Measurement Lists (SenML)](https://tools.ietf.org/html/rfc8428) defines a format for representing simple sensor measurements and device parameters in Sensor Measurement Lists.  

## Disco SenML
Disco SenML is a debloated fork of cisco/senml that produces programs 4 MB smaller. Other benefits include having all unit tests pass, resolving several important open issues, and having a smaller attack surface.  

Disco SenML was created on Nov 23, 2019 using cisco/senml (4d43ea8) dated Oct 10, 2019.

__Special thanks to Cullen Jennings__ at Cisco Systems for his work on RFC 8428 and creating cisco/senml. He did the heavy lifting and changes to his project were trivial by comparison. __And also to Faye Amacker__ for adding requested features to her CBOR library that made the changes easy.

## Improvements to cisco/senml
Primary improvements:
* __Reduce bloat__: compiled programs are 4 MB smaller (senmlCat and senmlServer).
* __Pass tests__:  [cisco/senml](https://github.com/cisco/senml) wasn't passing unit tests before any changes were made. Broken tests were fixed and a new test was added using the CBOR example from RFC 8428.
* __Improve safety__:  replaced a multi-codec library with a [CBOR library](https://github.com/fxamacker/cbor) (fxamacker/cbor) which properly handles tiny malicious CBOR messages.  Attack surface was also reduced by removing MessagePack for reasons stated below.
* __Close issues__: several important open issues at cisco/senml were resolved in this project.

Some resolved issues include:
* __CBOR does not encode or decode numeric field names__ cisco/senml #2 (2016).  SenML RFC 8428 requires this.
* __Base Value and Base Sum missing from the model__ cisco/senml #18 (2017).
* __CBOR support uses `go-codec` which adds bloat to the binary__ cisco/senml #22 (2019).

Other improvements:
* __Versioned__:  release versions to make using this library simpler for Go projects.
* __Go modules__: use Go modules. See [Using Go Modules](https://blog.golang.org/using-go-modules) at blog.golang.org for more info.

__WARNING: ONLY MINIMAL CHANGES WERE MADE__
* __Tests and Code Coverage__:  there's a lot of room for improvement here.
* __senmlCat and senmlServer__: these programs were only modified to remove MessagePack representation. They probably need other improvements I don't have time to make.

MessagePack was removed because it:
* increased bloat and attack surface.
* isn't mentioned in SenML RFC 8428.
* prevented having a [CBOR library](https://github.com/fxamacker/cbor) (fxamacker/cbor) as the only external dependency.

## Limitations and Requirements
Known limitations:
* __Go 1.12__: Go 1.12 or newer is required.
* __Security Audit__: I didn't conduct a security audit of cisco/senml or this project.  I just tackled the most obvious low-hanging fruit while helping fxamacker/cbor find a non-hobby project to use for a program size comparison.

Possible limitation (not sure if it matters yet):
* __CBOR Tags (major type 6)__:  this might be a non-issue because I didn't see any used in cisco/senml.  If CBOR tags are present in SenML, they'll be ignored until this project upgrades fxamacker/cbor v1.3.1 to v1.4 (when released).

# senmlCat
Tool to convert SenML between formats and act as gateway server to other services

# usage

## convert JSON SenML to XML 
senmlCat -json -i data.json > data.xml

## convert JSON SenML to CBOR
senmlCat.go -ijson -cbor data.json > data.cbor 

## convert to Excel spreadsheet CSV file
senmlCat -expand -ijsons -csv -print data.json > foo.csv

Note that this moves times to excel times that are days since 1900

## listen for posts of SenML in JSON and send to influxdb

This listens on port 880 then writes to an influx instance at localhost where to
the database called "junk"

The -expand is needed to expand base values into each line of the Line Protocol

senmlCat -ijsons -http 8880 -expand -linp -print -post http://localhost:8086/write?db=junk


