# AGENTS.md

AI Agent がこのリポジトリで実装・レビューを行う際のガイドラインです。

## プロジェクト概要

**原神アーティファクト管理API** - Gin Webフレームワークを使用したGoによるRESTful API

### 技術スタック
- **言語**: Go 1.24
- **Webフレームワーク**: [Gin](https://github.com/gin-gonic/gin)
- **データストレージ**: インメモリ（JSON永続化）
- **テスト**: 標準`testing`パッケージ、`google/go-cmp`

## アーキテクチャ

```
cmd/genshin-artifact-db/  # エントリポイント
pkg/
├── entity/       # ドメインモデル（バリデーション付き）
├── handler/      # HTTPハンドラ（リクエスト/レスポンス処理）
├── repository/   # データアクセス層（InterfaceとInMemory実装）
├── service/      # ビジネスロジック（DTOとコマンドパターン）
└── server/       # HTTPサーバー管理
```

### レイヤー間の依存関係
```
Handler → Service → Repository → Entity
```
- 各層はインターフェースを通じて疎結合
- 依存性注入（DI）パターンを採用

## 開発ルール

### コマンド
| タスク | コマンド |
|-------|---------|
| ビルド | `go build ./cmd/genshin-artifact-db` |
| 実行 | `go run ./cmd/genshin-artifact-db/main.go` |
| テスト | `go test ./...` |
| Docker | `docker build -t genshin-artifact-db .` |

### 命名規則
- **パッケージ**: 小文字単語（`entity`, `handler`）
- **型**: PascalCase（`ArtifactType`, `GetArtifactService`）
- **定数**: SCREAMING_SNAKE_CASE（`ARTIFACT_TYPE_FLOWER`）
- **インターフェース**: 動詞 + "Interface"（`GetArtifactServiceInterface`）

### 新機能追加時のパターン

**1. 新規エンティティ追加**
- `pkg/entity/`に型定義とバリデーション付きコンストラクタ
- 適切なエラー型を定義（例: `ErrInvalidArtifactType`）

**2. 新規エンドポイント追加**
1. `pkg/service/`にService Interface + 実装
2. `pkg/handler/`にハンドラ関数（クロージャパターン）
3. `main.go`でルーティング登録

**3. 新規Repository実装**
- `pkg/repository/repository.go`のインターフェースを実装
- 既存: `InMemoryArtifactRepository`

## API エンドポイント

| Method | Path | 説明 |
|--------|------|------|
| GET | `/artifact/:id` | ID指定で取得 |
| GET | `/artifacts/type/:type` | タイプで検索 |
| GET | `/artifacts/set/:set` | セットで検索 |
| GET | `/artifacts/type/:type/set/:set` | タイプ+セットで検索 |
| POST | `/artifact` | 新規作成 |

## テスト

### テストパターン
- 各パッケージにモックが用意済み
  - `pkg/repository/mock_repository.go`
  - `pkg/service/mock_service.go`
- Table-Drivenテストを採用

### モック使用例
```go
mockService := &service.MockGetArtifactService{
    MockArtifact: &service.ArtifactDTO{...},
    MockGetArtifactError: nil,
}
```

## ドメイン知識

### アーティファクトタイプ
`FLOWER`, `PLUME`, `SANDS`, `GOBLET`, `CIRCLET`

### アーティファクトセット
`Gladiator`, `Wanderer`, `Noblesse`, `Bloodstained`, `Maiden`, `Vermillion`

### ステータスタイプ
- **Primary**: `ATK_PERCENT`, `HP_PERCENT`, `DEF_PERCENT`, `ELEMENTAL_MASTERY`, `CRIT_RATE`, `CRIT_DMG`, `ENERGY_RECHARGE`, `PHYSICAL_DMG_BONUS`, `ELEMENTAL_DMG_BONUS`, `HEALING_BONUS`
- **Substat**: `ATK_PERCENT`, `HP_PERCENT`, `DEF_PERCENT`, `ELEMENTAL_MASTERY`, `CRIT_RATE`, `CRIT_DMG`, `ENERGY_RECHARGE`

## 注意事項

- データは設定されたパスに永続化（シャットダウン時自動保存）
- IDは `crypto/rand.Text()` で自動生成
- Graceful Shutdown対応（SIGTERM/SIGINT）

## 設定

### 設定ファイル
`/etc/config/genshin-artifact-db/config.yaml`:
```yaml
port: ":8080"
data_file_path: "/var/lib/genshin-artifact-db/artifacts.json"
```

### コマンドライン引数
| 引数 | デフォルト | 説明 |
|------|-----------|------|
| `-config` | `/etc/config/genshin-artifact-db/config.yaml` | 設定ファイルパス |
| `-port` | (設定ファイル優先) | サーバーポート |
| `-data` | (設定ファイル優先) | データファイルパス |

### 優先順位
1. コマンドライン引数
2. 設定ファイル
3. デフォルト値

## Docker

### docker compose での起動
```bash
task compose:up     # 起動
task compose:down   # 停止
task compose:logs   # ログ表示
```

### ボリュームマウント
- `./data:/var/lib/genshin-artifact-db` - データ永続化
- `./config:/etc/config/genshin-artifact-db` - 設定ファイル

## Task コマンド

| タスク | 説明 |
|--------|------|
| `task build` | アプリケーションビルド |
| `task run` | ローカル起動 |
| `task test` | テスト実行 |
| `task test:verbose` | 詳細テスト |
| `task test:coverage` | カバレッジ付きテスト |
| `task docker:build` | Dockerイメージビルド |
| `task compose:up` | docker compose起動 |
| `task compose:down` | docker compose停止 |
| `task compose:logs` | ログ表示 |
| `task clean` | ビルド成果物削除 |
