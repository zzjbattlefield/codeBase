<font size="4">

## Git Add
- git add 缓存更改
- git commit 提交项目历史
- git reset 撤销提交/缓存的快照

## 检查仓库状态
### git log
1. 用 <limit> 限制提交的数量。比如 git log -n 3 只会显示 3 个提交。
```git
git log -n <limit>
```
2. 将每个提交压缩到一行。
```
git log --oneline
```
3. 除了 git log 信息之外，包含哪些文件被更改了，以及每个文件相对的增删行数。
```
git log -stat
```
4. 显示代表每个提交的一堆信息。显示每个提交全部的差异（diff），这也是项目历史中最详细的视图。
```
git log -p
```
5. 搜索特定作者的提交/搜索提交信息匹配特定 
```
git log --author=<patten>/--grep=<patten>
```
6. 只显示包含特定文件的提交。查找特定文件的历史这样做会很方便。
```
git log <file>
```
7. --graph 标记会绘制一幅字符组成的图形，左边是提交，右边是提交信息。--decorate 标记会加上提交所在的分支名称和标签。--oneline 标记将提交信息显示在同一行，一目了然。
```
git log --graph --decorate --oneline
```

## git checkout
git checkout 这个命令有三个不同的作用：检出文件、检出提交和检出分支。

检出提交会使工作目录和这个提交完全匹配。你可以用它来查看项目之前的状态，而不改变当前的状态。检出文件使你能够查看某个特定文件的旧版本，而工作目录中剩下的文件不变。

1. 返回查看commit状态 一旦你回到 master 分支之后，你可以使用 git revert 或 git reset 来回滚任何不想要的更改。
```
git checkout <commit-id>
```
2. 如果你只对某个文件感兴趣，你也可以用 git checkout 来获取它的一个旧版本。比如说，如果你只想从之前的提交中查看 hello.py 文件，你可以使用下面的命令：
```
git checkout a1e8fb5 hello.py
```

## 回滚错误的修改
