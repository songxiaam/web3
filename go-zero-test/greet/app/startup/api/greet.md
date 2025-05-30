### 1. "创建订单"

1. route definition

- Url: /order/create
- Method: POST
- Request: `CreateOrderReq`
- Response: `CreateOrderRes`

2. request definition



```golang
type CreateOrderReq struct {
	ProductId uint `json:"productId"`
	Count uint `json:"count"`
}
```


3. response definition



```golang
type CreateOrderRes struct {
	OrderId uint `json:"orderId"`
}
```

### 2. "创建订单"

1. route definition

- Url: /order/detail
- Method: GET
- Request: `GetOrderDetailReq`
- Response: `GetOrderDetailRes`

2. request definition



```golang
type GetOrderDetailReq struct {
	OrderId uint `json:"orderId"`
}
```


3. response definition



```golang
type GetOrderDetailRes struct {
	OrderId uint `json:"orderId"`
	ProductId uint `json:"productId"`
	Count uint `json:"count"`
}
```

### 3. "创建用户"

1. route definition

- Url: /user/create
- Method: POST
- Request: `CreateUserReq`
- Response: `CreateUserRes`

2. request definition



```golang
type CreateUserReq struct {
	Name string `json:"name"`
}
```


3. response definition



```golang
type CreateUserRes struct {
	UserId uint `json:"userId"`
}
```

### 4. "获取用户"

1. route definition

- Url: /user/detail
- Method: GET
- Request: `GetUserDetailReq`
- Response: `GetUserDetailRes`

2. request definition



```golang
type GetUserDetailReq struct {
	UserId uint `form:"userId"`
}
```


3. response definition



```golang
type GetUserDetailRes struct {
	UserId uint `json:"userId"`
	Name string `json:"name"`
}
```

