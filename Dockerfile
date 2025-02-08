FROM golang:1.23 AS build
WORKDIR /build
COPY . .
RUN make build

FROM gcr.io/distroless/static-debian12 AS final
COPY --from=build /build/dist/gofile /
EXPOSE 8080
CMD [ "/gofile"]
