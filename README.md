# パーソナルフォトギャラリー

個人のフォトコレクションを管理・展示するためのサーバーレスパーソナルフォトギャラリーアプリケーション。最新のテクノロジースタックを採用し、サーバーレスアーキテクチャのベストプラクティスを取り入れています。

## 使用技術

- **バックエンド:** サーバーレスバックエンドはGoとEchoフレームワークを使用して実装しています。イベント駆動の処理にはAWS Lambdaを、RESTful APIの提供にはAWS API Gatewayを使用しています。
- **フロントエンド:** ダイナミックでレスポンシブなフロントエンドは、ReactとTypeScriptを用いて構築され、サーバーサイドレンダリングにはNext.jsフレームワークが使用されています。
- **データベース:** 構造化データにはAWS Aurora MySQLを使用し、NoSQLデータの保存にはAWS DynamoDBを使用しています。
- **ストレージ:** 写真の安全でスケーラブルなストレージとしてAWS S3を使用しています。
- **CI/CD:** コードベースにはGitHub Actionsを使用した堅牢なCI/CDパイプラインが組み込まれており、生産性の向上とエラーの削減に寄与しています。
- **ローカル開発:** アプリケーションが異なる環境でも一貫して動作するように、Dockerを使用しています。

## 機能

- 写真のアップロード、閲覧、削除が可能なユーザーフレンドリーなインターフェース。
- 写真の詳細情報を表示。
- 特定の写真の詳細を更新。（あなたの機能セットに基づいて拡張）

## はじめに

このプロジェクトはDocker化された環境で動作するように設定されています。始めるには、[開発環境のセットアップガイド](#開発環境のセットアップ)をご覧ください！
Docker를 설정하는 방법과 Docker 내에서 프로젝트를 빌드하고 실행하는 방법 등을 포함하면 유용 할 수 있습니다.

## 開発環境のセットアップ

_近日公開..._

## 貢献

このプロジェクトの改善提案がある場合や、バグを報告したい場合は、問題を開いてください！あらゆる貢献をお待ちしています。

## ライセンス

----

# docekr 실행

```

# TODO
1. aws 어필 포인트 스샷으로 첨부 cloudwatch , cognito 등
```
personal-photo-gallery
├─ .aws-sam
│  ├─ build
│  │  ├─ RegisterUserFunction
│  │  │  └─ register.go
│  │  └─ template.yaml
│  └─ build.toml
├─ .env

├─ .vscode
│  └─ launch.json
├─ README.md
├─ aws
│  └─ cmd
│     ├─ photo
│     │  ├─ delete.go
│     │  ├─ get_detail.go
│     │  ├─ get_list.go
│     │  ├─ update.go
│     │  └─ upload.go
│     └─ user
│        ├─ confirm
│        │  └─ main.go
│        ├─ login
│        │  └─ main.go
│        ├─ logout
│        │  └─ main.go
│        └─ register
│           └─ main.go
├─ go.mod
├─ go.sum
├─ handlers
│  ├─ photo_handler.go
│  └─ user_handler.go
├─ main.go
├─ middleware
│  └─ jwtvalidation.go
├─ models
│  └─ user.go
├─ samconfig.toml
├─ services
│  ├─ ssm_service.go
│  ├─ ssm_service_impl.go
│  ├─ user_service.go
│  └─ user_service_impl.go
├─ template.yml
└─ tests
   └─ aws
      └─ lambda
         └─ request
            └─ register.json

```