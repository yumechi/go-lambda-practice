# refer: https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/go-image.html
FROM alpine:latest as build

# install build tools
RUN apk add go git
RUN go env -w GOPROXY=direct

# cache dependencies
ADD go.mod go.sum ./
RUN go mod download

# build
ADD cmd/ cmd/
RUN go build -o /main cmd/lambda_functions.go

# copy artifacts to a clean image
FROM alpine:latest
COPY --from=build /main /main
# if WORKDIR is not set, overwrite WORKDIR in the AWS Lambda console
WORKDIR /

# (Optional) Add Lambda Runtime Interface Emulator and use a script in the ENTRYPOINT for simpler local runs
ADD https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie /usr/local/bin/aws-lambda-rie
RUN chmod 755 /usr/local/bin/aws-lambda-rie
COPY entry.sh /
RUN chmod 755 /entry.sh
ENTRYPOINT [ "/entry.sh" ]