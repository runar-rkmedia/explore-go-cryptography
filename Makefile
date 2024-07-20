testkey := "deaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddead"

test:
	gotestsum  ./... -- -count=1 | ,colored-logs
test_watch:
	fd | entr -rc  gotestsum  ./... -- -count=1 | ,colored-logs
cli_test:
	./cli_test.sh

