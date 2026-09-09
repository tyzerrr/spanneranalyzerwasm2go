# spanneranalyzerwasm2go

**自動生成物。手で編集しない。**

[Cloud Spanner Emulator](https://github.com/GoogleCloudPlatform/cloud-spanner-emulator) の DDL 検証部分を
wasm32-wasip1 にビルドし、[wasm2go](https://github.com/goccy/wasm2go) で純 Go に変換したもの。
cgo も wasm 実行環境も不要。呼び出し口は [go-spanner-analyzer](https://github.com/tyzerrr/go-spanner-analyzer) にある。

| 由来 | 版 |
|---|---|
| cloud-spanner-emulator | `fc811a1`（PostgreSQL 方言・gRPC・google-cloud-cpp を外すパッチ適用。go-spanner-analyzer の `patches/` を参照） |
| ICU | 76.1（データは正規化・照合の基本・root ロケールのみに絞ったもの。`tools/icu-keep.txt`） |
| go-spanner-analyzer（生成時の設定） | `d328734` |
| wasm | 17.4 MB → この Go |

生成手順は go-spanner-analyzer の `Makefile` と `tools/`。

## ライセンス

Apache-2.0。上流（Google）の帰属表示は `NOTICE` と `THIRD_PARTY_NOTICES.txt`。
"Spanner" は Google LLC の商標で、本プロジェクトは Google と無関係。
