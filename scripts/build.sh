#!/bin/bash

PROJECT=/home-assessment-kai-cyber
local_import=github.com/hamza-sharif/${PROJECT}

go build -buildvcs=false -o "bin/${PROJECT}-server" "${local_import}/cmd"
