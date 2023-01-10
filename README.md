## ポートフォリオ概要
- DUELSCORE
    - 海外サッカー試合情報サイト（試合日程、得点ランキング、順位表）
    - 試合情報はAPI（[football-data.org](https://www.football-data.org/)）を使い取得

## URL
※作成中

## テストアカウント
※作成中

## 主要技術
- 言語：Golang 1.17
- ORM：GORM 1.23.10
- JSONパーサー：GJSON 1.14.1
- DB：MySQL 8.0
- セッション管理：Redis 7.0.5
- バージョン管理：Github
- コンテナ管理：Docker/Docker-compose
- インフラ：AWS
    - VPC
    - Fargate
    - ECR
    - ELB
    - S3
    - IAM
    - Route53
    - CloudWatch
    - Secrets Manager
- IaC：Terraform

## 実装機能
- 試合日程取得機能
- 得点ランキング取得機能
- 順位表取得機能
- ユーザー管理機能(ユーザー登録/ログイン/ログアウト)

## 工夫した点
- 不要なAPIコールを防ぐ手段として、APIから取得した試合情報は30秒間Redisで一時保持し、使い回すようにした。
    
## 苦労した点
- Golangのプロジェクト構成に時間を費やした（参考書や参考サイトによって違うので迷った）。 
- Golangのエラーハンドリングの構成に時間を費やした（こちらも参考書や参考サイトによって違い、現在進行形で見直し中）。
- Redisの操作方法（セッション生成、破棄など）について記載されている日本語サイトがあまりなく、海外のコミュニティを参考にした。
- インフラ(AWS)やIaC（Terraform）の仕様理解に苦しむ。現在も試行錯誤中。

## 今後実装したい機能/実装中機能
- ユーザー管理機能（登録情報変更）
- チャット機能
- オッズ投票機能
- 取得情報追加（各国リーグ情報、選手情報など）

## 今後実施したい取り組み
- CI/CDパイプライン構築（CircleCI、またはGithubActionsを使用）
