#!/bin/bash

if [ $1 == 1 ]
then
echo
$GOPATH/src/github.com/soonkuk/stellar-brave-network/cli/bin/brave account balance --address $ADDRESS1
echo
$GOPATH/src/github.com/soonkuk/stellar-brave-network/cli/bin/brave account balance --address $ADDRESS2
echo
$GOPATH/src/github.com/soonkuk/stellar-brave-network/cli/bin/brave account balance --address $ADDRESS3
echo
elif [ $1 == 2 ]
then
echo
$GOPATH/src/github.com/soonkuk/stellar-brave-network/cli/bin/brave account balance2 --address $ADDRESS1
echo
$GOPATH/src/github.com/soonkuk/stellar-brave-network/cli/bin/brave account balance2 --address $ADDRESS2
echo
$GOPATH/src/github.com/soonkuk/stellar-brave-network/cli/bin/brave account balance2 --address $ADDRESS3
echo
elif [ $1 == 3 ]
then
echo
$GOPATH/src/github.com/soonkuk/stellar-brave-network/cli/bin/brave account balance3 --address $ADDRESS1
echo
$GOPATH/src/github.com/soonkuk/stellar-brave-network/cli/bin/brave account balance3 --address $ADDRESS2
echo
$GOPATH/src/github.com/soonkuk/stellar-brave-network/cli/bin/brave account balance3 --address $ADDRESS3
echo
fi
