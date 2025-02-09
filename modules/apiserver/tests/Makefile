.EXPORT_ALL_VARIABLES:

# private repo for images
IMG_REPO = quay.io/kuberlogic
IMG = $(IMG_REPO)/minio:1.0.0

KUBERLOGIC_APISERVER_LOG = kuberlogic.apiserver.log
KUBERLOGIC_OPERATOR_LOG = kuberlogic.operator.log

KUBERLOGIC_DEBUG_LOGS = True

ARGS ?= -short

INTERNAL_MINIO_IP=$(shell kubectl get pod -l app=minio -o go-template="{{range .items}}{{.status.podIP}}{{end}}")

# access/secret also added in the deployment.yaml
MINIO_ACCESS_KEY = BKIKJAA5BMMU2RHO6IBB
MINIO_SECRET_KEY = V7f1CwQqAcwo80UEIJEjc5gVQUSSx5ohQ9GSrr12
MINIO_ENDPOINT = http://$(INTERNAL_MINIO_IP):9000
BUCKET ?= test

SENTRY_DSN = https://50ec3b92240a4b4d97dbee02a07fc54f@kl.sentry.cloudlinux.com/5

POSTHOG_API_KEY =

KUBERLOGIC_AUTH_PROVIDER = keycloak
KUBERLOGIC_AUTH_KEYCLOAK_CLIENTID = apiserver-client
KUBERLOGIC_AUTH_KEYCLOAK_CLIENT_SECRET = apiserver-client-secret
KUBERLOGIC_AUTH_KEYCLOAK_REALM_NAME = kuberlogic_realm
KUBERLOGIC_AUTH_KEYCLOAK_URL=https://$(shell kubectl get svc keycloak -o jsonpath='{.spec.clusterIP}'):8443

INTERNAL_GRAFANA_IP=$(shell kubectl get pod -l app.kubernetes.io/component=grafana -o go-template="{{range .items}}{{.status.podIP}}{{end}}")

KUBERLOGIC_GRAFANA_ENABLED = true
KUBERLOGIC_GRAFANA_ENDPOINT = http://$(INTERNAL_GRAFANA_IP):3000/


test:
	POD_NAMESPACE=kuberlogic go test -v \
		-coverprofile=coverage.out \
		-coverpkg="github.com/kuberlogic/..." \
		-covermode=atomic \
		-run "$(RUN)" \
		-timeout 60m \
		-failfast \
		$(ARGS) ./...

remote-test:
	REMOTE_HOST=$(shell kubectl get pod -l app.kubernetes.io/instance=kuberlogic-apiserver -o jsonpath='{.items[].status.podIP}'):8001 \
	$(MAKE) test

coverage-report: test
	for pkg in apiserver operator; do \
	  	covfile=coverage-$${pkg}.out; \
		echo "mode: atomic" > $$covfile; \
		cat coverage.out | grep "github.com/kuberlogic/kuberlogic/modules/$${pkg}" >> $$covfile; \
	done; \


html-report: coverage-report
	mkdir -p html/ ;\
	for pkg in apiserver operator; do \
	  	covfile=coverage-$${pkg}.out; \
		go tool cover -html=$$covfile -o=html/$${pkg}.html; \
	done;

test-in-cluster:
	envsubst < config/in-cluster/job.yaml | kubectl apply -f -

watch-test-in-cluster:
	while :; do \
  		kubectl logs -l job-name=integration-tests -f; \
  		sleep 1; \
  	done;

clear-test-in-cluster:
	kubectl delete kls --all --all-namespaces
	kubectl delete klt --all
	kubectl delete job integration-tests

build-minio:
	docker build minio/ \
		-t $(IMG) \
		--build-arg MINIO_ACCESS_KEY=$(MINIO_ACCESS_KEY) \
		--build-arg MINIO_SECRET_KEY=$(MINIO_SECRET_KEY)

push-minio:
	docker push $(IMG)

deploy-minio:
	kubectl apply -f config/minio/deployment.yaml
	kubectl wait pods --for=condition=Ready  -l app=minio --timeout=5m

endpoint-minio:
	@echo $(MINIO_ENDPOINT)

destroy-minio:
	kubectl delete -f minio/deployment.yaml

create-bucket:
	[ -f "./mc" ]  || wget https://dl.min.io/client/mc/release/linux-amd64/mc -O mc && chmod +x mc; \
	./mc alias set minio $(MINIO_ENDPOINT) $(MINIO_ACCESS_KEY) $(MINIO_SECRET_KEY); \
	./mc mb minio/$(BUCKET); \

show-minio-credentials:
	@echo "AWS_ACCESS_KEY_ID:" $(MINIO_ACCESS_KEY)
	@echo "AWS_SECRET_ACCESS_KEY:" $(MINIO_SECRET_KEY)
	@echo "Endpoint:" $(MINIO_ENDPOINT)
	@echo "Default bucket:" $(BUCKET)

patch-endpoint:
	MINIKUBE_IP=$$(minikube ip) envsubst < config/webhook/endpoint.yaml | kubectl apply -f -

generate-local-webhook-certs:
	mkdir -p /tmp/k8s-webhook-server/serving-certs/; \
	kubectl get secret kuberlogic-webhook-cert -o=jsonpath="{.data['tls\.key']}" | base64 -d >  /tmp/k8s-webhook-server/serving-certs/tls.key
	kubectl get secret kuberlogic-webhook-cert -o=jsonpath="{.data['tls\.crt']}" | base64 -d >  /tmp/k8s-webhook-server/serving-certs/tls.crt

undeploy-operator:
	kubectl delete deployment -l app.kubernetes.io/instance=kuberlogic-operator
	kubectl wait pod -l app.kubernetes.io/instance=kuberlogic-operator --for=delete --timeout=5m

configure-logs:
	envsubst < config/logz.io/logzio-daemonset-rbac.yaml | kubectl apply -f -

follow-apiserver-logs:
	kubectl logs -l app.kubernetes.io/instance=kuberlogic-apiserver -f

follow-operator-logs:
	kubectl logs -l app.kubernetes.io/instance=kuberlogic-operator -c manager -f