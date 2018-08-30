# streamview
Convert UDP Streams to Configurable HTTP Presentation

:construction: :warning:	 :construction: :warning: Development/Prototyping phase.

# Build:
  make deps
  
  make

# Example Usage:
./streamview 7778 8080

# Connect via netcat

nc -l -u localhost 7778

# TODO

* Stats
* Configurable UDP Package Serializations

![(LITL)](diagram.svg)

# Motivation

Almost every device supports UDP.  Implementing complex protocols on devices is extremely expensive and difficult.  With this kind of function as service approach, external presentation layer can be used by this devices. 

# Use cases
* Displaying sensor data on web or mobile devices.
* Displaying realtime data on web pages.

