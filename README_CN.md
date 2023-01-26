# Wakatime Fetcher

这是一个简单的脚本，用于获取昨天的 wakatime 统计数据并将其保存到数据库中。

目前，它仅支持 sqlite3，但欢迎 pr。

## 用法

### 独立运行

如果要在自己的机器上运行，可以通过运行以下命令来实现：

```bash
git clone
cd wakatime-fetcher
go run wakatime <你的 api key>
```

这将在当前目录中创建一个名为 `wakatime.tar.gz` 的文件。其中包含 sqlite3 数据库。

### Workflow

如果要将其作为 workflow 运行，请先将其复制到私有存储库（或者如果你希望数据公开可用，则可以将其复制到公共存储库）。

#### GitHub 方式

你可以在存储库的右上角单击使用此模板按钮。这将创建一个与此存储库具有相同文件的新存储库。

#### Git 方式

你需要通过 [Github UI](https://github.com/new) 创建一个新的私有存储库。 （这里我们称之为 `private-repo`）

然后，你可以运行以下命令将此存储库复制到你的私有存储库：

```bash
git clone --bare https://github.com/foxwhite25/wakatime-fetcher.git
cd wakatime-fetcher.git
git push --mirror https://github.com/<your username>/private-repo.git
cd ..
rm -rf wakatime-fetcher.git
```

#### 设置 workflow

现在你有了私有存储库，你可以根据需要修改 workflow 文件。首先克隆你的私有存储库：

```bash
git clone https://github.com/<your username>/private-repo.git
cd private-repo
```

默认情况下，workflow 只在手动运行时运行。如果要将其设置为每天运行一次，请将以下行更改为：

```yaml
on:
  schedule:
    - cron: '0 0 * * *' # 每天凌晨运行
```

现在你可以提交更改并将其推送到你的私有存储库：

```bash
git add .
git commit -m "Change workflow to run on schedule"
git push origin master
```

你需要获取你的 api key。你可以在 [这里](https://wakatime.com/api-key) 找到它。

获得了 api key 后，你需要将其添加到你的存储库的 secrets 中。你可以在存储库的设置中找到它。将名称设置为 `WAKATIME_API_KEY` 并将值设置为你的 api key。

现在你可以去设定启用 workflow，并启用写入权限。

#### 从这个存储库拉取更新

如果你想从此存储库拉取更新，你可以运行以下命令：

```bash
cd private-repo
git remote add public https://github.com/foxwhite25/wakatime-fetcher.git
git pull public master # Creates a merge commit
git push origin master
```

## 数据库

数据库包含一个表 `heartbeats`，其中包含以下字段：

```sqlite
CREATE TABLE heartbeats (
    id TEXT PRIMARY KEY,
    branch TEXT, 
    category TEXT, 
    created_at TEXT,
    cursorpos INTEGER, 
    dependencies TEXT,
    entity TEXT, 
    is_write INTEGER, 
    language TEXT, 
    lineno INTEGER, 
    lines INTEGER, 
    machine_name_id TEXT, 
    project TEXT, 
    project_root_count INTEGER, 
    time REAL, 
    type TEXT, 
    user_agent_id TEXT, 
    user_id TEXT
);
```

## 数据分析

你可以使用数据库中的数据来做一些数据分析。

这些数据是心跳，这意味着它只包含你在编码时的时间，你需要做一些数据处理来得到你编码的总时间。

wakatime所做的是，它按项目和/或语言对心跳进行分组，然后对每个项目和语言所花费的时间进行汇总。

注意，这里的时间是指心跳的时间戳，而不是在项目上花费的时间。

在wakatime中，一个会话被定义为在同一个项目和/或语言上编码的时间段，当心跳停止超过5分钟时就会停止。
