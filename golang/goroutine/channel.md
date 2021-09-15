<font size="4">

# channel的特性
会panic的几种情况

1. 向已经关闭的channel发送数据

2. 关闭已经关闭的channel

3. 关闭未初始化的nil channel

会阻塞的情况：

1. 从未初始化nil channel中读数据

2. 向未初始化nil channel中发数据

3. 在没有读取的groutine时，向无缓冲channel发数据

4. 在没有数据时，从无缓冲channel读数据

返回零值:

从已经关闭的channe接收数据