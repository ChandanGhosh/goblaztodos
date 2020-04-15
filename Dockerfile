FROM mcr.microsoft.com/dotnet/core/sdk as dotnetbuilder
RUN mkdir -p /build
WORKDIR /build
ADD .  /build/
RUN dotnet publish blazfront/blazfront.csproj -c release -o dist

FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build/goback 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /build/main .
WORKDIR /build
FROM scratch
COPY --from=builder /build/main /app/
COPY --from=dotnetbuilder /build/dist /app/dist
EXPOSE 3000
WORKDIR /app
CMD ["./main"]