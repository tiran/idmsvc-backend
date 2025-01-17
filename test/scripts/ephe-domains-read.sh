#!/bin/bash

function error {
    local err=$?
    printf "%s\n" "$1" >&2
    exit $err
}

UUID="$1"
[ "${UUID}" != "" ] || error "UUID is empty"

export NAMESPACE="$(oc project -q)"
CREDS="$( oc get secrets/env-${NAMESPACE}-keycloak -o jsonpath='{.data.defaultUsername}' | base64 -d )"
CREDS="${CREDS}:$( oc get secrets/env-${NAMESPACE}-keycloak -o jsonpath='{.data.defaultPassword}' | base64 -d )"
export CREDS

unset X_RH_IDENTITY
unset X_RH_FAKE_IDENTITY
BASE_URL="https://$( oc get routes -l app=idmsvc-backend -o jsonpath='{.items[0].spec.host}' )/api/idmsvc/v1"

./scripts/curl.sh -i "${BASE_URL}/domains/${UUID}"
