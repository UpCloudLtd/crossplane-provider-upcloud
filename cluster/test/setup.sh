#!/usr/bin/env bash
set -aeuo pipefail

echo "Running setup.sh"
echo "Creating cloud credential secret..."
${KUBECTL} -n default create secret generic provider-secret --from-literal=credentials="{\"username\": \"${UPCLOUD_USERNAME}\", \"password\": \"${UPCLOUD_PASSWORD}\"}" --dry-run=client -o yaml | ${KUBECTL} apply -f -

echo "Waiting until provider is healthy..."
${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m

echo "Waiting for all pods to come online..."
${KUBECTL} -n upbound-system wait --for=condition=Available deployment --all --timeout=5m

echo "Creating a default provider config..."
cat <<EOF | ${KUBECTL} apply -f -
apiVersion: provider.upcloud.com/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: provider-secret
      namespace: default
      key: credentials
EOF

echo "Creating configs and secrets for test resources..."
${KUBECTL} apply -f ${RESOURCE_CONFIG}
