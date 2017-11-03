# https://vincent.bernat.im/en/blog/2017-makefile-build-golang
PACKAGE  = pizza-api
GOPATH   = $(CURDIR)/.gopath
BASE     = $(GOPATH)/src/$(PACKAGE)

# simple makefile to build and push docker container images
IMAGE_NAME = sebastianhutter/wallets

# If the GoCD label is set overwrite the commit ID env variable
ifneq ($(GO_PIPELINE_LABEL),"")
export COMMIT_ID := $(GO_PIPELINE_LABEL)
endif


# create necessary folder structure
$(BASE):
	@mkdir -p $(dir $@)
	@ln -sf $(CURDIR) $@

# compile go app
compile: | $(BASE)
	cd $(BASE) && go get && go build -o bin/$(PACKAGE) main.go

# build a new docker image
build_commit:
	docker build -t $(IMAGE_NAME):$(COMMIT_ID) .
	echo $(COMMIT_ID) > image_version
	echo $(IMAGE_NAME) > image_name

# latest
# set the latest tag for the image with the specified nextcloud version tag
build_latest:
	docker build -t $(IMAGE_NAME):latest .

# push the commit tag
push_commit:
	docker push $(IMAGE_NAME):$(COMMIT_ID)

# push the build containers
push_latest:
	docker push $(IMAGE_NAME):latest

# deploy commit id image to rancher
deploy_commit:
	rancher_upgrade.sh $(SERVICE_NAME) $(IMAGE_NAME) $(COMMIT_ID)