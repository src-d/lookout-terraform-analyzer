FROM alpine

ADD ./build/bin/lookout-terraform-analyzer /bin/lookout-terraform-analyzer

CMD ["/bin/lookout-terraform-analyzer"]
