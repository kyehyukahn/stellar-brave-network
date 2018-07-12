#!/bin/sh

$GOPATH/src/github.com/soonkuk/stellar-brave-network/cli/bin/brave transaction fund --seed $ROOT_SEED --address $ADDRESS1 --amount 10240
sleep 5s
$GOPATH/src/github.com/soonkuk/stellar-brave-network/cli/bin/brave transaction fund --seed $ROOT_SEED --address $ADDRESS2 --amount 10240
sleep 5s
$GOPATH/src/github.com/soonkuk/stellar-brave-network/cli/bin/brave transaction fund --seed $ROOT_SEED --address $ADDRESS3 --amount 10240
