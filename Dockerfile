FROM golang:alpine as build
RUN mkdir -p $GOPATH/src/github.com/srleyva/terraform-image
WORKDIR $GOPATH/src/github.com/srleyva/terraform-image
ADD . ./
RUN go build -o provider-gen main.go && cp provider-gen /

FROM sleyva97/base-layer:0.0.1
RUN curl -o terraform.zip https://releases.hashicorp.com/terraform/0.11.10/terraform_0.11.10_linux_amd64.zip && \
	unzip terraform.zip && \
	mv terraform /usr/local/bin && \
	rm -rf terraform*
COPY --from=build /provider-gen /usr/local/bin
ADD execute /usr/local/bin
