## 介绍

特殊时期要求每日打卡，抓包实现自动打卡。

**2022年12月13日** 已放开，本仓库仅用来保存代码。

<a href="https://hub.docker.com/r/zggsong/luxshare-daily-report">
  <img alt="Docker pull" src="https://img.shields.io/docker/pulls/zggsong/luxshare-daily-report">
</a>

## 使用

```shell
curl -X POST -F 'file=@xcm.jpeg' -F 'name=xcm.jpeg' 127.0.0.1:7201/upload
```

## 作者

**luxshare-daily-report** © [zggsong](https://github.com/zggsong), Released under the [MIT](https://github.com/ZGGSONG/luxshare-daily-report/blob/main/LICENSE) License.<br>
Thanks for the article [跨平台构建 Docker 镜像](https://cloud.tencent.com/developer/article/1543689).

> Website [Blog](https://www.zggsong.com) · GitHub [@zggsong](https://github.com/zggsong)
