# refer: https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/go-image.html
#
FROM alpine:3.13.2 as build

# install build tools
RUN apk add go git
RUN go env -w GOPROXY=direct

# cache dependencies
ADD go.mod go.sum ./
RUN go mod download

# build
ADD cmd/ cmd/
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app -o /main cmd/lambda_functions.go
RUN go build -o /main cmd/lambda_functions.go

# copy artifacts to a clean image
FROM alpine:3.13.2
# if WORKDIR is not set, overwrite WORKDIR in the AWS Lambda console
WORKDIR /
COPY --from=build /main /main
ENTRYPOINT [ "/main" ]
