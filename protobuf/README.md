## スキーマ以外の仕様

protoで明記されていないメタデータなどの仕様

### clientからserverのrpcで満たすべき仕様

|||required|
|---|---|---|
|認証|grpc のリクエストメタデータ(header)に jwt: `token: String` を付与する|○|

### server の実装が満たすべき仕様

#### 認証

サーバーは以下のエンドポイント以外は全て `token` による認証をする必要がある.  
- auth.v1.Authentication/PublishToken

以下の場合サーバーは UNAUTHENTICATED を返す  
- token が無効・不正な場合

### serverからclientへのレスポンスが満たすべき仕様

#### エラーなどのステータスコード

|grpc status code|例|
|---|---|
|OK = 0|リクエストが成功したケース|
|UNKNOWN = 2|エラーが発生するが、リトライで直るケース|
|INVALID_ARGUMENT = 3|クライアントのリクエストに問題がありリトライしても意味がないケース|
|NOT_FOUND = 4|クライアントのファイルリクエストで該当ファイルが存在しないケース|
|PERMISSION_DENIED = 7|権限不足|
|UNIMPLEMENTED = 12|サービスがまだ実装されていないケース|
|INTERNAL = 12|ザーバー内処理でのエラーのケース|
|UNAUTHENTICATED = 16|ユーザーの認証に失敗したケース|

[Status Codes and their use in gRPC](https://grpc.github.io/grpc/core/md_doc_statuscodes.html)