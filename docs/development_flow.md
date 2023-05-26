# 全体の開発作業の流れ

1. 新しいIssueの作成か、既存のIssueからやることを選ぶ  
2. devブランチから実装する機能毎にfeatureブランチを作成
3. feature, docs, hotfixなどのブランチでのPRはdevelopブランチにスカッシュマージ(PRでやったことをコメントとして記述)
4. リリース時点でdev → mainマージ

## 登場するブランチ

- main
　リリースした時点のソースコードを管理するブランチ
- dev (mainから派生)
　開発作業の主軸となるブランチ
- hotfix (mainから派生)
　公開中のもののバグ修正用ブランチ (hotfix-<Issue_number>-◯◯ 例: hotfix-1-xxなど)
- feature (devから派生)
　実装する機能毎のブランチ (feature-<Issue_number>-◯◯ 例: feature-1-xxなど)
- docs (devから派生)
　ドキュメントに対しての記述ブランチ (docs-<Issue_number>-◯◯ 例: docs-1-xxなど)
- fix (devから派生)
　開発中のもののバグ修正用ブランチ (fix-<Issue_number>-◯◯ 例: fix-1-xxなど)

## Projectについて

Projectはタスク管理アプリのようなものです。  
Snackでは5つの状態があります。  
- In Considering
    考案中や実際に行うかわからないIssueを割り当てる
- Todo
    Projectでやる必要があるIssueを割り当てる
- In Progress
    ある人がIssueに対して何かしらしている(ドキュメントなら記述、機能開発なら開発中)  
    Assigneesに誰かが割り当てられていいる
- In Review
    PRが作成されており、レビュー待ちの状態
- Done
    Issueが解決された

## Issueについて

### Issueの向き合い方

作成したIssueには適切なLabels, Projectを追加してください。  
あるIssueを消化したい場合はAssigneesに自分を追加してください。  
都度IssueのProjectの状態を変更してください。  

### Issueの書き方

以下のように記載してください

```issue.md
## Why

なぜその変更などがしたいのか

## What

どのように変更するのか

```

## Labelの種類

- backend
    backendに対しての提案など
- desktop
    desktopに対しての提案など
- mobile
    mobile(iOS, Android)に対しての提案など
- CI
    CI(Github Action)に対しての提案など
- documentation
    ドキュメントの整備など
- enhancement
    新しい機能の追加など
- good first issue
    開発に入ったときに最初に消化しやすいIssue
- help wanted
    何か他の人の意見や実装のやり方がわからない時など
- question
    ある機能やドキュメントに対してわからないことがあった時など
- search
    何かしら障害なので調べる必要がある時など
- duplicate
    IssueやPRがすでに作られていた時など
- bug
    バグを見つけた時など

