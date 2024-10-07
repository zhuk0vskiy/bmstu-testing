mkdir allure-results
go test .\service_test\ .\repository_test\
move .\repository_test\allure-results\* .\allure-results\
move .\service_test\allure-results\* .\allure-results\
allure generate allure-results -o allure-reports

allure serve allure-results -p 4000