.PHONY: test allure report

test:
	rm -rf allure-results
	export ALLURE_OUTPUT_PATH="/Users/maximhalitov/Desktop/IU7/bmstu_test/backend" && go test ./... --race --parallel 10
	cp environment.properties allure-results

allure:
	[ -d allure-reports ] && cp -rf allure-reports/history allure-results || :
	rm -rf allure-reports
	allure generate allure-results -o allure-reports
	allure serve allure-results -p 4000

report: test allure

ci-unit:
	export ALLURE_OUTPUT_PATH="${GITHUB_WORKSPACE}" && \
	export ALLURE_OUTPUT_FOLDER="unit-allure" && \
 	go test -tags=unit_test ./... --race --parallel 10

ci-integration:
	export ALLURE_OUTPUT_PATH="${GITHUB_WORKSPACE}" && \
	export ALLURE_OUTPUT_FOLDER="integration-allure" && \
	go test -tags=integration_test ./... --race --parallel 10

ci-e2e:
	export ALLURE_OUTPUT_PATH="${GITHUB_WORKSPACE}" && \
	export ALLURE_OUTPUT_FOLDER="e2e-allure" && \
	go test -tags=e2e_test ./... --race --parallel 10

integration-tests:
	export ALLURE_OUTPUT_PATH="/Users/maximhalitov/Desktop/IU7/bmstu_test/backend"
	go test ./... -tags=integration_test --race --parallel 10

e2e-tests:
	export ALLURE_OUTPUT_PATH="/Users/maximhalitov/Desktop/IU7/bmstu_test/backend"
	go test ./... -tags=e2e_test -v --race --parallel 10

ci-concat-reports:
	mkdir ${GITHUB_WORKSPACE}/allure-results
	cp ${GITHUB_WORKSPACE}/unit-allure/* ${GITHUB_WORKSPACE}/allure-results/
	cp ${GITHUB_WORKSPACE}/integration-allure/* ${GITHUB_WORKSPACE}/allure-results/
	cp ${GITHUB_WORKSPACE}/e2e-allure/* ${GITHUB_WORKSPACE}/allure-results/
	cp ${GITHUB_WORKSPACE}/backend/environment.properties ${GITHUB_WORKSPACE}/allure-results