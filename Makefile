# Builds project into local docker image
.PHONY: install
install: 
	docker build -t avdfd .

# Start the service in a docker container
.PHONY: start
start: 
	docker run -i -t avdfd bash -c '/observer || sleep 1000'

# Stops the container
.PHONY: stop
stop: 
	docker stop avdfd

# Runs the unit tests in the image
.PHONY: test
test: 
	docker build -f test/Dockerfile .