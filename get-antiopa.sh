#!/bin/bash
set -e

# Antiopa installer
#
# Usage:
# $ TOKEN=$(cat); curl -fLs -H "PRIVATE-TOKEN: $TOKEN"  https://github.com/deckhouse/deckhouse/raw/stable/get-antiopa.sh | bash -s -- --token $TOKEN
# Напишите token, нажмите `<Enter>` и `<Ctrl-D>`
#
# Подробную инструкцию по установке см. в knowledge base: https://fox.flant.com/docs/kb/blob/master/rfc/rfc-antiopa.md
#

main() {
  NAMESPACE='antiopa'
  IMAGE_REGISTRY='registry.flant.com'
  IMAGE_REPO='sys/antiopa'
  VERSION='stable'
  TOKEN=''
  PROJECT=''
  CLUSTER_NAME='main'
  CLUSTER_HOSTNAME=''
  DRY_RUN=0
  OUT_FILE=''
  LOG_LEVEL='DEBUG'

  parse_args "$@" || (usage && exit 1)

  if [[ $IMAGE_REGISTRY == ":minikube" ]]; then
    IMAGE_REGISTRY="localhost:5000"
  fi

  MANIFESTS=
  generate_yaml

  if [[ -n $OUT_FILE ]]; then
    echo "$MANIFESTS" > $OUT_FILE
  fi

  if [[ $DRY_RUN == 1 ]]; then
    if [[ -z $OUT_FILE ]]; then
      echo "$MANIFESTS"
    fi
  else
    install_yaml
  fi

  return $?
}


usage() {
printf " Usage: $0 --token <gitlab user auth token> [--dry-run]

    --version <version>
            Use specified antiopa image version. Notice recommendations:
              - master - for dev clusters
              - ea - early access, for production clusters with active work
              - stable - for production clusters
              - YYYY-MM-DD.N - images of fixed date - rejected and abondoned clusters, frozen on some version forever

            Default: stable

    --dev
            Use dev image (sys/antiopa/dev, instead of sys/antiopa).
            Name of a dev branch could be passed using --version option.

    --token <token>
            Auth token generated in gitlab user's profile.
            If no token specified, no imagePullSecret will generate.
            E.g. registry in minikube can be without auth (dapp kube minikube setup).

    --project <project>
            Name of the project, same as in bush.flant.com

    --cluster-name <name>
            Name of the cluster, same as in ***REMOVED***_registry.

    --cluster-hostname <hostname>
            This hostname will be used as a suffix for other system componentes (prometheus, dashboard, etc).

    -o, --out <filename>
            Put generated yaml into file.
            To disable installed modules - you should edit Antiopa's ConfigMaps.

    --dry-run
            Do not run kubectl apply.
            Print yaml to stdout or to -o file.

    --log-level <INFO|ERROR|DEBUG>
            Set RLOG_LOG_LEVEL.
            Default: DEBUG

    --help|-h
            Print this message.

"
}

parse_args() {
  while [ $# -gt 0 ]; do
    case "$1" in
      --version)
        VERSION="$2"
        shift
        ;;
      --dev)
        IMAGE_REPO="$IMAGE_REPO/dev"
        ;;
      --token)
        TOKEN="$2"
        shift
        ;;
      --project)
        PROJECT="$2"
        shift
        ;;
      --cluster-name)
        CLUSTER_NAME="$2"
        shift
        ;;
      --cluster-hostname)
        CLUSTER_HOSTNAME="$2"
        shift
        ;;
      -o|--out)
        OUT_FILE="$2"
        shift
        ;;
      --dry-run)
        DRY_RUN=1
        ;;
      --log-level)
        LOG_LEVEL="$2"
        shift
        ;;
      --help|-h)
        return 1
        ;;
      --*)
        echo "Illegal option $1"
        return 1
        ;;
    esac
    shift $(( $# > 0 ? 1 : 0 ))
  done

  if [[ "x$TOKEN" == "x" ]] ; then
    echo "--token required"
    return 1
  fi

  if [[ "x$PROJECT" == "x" ]] ; then
    echo "--project required"
    return 1
  fi
}

generate_yaml() {
  local SECRET
  local IMAGE_PULL_SECRETS

  if [ "x$TOKEN" != "x" ]
  then

    local AUTH_CFG_BASE64=$(cat <<- JSON | base64 -w0
{"$IMAGE_REGISTRY": {
  "username": "oauth2",
  "password": "${TOKEN}",
  "auth": "$(echo -n "oauth2:${TOKEN}" | base64 -w0)",
  "email": "some@email.com"
 }
}
JSON
)

    SECRET=$(cat <<- YAML
---
apiVersion: v1
kind: Secret
type: kubernetes.io/dockercfg
metadata:
  name: registrysecret
data:
  .dockercfg: ${AUTH_CFG_BASE64}
YAML
)
    IMAGE_PULL_SECRETS=$(cat <<- YAML
      imagePullSecrets:
        - name: registrysecret
YAML
)

  fi


  local DEPLOYMENT=$(cat <<- YAML
---
apiVersion: v1
kind: Secret
metadata:
  name: antiopa-secret
data:
  gitlab-token: $(echo -n ${TOKEN} | base64 -w0)
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: antiopa
spec:
  replicas: 1
  template:
    metadata:
      labels:
        service: antiopa
    spec:
      containers:
        - name: antiopa
          image: ${IMAGE_REGISTRY}/${IMAGE_REPO}:${VERSION}
          imagePullPolicy: Always
          command: ["/antiopa/antiopa"]
          resources:
            limits:
              # Важно!!!! Изменять синхронно с policy.cluster.antiopa
              cpu: 420m
              memory: 500Mi
          workingDir: /antiopa
          env:
            - name: KUBERNETES_DEPLOYED
              value: "$(date --rfc-3339=seconds)"
            - name: RLOG_LOG_LEVEL
              value: ${LOG_LEVEL}
            - name: GITLAB_TOKEN
              valueFrom:
                secretKeyRef:
                  name: antiopa-secret
                  key: gitlab-token
      serviceAccount: antiopa
${IMAGE_PULL_SECRETS}
YAML
)

  local RBAC=$(cat <<- YAML
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: antiopa
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: Role
metadata:
  name: antiopa
rules: []
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: RoleBinding
metadata:
  name: antiopa
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: antiopa
subjects:
  - kind: ServiceAccount
    name: antiopa
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: cluster-admin
rules:
- apiGroups:
    - "*"
  resources:
    - "*"
  verbs:
    - "*"
- nonResourceURLs:
    - "*"
  verbs:
    - "*"
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: ${NAMESPACE}
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
  - kind: ServiceAccount
    name: antiopa
    namespace: ${NAMESPACE}
YAML
)

  local VALUES_CONFIG_MAP=$(cat <<- YAML
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: antiopa
data:
  help: |
    Add values key to define global values yaml
    Add <module>-values key to define values yaml for module
    Add disable-modules to specify disabled modules (comma separated, may be globs), for example "disable-modules: test*, kube-dashboard"
  values: |
    project: "${PROJECT}"
    clusterName: "${CLUSTER_NAME}"
YAML
)
  if [[ "x$CLUSTER_HOSTNAME" != "x" ]] ; then
    VALUES_CONFIG_MAP="$VALUES_CONFIG_MAP"$(cat <<- YAML

    clusterHostname: "${CLUSTER_HOSTNAME}"
YAML
)
  fi

  MANIFESTS=$(cat << YAML
$SECRET
$DEPLOYMENT
$RBAC
$VALUES_CONFIG_MAP
YAML
)

  return 0

}

install_yaml() {
  if [ $DRY_RUN == 0 ]; then

    echo Begin install

    if [[ "$(kubectl get ns -a 2>/dev/null | cut -d' ' -f1 | grep "^$NAMESPACE\$")" == "" ]] ; then
        echo "  " Create namespace $NAMESPACE
        kubectl create ns $NAMESPACE
    fi

#    if [[ "$(kubectl -n $NAMESPACE get pod -a 2>/dev/null | cut -d' ' -f1 | grep "^deploy\$")" != "" ]] ; then
#        echo "  " Delete Install manifests...
#        kubectl -n $NAMESPACE delete pod deploy
#
#        while [[ "$(kubectl -n $NAMESPACE get pod -a 2>/dev/null | cut -d' ' -f1 | grep "^deploy\$")" != "" ]] ; do
#            sleep 1
#        done
#    fi

    echo "  " Apply manifests
    echo "$MANIFESTS" | kubectl -n $NAMESPACE apply -f -

    echo End install
  fi
  return $?
}

# wait for full file download if executed as
# $ curl | sh
main "$@"
