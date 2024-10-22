go-test-command:
	go test -v ./...

#test:
#	rm -rf allure-results
#	export ALLURE_OUTPUT_PATH="/Users/dmitry/Desktop/bmstu/7sem/bmstu-test/backend" && go test ./... --race --parallel 8
#	cp environment.properties allure-results

ci-unit:
	export ALLURE_OUTPUT_PATH="${GITHUB_WORKSPACE}" && \
 	export ALLURE_OUTPUT_FOLDER="unit-allure" && \
 	go test -tags=unit ./... --race --parallel 8

ci-integration:
	export ALLURE_OUTPUT_PATH="${GITHUB_WORKSPACE}" && \
	export ALLURE_OUTPUT_FOLDER="integration-allure" && \
	go test -tags=integration ./... --race --parallel 8

ci-e2e:
	export ALLURE_OUTPUT_PATH="${GITHUB_WORKSPACE}" && \
	export ALLURE_OUTPUT_FOLDER="e2e-allure" && \
	go test -tags=e2e ./... --race --parallel 8

ci-concat-reports:
	mkdir allure-results
	cp unit-allure/* allure-results/
	cp integration-allure/* allure-results/
	cp e2e-allure/* allure-results/
	cp environment.properties allure-results

integration-tests:
	export ALLURE_OUTPUT_PATH="/Users/dmitry/Desktop/bmstu/7sem/bmstu-test/backend"
	go test ./tests/integration_tests --race --parallel 8

e2e-tests:
	export ALLURE_OUTPUT_PATH="/Users/dmitry/Desktop/bmstu/7sem/bmstu-test/backend"
	go test ./tests/e2e --race --parallel 8
	cp environment.properties allure-results

allure:
	cp -R allure-reports/history allure-results
	rm -rf allure-reports
	allure generate allure-results -o allure-reports
	allure serve allure-results -p 4000

report: test allure

.PHONY: go-test-command test allure report