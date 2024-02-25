# 開発者ガイド

## goreleaser

`tag`をつけたら`goreleaser`を使ってリリースするような仕組みにしている。詳しくは以下の公式情報を確認すると良い。

- [goreleaser - github](https://github.com/goreleaser/goreleaser)
- [Deprecation notices](https://goreleaser.com/deprecations/)
  - 古い記法など廃止された仕組みについて確認できるページ
  - `brew`で`goreleaser`コマンドをインストールしてチェックするのも良き（詳しくは後述）
- [Builds](https://goreleaser.com/customization/builds/)
  - ビルド関連の設定に関する説明。たまには確認しておくと吉。

### チェック

久しぶりに使おうとすると記法が廃止になっていたり、Tokenが切れていたりといろいろ躓く。自分用にメモを残す。

- `goreleaser`コマンドをインストールしてチェックする。
  - インストール: `brew install goreleaser/tap/goreleaser`
  - チェック: `goreleaser check`
- [act](https://github.com/nektos/act)コマンドでCI設定のチェック
  - `missing GITHUB_TOKEN`というエラー以外は正常っぽかったら良い感じ
