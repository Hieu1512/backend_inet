# Sử dụng image golang để build ứng dụng
FROM golang:latest AS build

# Sao chép mã nguồn vào container
COPY . /app

# Đặt thư mục làm việc mặc định
WORKDIR /app

# Build ứng dụng
RUN go build -o myapp .

# Sử dụng image scratch để tạo image nhẹ
FROM scratch

# Sao chép file binary từ stage build vào container
COPY --from=build /app/myapp /

# Chạy ứng dụng khi container khởi động
CMD ["/myapp"]
