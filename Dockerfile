FROM sleyva97/base-layer:0.0.1
RUN curl -o terraform.zip https://releases.hashicorp.com/terraform/0.11.10/terraform_0.11.10_linux_amd64.zip && \
	unzip terraform.zip && \
	mv terraform /usr/local/bin && \
	rm -rf terraform*
ADD provider.tf /tmp/provider.tf
ADD cue-execute /usr/local/bin
