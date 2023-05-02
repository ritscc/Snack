# Snack


ファイル構成
```
├── .github                  --> GithubCI用の設定
├── Makefile
├── README.md
├── RDD.md                   --> 要件定義書
├── docs                     --> 自動生成されたDocumentationの保存フォルダ
├── backend 
├── desktop
├── mobile
├── protobuf                 --> .protoファイルの保存フォルダ
│   ├── authentication       --> authenticationという名のプロジェクト
│   │   └── v1               --> バージョンを管理
│   │       └── authentication.proto
│   ├── channel
│   │   └── v1
│   │       └── channel.proto
│   ├── event
│   │   └── v1
│   ├── message
│   │   └── v1
│   │       └── message.proto
│   ├── server_status.proto
│   ├── stamp
│   │   └── v1
│   │       └── stamp.proto
│   ├── user
│   │   └── v1
│   │       └── user.proto
│   ├── user_group
│   │   └── v1
│   │       └── user_group.proto
│   └── utils
└── web
```
