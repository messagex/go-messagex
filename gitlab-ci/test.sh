#!/bin/bash

printf "\Testing Branch:: $CI_COMMIT_REF_NAME"
printf "\n===================================\n"

printf "\n Tests triggered by : $CI_COMMIT_MESSAGE"
printf "\n=========================================\n"

printf "\n=============================\n"
printf "\nReplacing config variables"
printf "\n==========================\n"

go get -u ./...
go mod vendor

printf "\n===========================\n"
printf "\nRunning Tests"
printf "\n===========================\n"

/bin/bash dev-bin/test
EXIT_CODE=$?

exit $EXIT_CODE
