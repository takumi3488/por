# POR

idではなくポートでdockerコンテナを指定して操作したい場面が多かったので作りました。もちろんパブリックポート（ホスト側のポート）です。

## 使い方

コンテナ検索

```
por ps <PORT>
```

コンテナ名変更

```
por rename <PORT> <NAME>
```

その他`logs`,`pause`,`restart`,`stop`,`unpause`が使えます。
