# POR

idではなくポートでdockerコンテナを指定して操作したい場面が多かったので作りました。もちろんパブリックポート（ホスト側のポート）です。

## インストール

[リリースページ](https://github.com/takumi3488/por/releases/)からダウンロード・展開し、`por`ファイルをPATHが通ったフォルダに移動してください。

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
