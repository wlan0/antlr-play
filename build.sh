#!/bin/bash

java -Xmx500M -cp "/usr/local/lib/antlr-4.7.1-complete.jar:$CLASSPATH" org.antlr.v4.Tool -visitor -Dlanguage=Go $@ -o parser/

go build github.com/wlan0/antlr-play
