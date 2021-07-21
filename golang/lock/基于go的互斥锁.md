<font size="4"> 

# 互斥锁
## 它的作用是守护在临界区入口来确保同一时间只能有一个线程进入临界区。
### 假设 info 是一个需要上锁的放在共享内存中的变量
```go
    import  "sync"
    type Info struct {
        mu sync.Mutex
        // ... other fields, e.g.: Str string
    }
```
如果一个函数想要改变这个变量可以这样写:
```go
func Update(info *Info) {
	info.mu.Lock()
    info.Str = // new value
    info.mu.Unlock()
}
```

# 使用互斥锁解决商品扣减库存不准确
### (缺点:性能很差,是单机锁不适用分布式系统)
```go
//声明锁必须在全局声明
var m synce.Mutex
func (i *Inventort) Sell(ctx context.Context, req *proto.SellInfo) (*empty.Empty, error) {
    m.Lock()
    //开启事务
	tx := global.DB.Begin()
	for _, goodInfo := range req.GoodsInvInfo {
        var invModel model.Inventory
        //todo查看库存信息
        if res := global.DB.Where("goods=?", goodInfo.GoodsId).Find(&invModel); res.RowsAffected == 0 {
            tx.Rollback()
            return nil, status.Errorf(codes.InvalidArgument, "库存信息不存在")
        }
        if invModel.Stocks < goodInfo.Num {
            tx.Rollback()
            return nil, status.Errorf(codes.ResourceExhausted, "库存不足")
        }
        //扣减
        invModel.Stocks -= goodInfo.Num
        tx.save(&invModel)
	}
	tx.Commit()
    m.Unlock() //解锁要在事务完成之后
	return &emptypb.Empty{}, nil
}
```