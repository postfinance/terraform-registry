# 

# Variables for acceptance tests
ARTIFACTORY_BASE_URL=
ARTIFACTORY_USERNAME=
ARTIFACTORY_PASSWORD=

# acc-sweep: sweep objects used for acceptance tests
acc-sweep:
	TF_ACC=true go test -v -count=1 ./aixboms/ -sweep=aixboms

# acc-test: start acceptance tests
acc-test: acc-sweep
	TF_ACC=true go test -v -timeout=30m -count=1 ./... 2>&1 | tee local/acc-test.log

# acc-test-custom: start particular acceptance tests
acc-test-custom: acc-sweep
	TF_ACC=true go test -v -timeout=30m -count=1 -run=TestAccNode_create_vip ./... 2>&1 | tee local/acc-test.log

.PHONY: acc-sweep acc-test acc-test-custom
