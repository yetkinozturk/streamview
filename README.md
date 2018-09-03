# streamview
Convert UDP Streams to Configurable HTTP Presentation

:construction: :warning:	 :construction: :warning: Development/Prototyping phase. Expect dragons.

# Build:
cd examples
  
make deps

make

# Example Usage:
cd examples

./app 7778 8000

# See tests:

tests/push100K.go

# Connect via netcat

nc -u localhost 7778


# TODO

* Stats
* Configurable UDP Package Serializations

![(LITL)](diagram.svg)

# Motivation

Almost every device supports UDP.  Implementing complex protocols on devices is extremely expensive and difficult.  With this kind of function as a service approach, external presentation layer can be used by this devices. 

# Use cases
* Displaying sensor data on web or mobile devices.
* Displaying realtime data on web pages.

