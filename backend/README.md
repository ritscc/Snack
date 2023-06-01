# バックエンド

---

## ファイル構成
```
├── README.md
├── adaptor
│   ├── event
│   ├── repository
│   └── service_server
├── domain
│   ├── model
│   └── service
├── driver
└── logic

```

## 各層の説明
### domain
ドメイン層はビジネスロジックを扱うレイヤーです。原則、他の層はドメインに依存し、逆にdomainは他の層に依存してはいけません。
これはビジネスロジックと技術を分離するためであり、ビジネスロジックの流出は厳禁です。

ドメインは以下の要素を持っています。
- model: 各エンティティとビジネスロジック、リポジトリのinterfaceを扱う。
- service: エンティティが扱うには不自然なビジネスロジックを扱う。   ただし、ドメイン欠乏症にならないようにしてください。

### logic(use case)
いわゆるアプリケーションロジックを扱うレイヤーです。この層はドメイン層に依存します。    
基本的にビジネスロジックはこの層で実行します。また、この層もDBなどの詳細に関与することは避けましょう。

### adaptor
GatewayやControllerを扱うレイヤーです。  
代表的なGatewayにはリポジトリがあり、ドメイン層で定義したリポジトリの実装はここで行うと良いでしょう。

### driver
ルーティングやDBなどのコネクション等を生成するレイヤーです。

### 依存関係の図解
![依存関係](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)  
https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

## テストについて
基本的には対象ファイルと同じディレクトリにテストファイルを置いてください。  
例:
```
├── sample.go
└── sample_test.go
```

リポジトリの単体テストと統合テストはテスト用のデータベースを用意して行ってください。
その他の単体テストについては基本的にリポジトリのモックを使用することを推奨します。
