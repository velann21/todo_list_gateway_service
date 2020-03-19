setup:
	if [ -e ./dist ]; then rm -rf ./dist; fi; mkdir ./dist; mkdir -p ./dist/darwin/bin; mkdir -p ./dist/linux/bin

build:
	@echo "******************"
	@echo "**** Building ****"
	@echo "******************"

	GO111MODULE=on GOOS=linux go build -mod=vendor ../../cmd/gateway_bootstrap.go;cp dist/linux/bin/gateway_bootstrap ../docker;mv ./gateway_bootstrap dist/linux/bin/gateway_bootstrap


docker-build-prep:
	cd docker;cp ../dist/linux/bin/gateway_bootstrap .

docker-build-cleanup:
	cd ../docker;rm gateway_bootstrap

release-tag: build

	cd ../; git tag `cat .version`
	git push --tags

release: release-tag
	cd ../docker; docker build -t "singaravelan21/gateway_bootstrap" . --no-cache;
	cd ../docker; docker tag "singaravelan21/gateway_bootstrap" "singaravelan21/gateway_bootstrap:`cat ../.version`";
	cd ../docker; docker push "singaravelan21/gateway_bootstrap:`cat ../.version`"
	make docker-build-cleanup

